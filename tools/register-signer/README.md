# Signer registration utility

This utility can be used to register an op-enclave signer key with the
[SystemConfigGlobal](../contracts/src/SystemConfigGlobal.sol) contract.

## Installation

```
go install github.com/base/op-enclave/register-signer
```

## Usage

Query an AWS Nitro attestation from op-enclave server:
```bash
curl -d '{"id":0,"jsonrpc":"2.0","method":"enclave_signerAttestation"}' -H "Content-Type: application/json" http://op-enclave:7333
```

```
Usage of register-signer:
  -attestation string
    	attestation hex
  -deployment string
    	deployment file (default "deployments/84532-deploy.json")
  -private-key string
    	private key
  -rpc string
    	rpc url (default "https://sepolia.base.org")
```
