package hpp

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	var tests = []struct {
		//given
		description string
		request     Request

		//expected
		err error
	}{
		{
			"Given the merchant ID is missing",
			Request{MerchantID: ""},

			fmt.Errorf("MERCHANT_ID: is required"),
		},
		{
			"Given the merchant ID is too long",
			Request{MerchantID: randomString(51)},

			fmt.Errorf("MERCHANT_ID: %s", merchantIDSize),
		},
		{
			"Given the merchant ID is incorrect",
			Request{MerchantID: "test%"},

			fmt.Errorf("MERCHANT_ID: %s", merchantIDPattern),
		},
	}

	for _, test := range tests {
		// Subject
		err := test.request.Validate()

		// Assertions
		if err != nil && assert.NotNil(t, test.err, test.description) {
			assert.Contains(t, err.Error(), test.err.Error(), test.description)
		} else {
			assert.Nil(t, test.err, test.description)
		}
	}
}

func TestConvertableBoolMarshalJSON(t *testing.T) {
	var tests = []struct {
		//given
		description string
		cb          ConvertibleBoolean

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

func TestConvertableBoolUnmarshalJSON(t *testing.T) {
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
		var cb ConvertibleBoolean
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

func TestBuildHash(t *testing.T) {
	var tests = []struct {
		//given
		description string
		request     Request

		//expected
		hash string
	}{
		{
			"Given blank HPP, a valid hash is returned",
			Request{},

			"5ece5764864e9afac4cd0c9560055f7598e3af42",
		},
		{
			"Given basic details the hash is built correctly",
			testRequest(false, false, false),

			"cc72c08e529b3bc153481eda9533b815cef29de3",
		},
		{
			"Given the enable card storage flag, a hash is returned with the payer details",
			testRequest(true, false, false),

			"4106afc4666c6145b623089b1ad4098846badba2",
		},
		{
			"Given the select stored card, a hash is returned with the payer details",
			testRequest(false, true, false),

			"4106afc4666c6145b623089b1ad4098846badba2",
		},
		{
			"Given the fraud filter mode flag, the fraud filter mode is included in the hash",
			testRequest(false, false, true),

			"b7b3cbb60129a1c169a066afa09ce7cc843ff1c1",
		},
		{
			"Given the fraud filter mode flag, and stored card flag",
			testRequest(true, false, true),

			"39f637a321da4ebc3a433ed327b2c2921ad58fdb",
		},
	}

	for _, test := range tests {
		// Subject
		hash := test.request.BuildHash("mysecret")

		// Assertions
		assert.Equal(t, hash, test.hash, test.description)
	}
}

func testRequest(cardStorage, selectStoredCard, fraudFilterMode bool) Request {
	timestamp := time.Date(2013, 8, 14, 12, 22, 39, 0, time.UTC)
	t := JSONTime(timestamp)

	r := Request{
		TimeStamp:  &t,
		MerchantID: "thestore",
		OrderID:    "ORD453-11",
		Amount:     29900,
		Currency:   "EUR",
	}

	if cardStorage {
		r.EnableCardStorage = true
	}

	if selectStoredCard {
		r.SelectStoredCard = "2b8de093-0241-4985-ad96-76ca0b26b478"
	}

	if cardStorage || selectStoredCard {
		r.PayerReference = "newpayer1"
		r.PaymentReference = "mycard1"
	}

	if fraudFilterMode {
		r.FraudFilterMode = "ACTIVE"
	}

	return r
}

func randomString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
