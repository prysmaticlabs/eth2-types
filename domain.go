package types

import (
	"fmt"

	fssz "github.com/ferranbt/fastssz"
)

var _ fssz.HashRoot = (Domain)([32]byte{})
var _ fssz.Marshaler = (*Domain)(nil)
var _ fssz.Unmarshaler = (*Domain)(nil)

// Domain represents a 32 bytes domain object in Ethereum beacon chain consensus.
type Domain [32]byte

// HashTreeRoot returns calculated hash root.
func (e Domain) HashTreeRoot() ([32]byte, error) {
	return fssz.HashWithDefaultHasher(e)
}

// HashTreeRootWith hashes a Domain object with a Hasher from the default HasherPool.
func (e Domain) HashTreeRootWith(hh *fssz.Hasher) error {
	hh.PutBytes(e[:])
	return nil
}

// UnmarshalSSZ deserializes the provided bytes buffer into the Domain object.
func (e *Domain) UnmarshalSSZ(buf []byte) error {
	if len(buf) != e.SizeSSZ() {
		return fmt.Errorf("expected buffer of length %d received %d", e.SizeSSZ(), len(buf))
	}
	var item Domain
	copy(item[:], buf)
	*e = item
	return nil
}

// MarshalSSZTo marshals Domain with the provided byte slice.
func (e *Domain) MarshalSSZTo(dst []byte) ([]byte, error) {
	marshalled, err := e.MarshalSSZ()
	if err != nil {
		return nil, err
	}
	return append(dst, marshalled...), nil
}

// MarshalSSZ marshals Domain into a serialized object.
func (e *Domain) MarshalSSZ() ([]byte, error) {
	return e[:], nil
}

// SizeSSZ returns the size of the serialized object.
func (e *Domain) SizeSSZ() int {
	return 32
}
