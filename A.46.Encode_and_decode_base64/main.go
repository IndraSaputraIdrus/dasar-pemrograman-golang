package main

import (
	"encoding/base64"
	"fmt"
)

func encodeToStringAndDecodeString() {
	var data = "john wick"

	var encodedString = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(encodedString)

	var decodedByte, _ = base64.StdEncoding.DecodeString(encodedString)
	var decodedString = string(decodedByte)
	fmt.Println("decoded", decodedString)
}

func encodeAndDecode() {
	var data = "john wick"

	var encoded = make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded, []byte(data))

	var encodedString = string(encoded)
	fmt.Println(encodedString)

	var decoded = make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	var _, err = base64.StdEncoding.Decode(decoded, encoded)
	if err != nil {
		panic(err.Error())
	}

	var decodedString = string(decoded)
	fmt.Println(decodedString)
}

func encodeAndDecodeURL() {
	var data = "https://kalipare.com"

	var encodedString = base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(encodedString)

	var decodedByte, _ = base64.StdEncoding.DecodeString(encodedString)
	var decodedString = string(decodedByte)
	fmt.Println("decoded", decodedString)
}

func main() {
	// encodeToStringAndDecodeString()

	// encodeAndDecode()

	encodeAndDecodeURL()
}
