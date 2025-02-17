# PCR0 Extractor

This tool extracts the PCR0 measurement from an op-enclave EIF (Enclave Image Format) file. The PCR0 measurement is a cryptographic hash that represents the initial state of the enclave, which is crucial for attestation and verification purposes.

## Prerequisites

- Docker installed on your system
- Access to the op-enclave container registry

## Building and Running

1. Build the PCR0 extractor container:
```bash
docker build -f Dockerfile -t pcr0-extractor .
```

2. Run the container to extract the PCR0:
```bash
docker run --rm pcr0-extractor
```

The tool will:
1. Download the specified op-enclave EIF
2. Extract it using AWS Nitro CLI tools
3. Output the PCR0 measurement

## How it Works

The tool uses a multi-stage Docker build to:
1. Build required tools (skopeo and umoci)
2. Download and extract the op-enclave EIF
3. Use AWS Nitro CLI tools to extract the PCR0 measurement

The output will be a hex string representing the PCR0 measurement of the enclave.

## Note

The PCR0 measurement is specific to the version of the op-enclave EIF being examined. The current version being used is specified in the Dockerfile as `TAG=v0.0.1-rc5`.
