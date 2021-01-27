package types

import (
	"fmt"

	fssz "github.com/ferranbt/fastssz"
)

var _ fssz.HashRoot = (CommitteeIndex)(0)
var _ fssz.Marshaler = (*CommitteeIndex)(nil)
var _ fssz.Unmarshaler = (*CommitteeIndex)(nil)

// CommitteeIndex in eth2.
type CommitteeIndex uint64

// HashTreeRoot returns calculated hash root.
func (c CommitteeIndex) HashTreeRoot() ([32]byte, error) {
	return fssz.HashWithDefaultHasher(c)
}

// HashWithDefaultHasher hashes a HashRoot object with a Hasher from the default HasherPool.
func (c CommitteeIndex) HashTreeRootWith(hh *fssz.Hasher) error {
	hh.PutUint64(uint64(c))
	return nil
}

// UnmarshalSSZ deserializes the provided bytes buffer into the committee index object.
func (c *CommitteeIndex) UnmarshalSSZ(buf []byte) error {
	if len(buf) != c.SizeSSZ() {
		return fmt.Errorf("expected buffer of length %d receiced %d", c.SizeSSZ(), len(buf))
	}
	*c = CommitteeIndex(fssz.UnmarshallUint64(buf))
	return nil
}

// MarshalSSZTo marshals committee index with the provided byte slice.
func (c *CommitteeIndex) MarshalSSZTo(dst []byte) ([]byte, error) {
	marshalled, err := c.MarshalSSZ()
	if err != nil {
		return nil, err
	}
	return append(dst, marshalled...), nil
}

// MarshalSSZ marshals committee index into a serialized object.
func (c *CommitteeIndex) MarshalSSZ() ([]byte, error) {
	marshalled := fssz.MarshalUint64([]byte{}, uint64(*c))
	return marshalled, nil
}

// SizeSSZ returns the size of the serialized object.
func (c *CommitteeIndex) SizeSSZ() int {
	return 8
}
