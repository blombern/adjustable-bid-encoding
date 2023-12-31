// Code generated by fastssz. DO NOT EDIT.
// Hash: e755a279d96af9adfab03bb68b25db799ff79a53a1dfe0f09691199725453372
// Version: 0.1.3
package types

import (
	v1 "github.com/attestantio/go-builder-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/capella"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the AdjustableSubmitBlockRequest object
func (a *AdjustableSubmitBlockRequest) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(a)
}

// MarshalSSZTo ssz marshals the AdjustableSubmitBlockRequest object to a target array
func (a *AdjustableSubmitBlockRequest) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(340)

	// Field (0) 'Message'
	if a.Message == nil {
		a.Message = new(v1.BidTrace)
	}
	if dst, err = a.Message.MarshalSSZTo(dst); err != nil {
		return
	}

	// Offset (1) 'ExecutionPayload'
	dst = ssz.WriteOffset(dst, offset)
	if a.ExecutionPayload == nil {
		a.ExecutionPayload = new(capella.ExecutionPayload)
	}
	offset += a.ExecutionPayload.SizeSSZ()

	// Field (2) 'Signature'
	dst = append(dst, a.Signature[:]...)

	// Offset (3) 'AdjustmentData'
	dst = ssz.WriteOffset(dst, offset)
	if a.AdjustmentData == nil {
		a.AdjustmentData = new(AdjustmentData)
	}
	offset += a.AdjustmentData.SizeSSZ()

	// Field (1) 'ExecutionPayload'
	if dst, err = a.ExecutionPayload.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (3) 'AdjustmentData'
	if dst, err = a.AdjustmentData.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the AdjustableSubmitBlockRequest object
func (a *AdjustableSubmitBlockRequest) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 340 {
		return ssz.ErrSize
	}

	tail := buf
	var o1, o3 uint64

	// Field (0) 'Message'
	if a.Message == nil {
		a.Message = new(v1.BidTrace)
	}
	if err = a.Message.UnmarshalSSZ(buf[0:236]); err != nil {
		return err
	}

	// Offset (1) 'ExecutionPayload'
	if o1 = ssz.ReadOffset(buf[236:240]); o1 > size {
		return ssz.ErrOffset
	}

	if o1 < 340 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (2) 'Signature'
	copy(a.Signature[:], buf[240:336])

	// Offset (3) 'AdjustmentData'
	if o3 = ssz.ReadOffset(buf[336:340]); o3 > size || o1 > o3 {
		return ssz.ErrOffset
	}

	// Field (1) 'ExecutionPayload'
	{
		buf = tail[o1:o3]
		if a.ExecutionPayload == nil {
			a.ExecutionPayload = new(capella.ExecutionPayload)
		}
		if err = a.ExecutionPayload.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (3) 'AdjustmentData'
	{
		buf = tail[o3:]
		if a.AdjustmentData == nil {
			a.AdjustmentData = new(AdjustmentData)
		}
		if err = a.AdjustmentData.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the AdjustableSubmitBlockRequest object
func (a *AdjustableSubmitBlockRequest) SizeSSZ() (size int) {
	size = 340

	// Field (1) 'ExecutionPayload'
	if a.ExecutionPayload == nil {
		a.ExecutionPayload = new(capella.ExecutionPayload)
	}
	size += a.ExecutionPayload.SizeSSZ()

	// Field (3) 'AdjustmentData'
	if a.AdjustmentData == nil {
		a.AdjustmentData = new(AdjustmentData)
	}
	size += a.AdjustmentData.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the AdjustableSubmitBlockRequest object
func (a *AdjustableSubmitBlockRequest) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(a)
}

// HashTreeRootWith ssz hashes the AdjustableSubmitBlockRequest object with a hasher
func (a *AdjustableSubmitBlockRequest) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Message'
	if a.Message == nil {
		a.Message = new(v1.BidTrace)
	}
	if err = a.Message.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'ExecutionPayload'
	if err = a.ExecutionPayload.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (2) 'Signature'
	hh.PutBytes(a.Signature[:])

	// Field (3) 'AdjustmentData'
	if err = a.AdjustmentData.HashTreeRootWith(hh); err != nil {
		return
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the AdjustableSubmitBlockRequest object
func (a *AdjustableSubmitBlockRequest) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(a)
}
