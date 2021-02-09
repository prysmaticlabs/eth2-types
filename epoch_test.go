package types_test

import (
	"fmt"
	"math"
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

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Epoch(%v).Mul(%v) = %v", tt.a, tt.b, tt.res), func(t *testing.T) {
			var res types.Epoch
			if tt.panicMsg != "" {
				assertPanic(t, tt.panicMsg, func() {
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

func TestEpoch_Div(t *testing.T) {
	tests := []struct {
		a, b     uint64
		res      types.Epoch
		panicMsg string
	}{
		{a: 0, b: 1, res: 0},
		{a: 1, b: 0, res: 0, panicMsg: types.ErrDivByZero.Error()},
		{a: 1 << 32, b: 1 << 32, res: 1},
		{a: 429496729600, b: 1 << 32, res: 100},
		{a: 9223372036854775808, b: 1 << 32, res: 1 << 31},
		{a: 1 << 32, b: 1 << 32, res: 1},
		{a: 9223372036854775808, b: 1 << 62, res: 2},
		{a: 9223372036854775808, b: 1 << 63, res: 1},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Epoch(%v).Div(%v) = %v", tt.a, tt.b, tt.res), func(t *testing.T) {
			var res types.Epoch
			if tt.panicMsg != "" {
				assertPanic(t, tt.panicMsg, func() {
					res = types.Epoch(tt.a).Div(tt.b)
				})
			} else {
				res = types.Epoch(tt.a).Div(tt.b)
			}
			if tt.res != res {
				t.Errorf("Epoch.Div() = %v, want %v", res, tt.res)
			}
		})
	}
}

func TestEpoch_Add(t *testing.T) {
	tests := []struct {
		a, b     uint64
		res      types.Epoch
		panicMsg string
	}{
		{a: 0, b: 1, res: 1},
		{a: 1 << 32, b: 1, res: 4294967297},
		{a: 1 << 32, b: 100, res: 4294967396},
		{a: 1 << 31, b: 1 << 31, res: 4294967296},
		{a: 1 << 63, b: 1 << 63, res: 0, panicMsg: types.ErrAddOverflow.Error()},
		{a: 1 << 63, b: 1, res: 9223372036854775809},
		{a: math.MaxUint64, b: 1, res: 0, panicMsg: types.ErrAddOverflow.Error()},
		{a: math.MaxUint64, b: 0, res: math.MaxUint64},
		{a: 1 << 63, b: 2, res: 9223372036854775810},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Epoch(%v).Add(%v) = %v", tt.a, tt.b, tt.res), func(t *testing.T) {
			var res types.Epoch
			if tt.panicMsg != "" {
				assertPanic(t, tt.panicMsg, func() {
					res = types.Epoch(tt.a).Add(tt.b)
				})
			} else {
				res = types.Epoch(tt.a).Add(tt.b)
			}
			if tt.res != res {
				t.Errorf("Epoch.Add() = %v, want %v", res, tt.res)
			}
		})
		t.Run(fmt.Sprintf("Epoch(%v).AddEpoch(%v) = %v", tt.a, tt.b, tt.res), func(t *testing.T) {
			var res types.Epoch
			if tt.panicMsg != "" {
				assertPanic(t, tt.panicMsg, func() {
					res = types.Epoch(tt.a).AddEpoch(types.Epoch(tt.b))
				})
			} else {
				res = types.Epoch(tt.a).AddEpoch(types.Epoch(tt.b))
			}
			if tt.res != res {
				t.Errorf("Epoch.AddEpoch() = %v, want %v", res, tt.res)
			}
		})
	}
}

func assertPanic(t *testing.T, panicMessage string, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic not thrown")
		} else if r != panicMessage {
			t.Errorf("Unexpected panic thrown, want: %#v, got: %#v", panicMessage, r)
		}
	}()
	f()
}
