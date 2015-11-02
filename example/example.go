package main

import (
	"fmt"
	"github.com/gorward/cipher"
	"log"
)

func main() {
	key := []byte("thiskeyforuser1234567890") // 32 bytes or 24 characters
	text := []byte("Hi this is the text to be encrypted")

	fmt.Printf("\n%s\n", text)

	ciphertext, err := cipher.Encrypt(key, text)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%0x\n", ciphertext)

	result, err := cipher.Decrypt(key, ciphertext)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%s\n", result)
}
