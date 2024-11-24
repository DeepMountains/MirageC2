package exchange

import (
	"crypto/aes"
	"crypto/cipher"
)

// 解密命令执行的内容
func ResultDecode(ciphertext []byte, crkey string) string {
	block, _ := aes.NewCipher([]byte(crkey))
	commonIV := []byte("Mirage&&DeepMoun")
	cfbdec := cipher.NewCFBDecrypter(block, commonIV)
	plaintext := make([]byte, len(ciphertext))
	cfbdec.XORKeyStream(plaintext, ciphertext)
	return string(plaintext)
}

// 加密命令执行的内容
func CommandCrypto(org string, crkey string) []byte {
	plaintext := []byte(org)
	block, _ := aes.NewCipher([]byte(crkey))
	commonIV := []byte("Mirage&&DeepMoun")
	cfb := cipher.NewCFBEncrypter(block, commonIV)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	return ciphertext
}
