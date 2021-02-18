package types_test

import (
	"reflect"
	"strings"
	"testing"

	types "github.com/prysmaticlabs/eth2-types"
)

func TestErrorMessage_Limit(t *testing.T) {
	errorMessage := make([]byte, 0)
	// Provide a message of size 6400 bytes.
	for i := uint64(0); i < 200; i++ {
		byteArr := [32]byte{byte(i)}
		errorMessage = append(errorMessage, byteArr[:]...)
	}
	errMsg := types.ErrorMessage{}
	err := errMsg.UnmarshalSSZ(errorMessage)
	if err == nil || !strings.Contains(err.Error(), "expected buffer of length upto") {
		t.Errorf("Expected Error = %s, got: %v", "expected buffer of length upto", err)
	}
}

func TestTestErrorMessage_RoundTrip(t *testing.T) {
	errMsg := []byte{'e', 'r', 'r', 'o', 'r'}
	sszErr := make(types.ErrorMessage, len(errMsg))
	copy(sszErr, errMsg)

	marshalledObj, err := sszErr.MarshalSSZ()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	newVal := types.ErrorMessage(nil)

	err = newVal.UnmarshalSSZ(marshalledObj)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual([]byte(newVal), errMsg) {
		t.Errorf("Unequal: %v = %v", []byte(newVal), errMsg)
	}
}
