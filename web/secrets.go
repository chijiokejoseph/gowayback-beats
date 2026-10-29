package web

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"log"
)

// NewCode creates a randomized cryptographically secure
// byte string
func NewCode(size ...int) string {
	num := 0
	if len(size) > 0 {
		// if more than one argument is passed, set num to the first argument
		num = size[0]
	} else {
		// num should default to 43
		num = 43
	}

	// make random bytes of length `num`
	byteArr := make([]byte, num)
	// fill `byteArr` with cryptographically secure bytes
	_, err := rand.Read(byteArr)
	if err != nil {
		log.Fatalf("Randomization failed with error %s\n", err)
	}
	// return bytes as a url encoded string
	return base64.RawURLEncoding.EncodeToString(byteArr)
}

// NewSecret creates a url encoded password or secret string
// from a verification or source string.
func NewSecret(verifier string) string {
	// create byte array from sha256-Sum256 algorithm from the verification string i.e., `verifier`
	byteArr := sha256.Sum256([]byte(verifier))
	// url encode the resulting byte array to a string
	secret := base64.RawURLEncoding.EncodeToString(byteArr[:])
	return secret
}
