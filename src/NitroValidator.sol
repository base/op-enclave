// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import {CBORDecoding} from "@solidity-cbor/CBORDecoding.sol";
import {INitroValidator} from "./INitroValidator.sol";
import {CertManager} from "@marlinprotocol/CertManager.sol";
import {NitroProver} from "@marlinprotocol/NitroProver.sol";

contract NitroValidator is NitroProver, INitroValidator {
    constructor(CertManager certManager) NitroProver(certManager) {}

    function validateAttestation(bytes memory attestation, uint256 maxAge)
        external view
        returns (bytes memory, bytes memory)
    {
        (bytes memory enclaveKey, , bytes memory rawPcrs) = verifyAttestation(attestation, maxAge);
        bytes memory pcr0 = CBORDecoding.decodeMappingGetValue(rawPcrs, hex"00");
        return (enclaveKey, pcr0);
    }
}
