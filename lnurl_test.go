package lnurl

import (
	"testing"
)

var (
	ExampleDecoded = "https://example.com"
	ExampleEncoded = "LNURL1DP68GURN8GHJ7ETCV9KHQMR99E3K7MGMQGN24"
)

func TestEncode(test *testing.T) {
	encoded, err := Encode(ExampleDecoded)

	if err != nil {
		test.Fatalf("failed to encode with: %v\n", err)
	}

	if encoded != ExampleEncoded {
		test.Fatalf("failed to encode, expected %s got %s\n", ExampleEncoded, encoded)
	}

	test.Logf("encoded: %s\n", encoded)
}

func TestDecode(test *testing.T) {
	decoded, err := Decode(ExampleEncoded)

	if err != nil {
		test.Fatalf("failed to decode with: %v\n", err)
	}

	if decoded != ExampleDecoded {
		test.Fatalf("failed to decode, expected %s got %s\n", ExampleDecoded, decoded)
	}

	test.Logf("decoded: %s\n", decoded)
}
