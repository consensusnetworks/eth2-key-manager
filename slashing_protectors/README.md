# Blox KeyVault - Slashing Protector


[![blox.io](https://s3.us-east-2.amazonaws.com/app-files.blox.io/static/media/powered_by.png)](https://blox.io)

Slashing protection implementation for eth 2.0 

### Slashing Rules
Slashing can occur for validators proposing a block or signing attestations which can conflict with a previous signature.

#### Attestation - Double Vote
Description: Do not sign 2 attestations for the same block height. [eth 2 spec](https://github.com/ethereum/eth2.0-specs/blob/dev/specs/phase0/validator.md#attester-slashing).

 ```python
def is_slashable_attestation_data(data_1: AttestationData, data_2: AttestationData) -> bool:
    """
    Check if ``data_1`` and ``data_2`` are slashable according to Casper FFG rules.
    """
    return (
        # Double vote
        (data_1 != data_2 and data_1.target.epoch == data_2.target.epoch) or
    )
 ```

#### Attestation - Surrounding/Surrounded Vote
Description: Do not surround an already existing attestation/s, a.k.a do not forget them. [eth 2 spec](https://github.com/ethereum/eth2.0-specs/blob/dev/specs/phase0/validator.md#attester-slashing).

 ```python
def is_slashable_attestation_data(data_1: AttestationData, data_2: AttestationData) -> bool:
    """
    Check if ``data_1`` and ``data_2`` are slashable according to Casper FFG rules.
    """
    return (
        # Surround vote
        (data_1.source.epoch < data_2.source.epoch and data_2.target.epoch < data_1.target.epoch)
    )
 ```

#### Proposal - Duplicate
Description: Do not propose 2 blocks for the same block height. [eth 2 spec](https://github.com/ethereum/eth2.0-specs/blob/dev/specs/phase0/validator.md#proposer-slashing).