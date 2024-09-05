package compress

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompress(t *testing.T) {

	data := []byte(strings.Repeat("This is a test message", 20))
	b, err := Compress(data)

	if err != nil {
		t.Fatalf("%q", err.Error())
	}

	t.Logf("%d bytes has been compressed to %d bytes\r\n", len(data), len(b))

	out, err := Decompress(b)

	if err != nil {
		t.Fatalf("%q", err.Error())
	}

	if !bytes.Equal(data, out) {
		t.Fatalf("origin data != decompressed data")
	}

	require.Equal(t, data, out)
}
