package hashblock

import (
	"crypto/sha256"
	"hash"
	"io"
	"log"
)

type HashBlocker struct {
	h hash.Hash
}

func (hb HashBlocker) Write(b []byte) (int, error) {
	return hb.h.Write(b)
}

func (hb HashBlocker) Add(r io.Reader) int64 {
	n, err := io.Copy(hb, r)

	if err != nil {
		log.Printf("hb.Add error: %v\n", err)
	}

	return n
}

func New() HashBlocker {
	return HashBlocker{
		sha256.New(),
	}
}

func (hb HashBlocker) Generate() []byte {
	sum := hb.h.Sum(nil)
	hb.h.Reset()

	return sum
}