package util

import (
	cryptorand "crypto/rand"
	"encoding/binary"
	"math/rand"
)

// InitRand は安全な乱数シードを設定します。
func InitRand() error {
	var seed int64
	err := binary.Read(cryptorand.Reader, binary.LittleEndian, &seed)
	if err != nil {
		return err
	}
	rand.Seed(seed)
	return nil
}
