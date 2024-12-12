// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {Script} from "forge-std/Script.sol";
import {console2 as console} from "forge-std/console2.sol";
import {Config} from "@eth-optimism-bedrock/scripts/libraries/Config.sol";
import {CertManager} from "@nitro-validator/CertManager.sol";

/// @notice will deploy the singleton CertManager to a deterministic address
contract DeployCertManager is Script {
    function run() public {
        vm.startBroadcast();

        CertManager manager = new CertManager{salt: _implSalt()}();

        console.log("CertManager deployed at:", address(manager));

        string memory deploymentOutfile =
            string.concat(vm.projectRoot(), "/deployments/", vm.toString(block.chainid), "-certmanager.json");
        vm.writeJson({json: vm.serializeAddress("", "CertManager", address(manager)), path: deploymentOutfile});

        vm.stopBroadcast();
    }

    function _implSalt() internal view returns (bytes32) {
        return keccak256(bytes(Config.implSalt()));
    }
}
