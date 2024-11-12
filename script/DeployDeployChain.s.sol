// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {Script} from "forge-std/Script.sol";
import {console2 as console} from "forge-std/console2.sol";
import {DeployChain} from "../src/DeployChain.sol";
import {Artifacts} from "@eth-optimism-bedrock/scripts/Artifacts.s.sol";
import {Config} from "@eth-optimism-bedrock/scripts/libraries/Config.sol";

contract DeployDeployChain is Script, Artifacts {
    function run() public {
        _loadAddresses(deploymentOutfile);

        console.log("Deploying DeployChain implementation");
        vm.broadcast();
        DeployChain deployChain = new DeployChain{salt: _implSalt()}({
            _proxyAdmin: mustGetAddress("ProxyAdmin"),
            _optimismPortal: mustGetAddress("OptimismPortalProxy"),
            _systemConfig: mustGetAddress("SystemConfigProxy"),
            _l1StandardBridge: mustGetAddress("L1StandardBridgeProxy"),
            _l1ERC721Bridge: mustGetAddress("L1ERC721BridgeProxy"),
            _optimismMintableERC20Factory: mustGetAddress("OptimismMintableERC20FactoryProxy"),
            _l1CrossDomainMessenger: mustGetAddress("L1CrossDomainMessengerProxy"),
            _l2OutputOracle: mustGetAddress("L2OutputOracleProxy"),
            _superchainConfig: mustGetAddress("SuperchainConfigProxy"),
            _protocolVersions: mustGetAddress("ProtocolVersionsProxy")
        });

        delete _namedDeployments["DeployChain"];
        save("DeployChain", address(deployChain));
    }

    function _implSalt() internal view returns (bytes32 _env) {
        _env = keccak256(bytes(vm.envOr("IMPL_SALT", string("ethers phoenix"))));
    }
}
