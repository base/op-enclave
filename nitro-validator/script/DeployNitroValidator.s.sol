// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {Script} from "forge-std/Script.sol";
import {console2 as console} from "forge-std/console2.sol";
import {NitroValidator} from "src/NitroValidator.sol";

/// @notice will deploy the singleton NitroValidatorContract to a deterministic address
contract DeployNitroValidator is Script {
    bytes32 constant SALT = bytes32(uint256(0x4E313752304F5)); // todo

    function run() public returns (address addr_) {
        vm.startBroadcast();

        addr_ = address(new NitroValidator{salt: SALT}());

        console.log("NitroValidator to be deployed at:", addr_);

        // Save the address to the deployment file
        string memory deploymentJson = string.concat("{", '"address": "', vm.toString(addr_), '"}');

        vm.writeFile(
            string.concat(vm.projectRoot(), "/deployments/", vm.toString(block.chainid), "-nitro-validator-deploy.json"),
            deploymentJson
        );

        vm.stopBroadcast();
    }
}
