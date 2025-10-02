package main

import (
	"crypto/sha1"
	"fmt"
	"hash"
	"time"
)

func basicHash(sha hash.Hash, text string) string {
	sha.Write([]byte(text))
	var encrypted = sha.Sum(nil)
	var encryptedString = fmt.Sprintf("%x", encrypted)

	return encryptedString
}

func hashUsingSalt(sha hash.Hash, text string) (string, string) {
	var salt = fmt.Sprintf("%d", time.Now().UnixNano())
	var saltedText = fmt.Sprintf("text: '%s', salt: %s", text, salt)
	sha.Write([]byte(saltedText))
	var encrypted = sha.Sum(nil)

	return fmt.Sprintf("%x", encrypted), salt
}

func main() {
	var text = "this is secret"
	var sha = sha1.New()

	var encrypt = basicHash(sha, text)
	fmt.Println(encrypt)

	var encryptWithSalt, salt = hashUsingSalt(sha, text)
	fmt.Println(encryptWithSalt, salt)
}
