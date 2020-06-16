package hashblock

import (
	"bytes"
	"testing"
)

func TestHashBlocker_Write(t *testing.T) {
	hb := New()

	emptyHash := hb.Generate()

	b := []byte("hello")

	_, err := hb.Write(b)

	if err != nil {
		t.Error(err)
	}

	writtenHash := hb.Generate()

	if bytes.Equal(writtenHash, emptyHash) {
		t.Error("HashBlocker not written to.")
	}
}

func TestHashBlocker_Add(t *testing.T) {
	hb := New()

	emptyHash := hb.Generate()

	b := []byte("hello world")
	r := bytes.NewBuffer(b)

	hb.Add(r)

	writtenHash := hb.Generate()

	if bytes.Equal(writtenHash, emptyHash) {
		t.Error("HashBlocker not written to.")
	}
}

func TestHashBlocker_Generate(t *testing.T) {
	hb1 := New()
	hb2 := New()

	s := "hola"

	hb1.h.Write([]byte(s))
	hb2.h.Write([]byte(s))

	if !bytes.Equal(hb1.Generate(), hb2.Generate()) {
		t.Error("Generate doesn't generate identical hashes for the same string")
	}
}
