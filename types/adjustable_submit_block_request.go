package types

import (
	d "github.com/attestantio/go-builder-client/api/deneb"
	v1 "github.com/attestantio/go-builder-client/api/v1"
	deneb "github.com/attestantio/go-eth2-client/spec/deneb"
	phase0 "github.com/attestantio/go-eth2-client/spec/phase0"
)

type AdjustableSubmitBlockRequest struct {
	Message          *v1.BidTrace
	ExecutionPayload *deneb.ExecutionPayload
	BlobsBundle      *d.BlobsBundle
	Signature        phase0.BLSSignature `ssz-size:"96"`
	AdjustmentData   *AdjustmentData
}
