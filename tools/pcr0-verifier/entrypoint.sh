#!/bin/bash

echo "Starting PCR0 extraction..."
echo "Checking if EIF file exists:"
ls -l /app/eif.bin

echo "Command used: nitro-cli describe-eif --eif-path /app/eif.bin"
echo "PCR0 measurement:"
PCR0=$(nitro-cli describe-eif --eif-path /app/eif.bin | tee /dev/stderr | jq -r ".Measurements.PCR0")
PCR0_WITH_PREFIX="0x${PCR0}"

echo -e "\nTo verify against SystemConfigGlobal contract, run:"
echo "# First set your environment variables:"
echo "export SYSTEM_CONFIG_GLOBAL_ADDRESS=<contract_address>"
echo "export RPC_URL=<rpc_url>"
echo -e "\n# Then run these commands to verify:"
echo "# To register a new PCR0 (requires owner access):"
echo "cast send \$SYSTEM_CONFIG_GLOBAL_ADDRESS 'registerPCR0(bytes)' ${PCR0_WITH_PREFIX} --rpc-url \$RPC_URL"
echo -e "\n# To check if a PCR0 is valid:"
echo "cast call \$SYSTEM_CONFIG_GLOBAL_ADDRESS 'validPCR0s(bytes32)' 0x\$(cast keccak \${PCR0}) --rpc-url \$RPC_URL"
