package types_test

import (
	"reflect"
	"testing"

	types "github.com/prysmaticlabs/eth2-types"
)

func TestSSZBytes_HashTreeRoot(t *testing.T) {
	tests := []struct {
		name        string
		actualValue []byte
		root        []byte
		wantErr     bool
	}{
		{
			name:        "random1",
			actualValue: types.HexDecodeOrDie(t, "844e1063e0b396eed17be8eddb7eecd1fe3ea46542a4b72f7466e77325e5aa6d"),
			root:        types.HexDecodeOrDie(t, "844e1063e0b396eed17be8eddb7eecd1fe3ea46542a4b72f7466e77325e5aa6d"),
			wantErr:     false,
		},
		{
			name:        "random1",
			actualValue: types.HexDecodeOrDie(t, "7b16162ecd9a28fa80a475080b0e4fff4c27efe19ce5134ce3554b72274d59fd534400ba4c7f699aa1c307cd37c2b103"),
			root:        types.HexDecodeOrDie(t, "128ed34ee798b9f00716f9ba5c000df5c99443dabc4d3f2e9bb86c77c732e007"),
			wantErr:     false,
		},
		{
			name:        "random2",
			actualValue: []byte{},
			root:        types.HexDecodeOrDie(t, "0000000000000000000000000000000000000000000000000000000000000000"),
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := types.SSZBytes(tt.actualValue)
			htr, err := s.HashTreeRoot()
			if err != nil {
				t.Errorf("SSZBytes.HashTreeRoot() unexpected error = %v", err)
			}
			if !reflect.DeepEqual(tt.root, htr[:]) {
				t.Errorf("SSZBytes.HashTreeRoot() = %v, want %v", htr[:], tt.root)
			}
		})
	}
}
