// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {Script} from "forge-std/Script.sol";
import {console2 as console} from "forge-std/console2.sol";
import {Config} from "@eth-optimism-bedrock/scripts/libraries/Config.sol";
import {CertManager} from "@marlinprotocol/CertManager.sol";
import {NitroValidator} from "../src/NitroValidator.sol";

/// @notice will deploy the singleton NitroValidatorContract to a deterministic address
contract DeployNitroValidator is Script {
    function run() public {
        vm.startBroadcast();

        CertManager manager = new CertManager{salt: _implSalt()}();
        NitroValidator validator = new NitroValidator{salt: _implSalt()}(manager);

        console.log("CertManager deployed at:", address(manager));
        console.log("NitroValidator deployed at:", address(validator));

        string memory deploymentOutfile =
            string.concat(vm.projectRoot(), "/deployments/", vm.toString(block.chainid), "-validator.json");
        vm.writeJson({json: vm.serializeAddress("", "CertManager", address(manager)), path: deploymentOutfile});
        vm.writeJson({json: vm.serializeAddress("", "NitroValidator", address(validator)), path: deploymentOutfile});

        vm.stopBroadcast();
    }

    function _implSalt() internal view returns (bytes32) {
        return keccak256(bytes(Config.implSalt()));
    }
}
