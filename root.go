package types

import (
	"fmt"

	fssz "github.com/ferranbt/fastssz"
)

var _ fssz.HashRoot = (Root)([32]byte{})
var _ fssz.Marshaler = (*Root)(nil)
var _ fssz.Unmarshaler = (*Root)(nil)

// Root represents a 32 bytes root object in Ethereum beacon chain consensus.
type Root [32]byte

// HashTreeRoot returns calculated hash root.
func (e Root) HashTreeRoot() ([32]byte, error) {
	return fssz.HashWithDefaultHasher(e)
}

// HashTreeRootWith hashes a HashRoot object with a Hasher from the default HasherPool.
func (e Root) HashTreeRootWith(hh *fssz.Hasher) error {
	hh.PutBytes(e[:])
	return nil
}

// UnmarshalSSZ deserializes the provided bytes buffer into the Root object.
func (e *Root) UnmarshalSSZ(buf []byte) error {
	if len(buf) != e.SizeSSZ() {
		return fmt.Errorf("expected buffer of length %d received %d", e.SizeSSZ(), len(buf))
	}
	var item Root
	copy(item[:], buf)
	*e = item
	return nil
}

// MarshalSSZTo marshals Root with the provided byte slice.
func (e *Root) MarshalSSZTo(dst []byte) ([]byte, error) {
	marshalled, err := e.MarshalSSZ()
	if err != nil {
		return nil, err
	}
	return append(dst, marshalled...), nil
}

// MarshalSSZ marshals Root into a serialized object.
func (e *Root) MarshalSSZ() ([]byte, error) {
	return e[:], nil
}

// SizeSSZ returns the size of the serialized object.
func (e *Root) SizeSSZ() int {
	return 32
}
