package types

// ValidatorIndex in eth2.
type ValidatorIndex SSZUint64

// Div divides validator index by x.
func (v ValidatorIndex) Div(x uint64) ValidatorIndex {
	if x == 0 {
		panic("divbyzero")
	}
	return ValidatorIndex(uint64(v) / x)
}

// Add increases validator index by x.
func (v ValidatorIndex) Add(x uint64) ValidatorIndex {
	return ValidatorIndex(uint64(v) + x)
}

// Sub subtracts x from the validator index.
func (v ValidatorIndex) Sub(x uint64) ValidatorIndex {
	if uint64(v) < x {
		panic("underflow")
	}
	return ValidatorIndex(uint64(v) - x)
}

// Mod returns result of `validator index % x`.
func (v ValidatorIndex) Mod(x uint64) ValidatorIndex {
	return ValidatorIndex(uint64(v) % x)
}
