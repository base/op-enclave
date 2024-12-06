// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import {Vm} from "forge-std/Vm.sol";
import {Test, console} from "forge-std/Test.sol";
import {ResolvingProxyFactory} from "../src/ResolvingProxyFactory.sol";
import {ProxyAdmin} from "@eth-optimism-bedrock/src/universal/ProxyAdmin.sol";
import {Proxy} from "@eth-optimism-bedrock/src/universal/Proxy.sol";

contract Implementation {
    function foo() public pure returns (string memory) {
        return "bar";
    }
}

contract ResolvingProxyFactoryTest is Test {
    Implementation public implementation;
    ProxyAdmin public admin;
    Proxy public proxy;
    address public resolvingProxy;

    function setUp() public {
        implementation = new Implementation();
        admin = new ProxyAdmin(address(this));
        proxy = new Proxy(address(admin));
        admin.upgrade(payable(address(proxy)), address(implementation));
        resolvingProxy = ResolvingProxyFactory.setupProxy(address(proxy), address(admin), 0x00);
    }

    function test_setupProxy() public view {
        string memory foo = Implementation(resolvingProxy).foo();
        assertEq(foo, "bar");
    }
}
