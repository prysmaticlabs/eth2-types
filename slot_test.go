package types

import (
	"testing"
	"time"
)

func TestSlot_Casting(t *testing.T) {
	slot := Slot(42)

	t.Run("time.Duration", func(t *testing.T) {
		if uint64(time.Duration(slot)) != uint64(slot) {
			t.Error("Slot should produce the same result with time.Duration")
		}
	})

	t.Run("floats", func(t *testing.T) {
		var x1 float32 = 42.2
		if Slot(x1) != slot {
			t.Errorf("Unequal: %v = %v", Slot(x1), slot)
		}

		var x2 float64 = 42.2
		if Slot(x2) != slot {
			t.Errorf("Unequal: %v = %v", Slot(x2), slot)
		}
	})

	t.Run("int", func(t *testing.T) {
		var x int = 42
		if Slot(x) != slot {
			t.Errorf("Unequal: %v = %v", Slot(x), slot)
		}
	})
}
