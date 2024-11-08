// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {Script} from "forge-std/Script.sol";
import {console2 as console} from "forge-std/console2.sol";
import {Config} from "@eth-optimism-bedrock/scripts/libraries/Config.sol";
import {ProxyAdmin} from "@eth-optimism-bedrock/src/universal/ProxyAdmin.sol";
import {Proxy} from "@eth-optimism-bedrock/src/universal/Proxy.sol";
import {SystemConfigOwnable} from "../src/SystemConfigOwnable.sol";
import {OwnerConfig} from "../src/OwnerConfig.sol";
import {IGnosisSafe, Enum} from "@eth-optimism-bedrock/scripts/interfaces/IGnosisSafe.sol";

contract UpgradeSystemConfigOwnable is Script {
    function run() public {
        address safe = 0xFfe2Ae7C40F70fEa4773F58d39A67526238c2750;
        address admin = 0x07D583226bE1636d1b00B479F994a264E183298E;
        address proxy = 0x8E248445a3943f5c680707dEf0AFAd8D24A29248;
        address ownerConfig = 0xD4110A249FAB4f20eaa3Aba2eaa5fE196a33C9e6;
        bytes memory signature = abi.encodePacked(uint256(uint160(msg.sender)), bytes32(0), uint8(1));

        vm.startBroadcast();

        SystemConfigOwnable config = new SystemConfigOwnable(OwnerConfig(ownerConfig));
        bytes memory data = abi.encodeCall(ProxyAdmin.upgrade, (payable(proxy), address(config)));
        IGnosisSafe(safe).execTransaction({
            to: admin,
            value: 0,
            data: data,
            operation: Enum.Operation.Call,
            safeTxGas: 0,
            baseGas: 0,
            gasPrice: 0,
            gasToken: address(0),
            refundReceiver: payable(address(0)),
            signatures: signature
        });

        vm.stopBroadcast();
    }
}
