package core

import (
	"time"

	types "github.com/prysmaticlabs/eth2-types"

	prysmTime "github.com/prysmaticlabs/prysm/time"
	"github.com/sirupsen/logrus"
)

// Network represents the network.
type Network string

// NetworkFromString returns network from the given string value
func NetworkFromString(n string) Network {
	switch n {
	case string(PyrmontNetwork):
		return PyrmontNetwork
	case string(PraterNetwork):
		return PraterNetwork
	case string(MainNetwork):
		return MainNetwork
	case string(DevNetwork):
		return DevNetwork
	default:
		return ""
	}
}

// ForkVersion returns the fork version of the network.
func (n Network) ForkVersion() []byte {
	switch n {
	case PyrmontNetwork:
		return []byte{0, 0, 32, 9}
	case PraterNetwork:
		return []byte{0x00, 0x00, 0x10, 0x20}
	case MainNetwork:
		return []byte{0, 0, 0, 0}
	case DevNetwork:
		return []byte{0x10, 0x00, 0x00, 0x38}
	default:
		logrus.WithField("network", n).Fatal("undefined network")
		return nil
	}
}

// DepositContractAddress returns the deposit contract address of the network.
func (n Network) DepositContractAddress() string {
	switch n {
	case PyrmontNetwork:
		return "0x8c5fecdC472E27Bc447696F431E425D02dd46a8c"
	case PraterNetwork:
		return "0xff50ed3d0ec03ac01d4c79aad74928bff48a7b2b"
	case MainNetwork:
		return "0x00000000219ab540356cBB839Cbe05303d7705Fa"
	case DevNetwork:
		return "0x4242424242424242424242424242424242424242"
	default:
		logrus.WithField("network", n).Fatal("undefined network")
		return ""
	}
}

// FullPath returns the full path of the network.
func (n Network) FullPath(relativePath string) string {
	return BaseEIP2334Path + relativePath
}

// MinGenesisTime returns min genesis time value
func (n Network) MinGenesisTime() uint64 {
	switch n {
	case PyrmontNetwork:
		return 1605700807
	case PraterNetwork:
		return 1616508000
	case MainNetwork:
		return 1606824023
	case DevNetwork:
		return 1668789575
	default:
		logrus.WithField("network", n).Fatal("undefined network")
		return 0
	}
}

// SlotDurationSec returns slot duration
func (n Network) SlotDurationSec() time.Duration {
	return 12 * time.Second
}

// SlotsPerEpoch returns number of slots per one epoch
func (n Network) SlotsPerEpoch() uint64 {
	return 32
}

// EstimatedCurrentSlot returns the estimation of the current slot
func (n Network) EstimatedCurrentSlot() types.Slot {
	return n.EstimatedSlotAtTime(prysmTime.Now().Unix())
}

// EstimatedSlotAtTime estimates slot at the given time
func (n Network) EstimatedSlotAtTime(time int64) types.Slot {
	genesis := int64(n.MinGenesisTime())
	if time < genesis {
		return 0
	}
	return types.Slot(uint64(time-genesis) / uint64(n.SlotDurationSec().Seconds()))
}

// EstimatedCurrentEpoch estimates the current epoch
// https://github.com/ethereum/eth2.0-specs/blob/dev/specs/phase0/beacon-chain.md#compute_start_slot_at_epoch
func (n Network) EstimatedCurrentEpoch() types.Epoch {
	return n.EstimatedEpochAtSlot(n.EstimatedCurrentSlot())
}

// EstimatedEpochAtSlot estimates epoch at the given slot
func (n Network) EstimatedEpochAtSlot(slot types.Slot) types.Epoch {
	return types.Epoch(slot / types.Slot(n.SlotsPerEpoch()))
}

// Available networks.
const (
	// PyrmontNetwork represents the Pyrmont test network.
	PyrmontNetwork Network = "pyrmont"

	// PraterNetwork represents the Prater test network.
	PraterNetwork Network = "prater"

	// MainNetwork represents the main network.
	MainNetwork Network = "mainnet"

	// DevNetwork represents the dev network.
	DevNetwork Network = "devnet"
)
