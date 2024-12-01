// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {ISemver} from "@eth-optimism-bedrock/src/universal/interfaces/ISemver.sol";
import {NitroValidator} from "@nitro-validator/NitroValidator.sol";
import {CertManager} from "@nitro-validator/CertManager.sol";
import {NodePtr, LibNodePtr} from "@nitro-validator/NodePtr.sol";
import {LibBytes} from "@nitro-validator/LibBytes.sol";

contract SystemConfigGlobal is OwnableUpgradeable, ISemver, NitroValidator {
    using LibNodePtr for NodePtr;
    using LibBytes for bytes;

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

    constructor(CertManager certManager) NitroValidator(certManager) {
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
        bytes memory pcr0 = attestationTbs.slice(ptrs.pcrs[0].content(), ptrs.pcrs[0].length());
        require(validPCR0s[keccak256(pcr0)], "invalid pcr0 in attestation");

        require(ptrs.timestamp + MAX_AGE > block.timestamp, "attestation too old");

        bytes memory publicKey = attestationTbs.slice(ptrs.publicKey.content(), ptrs.publicKey.length());
        address enclaveAddress = address(uint160(uint256(keccak256(publicKey))));
        validSigners[enclaveAddress] = true;
    }

    function deregisterSigner(address signer) external onlyOwner {
        delete validSigners[signer];
    }
}
