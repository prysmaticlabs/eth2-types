package types

import (
	"encoding/binary"
	"fmt"

	fssz "github.com/ferranbt/fastssz"
)

var _ fssz.HashRoot = (Epoch)(0)
var _ fssz.Marshaler = (*Epoch)(nil)
var _ fssz.Unmarshaler = (*Epoch)(nil)

// SSZUint64 is a uint64 type that satisfies the fast-ssz interface.
type SSZUint64 uint64

// SizeSSZ returns the size of the serialized representation.
func (s *SSZUint64) SizeSSZ() int {
	return 8
}

// MarshalSSZTo marshals the uint64 with the provided byte slice.
func (s *SSZUint64) MarshalSSZTo(dst []byte) ([]byte, error) {
	marshalled, err := s.MarshalSSZ()
	if err != nil {
		return nil, err
	}
	return append(dst, marshalled...), nil
}

// MarshalSSZ marshals uin64 into a serialized object.
func (s *SSZUint64) MarshalSSZ() ([]byte, error) {
	marshalled := fssz.MarshalUint64([]byte{}, uint64(*s))
	return marshalled, nil
}

// UnmarshalSSZ deserializes the provided bytes buffer into the uint64 object.
func (s *SSZUint64) UnmarshalSSZ(buf []byte) error {
	if len(buf) != s.SizeSSZ() {
		return fmt.Errorf("expected buffer of length %d received %d", s.SizeSSZ(), len(buf))
	}
	*s = SSZUint64(fssz.UnmarshallUint64(buf))
	return nil
}

// HashTreeRoot returns calculated hash root.
func (s *SSZUint64) HashTreeRoot() ([32]byte, error) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(*s))
	var root [32]byte
	copy(root[:], buf)
	return root, nil
}

// HashWithDefaultHasher hashes a HashRoot object with a Hasher from the default HasherPool.
func (s *SSZUint64) HashTreeRootWith(hh *fssz.Hasher) error {
	indx := hh.Index()
	hh.PutUint64(uint64(*s))
	hh.Merkleize(indx)
	return nil
}
