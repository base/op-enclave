// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import {Test, console} from "forge-std/Test.sol";
import "../src/INitroValidator.sol";
import "../src/NitroValidator.sol";

contract NitroValidatorTest is Test {
    INitroValidator validator;

    function setUp() public {
        vm.warp(1708930774);
        validator = new NitroValidator(); 
    }

    function test_validateAttestation() public {
        bytes memory attestation = vm.readFileBinary("./test/nitro-attestation/sample_attestation.bin");
        
        (bytes memory enclavePubKey, bytes memory pcr0) = validator.validateAttestation(attestation, 365 days);

        assertEq(enclavePubKey, hex"d239fd059dd0e0a01e280bec44903bb8143bae7e578b9844c6df5fd6351eddc0");
        assertEq(pcr0, hex"17BF8F048519797BE90497001A7559A3D555395937117D76F8BAAEDF56CA6D97952DE79479BC0C76E5D176D20F663790");
    }

    function test_validateAttestation_RevertOnExpiredTime() public {
        bytes memory attestation = vm.readFileBinary("./test/nitro-attestation/sample_attestation.bin");
        
        // Warp time to 366 days in the future
        vm.warp(block.timestamp + 366 days);

        vm.expectRevert("certificate not valid anymore");
        validator.validateAttestation(attestation, 365 days);
    }
}
