// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import {ResolvingProxy} from "./ResolvingProxy.sol";

library ResolvingProxyFactory {
    function setupProxy(address proxy, address admin, bytes32 salt) internal returns (address instance) {
        /// @solidity memory-safe-assembly
        assembly {
            let ptr := mload(0x40)
            mstore(ptr, 0x600661010d565b73000000000000000000000000000000000000000000000000)
            mstore(add(ptr, 0x8), shl(0x60, proxy))
            mstore(add(ptr, 0x1c), 0x9055730000000000000000000000000000000000000000000000000000000000)
            mstore(add(ptr, 0x1f), shl(0x60, admin))
            mstore(add(ptr, 0x33), 0x905561011380603f5f395ff3365f600760ce565b805490918054803314331517)
            mstore(add(ptr, 0x53), 0x1560535760045f5f375f5160e01c8063f851a4401460975780635c60da1b1460)
            mstore(add(ptr, 0x73), 0x945780638f2839701460a45780633659cfe61460a157634f1ef28614609f575b)
            mstore(add(ptr, 0x93), 0x63204e1c7a60e01b5f52826004525f5f60245f845afa3d5f5f3e3d6020141680)
            mstore(add(ptr, 0xb3), 0x5f510290158402015f875f89895f375f935af43d5f5f3e5f3d91609257fd5bf3)
            mstore(add(ptr, 0xd3), 0x5b50505b505f5260205ff35b5f5b93915b5050602060045f375f518091559160)
            mstore(add(ptr, 0xf3), 0xca57903333602060445f375f519560649550506053565b5f5ff35b7f360894a1)
            mstore(add(ptr, 0x113), 0x3ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc7fb53127)
            mstore(add(ptr, 0x133), 0x684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103915600)
            instance := create2(0, ptr, 0x152, salt)
        }
        require(instance != address(0), "Proxy: create2 failed");
    }

    function proxyAddress(address proxy, address admin, bytes32 salt) internal view returns (address predicted) {
        address deployer = address(this);
        /// @solidity memory-safe-assembly
        assembly {
            let ptr := mload(0x40)
            mstore(ptr, 0x600661010d565b73000000000000000000000000000000000000000000000000)
            mstore(add(ptr, 0x8), shl(0x60, proxy))
            mstore(add(ptr, 0x1c), 0x9055730000000000000000000000000000000000000000000000000000000000)
            mstore(add(ptr, 0x1f), shl(0x60, admin))
            mstore(add(ptr, 0x33), 0x905561011380603f5f395ff3365f600760ce565b805490918054803314331517)
            mstore(add(ptr, 0x53), 0x1560535760045f5f375f5160e01c8063f851a4401460975780635c60da1b1460)
            mstore(add(ptr, 0x73), 0x945780638f2839701460a45780633659cfe61460a157634f1ef28614609f575b)
            mstore(add(ptr, 0x93), 0x63204e1c7a60e01b5f52826004525f5f60245f845afa3d5f5f3e3d6020141680)
            mstore(add(ptr, 0xb3), 0x5f510290158402015f875f89895f375f935af43d5f5f3e5f3d91609257fd5bf3)
            mstore(add(ptr, 0xd3), 0x5b50505b505f5260205ff35b5f5b93915b5050602060045f375f518091559160)
            mstore(add(ptr, 0xf3), 0xca57903333602060445f375f519560649550506053565b5f5ff35b7f360894a1)
            mstore(add(ptr, 0x113), 0x3ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc7fb53127)
            mstore(add(ptr, 0x133), 0x684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103915600)
            mstore(add(ptr, 0x152), shl(0x60, deployer))
            mstore(add(ptr, 0x166), salt)
            mstore(add(ptr, 0x186), keccak256(ptr, 0x152))
            predicted := keccak256(add(ptr, 0x152), 0x55)
        }
    }

    function setupExpensiveProxy(address proxy, address admin, bytes32 salt) internal returns (address instance) {
        return address(new ResolvingProxy{salt: salt}(proxy, admin));
    }

    function expensiveProxyAddress(address proxy, address admin, bytes32 salt) internal view returns (address predicted) {
        bytes memory bytecode = abi.encodePacked(type(ResolvingProxy).creationCode, abi.encode(proxy, admin));
        bytes32 hash = keccak256(abi.encodePacked(bytes1(0xff), address(this), salt, keccak256(bytecode)));
        return address(uint160(uint256(hash)));
    }
}
