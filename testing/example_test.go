package testing

import (
	"encoding/hex"
	"fmt"
	"log"
	"testing"
)

func TestHex(t *testing.T) {
	b := []byte("abcdhh")
	var result [3]byte
	c, err := hex.Decode(result[:], b)
	if err != nil {
		fmt.Errorf("error:%w", err)
	}

	log.Printf("%v", c)
}
