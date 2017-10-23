package hpp

import (
	"crypto/sha1"
	"fmt"
	"strings"
)

// Separator used in generating sha1 hashes
const Separator = "."

// HPP container that we pass to requests / responses
type HPP struct {
	Secret string
}

// New builds a new HPP
func New(s string) HPP {
	return HPP{Secret: s}
}

// Request initialise a new request with this HPP set
func (hpp *HPP) Request() Request {
	return Request{hpp: hpp}
}

// Response initialise a new request with this HPP set
func (hpp *HPP) Response() Response {
	return Response{hpp: hpp}
}

// GenerateHash ...
// Each message sent to Realex should have a hash, attached. For a message using the remote
// interface this is generated using the This is generated from the TIMESTAMP, MERCHANT_ID,
// ORDER_ID, AMOUNT, and CURRENCY fields concatenated together with "." in between each field.
// This confirms the message comes
// from the client and
// Generate a hash, required for all messages sent to IPS to prove it was not tampered with.
// <p>
// Hashing takes a string as input, and produce a fixed size number (160 bits for SHA-1 which
// this implementation uses). This number is a hash of the input, and a small change in the
// input results in a substantial change in the output. The functions are thought to be secure
// in the sense that it requires an enormous amount of computing power and time to find a string
// that hashes to the same value. In others words, there's no way to decrypt a secure hash.
// Given the larger key size, this implementation uses SHA-1 which we prefer that you, but Realex
// has retained compatibilty with MD5 hashing for compatibility with older systems.
// <p>
// <p>
// To construct the hash for the remote interface follow this procedure:
//
// To construct the hash for the remote interface follow this procedure:
// Form a string by concatenating the above fields with a period ('.') in the following order
// <p>
// (TIMESTAMP.MERCHANT_ID.ORDER_ID.AMOUNT.CURRENCY)
// <p>
// Like so (where a field is empty an empty string "" is used):
// <p>
// (20120926112654.thestore.ORD453-11.29900.EUR)
// <p>
// Get the hash of this string (SHA-1 shown below).
// <p>
// (b3d51ca21db725f9c7f13f8aca9e0e2ec2f32502)
// <p>
// Create a new string by concatenating this string and your shared secret using a period.
// <p>
// (b3d51ca21db725f9c7f13f8aca9e0e2ec2f32502.mysecret )
// <p>
// Get the hash of this value. This is the value that you send to Realex Payments.
// <p>
// (3c3cac74f2b783598b99af6e43246529346d95d1)
//
// This method takes the pre-built string of concatenated fields and the secret and returns the
// SHA-1 hash to be placed in the request sent to Realex.
func GenerateHash(str, secret string) string {
	firstHash := sha1.Sum([]byte(str))
	firstHashStr := fmt.Sprintf("%x", firstHash)

	//second pass takes the first hash, adds the secret and hashes again
	firstWithSecret := strings.Join([]string{firstHashStr, secret}, Separator)
	secondHash := sha1.Sum([]byte(firstWithSecret))

	return fmt.Sprintf("%x", secondHash)
}
