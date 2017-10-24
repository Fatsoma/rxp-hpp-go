package hpp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestResponseUnmarshalJSON(t *testing.T) {
	hpp := New("mysecret")
	timestamp := JSONTime(time.Date(2013, 8, 14, 12, 22, 39, 0, time.UTC))

	var tests = []struct {
		//given
		description string
		json        json.RawMessage

		//expected
		response Response
		err      error
	}{
		{
			"Given the data can be unmarshalled into a response",
			readSampleResponse("unknown-data"),

			Response{
				hpp:        &hpp,
				MerchantID: "thestore",
				Account:    "myAccount",
				OrderID:    "ORD453-11",
				Amount:     100,
				AuthCode:   "79347",
				TimeStamp:  &timestamp,
				Hash:       "f093a0b233daa15f2bf44888f4fe75cb652e7bf0",
				Result:     "00",
				Message:    "Successful",
				CvnResult:  "1",
				PasRef:     "3737468273643",
				BatchID:    "654321",
				ECI:        "1",
				CAVV:       "123",
				XID:        "654564564",
				CommentOne: `a-z A-Z 0-9 ' ", + “” ._ - & \ / @ ! ? % ( )* : £ $ & € # [ ] | = ;ÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏÐÑÒÓÔÕÖ×ØÙÚÛÜÝÞßàáâãäåæçèéêëìíîïðñòóôõö÷ø¤ùúûüýþÿŒŽšœžŸ¥`,
				CommentTwo: "Comment Two",
				TSS: map[string]string{
					"TSS_1": "TSS_1_VALUE",
					"TSS_2": "TSS_2_VALUE",
				},
				SupplementaryData: map[string]interface{}{
					"UNKNOWN_1": "Unknown value 1",
					"UNKNOWN_2": "Unknown value 2",
					"UNKNOWN_3": "Unknown value 3",
					"UNKNOWN_4": "Unknown value 4",
				},
			},
			nil,
		},
		{
			"Given the data cannot be unmarshalled into a response",
			[]byte(`{"AMOUNT": "test"}`),

			Response{},
			fmt.Errorf("unable to unmarshal response: json: invalid use of ,string struct tag, trying to unmarshal \"test\" into int"),
		},
	}

	for _, test := range tests {
		// Subject
		r := Response{hpp: &hpp}
		err := json.Unmarshal(test.json, &r)

		// Assertions
		if err != nil {
			if assert.NotNil(t, test.err, test.description) {
				assert.EqualError(t, err, test.err.Error(), test.description)
			}
		} else {
			assert.Nil(t, err, test.description)
			assert.Equal(t, test.response, r, test.description)
		}
	}
}

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

func TestUnmarshalJSONEncoded(t *testing.T) {
	hpp := New("mysecret")

	type TestStruct struct {
		Test bool `json:"TEST"`
	}

	resp := Response{hpp: &hpp}

	var tests = []struct {
		//given
		description string
		response    interface{}
		data        []byte

		//expected
		err error
	}{
		{
			"Given valid response",
			&resp,
			[]byte(`{}`),

			nil,
		},
		{
			"Given a structure that cannot be unmarshalled",
			"test",
			[]byte(`{"MERCHANT_ID": "TEST@"}`),

			fmt.Errorf("failed to decode string from json response: illegal base64 data at input byte 4"),
		},
		{
			"Given a nested structure that cannot be unmarshalled",
			"test",
			[]byte(`{"TSS": {"TEST": "TEST@"}}`),

			fmt.Errorf("failed to decode map from json response: cannot decode val: illegal base64 data at input byte 4"),
		},
	}

	for _, test := range tests {
		err := UnmarshalJSONEncoded(test.response, test.data)

		if err != nil && assert.NotNil(t, test.err, test.description) {
			assert.EqualError(t, err, test.err.Error())
		} else {
			assert.Nil(t, test.err)
		}
	}
}

func testResponse() Response {
	t := JSONTime(time.Date(2013, 8, 14, 12, 22, 39, 0, time.UTC))

	return Response{TimeStamp: &t}
}

func readSampleResponse(file string) []byte {
	js, err := ioutil.ReadFile("./sample-json/hpp-response-" + file + ".json")
	if err != nil {
		log.Fatal("Could not read " + file)
	}

	return js
}
