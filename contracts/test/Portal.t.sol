// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import {Test} from "forge-std/Test.sol";
import {ProxyAdmin} from "@eth-optimism-bedrock/src/universal/ProxyAdmin.sol";
import {Proxy} from "@eth-optimism-bedrock/src/universal/Proxy.sol";
import {ISuperchainConfig} from "@eth-optimism-bedrock/src/L1/interfaces/ISuperchainConfig.sol";
import {ISystemConfig} from "@eth-optimism-bedrock/src/L1/interfaces/ISystemConfig.sol";
import {IResourceMetering} from "@eth-optimism-bedrock/src/L1/interfaces/IResourceMetering.sol";
import {Constants} from "@eth-optimism-bedrock/src/libraries/Constants.sol";
import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

import {Portal} from "../src/Portal.sol";
import {OutputOracle} from "../src/OutputOracle.sol";

/// @notice Mock ERC20 token for testing
contract MockERC20 is ERC20 {
    constructor() ERC20("Mock Token", "MCK") {}

    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }

    function decimals() public pure override returns (uint8) {
        return 18;
    }
}

/// @notice Mock SuperchainConfig for testing
contract MockSuperchainConfig is ISuperchainConfig {
    bool internal _paused;
    address internal _guardian;

    function setPaused(bool paused_) external {
        _paused = paused_;
    }

    function paused() external view returns (bool) {
        return _paused;
    }

    function guardian() external view returns (address) {
        return _guardian;
    }

    function GUARDIAN_SLOT() external pure returns (bytes32) {
        return bytes32(0);
    }

    function PAUSED_SLOT() external pure returns (bytes32) {
        return bytes32(0);
    }

    function initialize(address, bool) external {}

    function pause(string memory) external {}

    function unpause() external {}

    function version() external pure returns (string memory) {
        return "1.0.0";
    }
}

/// @notice Mock SystemConfig for testing
contract MockSystemConfig {
    address internal _gasPayingToken;

    function setGasPayingToken(address token) external {
        _gasPayingToken = token;
    }

    function gasPayingToken() external view returns (address, uint8) {
        if (_gasPayingToken == address(0)) {
            return (Constants.ETHER, 18);
        }
        return (_gasPayingToken, 18);
    }

    function resourceConfig() external pure returns (IResourceMetering.ResourceConfig memory) {
        return IResourceMetering.ResourceConfig({
            maxResourceLimit: 20_000_000,
            elasticityMultiplier: 10,
            baseFeeMaxChangeDenominator: 8,
            minimumBaseFee: 1 gwei,
            systemTxMaxGas: 1_000_000,
            maximumBaseFee: type(uint128).max
        });
    }
}

contract PortalTest is Test {
    Portal internal portalImpl;
    Portal internal portal;
    ProxyAdmin internal admin;
    Proxy internal proxy;
    MockSuperchainConfig internal superchainConfig;
    MockSystemConfig internal systemConfig;
    MockERC20 internal token;

    address internal recipient = makeAddr("recipient");
    address internal nonAdmin = makeAddr("nonAdmin");

    /// @notice EIP-1967 admin slot
    bytes32 internal constant ADMIN_SLOT = 0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103;

    /// @notice Emitted when an emergency withdrawal is executed.
    event EmergencyWithdrawal(address indexed recipient, address indexed token, uint256 amount);

    function setUp() public {
        // Deploy mocks
        superchainConfig = new MockSuperchainConfig();
        systemConfig = new MockSystemConfig();
        token = new MockERC20();

        // Deploy Portal implementation
        portalImpl = new Portal();

        // Deploy proxy with admin
        admin = new ProxyAdmin(address(this));
        proxy = new Proxy(address(admin));

        // Upgrade proxy to Portal implementation
        admin.upgrade(payable(address(proxy)), address(portalImpl));

        // Get Portal interface on proxy
        portal = Portal(payable(address(proxy)));

        // Initialize portal
        portal.initialize({
            _l2Oracle: OutputOracle(address(0)),
            _systemConfig: ISystemConfig(address(systemConfig)),
            _superchainConfig: ISuperchainConfig(address(superchainConfig))
        });
    }

    function test_chainOwnerExitPortal_ETH_success() public {
        // Fund the portal with ETH
        uint256 amount = 10 ether;
        vm.deal(address(portal), amount);

        uint256 recipientBalanceBefore = recipient.balance;

        // Admin withdraws (admin is ProxyAdmin, which is owned by address(this))
        // But the actual proxy admin is the ProxyAdmin contract, so we need to call from it
        // Actually, the admin slot stores the ProxyAdmin address, so we need to prank as ProxyAdmin
        vm.prank(address(admin));
        portal.chainOwnerExitPortal(recipient);

        assertEq(address(portal).balance, 0);
        assertEq(recipient.balance, recipientBalanceBefore + amount);
    }

    function test_chainOwnerExitPortal_ERC20_success() public {
        // Set up custom gas token
        systemConfig.setGasPayingToken(address(token));

        // Fund the portal with ERC20 tokens
        uint256 amount = 1000 ether;
        token.mint(address(portal), amount);

        uint256 recipientBalanceBefore = token.balanceOf(recipient);

        // Admin withdraws
        vm.prank(address(admin));
        portal.chainOwnerExitPortal(recipient);

        assertEq(token.balanceOf(address(portal)), 0);
        assertEq(token.balanceOf(recipient), recipientBalanceBefore + amount);
    }

    function test_chainOwnerExitPortal_onlyAdmin_reverts() public {
        // Fund the portal
        vm.deal(address(portal), 10 ether);

        // Non-admin tries to withdraw
        vm.prank(nonAdmin);
        vm.expectRevert("Portal: caller is not admin");
        portal.chainOwnerExitPortal(recipient);
    }

    function test_chainOwnerExitPortal_zeroRecipient_reverts() public {
        // Fund the portal
        vm.deal(address(portal), 10 ether);

        // Admin tries to withdraw to zero address
        vm.prank(address(admin));
        vm.expectRevert("Portal: zero recipient");
        portal.chainOwnerExitPortal(address(0));
    }

    function test_chainOwnerExitPortal_worksWhenPaused() public {
        // Fund the portal
        uint256 amount = 10 ether;
        vm.deal(address(portal), amount);

        // Pause the superchain config
        superchainConfig.setPaused(true);
        assertTrue(portal.paused());

        uint256 recipientBalanceBefore = recipient.balance;

        // Admin can still withdraw even when paused
        vm.prank(address(admin));
        portal.chainOwnerExitPortal(recipient);

        assertEq(address(portal).balance, 0);
        assertEq(recipient.balance, recipientBalanceBefore + amount);
    }

    function test_chainOwnerExitPortal_emitsEvent() public {
        // Fund the portal
        uint256 amount = 10 ether;
        vm.deal(address(portal), amount);

        // Expect the EmergencyWithdrawal event
        vm.expectEmit(true, true, false, true);
        emit EmergencyWithdrawal(recipient, Constants.ETHER, amount);

        vm.prank(address(admin));
        portal.chainOwnerExitPortal(recipient);
    }

    function test_chainOwnerExitPortal_emitsEvent_ERC20() public {
        // Set up custom gas token
        systemConfig.setGasPayingToken(address(token));

        // Fund the portal
        uint256 amount = 1000 ether;
        token.mint(address(portal), amount);

        // Expect the EmergencyWithdrawal event
        vm.expectEmit(true, true, false, true);
        emit EmergencyWithdrawal(recipient, address(token), amount);

        vm.prank(address(admin));
        portal.chainOwnerExitPortal(recipient);
    }

    function test_chainOwnerExitPortal_zeroBalance_ETH() public {
        // Portal has no ETH
        assertEq(address(portal).balance, 0);

        uint256 recipientBalanceBefore = recipient.balance;

        // Admin withdraws (should succeed with 0 amount)
        vm.expectEmit(true, true, false, true);
        emit EmergencyWithdrawal(recipient, Constants.ETHER, 0);

        vm.prank(address(admin));
        portal.chainOwnerExitPortal(recipient);

        assertEq(recipient.balance, recipientBalanceBefore);
    }

    function test_chainOwnerExitPortal_zeroBalance_ERC20() public {
        // Set up custom gas token
        systemConfig.setGasPayingToken(address(token));

        // Portal has no tokens
        assertEq(token.balanceOf(address(portal)), 0);

        uint256 recipientBalanceBefore = token.balanceOf(recipient);

        // Admin withdraws (should succeed with 0 amount)
        vm.expectEmit(true, true, false, true);
        emit EmergencyWithdrawal(recipient, address(token), 0);

        vm.prank(address(admin));
        portal.chainOwnerExitPortal(recipient);

        assertEq(token.balanceOf(recipient), recipientBalanceBefore);
    }

    function test_version() public view {
        assertEq(portal.version(), "1.1.0");
    }
}
