// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {Script} from "forge-std/Script.sol";
import {console2 as console} from "forge-std/console2.sol";
import {DeployChain} from "../src/DeployChain.sol";
import {Artifacts} from "@eth-optimism-bedrock/scripts/Artifacts.s.sol";
import {ProxyAdmin} from "@eth-optimism-bedrock/src/universal/ProxyAdmin.sol";
import {SystemConfigGlobal} from "../src/SystemConfigGlobal.sol";
import {ICertManager} from "@nitro-validator/ICertManager.sol";
import {IGnosisSafe, Enum} from "@eth-optimism-bedrock/scripts/interfaces/IGnosisSafe.sol";

contract UpgradeSystemConfigGlobal is Script, Artifacts {
    function run() public {
        _loadAddresses(deploymentOutfile);

        bytes memory signature = abi.encodePacked(uint256(uint160(msg.sender)), bytes32(0), uint8(1));

        console.log("Deploying SystemConfigGlobal implementation");
        vm.startBroadcast();

        address addr_ = address(new SystemConfigGlobal{salt: _implSalt()}(ICertManager(mustGetAddress("CertManager"))));
        bytes memory data = abi.encodeCall(ProxyAdmin.upgrade, (payable(mustGetAddress("SystemConfigGlobalProxy")), address(addr_)));
        IGnosisSafe(mustGetAddress("SystemOwnerSafe")).execTransaction({
            to: mustGetAddress("ProxyAdmin"),
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

        delete _namedDeployments["SystemConfigGlobal"];
        save("SystemConfigGlobal", addr_);
    }

    function _implSalt() internal view returns (bytes32 _env) {
        _env = keccak256(bytes(vm.envOr("IMPL_SALT", string("ethers phoenix"))));
    }
}
