package hpp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateHash(t *testing.T) {
	hash := GenerateHash("test", "secret")

	assert.Equal(
		t,
		hash,
		"c6f07ec4e93a4fbd1a0ef1be168dabf7c2106106",
		"generated hash matches expected",
	)
}
