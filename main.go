package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/blombern/adjustment-encoding/types"
)

func main() {
	filename := "./data/adjustableSubmitBlockPayloadDeneb.ssz"

	sszBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed reading file: %s", err)
	}

	req := new(types.AdjustableSubmitBlockRequest)
	err = req.UnmarshalSSZ(sszBytes)

	if err != nil {
		log.Fatalf("failed unmarshalling ssz: %s", err)
	}

	roundtripped, err := req.MarshalSSZ()

	if err != nil {
		log.Fatalf("failed marshalling ssz: %s", err)
	}

	if !bytes.Equal(roundtripped, sszBytes) {
		log.Fatalf("failed roundtrip")
	}

	fmt.Println("success")
}
