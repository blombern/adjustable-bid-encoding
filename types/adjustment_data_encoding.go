// Code generated by fastssz. DO NOT EDIT.
// Hash: 5a0ca77ff7a9cf370182a3600831526046fe566da0897a655ee542a11ba5ef60
// Version: 0.1.3
package types

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the AdjustmentData object
func (a *AdjustmentData) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(a)
}

// MarshalSSZTo ssz marshals the AdjustmentData object to a target array
func (a *AdjustmentData) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(176)

	// Field (0) 'StateRoot'
	dst = append(dst, a.StateRoot[:]...)

	// Field (1) 'TransactionsRoot'
	dst = append(dst, a.TransactionsRoot[:]...)

	// Field (2) 'ReceiptsRoot'
	dst = append(dst, a.ReceiptsRoot[:]...)

	// Field (3) 'BuilderAddress'
	dst = append(dst, a.BuilderAddress[:]...)

	// Offset (4) 'BuilderProof'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(a.BuilderProof); ii++ {
		offset += 4
		offset += len(a.BuilderProof[ii])
	}

	// Field (5) 'FeeRecipientAddress'
	dst = append(dst, a.FeeRecipientAddress[:]...)

	// Offset (6) 'FeeRecipientProof'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(a.FeeRecipientProof); ii++ {
		offset += 4
		offset += len(a.FeeRecipientProof[ii])
	}

	// Field (7) 'FeePayerAddress'
	dst = append(dst, a.FeePayerAddress[:]...)

	// Offset (8) 'FeePayerProof'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(a.FeePayerProof); ii++ {
		offset += 4
		offset += len(a.FeePayerProof[ii])
	}

	// Offset (9) 'PlaceholderTxProof'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(a.PlaceholderTxProof); ii++ {
		offset += 4
		offset += len(a.PlaceholderTxProof[ii])
	}

	// Offset (10) 'PlaceholderReceiptProof'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(a.PlaceholderReceiptProof); ii++ {
		offset += 4
		offset += len(a.PlaceholderReceiptProof[ii])
	}

	// Field (4) 'BuilderProof'
	if size := len(a.BuilderProof); size > 64 {
		err = ssz.ErrListTooBigFn("AdjustmentData.BuilderProof", size, 64)
		return
	}
	{
		offset = 4 * len(a.BuilderProof)
		for ii := 0; ii < len(a.BuilderProof); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += len(a.BuilderProof[ii])
		}
	}
	for ii := 0; ii < len(a.BuilderProof); ii++ {
		if size := len(a.BuilderProof[ii]); size > 1073741824 {
			err = ssz.ErrBytesLengthFn("AdjustmentData.BuilderProof[ii]", size, 1073741824)
			return
		}
		dst = append(dst, a.BuilderProof[ii]...)
	}

	// Field (6) 'FeeRecipientProof'
	if size := len(a.FeeRecipientProof); size > 64 {
		err = ssz.ErrListTooBigFn("AdjustmentData.FeeRecipientProof", size, 64)
		return
	}
	{
		offset = 4 * len(a.FeeRecipientProof)
		for ii := 0; ii < len(a.FeeRecipientProof); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += len(a.FeeRecipientProof[ii])
		}
	}
	for ii := 0; ii < len(a.FeeRecipientProof); ii++ {
		if size := len(a.FeeRecipientProof[ii]); size > 1073741824 {
			err = ssz.ErrBytesLengthFn("AdjustmentData.FeeRecipientProof[ii]", size, 1073741824)
			return
		}
		dst = append(dst, a.FeeRecipientProof[ii]...)
	}

	// Field (8) 'FeePayerProof'
	if size := len(a.FeePayerProof); size > 64 {
		err = ssz.ErrListTooBigFn("AdjustmentData.FeePayerProof", size, 64)
		return
	}
	{
		offset = 4 * len(a.FeePayerProof)
		for ii := 0; ii < len(a.FeePayerProof); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += len(a.FeePayerProof[ii])
		}
	}
	for ii := 0; ii < len(a.FeePayerProof); ii++ {
		if size := len(a.FeePayerProof[ii]); size > 1073741824 {
			err = ssz.ErrBytesLengthFn("AdjustmentData.FeePayerProof[ii]", size, 1073741824)
			return
		}
		dst = append(dst, a.FeePayerProof[ii]...)
	}

	// Field (9) 'PlaceholderTxProof'
	if size := len(a.PlaceholderTxProof); size > 64 {
		err = ssz.ErrListTooBigFn("AdjustmentData.PlaceholderTxProof", size, 64)
		return
	}
	{
		offset = 4 * len(a.PlaceholderTxProof)
		for ii := 0; ii < len(a.PlaceholderTxProof); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += len(a.PlaceholderTxProof[ii])
		}
	}
	for ii := 0; ii < len(a.PlaceholderTxProof); ii++ {
		if size := len(a.PlaceholderTxProof[ii]); size > 1073741824 {
			err = ssz.ErrBytesLengthFn("AdjustmentData.PlaceholderTxProof[ii]", size, 1073741824)
			return
		}
		dst = append(dst, a.PlaceholderTxProof[ii]...)
	}

	// Field (10) 'PlaceholderReceiptProof'
	if size := len(a.PlaceholderReceiptProof); size > 64 {
		err = ssz.ErrListTooBigFn("AdjustmentData.PlaceholderReceiptProof", size, 64)
		return
	}
	{
		offset = 4 * len(a.PlaceholderReceiptProof)
		for ii := 0; ii < len(a.PlaceholderReceiptProof); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += len(a.PlaceholderReceiptProof[ii])
		}
	}
	for ii := 0; ii < len(a.PlaceholderReceiptProof); ii++ {
		if size := len(a.PlaceholderReceiptProof[ii]); size > 1073741824 {
			err = ssz.ErrBytesLengthFn("AdjustmentData.PlaceholderReceiptProof[ii]", size, 1073741824)
			return
		}
		dst = append(dst, a.PlaceholderReceiptProof[ii]...)
	}

	return
}

// UnmarshalSSZ ssz unmarshals the AdjustmentData object
func (a *AdjustmentData) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 176 {
		return ssz.ErrSize
	}

	tail := buf
	var o4, o6, o8, o9, o10 uint64

	// Field (0) 'StateRoot'
	copy(a.StateRoot[:], buf[0:32])

	// Field (1) 'TransactionsRoot'
	copy(a.TransactionsRoot[:], buf[32:64])

	// Field (2) 'ReceiptsRoot'
	copy(a.ReceiptsRoot[:], buf[64:96])

	// Field (3) 'BuilderAddress'
	copy(a.BuilderAddress[:], buf[96:116])

	// Offset (4) 'BuilderProof'
	if o4 = ssz.ReadOffset(buf[116:120]); o4 > size {
		return ssz.ErrOffset
	}

	if o4 < 176 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (5) 'FeeRecipientAddress'
	copy(a.FeeRecipientAddress[:], buf[120:140])

	// Offset (6) 'FeeRecipientProof'
	if o6 = ssz.ReadOffset(buf[140:144]); o6 > size || o4 > o6 {
		return ssz.ErrOffset
	}

	// Field (7) 'FeePayerAddress'
	copy(a.FeePayerAddress[:], buf[144:164])

	// Offset (8) 'FeePayerProof'
	if o8 = ssz.ReadOffset(buf[164:168]); o8 > size || o6 > o8 {
		return ssz.ErrOffset
	}

	// Offset (9) 'PlaceholderTxProof'
	if o9 = ssz.ReadOffset(buf[168:172]); o9 > size || o8 > o9 {
		return ssz.ErrOffset
	}

	// Offset (10) 'PlaceholderReceiptProof'
	if o10 = ssz.ReadOffset(buf[172:176]); o10 > size || o9 > o10 {
		return ssz.ErrOffset
	}

	// Field (4) 'BuilderProof'
	{
		buf = tail[o4:o6]
		num, err := ssz.DecodeDynamicLength(buf, 64)
		if err != nil {
			return err
		}
		a.BuilderProof = make([][]byte, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if len(buf) > 1073741824 {
				return ssz.ErrBytesLength
			}
			if cap(a.BuilderProof[indx]) == 0 {
				a.BuilderProof[indx] = make([]byte, 0, len(buf))
			}
			a.BuilderProof[indx] = append(a.BuilderProof[indx], buf...)
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (6) 'FeeRecipientProof'
	{
		buf = tail[o6:o8]
		num, err := ssz.DecodeDynamicLength(buf, 64)
		if err != nil {
			return err
		}
		a.FeeRecipientProof = make([][]byte, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if len(buf) > 1073741824 {
				return ssz.ErrBytesLength
			}
			if cap(a.FeeRecipientProof[indx]) == 0 {
				a.FeeRecipientProof[indx] = make([]byte, 0, len(buf))
			}
			a.FeeRecipientProof[indx] = append(a.FeeRecipientProof[indx], buf...)
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (8) 'FeePayerProof'
	{
		buf = tail[o8:o9]
		num, err := ssz.DecodeDynamicLength(buf, 64)
		if err != nil {
			return err
		}
		a.FeePayerProof = make([][]byte, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if len(buf) > 1073741824 {
				return ssz.ErrBytesLength
			}
			if cap(a.FeePayerProof[indx]) == 0 {
				a.FeePayerProof[indx] = make([]byte, 0, len(buf))
			}
			a.FeePayerProof[indx] = append(a.FeePayerProof[indx], buf...)
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (9) 'PlaceholderTxProof'
	{
		buf = tail[o9:o10]
		num, err := ssz.DecodeDynamicLength(buf, 64)
		if err != nil {
			return err
		}
		a.PlaceholderTxProof = make([][]byte, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if len(buf) > 1073741824 {
				return ssz.ErrBytesLength
			}
			if cap(a.PlaceholderTxProof[indx]) == 0 {
				a.PlaceholderTxProof[indx] = make([]byte, 0, len(buf))
			}
			a.PlaceholderTxProof[indx] = append(a.PlaceholderTxProof[indx], buf...)
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (10) 'PlaceholderReceiptProof'
	{
		buf = tail[o10:]
		num, err := ssz.DecodeDynamicLength(buf, 64)
		if err != nil {
			return err
		}
		a.PlaceholderReceiptProof = make([][]byte, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if len(buf) > 1073741824 {
				return ssz.ErrBytesLength
			}
			if cap(a.PlaceholderReceiptProof[indx]) == 0 {
				a.PlaceholderReceiptProof[indx] = make([]byte, 0, len(buf))
			}
			a.PlaceholderReceiptProof[indx] = append(a.PlaceholderReceiptProof[indx], buf...)
			return nil
		})
		if err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the AdjustmentData object
func (a *AdjustmentData) SizeSSZ() (size int) {
	size = 176

	// Field (4) 'BuilderProof'
	for ii := 0; ii < len(a.BuilderProof); ii++ {
		size += 4
		size += len(a.BuilderProof[ii])
	}

	// Field (6) 'FeeRecipientProof'
	for ii := 0; ii < len(a.FeeRecipientProof); ii++ {
		size += 4
		size += len(a.FeeRecipientProof[ii])
	}

	// Field (8) 'FeePayerProof'
	for ii := 0; ii < len(a.FeePayerProof); ii++ {
		size += 4
		size += len(a.FeePayerProof[ii])
	}

	// Field (9) 'PlaceholderTxProof'
	for ii := 0; ii < len(a.PlaceholderTxProof); ii++ {
		size += 4
		size += len(a.PlaceholderTxProof[ii])
	}

	// Field (10) 'PlaceholderReceiptProof'
	for ii := 0; ii < len(a.PlaceholderReceiptProof); ii++ {
		size += 4
		size += len(a.PlaceholderReceiptProof[ii])
	}

	return
}

// HashTreeRoot ssz hashes the AdjustmentData object
func (a *AdjustmentData) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(a)
}

// HashTreeRootWith ssz hashes the AdjustmentData object with a hasher
func (a *AdjustmentData) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'StateRoot'
	hh.PutBytes(a.StateRoot[:])

	// Field (1) 'TransactionsRoot'
	hh.PutBytes(a.TransactionsRoot[:])

	// Field (2) 'ReceiptsRoot'
	hh.PutBytes(a.ReceiptsRoot[:])

	// Field (3) 'BuilderAddress'
	hh.PutBytes(a.BuilderAddress[:])

	// Field (4) 'BuilderProof'
	{
		subIndx := hh.Index()
		num := uint64(len(a.BuilderProof))
		if num > 64 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range a.BuilderProof {
			{
				elemIndx := hh.Index()
				byteLen := uint64(len(elem))
				if byteLen > 1073741824 {
					err = ssz.ErrIncorrectListSize
					return
				}
				hh.AppendBytes32(elem)
				hh.MerkleizeWithMixin(elemIndx, byteLen, (1073741824+31)/32)
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 64)
	}

	// Field (5) 'FeeRecipientAddress'
	hh.PutBytes(a.FeeRecipientAddress[:])

	// Field (6) 'FeeRecipientProof'
	{
		subIndx := hh.Index()
		num := uint64(len(a.FeeRecipientProof))
		if num > 64 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range a.FeeRecipientProof {
			{
				elemIndx := hh.Index()
				byteLen := uint64(len(elem))
				if byteLen > 1073741824 {
					err = ssz.ErrIncorrectListSize
					return
				}
				hh.AppendBytes32(elem)
				hh.MerkleizeWithMixin(elemIndx, byteLen, (1073741824+31)/32)
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 64)
	}

	// Field (7) 'FeePayerAddress'
	hh.PutBytes(a.FeePayerAddress[:])

	// Field (8) 'FeePayerProof'
	{
		subIndx := hh.Index()
		num := uint64(len(a.FeePayerProof))
		if num > 64 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range a.FeePayerProof {
			{
				elemIndx := hh.Index()
				byteLen := uint64(len(elem))
				if byteLen > 1073741824 {
					err = ssz.ErrIncorrectListSize
					return
				}
				hh.AppendBytes32(elem)
				hh.MerkleizeWithMixin(elemIndx, byteLen, (1073741824+31)/32)
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 64)
	}

	// Field (9) 'PlaceholderTxProof'
	{
		subIndx := hh.Index()
		num := uint64(len(a.PlaceholderTxProof))
		if num > 64 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range a.PlaceholderTxProof {
			{
				elemIndx := hh.Index()
				byteLen := uint64(len(elem))
				if byteLen > 1073741824 {
					err = ssz.ErrIncorrectListSize
					return
				}
				hh.AppendBytes32(elem)
				hh.MerkleizeWithMixin(elemIndx, byteLen, (1073741824+31)/32)
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 64)
	}

	// Field (10) 'PlaceholderReceiptProof'
	{
		subIndx := hh.Index()
		num := uint64(len(a.PlaceholderReceiptProof))
		if num > 64 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range a.PlaceholderReceiptProof {
			{
				elemIndx := hh.Index()
				byteLen := uint64(len(elem))
				if byteLen > 1073741824 {
					err = ssz.ErrIncorrectListSize
					return
				}
				hh.AppendBytes32(elem)
				hh.MerkleizeWithMixin(elemIndx, byteLen, (1073741824+31)/32)
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 64)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the AdjustmentData object
func (a *AdjustmentData) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(a)
}
