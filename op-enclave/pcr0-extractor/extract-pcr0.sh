#!/bin/bash
set -x

echo "Starting PCR0 extraction..."
echo "Checking if EIF file exists:"
ls -l /app/eif.bin

echo "Command used: nitro-cli describe-eif --eif-path /app/eif.bin"
echo "PCR0 measurement:"
nitro-cli describe-eif --eif-path /app/eif.bin | tee /dev/stderr | jq -r ".Measurements.PCR0"