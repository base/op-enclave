// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import {Test, console} from "forge-std/Test.sol";

import "@nitro-validator/src/INitroValidator.sol";
import "../src/SystemConfigGlobal.sol";

// Mock NitroValidator contract
contract MockNitroValidator is INitroValidator {
    bytes public pcr0;
    bytes public publicKey;

    function setValidationResult(bytes memory _publicKey, bytes memory _pcr0) external {
        pcr0 = _pcr0;
        publicKey = _publicKey;
    }
    function validateAttestation(bytes memory, uint256) external view returns (bytes memory, bytes memory) {
        return (publicKey, pcr0);
    }
}

contract NitroValidatorTest is Test {
    MockNitroValidator mockValidator;
    SystemConfigGlobal systemConfigGlobal;

    function setUp() public {
        // Create a mock contract
        mockValidator = new MockNitroValidator();
    
        // create system config global
        systemConfigGlobal = new SystemConfigGlobal(mockValidator);
    }

    function test_validateAttestation() public {
        vm.startPrank(systemConfigGlobal.owner());

        mockValidator.setValidationResult(hex"d239fd059dd0e0a01e280bec44903bb8143bae7e578b9844c6df5fd6351eddc0", hex"17BF8F048519797BE90497001A7559A3D555395937117D76F8BAAEDF56CA6D97952DE79479BC0C76E5D176D20F663790");

        string memory attestationFile = vm.readFile("./test/nitro-attestation/sample_attestation.json");
        bytes memory attestation = abi.decode(vm.parseJson(attestationFile, ".attestation"), (bytes));

        systemConfigGlobal.registerPCR0(hex"17BF8F048519797BE90497001A7559A3D555395937117D76F8BAAEDF56CA6D97952DE79479BC0C76E5D176D20F663790");
        
        systemConfigGlobal.registerSigner(attestation);

        address expectedSigner = 0xe04d808785d2BBdE18E9D0C01c05FB8CE0711f2d;
        assertTrue(systemConfigGlobal.validSigners(expectedSigner));
    }
}
