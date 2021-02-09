package types_test

import (
	"math"
	"testing"

	"github.com/prysmaticlabs/eth2-types"
)

func TestMul64(t *testing.T) {
	type args struct {
		a uint64
		b uint64
	}
	tests := []struct {
		args args
		res  uint64
		err  bool
	}{
		{args: args{0, 1}, res: 0, err: false},
		{args: args{1 << 32, 1}, res: 1 << 32, err: false},
		{args: args{1 << 32, 100}, res: 429496729600, err: false},
		{args: args{1 << 32, 1 << 31}, res: 9223372036854775808, err: false},
		{args: args{1 << 32, 1 << 32}, res: 0, err: true},
		{args: args{1 << 62, 2}, res: 9223372036854775808, err: false},
		{args: args{1 << 62, 4}, res: 0, err: true},
		{args: args{1 << 63, 1}, res: 9223372036854775808, err: false},
		{args: args{1 << 63, 2}, res: 0, err: true},
	}
	for _, tt := range tests {
		got, err := types.Mul64(tt.args.a, tt.args.b)
		if tt.err && err == nil {
			t.Errorf("Mul64() Expected Error = %v, want error", tt.err)
			continue
		}
		if tt.res != got {
			t.Errorf("Mul64() %v, want %v", got, tt.res)
		}
	}
}

func TestAdd64(t *testing.T) {
	type args struct {
		a uint64
		b uint64
	}
	tests := []struct {
		args args
		res  uint64
		err  bool
	}{
		{args: args{0, 1}, res: 1, err: false},
		{args: args{1 << 32, 1}, res: 4294967297, err: false},
		{args: args{1 << 32, 100}, res: 4294967396, err: false},
		{args: args{1 << 31, 1 << 31}, res: 4294967296, err: false},
		{args: args{1 << 63, 1 << 63}, res: 0, err: true},
		{args: args{1 << 63, 1}, res: 9223372036854775809, err: false},
		{args: args{math.MaxUint64, 1}, res: 0, err: true},
		{args: args{math.MaxUint64, 0}, res: math.MaxUint64, err: false},
		{args: args{1 << 63, 2}, res: 9223372036854775810, err: false},
	}
	for _, tt := range tests {
		got, err := types.Add64(tt.args.a, tt.args.b)
		if tt.err && err == nil {
			t.Errorf("Add64() Expected Error = %v, want error", tt.err)
			continue
		}
		if tt.res != got {
			t.Errorf("Add64() %v, want %v", got, tt.res)
		}
	}
}
