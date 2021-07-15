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

// Size of the array.
func (d *Domain) Size() int {
	return len(d)
}

// MarshalTo serializes the array to the buffer.
func (d *Domain) MarshalTo(data []byte) (int, error) {
	return copy(data[:d.Size()], d[:]), nil
}

// Unmarshal deserializes into the provided array type.
func (d *Domain) Unmarshal(data []byte) error {
	copy(d[:], data)
	return nil
}

// HashTreeRoot returns calculated hash root.
func (d Domain) HashTreeRoot() ([32]byte, error) {
	return fssz.HashWithDefaultHasher(d)
}

// HashTreeRootWith hashes a Domain object with a Hasher from the default HasherPool.
func (d Domain) HashTreeRootWith(hh *fssz.Hasher) error {
	hh.PutBytes(d[:])
	return nil
}

// UnmarshalSSZ deserializes the provided bytes buffer into the Domain object.
func (d *Domain) UnmarshalSSZ(buf []byte) error {
	if len(buf) != d.SizeSSZ() {
		return fmt.Errorf("expected buffer of length %d received %d", d.SizeSSZ(), len(buf))
	}

	var b [32]byte
	item := Domain(b)
	copy(item[:], buf)
	*d = item
	return nil
}

// MarshalSSZTo marshals Domain with the provided byte slice.
func (d *Domain) MarshalSSZTo(dst []byte) ([]byte, error) {
	marshalled, err := d.MarshalSSZ()
	if err != nil {
		return nil, err
	}
	return append(dst, marshalled...), nil
}

// MarshalSSZ marshals Domain into a serialized object.
func (d *Domain) MarshalSSZ() ([]byte, error) {
	b := [32]byte(*d)
	return b[:], nil
}

// SizeSSZ returns the size of the serialized object.
func (d *Domain) SizeSSZ() int {
	return 32
}
