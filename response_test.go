package hpp

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestResponseBuildHash(t *testing.T) {
	var tests = []struct {
		//given
		description string
		response    Response

		//expected
		hash string
	}{
		{
			"Given blank HPP Response, a valid hash is returned",
			Response{},

			"aca0089a38f647d3dae1c1fae9fa0a1c642151f0",
		},
		{
			"Given basic details the hash is built correctly",
			testResponse(),

			"43f6065bede40f3e0d7d732352b832c0136189e4",
		},
	}

	for _, test := range tests {
		// Subject
		hash := test.response.BuildHash("mysecret")

		// Assertions
		assert.Equal(t, hash, test.hash, test.description)
	}
}

func testResponse() Response {
	t := JSONTime(time.Date(2013, 8, 14, 12, 22, 39, 0, time.UTC))

	return Response{TimeStamp: &t}
}
