// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import {OwnableManagedUpgradeable} from "./OwnableManagedUpgradeable.sol";
import {ISemver} from "@eth-optimism-bedrock/src/universal/interfaces/ISemver.sol";
import {NitroValidator} from "@nitro-validator/NitroValidator.sol";
import {LibBytes} from "@nitro-validator/LibBytes.sol";
import {LibCborElement, CborElement, CborDecode} from "@nitro-validator/CborDecode.sol";
import {ICertManager} from "@nitro-validator/ICertManager.sol";

contract SystemConfigGlobal is OwnableManagedUpgradeable, ISemver, NitroValidator {
    using LibBytes for bytes;
    using CborDecode for bytes;
    using LibCborElement for CborElement;

    uint256 public constant MAX_AGE = 60 minutes;

    /// @notice The address of the proposer.
    address public proposer;

    /// @notice Mapping of valid PCR0s attested from AWS Nitro.
    mapping(bytes32 => bool) public validPCR0s;

    /// @notice Mapping of valid signers attested from AWS Nitro.
    mapping(address => bool) public validSigners;

    /// @notice Semantic version.
    /// @custom:semver 0.0.1
    function version() public pure virtual returns (string memory) {
        return "0.0.1";
    }

    constructor(ICertManager certManager) NitroValidator(certManager) {
        initialize({_owner: address(0xdEaD), _manager: address(0xdEaD)});
    }

    function initialize(address _owner, address _manager) public initializer {
        __OwnableManaged_init();
        transferOwnership(_owner);
        transferManagement(_manager);
    }

    function setProposer(address _proposer) external onlyOwner {
        proposer = _proposer;
    }

    function registerPCR0(bytes calldata pcr0) external onlyOwner {
        validPCR0s[keccak256(pcr0)] = true;
    }

    function deregisterPCR0(bytes calldata pcr0) external onlyOwner {
        delete validPCR0s[keccak256(pcr0)];
    }

    function registerSigner(bytes calldata attestationTbs, bytes calldata signature) external onlyOwnerOrManager {
        Ptrs memory ptrs = validateAttestation(attestationTbs, signature);
        bytes32 pcr0 = attestationTbs.keccak(ptrs.pcrs[0]);
        require(validPCR0s[pcr0], "invalid pcr0 in attestation");

        require(ptrs.timestamp + MAX_AGE > block.timestamp, "attestation too old");

        // The publicKey is encoded in the form specified in section 4.3.6 of ANSI X9.62, which is a
        // 0x04 byte followed by the x and y coordinates of the public key. We ignore the first byte.
        bytes32 publicKeyHash = attestationTbs.keccak(ptrs.publicKey.start() + 1, ptrs.publicKey.length() - 1);
        address enclaveAddress = address(uint160(uint256(publicKeyHash)));
        validSigners[enclaveAddress] = true;
    }

    function deregisterSigner(address signer) external onlyOwnerOrManager {
        delete validSigners[signer];
    }
}
