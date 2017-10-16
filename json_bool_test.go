package hpp

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONBoolMarshalJSON(t *testing.T) {
	var tests = []struct {
		//given
		description string
		cb          JSONBool

		//expected
		expected string
	}{
		{
			"Given a true bool it marshals as 1",
			true,

			"\"1\"",
		},
		{
			"Given a false bool it marshals as 0",
			false,

			"\"0\"",
		},
	}

	for _, test := range tests {
		// Subject
		res, err := test.cb.MarshalJSON()

		// Assertions
		assert.Nil(t, err, test.description)
		assert.Equal(t, test.expected, string(res), test.description)
	}
}

func TestJSONBoolUnmarshalJSON(t *testing.T) {
	var tests = []struct {
		//given
		description string
		data        []byte

		//expected
		expected bool
		err      error
	}{
		{
			"Given \"1\" it unmarshals to true",
			[]byte("1"),

			true,
			nil,
		},
		{
			"Given \"0\" it unmarshals to be false",
			[]byte("0"),

			false,
			nil,
		},
		{
			"Given \"true\" it unmarshals to true",
			[]byte("1"),

			true,
			nil,
		},
		{
			"Given \"false\" it unmarshals to be false",
			[]byte("0"),

			false,
			nil,
		},
		{
			"Given \"unknown\" it returns an error",
			[]byte("unknown"),

			false,
			errors.New("Boolean unmarshal error: invalid input unknown"),
		},
	}

	for _, test := range tests {
		// Subject
		var cb JSONBool
		err := cb.UnmarshalJSON(test.data)

		// Assertions
		if err != nil {
			if assert.NotNil(t, test.err, test.description) {
				assert.Equal(t, test.err, err, test.description)
			}
		} else {
			assert.Nil(t, err, test.description)
			assert.Equal(t, test.expected, bool(cb), test.description)
		}
	}
}
