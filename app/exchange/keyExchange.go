package exchange

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
)

func RandomOriginalkey() string {
	randomBytes := make([]byte, 32)
	rand.Read(randomBytes)
	Originalkey := base64.URLEncoding.EncodeToString(randomBytes)[:32]
	return Originalkey
}

func Cryptokey(originalkey string) string {
	orgkey32 := originalkey
	first16 := orgkey32[:16]
	last16 := orgkey32[len(orgkey32)-16:]
	hash := md5.Sum([]byte(last16))
	hashString := hex.EncodeToString(hash[:])
	crkey32 := first16 + hashString[:16]
	return crkey32
}
