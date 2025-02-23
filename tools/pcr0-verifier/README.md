# PCR0 Verifier

This tool gets the PCR0 measurement from an op-enclave EIF (Enclave Image Format) file. The PCR0 measurement is a cryptographic hash that represents the initial state of the enclave, which is crucial for attestation and verification purposes.

## Prerequisites

- Docker installed on your system
- [cast](https://book.getfoundry.sh/cast/) (for contract verification)

## Building and Running

1. Build the PCR0 verifier container:
```bash
docker build -t pcr0-verifier .
```

2. Run the container to extract the PCR0:
```bash
docker run --rm pcr0-verifier
```

The tool will:
1. Download the specified op-enclave EIF
2. Extract it using AWS Nitro CLI tools
3. Output the PCR0 measurement
4. Generate a cast command to verify the PCR0 against the SystemConfigGlobal contract

## Verifying PCR0 Against SystemConfigGlobal Contract

The tool will output a `cast` command that you can use to verify if the PCR0 is registered in the SystemConfigGlobal contract. The command will look like:

```bash
export SYSTEM_CONFIG_GLOBAL_ADDRESS=<contract_address>
cast call $SYSTEM_CONFIG_GLOBAL_ADDRESS 'validPCR0s(bytes32)' <keccak256_hash_of_pcr0>
```

Note that the contract stores the keccak256 hash of the PCR0 value, not the raw PCR0 measurement. The tool automatically converts the PCR0 to the correct format for verification.

## How it Works

The tool uses a multi-stage Docker build to:
1. Build required tools (skopeo and umoci)
2. Download and extract the op-enclave EIF
3. Use AWS Nitro CLI tools to extract the PCR0 measurement
4. Convert the PCR0 to a keccak256 hash for contract verification

The output will include both the raw PCR0 measurement and instructions for verifying it against the contract.

## Note

The PCR0 measurement is specific to the version of the op-enclave EIF being examined. The current version being used is specified in the Dockerfile as `TAG=v0.0.1-rc5`. You can perform the same measurement on other EIF files by modifying the Dockerfile.
