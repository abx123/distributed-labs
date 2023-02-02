package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

const (
	keyList string = "abcdefghijklmnopqrstuvwxyzABCDEFHFGHIJKLMNOPQRSTUVWXYZ1234567890"
)

func main() {
	var bint big.Int
	base, pow8, pow16, pow32, pow64, pow128, pow256, pow512, pow1024, pow2048, pow4196 := big.NewInt(2), big.NewInt(8), big.NewInt(16), big.NewInt(32), big.NewInt(64), big.NewInt(128), big.NewInt(256), big.NewInt(512), big.NewInt(1024), big.NewInt(2048), big.NewInt(4196)
	key8 := genkey(8)
	key16 := genkey(16)
	key32 := genkey(32)
	key64 := genkey(64)
	key128 := genkey(128)
	key256 := genkey(256)
	key512 := genkey(512)
	key1024 := genkey(1024)
	key2048 := genkey(2048)
	key4196 := genkey(4196)

	fmt.Printf("If the key length is %s bits, then the key space is %d. Generated key: %s \n", pow8.String(), bint.Exp(pow8, base, nil), key8)
	fmt.Printf("If the key length is %s bits, then the key space is %d. Generated key: %s \n", pow16.String(), bint.Exp(pow16, base, nil), key16)
	fmt.Printf("If the key length is %s bits, then the key space is %d. Generated key: %s \n", pow32.String(), bint.Exp(pow32, base, nil), key32)
	fmt.Printf("If the key length is %s bits, then the key space is %d. Generated key: %s \n", pow64.String(), bint.Exp(pow64, base, nil), key64)
	fmt.Printf("If the key length is %s bits, then the key space is %d. Generated key: %s \n", pow128.String(), bint.Exp(pow128, base, nil), key128)
	fmt.Printf("If the key length is %s bits, then the key space is %d. Generated key: %s \n", pow256.String(), bint.Exp(pow256, base, nil), key256)
	fmt.Printf("If the key length is %s bits, then the key space is %d. Generated key: %s \n", pow512.String(), bint.Exp(pow512, base, nil), key512)
	fmt.Printf("If the key length is %s bits, then the key space is %d. Generated key: %s \n", pow1024.String(), bint.Exp(pow1024, base, nil), key1024)
	fmt.Printf("If the key length is %s bits, then the key space is %d. Generated key: %s \n", pow2048.String(), bint.Exp(pow2048, base, nil), key2048)
	fmt.Printf("If the key length is %s bits, then the key space is %d. Generated key: %s \n", pow4196.String(), bint.Exp(pow4196, base, nil), key4196)

	bruteForce(key8, pow8.String())
	bruteForce(key16, pow16.String())
	bruteForce(key32, pow32.String())
	bruteForce(key64, pow64.String())
	bruteForce(key128, pow128.String())
	bruteForce(key256, pow256.String())
	bruteForce(key512, pow512.String())
	bruteForce(key1024, pow1024.String())
	bruteForce(key2048, pow2048.String())
	bruteForce(key4196, pow4196.String())
}

func genkey(size int) []byte {
	key := []byte{}
	for k := 1; k <= size; k++ {
		res, _ := rand.Int(rand.Reader, big.NewInt(64))
		keyGen := keyList[res.Int64()]
		key = append(key, keyGen)
	}

	return key
}

func bruteForce(key []byte, bits string) {
	start := time.Now()
	genKey := []byte{}
	for i := 0; i < len(key); i++ {
		for j := 0; j < len(keyList); j++ {
			if keyList[j] == key[i] {
				genKey = append(genKey, keyList[j])
			}
		}
	}
	fmt.Printf("Took %d ns to brute force key with %s length bits \n", time.Since(start).Nanoseconds(), bits)
}
