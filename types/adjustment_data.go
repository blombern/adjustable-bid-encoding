package types

// I couldn't find any info on merkle proof max sizes, so they're set to "large enough"
type AdjustmentData struct {
	StateRoot               [32]byte `ssz-size:"32"`
	TransactionsRoot        [32]byte `ssz-size:"32"`
	ReceiptsRoot            [32]byte `ssz-size:"32"`
	BuilderAddress          [20]byte `ssz-size:"20"`
	BuilderProof            [][]byte `ssz-size:"?,?" ssz-max:"64,1073741824"`
	FeeRecipientAddress     [20]byte `ssz-size:"20"`
	FeeRecipientProof       [][]byte `ssz-size:"?,?" ssz-max:"64,1073741824"`
	FeePayerAddress         [20]byte `ssz-size:"20"`
	FeePayerProof           [][]byte `ssz-size:"?,?" ssz-max:"64,1073741824"`
	PlaceholderTxProof      [][]byte `ssz-size:"?,?" ssz-max:"64,1073741824"`
	PlaceholderReceiptProof [][]byte `ssz-size:"?,?" ssz-max:"64,1073741824"`
}
