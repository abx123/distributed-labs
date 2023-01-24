package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("encrypting plain text: super secret plain text")

	text := []byte("super secret plain text")
	key := []byte("passphrasewhichneedstobe32bytes!")

	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	encrypted := gcm.Seal(nonce, nonce, text, nil)
	fmt.Println("encrypted text: " + string(encrypted))

	nonceSize := gcm.NonceSize()
	if len(encrypted) < nonceSize {
		fmt.Println(err)
	}

	nonce, ciphertext := encrypted[:nonceSize], encrypted[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("decrypted text:" + string(plaintext))

	if string(plaintext) != string(text) {
		fmt.Println("input and output text does not match")
		os.Exit(1)
	} else {
		fmt.Println("input and output text matches")
	}
}
