package types

import (
	fmt "fmt"

	fssz "github.com/ferranbt/fastssz"
)

var _ fssz.HashRoot = (Slot)(0)
var _ fssz.Marshaler = (*Slot)(nil)
var _ fssz.Unmarshaler = (*Slot)(nil)

// Slot represents a single slot.
type Slot uint64

// Mul multiplies slot by x.
func (s Slot) Mul(x uint64) Slot {
	return Slot(uint64(s) * x)
}

// MulSlot multiplies slot by another slot.
func (s Slot) MulSlot(x Slot) Slot {
	return s * x
}

// MulEpoch multiplies slot using epoch value.
func (s Slot) MulEpoch(x Epoch) Slot {
	return Slot(uint64(s) * uint64(x))
}

// Div divides slot by x.
func (s Slot) Div(x uint64) Slot {
	if x == 0 {
		panic("divbyzero")
	}
	return Slot(uint64(s) / x)
}

// DivSlot divides slot by another slot.
func (s Slot) DivSlot(x Slot) Slot {
	if x == 0 {
		panic("divbyzero")
	}
	return s / x
}

// DivEpoch divides slot using epoch value.
func (s Slot) DivEpoch(x Epoch) Slot {
	if x == 0 {
		panic("divbyzero")
	}
	return Slot(uint64(s) / uint64(x))
}

// Add increases slot by x.
func (s Slot) Add(x uint64) Slot {
	return Slot(uint64(s) + x)
}

// AddSlot increases slot by another slot.
func (s Slot) AddSlot(x Slot) Slot {
	return s + x
}

// AddEpoch increases slot using epoch value.
func (s Slot) AddEpoch(x Epoch) Slot {
	return Slot(uint64(s) + uint64(x))
}

// Sub subtracts x from the slot.
func (s Slot) Sub(x uint64) Slot {
	if uint64(s) < x {
		panic("underflow")
	}
	return Slot(uint64(s) - x)
}

// SubSlot finds difference between two slot values.
func (s Slot) SubSlot(x Slot) Slot {
	if s < x {
		panic("underflow")
	}
	return s - x
}

// SubEpoch subtracts value of epoch type from the slot.
func (s Slot) SubEpoch(x Epoch) Slot {
	if uint64(s) < uint64(x) {
		panic("underflow")
	}
	return Slot(uint64(s) - uint64(x))
}

// Mod returns result of `slot % x`.
func (s Slot) Mod(x uint64) Slot {
	return Slot(uint64(s) % x)
}

// ModSlot returns result of `slot % slot`.
func (s Slot) ModSlot(x Slot) Slot {
	return s % x
}

// ModEpoch returns result of `slot % epoch`.
func (s Slot) ModEpoch(x Epoch) Slot {
	return Slot(uint64(s) % uint64(x))
}

// HashTreeRoot returns calculated hash root.
func (s Slot) HashTreeRoot() ([32]byte, error) {
	return fssz.HashWithDefaultHasher(s)
}

// HashWithDefaultHasher hashes a HashRoot object with a Hasher from the default HasherPool.
func (s Slot) HashTreeRootWith(hh *fssz.Hasher) error {
	hh.PutUint64(uint64(s))
	return nil
}

// UnmarshalSSZ deserializes the provided bytes buffer into the slot object.
func (s *Slot) UnmarshalSSZ(buf []byte) error {
	if len(buf) != s.SizeSSZ() {
		return fmt.Errorf("expected buffer of length %d received %d", s.SizeSSZ(), len(buf))
	}
	*s = Slot(fssz.UnmarshallUint64(buf))
	return nil
}

// MarshalSSZTo marshals slot with the provided byte slice.
func (s *Slot) MarshalSSZTo(dst []byte) ([]byte, error) {
	marshalled, err := s.MarshalSSZ()
	if err != nil {
		return nil, err
	}
	return append(dst, marshalled...), nil
}

// MarshalSSZ marshals slot into a serialized object.
func (s *Slot) MarshalSSZ() ([]byte, error) {
	marshalled := fssz.MarshalUint64([]byte{}, uint64(*s))
	return marshalled, nil
}

// SizeSSZ returns the size of the serialized object.
func (s *Slot) SizeSSZ() int {
	return 8
}
