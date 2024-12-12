guard-%:
	@ if [ "${${*}}" = "" ]; then echo "Environment variable $* not set" && exit 1; fi

define abigen
	echo "Generating bindings for $(1)"
	cp contracts/out/$(1).sol/$(1).$(3).json contracts/out/$(1).sol/$(1).json 2>/dev/null || true
	jq -r '.bytecode.object' contracts/out/$(1).sol/$(1).json > contracts/out/$(1).sol/$(1).bin
	jq -r '.abi' contracts/out/$(1).sol/$(1).json > contracts/out/$(1).sol/$(1).abi
	abigen --abi contracts/out/$(1).sol/$(1).abi --bin contracts/out/$(1).sol/$(1).bin --pkg bindings --type $(1) --out bindings/$(2).go
endef

define verify
	deploy=$(1); \
	version=$(2); \
	addresses=$$(jq -r '.transactions[] | select(.transactionType=="CREATE" or .transactionType=="CREATE2") | .contractAddress' $$deploy); \
	for address in $$addresses; do \
		name=$$(jq -r --arg address "$$address" '.transactions[] | select((.transactionType=="CREATE" or .transactionType=="CREATE2") and .contractAddress==$$address) | .contractName' $$deploy); \
		arguments=$$(jq -r --arg address "$$address" '.transactions[] | select((.transactionType=="CREATE" or .transactionType=="CREATE2") and .contractAddress==$$address) | .arguments // [] | join(" ")' $$deploy); \
		namewithoutversion=$${name%.*.*.*}; \
		constructor=$$(jq '.abi[] | select(.type=="constructor")' contracts/out/$$namewithoutversion.sol/$$name.json | jq -r '.inputs | map(.type) | join(",")'); \
		echo; \
		echo "Verifying $$namewithoutversion @ $$address using constructor($$constructor) $$arguments"; \
		constructor_args=$$(cast abi-encode "constructor($$constructor)" $$arguments); \
		cd contracts; \
		forge verify-contract --compiler-version $$version --watch --verifier-url https://api-sepolia.basescan.org/api --constructor-args $$constructor_args $$address $$namewithoutversion; \
		cd ..; \
	done
endef

.PHONY: bindings
bindings:
	go install github.com/ethereum/go-ethereum/cmd/abigen@v1.14.11
	cd contracts && forge build
	mkdir -p bindings
	@$(call abigen,"OutputOracle","output_oracle","0.8.15")
	@$(call abigen,"Portal","portal","0.8.15")
	@$(call abigen,"DeployChain","deploy_chain","0.8.15")
	@$(call abigen,"CertManager","cert_manager","0.8.15")
	@$(call abigen,"SystemConfigGlobal","system_config_global","0.8.15")
	@$(call abigen,"GnosisSafe","gnosis_safe","0.8.15")

.PHONY: deploy-cert-manager
deploy-cert-manager: guard-IMPL_SALT guard-DEPLOY_PRIVATE_KEY guard-RPC_URL
	@cd contracts && forge script DeployCertManager --rpc-url $(RPC_URL) \
		--private-key $(DEPLOY_PRIVATE_KEY) --broadcast

.PHONY: deploy
deploy: guard-IMPL_SALT guard-DEPLOY_CONFIG_PATH guard-DEPLOY_PRIVATE_KEY guard-RPC_URL
	@cd contracts && forge script DeploySystem --sig deploy --rpc-url $(RPC_URL) \
		--private-key $(DEPLOY_PRIVATE_KEY) --broadcast

.PHONY: deploy-deploy-chain
deploy-deploy-chain: guard-IMPL_SALT guard-DEPLOY_PRIVATE_KEY guard-RPC_URL
	@cd contracts && forge script DeployDeployChain --rpc-url $(RPC_URL) \
		--private-key $(DEPLOY_PRIVATE_KEY) --broadcast

.PHONY: testnet
testnet: guard-L1_URL guard-DEPLOY_PRIVATE_KEY
	DEPLOY_CHAIN_ADDRESS=$${DEPLOY_CHAIN_ADDRESS:-$$(jq -r ".DeployChain" contracts/deployments/84532-deploy.json)} \
		go run ./testnet

.PHONY: verify
verify:
	@$(call verify,"contracts/broadcast/DeployCertManager.s.sol/84532/run-1733890597.json","0.8.24")
	@$(call verify,"contracts/broadcast/DeploySystem.s.sol/84532/run-1733867021.json","0.8.15")
	@$(call verify,"contracts/broadcast/DeployDeployChain.s.sol/84532/run-1733884066.json","0.8.15")
