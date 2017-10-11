package hpp

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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
			_ = assert.NotNil(t, test.err, test.description) &&
				assert.Equal(t, test.err, err, test.description)
		} else {
			assert.Nil(t, err, test.description)
			assert.Equal(t, test.expected, bool(cb), test.description)
		}
	}
}

func TestBuildHash(t *testing.T) {
	timestamp := time.Date(2013, 8, 14, 12, 22, 39, 0, time.UTC)
	merchantID := "thestore"
	orderID := "ORD453-11"
	amount := 29900
	currency := "EUR"
	payerRef := "newpayer1"
	paymentRef := "mycard1"
	fraudFilterMode := "ACTIVE"
	selectStoredCard := "2b8de093-0241-4985-ad96-76ca0b26b478"

	var tests = []struct {
		//given
		description string
		hpp         HPP

		//expected
		hash string
	}{
		{
			"Given blank HPP, a valid hash is returned",
			HPP{},

			"5ece5764864e9afac4cd0c9560055f7598e3af42",
		},
		{
			"Given basic details the hash is built correctly",
			HPP{
				TimeStamp:  &timestamp,
				MerchantID: merchantID,
				OrderID:    orderID,
				Amount:     amount,
				Currency:   currency,
			},

			"cc72c08e529b3bc153481eda9533b815cef29de3",
		},
		{
			"Given the enable card storage flag, a hash is returned with the payer details",
			HPP{
				EnableCardStorage: true,
				TimeStamp:         &timestamp,
				MerchantID:        merchantID,
				OrderID:           orderID,
				Amount:            amount,
				Currency:          currency,
				PayerReference:    payerRef,
				PaymentReference:  paymentRef,
			},

			"4106afc4666c6145b623089b1ad4098846badba2",
		},
		{
			"Given the select stored card, a hash is returned with the payer details",
			HPP{
				SelectStoredCard: selectStoredCard,
				TimeStamp:        &timestamp,
				MerchantID:       merchantID,
				OrderID:          orderID,
				Amount:           amount,
				Currency:         currency,
				PayerReference:   payerRef,
				PaymentReference: paymentRef,
			},

			"4106afc4666c6145b623089b1ad4098846badba2",
		},
		{
			"Given the fraud filter mode flag, the fraud filter mode is included in the hash",
			HPP{
				TimeStamp:       &timestamp,
				MerchantID:      merchantID,
				OrderID:         orderID,
				Amount:          amount,
				Currency:        currency,
				FraudFilterMode: fraudFilterMode,
			},

			"b7b3cbb60129a1c169a066afa09ce7cc843ff1c1",
		},
		{
			"Given the fraud filter mode flag, and stored card",
			HPP{
				EnableCardStorage: true,
				TimeStamp:         &timestamp,
				MerchantID:        merchantID,
				OrderID:           orderID,
				Amount:            amount,
				Currency:          currency,
				PayerReference:    payerRef,
				PaymentReference:  paymentRef,
				FraudFilterMode:   fraudFilterMode,
			},

			"39f637a321da4ebc3a433ed327b2c2921ad58fdb",
		},
	}

	for _, test := range tests {
		// Subject
		hash := test.hpp.BuildHash("mysecret")

		// Assertions
		assert.Equal(t, hash, test.hash, test.description)
	}
}
