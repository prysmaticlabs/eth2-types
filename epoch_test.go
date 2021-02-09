package types_test

import (
	"fmt"
	"testing"

	types "github.com/prysmaticlabs/eth2-types"
)

func TestEpoch_Mul(t *testing.T) {
	tests := []struct {
		a, b     uint64
		res      types.Epoch
		panicMsg string
	}{
		{a: 0, b: 1, res: 0},
		{a: 1 << 32, b: 1, res: 1 << 32},
		{a: 1 << 32, b: 100, res: 429496729600},
		{a: 1 << 32, b: 1 << 31, res: 9223372036854775808},
		{a: 1 << 32, b: 1 << 32, res: 0, panicMsg: types.ErrMulOverflow.Error()},
		{a: 1 << 62, b: 2, res: 9223372036854775808},
		{a: 1 << 62, b: 4, res: 0, panicMsg: types.ErrMulOverflow.Error()},
		{a: 1 << 63, b: 1, res: 9223372036854775808},
		{a: 1 << 63, b: 2, res: 0, panicMsg: types.ErrMulOverflow.Error()},
	}

	assertPanic := func(panicMessage string, f func()) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic not thrown")
			} else if r != panicMessage {
				t.Errorf("Unexpected panic thrown, want: %#v, got: %#v", panicMessage, r)
			}
		}()
		f()
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v x %v = %v", tt.a, tt.b, tt.res), func(t *testing.T) {
			var res types.Epoch
			if tt.panicMsg != "" {
				assertPanic(tt.panicMsg, func() {
					res = types.Epoch(tt.a).Mul(tt.b)
				})
			} else {
				res = types.Epoch(tt.a).Mul(tt.b)
			}
			if tt.res != res {
				t.Errorf("Epoch.Mul() = %v, want %v", res, tt.res)
			}
		})
	}
}
