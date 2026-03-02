// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {Script} from "forge-std/Script.sol";
import {PortalWithChainExit} from "../src/PortalWithChainExit.sol";
import {console2 as console} from "forge-std/console2.sol";

contract DeployPortalWithChainExit is Script {
    function run() public {
        vm.startBroadcast();
        PortalWithChainExit portalWithChainExit = new PortalWithChainExit();
        vm.stopBroadcast();
        console.log("Deploying PortalWithChainOwnerExit at:", address(portalWithChainExit));
    }
}
