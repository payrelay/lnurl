package lnurl

import (
	"bytes"
	"errors"
	"strings"
	_ "unsafe"

	"github.com/btcsuite/btcd/btcutil/bech32"
	_ "github.com/lightningnetwork/lnd/zpay32"
)

//go:linkname decodeBech32 github.com/lightningnetwork/lnd/zpay32.decodeBech32
//go:noescape
func decodeBech32(bech string) (string, []byte, error)

func Encode(raw string) (string, error) {
	body, err := bech32.ConvertBits([]byte(raw), 8, 5, true)

	if err != nil {
		return "", err
	}

	encoded, err := bech32.Encode("lnurl", body)

	if err != nil {
		return "", err
	}

	uppercase := strings.ToUpper(encoded)

	return uppercase, nil
}

func Decode(raw string) (string, error) {
	prefix, content, err := decodeBech32(raw)

	if err != nil {
		return "", err
	}

	if prefix != "lnurl" {
		return "", errors.New("unexpected prefix")
	}

	decoded, err := bech32.ConvertBits(content, 5, 8, true)

	if err != nil {
		return "", err
	}

	trimmed := bytes.Trim(decoded, "\x00")

	return string(trimmed), nil
}
