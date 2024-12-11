// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {ISemver} from "@eth-optimism-bedrock/src/universal/interfaces/ISemver.sol";
import {NitroValidator} from "@nitro-validator/NitroValidator.sol";
import {LibBytes} from "@nitro-validator/LibBytes.sol";
import {LibCborElement, CborElement, CborDecode} from "@nitro-validator/CborDecode.sol";
import {ICertManager} from "@nitro-validator/ICertManager.sol";

contract SystemConfigGlobal is OwnableUpgradeable, ISemver, NitroValidator {
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
        initialize({_owner: address(0xdEaD)});
    }

    function initialize(address _owner) public initializer {
        __Ownable_init();
        transferOwnership(_owner);
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

    function registerSigner(bytes calldata attestationTbs, bytes calldata signature) external onlyOwner {
        Ptrs memory ptrs = validateAttestation(attestationTbs, signature);
        bytes32 pcr0 = attestationTbs.keccak(ptrs.pcrs[0]);
        require(validPCR0s[pcr0], "invalid pcr0 in attestation");

        require(ptrs.timestamp + MAX_AGE > block.timestamp, "attestation too old");

        bytes32 publicKeyHash = attestationTbs.keccak(ptrs.publicKey.start() + 1, ptrs.publicKey.length() - 1);
        address enclaveAddress = address(uint160(uint256(publicKeyHash)));
        validSigners[enclaveAddress] = true;
    }

    function deregisterSigner(address signer) external onlyOwner {
        delete validSigners[signer];
    }
}
