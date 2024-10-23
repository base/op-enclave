// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {Script} from "forge-std/Script.sol";
import {console2 as console} from "forge-std/console2.sol";
import {CertManager} from "@marlinprotocol/CertManager.sol";
import {NitroValidator} from "src/NitroValidator.sol";

/// @notice will deploy the singleton NitroValidatorContract to a deterministic address
contract DeployNitroValidator is Script {
    bytes32 constant SALT = bytes32(uint256(0x4E313752304F5)); // todo

    function run() public {
        vm.startBroadcast();

        CertManager manager = new CertManager{salt: SALT}();
        NitroValidator validator = new NitroValidator{salt: SALT}(manager);

        console.log("CertManager deployed at:", address(manager));
        console.log("NitroValidator deployed at:", address(validator));

        // Save the address to the deployment file
        string memory deploymentJson = string.concat(
            "{", '"certManager": "', vm.toString(address(manager)),
            '", "nitroValidator": "', vm.toString(address(validator)), '"}');

        vm.writeFile(
            string.concat(vm.projectRoot(), "/deployments/", vm.toString(block.chainid), "-deploy.json"),
            deploymentJson
        );

        vm.stopBroadcast();
    }
}
