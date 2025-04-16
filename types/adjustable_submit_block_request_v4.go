package types

import (
	d "github.com/attestantio/go-builder-client/api/deneb"
	v1 "github.com/attestantio/go-builder-client/api/v1"
	deneb "github.com/attestantio/go-eth2-client/spec/deneb"
	electra "github.com/attestantio/go-eth2-client/spec/electra"
	phase0 "github.com/attestantio/go-eth2-client/spec/phase0"
)

type AdjustableSubmitBlockRequestV4 struct {
	Message           *v1.BidTrace
	ExecutionPayload  *deneb.ExecutionPayload
	BlobsBundle       *d.BlobsBundle
	ExecutionRequests *electra.ExecutionRequests
	Signature         phase0.BLSSignature `ssz-size:"96"`
	AdjustmentData    *AdjustmentData
}
