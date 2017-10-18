package hpp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHPPNew(t *testing.T) {
	hpp := New("mysecret")

	assert.Equal(t, hpp, HPP{Secret: "mysecret"}, "builds a new HPP")
}

func TestHPPRequest(t *testing.T) {
	hpp := New("mysecret")
	req := hpp.Request()

	assert.Equal(t, req, Request{hpp: &hpp}, "builds a new HPP request")
}

func TestHPPResponse(t *testing.T) {
	hpp := New("mysecret")
	resp := hpp.Response()

	assert.Equal(t, resp, Response{hpp: &hpp}, "builds a new HPP response")
}

func TestGenerateHash(t *testing.T) {
	hash := GenerateHash("test", "secret")

	assert.Equal(
		t,
		hash,
		"c6f07ec4e93a4fbd1a0ef1be168dabf7c2106106",
		"generated hash matches expected",
	)
}
