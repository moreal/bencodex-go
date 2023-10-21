package main

import (
	"encoding/json"
	"fmt"
	bencode "github.com/moreal/bencodex-go"
	"github.com/moreal/bencodex-go/internal/decoder"
	"io"
	"os"
)

func main() {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	decoder := decoder.Decoder{}
	result, err := decoder.Decode(bytes)
	if err != nil {
		panic(err)
	}

	converted, err := bencodex.ConvertToBencodexJson(result)
	if err != nil {
		panic(err)
	}

	jsonEncoder := json.NewEncoder(os.Stdout)
	err = jsonEncoder.Encode(converted)
	if err != nil {
		fmt.Errorf("%v", err)
	}
}
