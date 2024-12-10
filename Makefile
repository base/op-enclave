guard-%:
	@ if [ "${${*}}" = "" ]; then echo "Environment variable $* not set" && exit 1; fi

define abigen
	echo "Generating bindings for $(1)"
    jq -r '.bytecode.object' out/$(1).sol/$(1).json > out/$(1).sol/$(1).bin
	jq -r '.abi' out/$(1).sol/$(1).json > out/$(1).sol/$(1).abi
	abigen --abi out/$(1).sol/$(1).abi --bin out/$(1).sol/$(1).bin --pkg bindings --type $(1) --out bindings/$(2).go
endef

.PHONY: bindings
bindings:
	go install github.com/ethereum/go-ethereum/cmd/abigen@v1.14.11
	forge build
	mkdir -p bindings
	@$(call abigen,"OutputOracle","output_oracle")
	@$(call abigen,"Portal","portal")
	@$(call abigen,"DeployChain","deploy_chain")

.PHONY: deploy-cert-manager
deploy-cert-manager: guard-IMPL_SALT guard-DEPLOY_PRIVATE_KEY guard-RPC_URL
	@forge script DeployCertManager --rpc-url $(RPC_URL) \
		--private-key $(DEPLOY_PRIVATE_KEY) --broadcast

.PHONY: deploy
deploy: guard-IMPL_SALT guard-DEPLOY_CONFIG_PATH guard-DEPLOY_PRIVATE_KEY guard-RPC_URL
	@forge script DeploySystem --sig deploy --rpc-url $(RPC_URL) \
		--private-key $(DEPLOY_PRIVATE_KEY) --broadcast

.PHONY: testnet
testnet: guard-L1_URL guard-DEPLOY_PRIVATE_KEY
	DEPLOY_CHAIN_ADDRESS=$${DEPLOY_CHAIN_ADDRESS:-$$(jq -r ".DeployChain" deployments/84532-deploy.json)} \
		go run ./testnet

.PHONY: verify
verify:
	deploy=broadcast/DeploySystem.s.sol/84532/run-1733867021.json; \
	addresses=$$(jq -r '.transactions[] | select(.transactionType=="CREATE" or .transactionType=="CREATE2") | .contractAddress' $$deploy); \
	for address in $$addresses; do \
		name=$$(jq -r --arg address "$$address" '.transactions[] | select((.transactionType=="CREATE" or .transactionType=="CREATE2") and .contractAddress==$$address) | .contractName' $$deploy); \
		arguments=$$(jq -r --arg address "$$address" '.transactions[] | select((.transactionType=="CREATE" or .transactionType=="CREATE2") and .contractAddress==$$address) | .arguments // [] | join(" ")' $$deploy); \
		constructor=$$(jq '.abi[] | select(.type=="constructor")' out/$$name.sol/$$name.json | jq -r '.inputs | map(.type) | join(",")'); \
		echo "Verifying $$name @ $$address using constructor($$constructor) $$arguments"; \
		constructor_args=$$(cast abi-encode "constructor($$constructor)" $$arguments); \
		forge verify-contract --watch --verifier-url https://api-sepolia.basescan.org/api --constructor-args $$constructor_args $$address $$name ; \
	done
