package types

import (
	"errors"
	"math/bits"
)

var (
	ErrMulOverflow = errors.New("multiplication overflows")
	ErrAddOverflow = errors.New("addition overflows")
)

// MaxSlot returns the larger of the two slots.
func MaxSlot(a, b Slot) Slot {
	if a > b {
		return a
	}
	return b
}

// MinSlot returns the smaller of the two slots.
func MinSlot(a, b Slot) Slot {
	if a < b {
		return a
	}
	return b
}

// MaxEpoch returns the larger of the two epochs.
func MaxEpoch(a, b Epoch) Epoch {
	if a > b {
		return a
	}
	return b
}

// MinEpoch returns the smaller of the two epochs.
func MinEpoch(a, b Epoch) Epoch {
	if a < b {
		return a
	}
	return b
}

// Mul64 multiples 2 64-bit unsigned integers and checks if they
// lead to an overflow. If they do not, it returns the result
// without an error.
func Mul64(a, b uint64) (uint64, error) {
	overflows, val := bits.Mul64(a, b)
	if overflows > 0 {
		return 0, ErrMulOverflow
	}
	return val, nil
}

// Add64 adds 2 64-bit unsigned integers and checks if they
// lead to an overflow. If they do not, it returns the result
// without an error.
func Add64(a, b uint64) (uint64, error) {
	res, carry := bits.Add64(a, b, 0 /* carry */)
	if carry > 0 {
		return 0, ErrAddOverflow
	}
	return res, nil
}
