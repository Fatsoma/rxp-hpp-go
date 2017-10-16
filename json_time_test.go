package hpp

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestJSONTimeString(t *testing.T) {
	timestamp := time.Date(2013, 8, 14, 12, 22, 39, 0, time.UTC)
	jt := JSONTime(timestamp)
	jsonTime := jt.String()

	assert.Equal(t, "20130814122239", jsonTime, "json should match")
}

func TestJSONTimeMarshalJSON(t *testing.T) {
	timestamp := time.Date(2013, 8, 14, 12, 22, 39, 0, time.UTC)
	jt := JSONTime(timestamp)
	json, err := jt.MarshalJSON()

	assert.Equal(t, []byte("20130814122239"), json, "json should match")
	assert.Nil(t, err, "no error is presnet")
}

func TestJSONTimeUnmarshalJSON(t *testing.T) {
	jsonTime := JSONTime{}

	var tests = []struct {
		//given
		description string
		jt          *JSONTime
		json        []byte

		//expected
		err error
	}{
		{
			"Given a valid json time",
			&jsonTime,
			[]byte("20130814122239"),

			nil,
		},
		{
			"Given a null pointer",
			nil,
			nil,

			fmt.Errorf("json.RawMessage: UnmarshalJSON on nil pointer"),
		},
		{
			"Given an invalid time",
			&jsonTime,
			[]byte("test"),

			fmt.Errorf("parsing time \"test\" as \"20060102150405\": cannot parse \"test\" as \"2006\""),
		},
	}

	for _, test := range tests {
		// Subject
		err := test.jt.UnmarshalJSON(test.json)

		// Assertions
		if err != nil && assert.NotNil(t, test.err, test.description) {
			assert.EqualError(t, test.err, err.Error(), test.description)
		} else {
			assert.Nil(t, test.err, test.description)
		}
	}
}
