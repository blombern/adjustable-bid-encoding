package types

import (
	v1 "github.com/attestantio/go-builder-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/capella"
)

type AdjustableSubmitBlockRequest struct {
	Message          *v1.BidTrace
	ExecutionPayload *capella.ExecutionPayload
	Signature        [96]byte `ssz-size:"96"`
	AdjustmentData   *AdjustmentData
}
