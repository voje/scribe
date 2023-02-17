package scribe

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
    b, err := decodeSvcaccJSON("./test_data/encoded.gpg", "secretpass")
    assert.NoError(t, err)

    t.Logf("Result: %s", b)
}

