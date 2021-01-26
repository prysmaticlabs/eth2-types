package types

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
