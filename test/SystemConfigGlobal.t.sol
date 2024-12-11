// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import {Test, console} from "forge-std/Test.sol";
import {CertManager} from "@nitro-validator/CertManager.sol";

import "../src/SystemConfigGlobal.sol";

contract SystemConfigGlobalTest is Test {
    SystemConfigGlobal systemConfigGlobal;

    function setUp() public {
        vm.warp(1708930774);
        CertManager certManager = new CertManager();
        systemConfigGlobal = new SystemConfigGlobal(certManager);
    }

    function test_validateAttestation() public {
        vm.startPrank(systemConfigGlobal.owner());

        systemConfigGlobal.registerPCR0(
            hex"17BF8F048519797BE90497001A7559A3D555395937117D76F8BAAEDF56CA6D97952DE79479BC0C76E5D176D20F663790"
        );

        bytes memory attestation = vm.readFileBinary("./test/nitro-attestation/sample_attestation.bin");
        (bytes memory attestationTbs, bytes memory signature) = systemConfigGlobal.decodeAttestationTbs(attestation);
        systemConfigGlobal.registerSigner(attestationTbs, signature);

        address expectedSigner = 0x874a4c5675cd4850dB08bD9A1e3184ED239087e4;
        assertTrue(systemConfigGlobal.validSigners(expectedSigner));
    }
}
