package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"golang.org/x/crypto/sha3"
)

func main() {
	text := []byte("super secret plain text")
	hashVal := "3ad45f1a830758aa15d0e34877db6d5275a72775e91b10f4b3627d7d2cb8564a"
	for i := 0; i < 999; i++ {
		hash := sha3.NewLegacyKeccak256()

		var buf []byte
		hash.Write(text)
		buf = hash.Sum(nil)
		hashed := hex.EncodeToString(buf)
		fmt.Println(fmt.Printf("round %d - hashed val: %s", i, hashed))
		if hashed != hashVal {
			fmt.Println("hashed not expected value, exiting")
			os.Exit(1)
		}
	}

}
