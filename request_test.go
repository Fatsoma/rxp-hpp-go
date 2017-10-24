package hpp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRequestMarshalJSON(t *testing.T) {
	timestamp := JSONTime(time.Date(2099, 1, 1, 12, 0, 0, 0, time.UTC))
	hpp := New("mysecret")

	var tests = []struct {
		//given
		description string
		request     Request

		//expected
		json json.RawMessage
		err  error
	}{
		{
			"Given the request can be marshalled",
			Request{
				hpp:               &hpp,
				Account:           "myAccount",
				Amount:            100,
				AutoSettleFlag:    "1",
				BillingCountry:    "IRELAND",
				BillingCode:       "123|56",
				CardPaymentButton: "Submit Payment",
				CommentOne:        "a-z A-Z 0-9 ' \", + “” ._ - & \\ / @ ! ? % ( )* : £ $ & € # [ ] | = ;ÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏÐÑÒÓÔÕÖ×ØÙÚÛÜÝÞßàáâãäåæçèéêëìíîïðñòóôõö÷ø¤ùúûüýþÿŒŽšœžŸ¥",
				CommentTwo:        "Comment Two",
				Currency:          "EUR",
				CustomerNumber:    "123456",
				Language:          "EN",
				MerchantID:        "MerchantID",
				PayerReference:    "PayerRef",
				PaymentReference:  "PaymentRef",
				OrderID:           "OrderID",
				Hash:              "5d8f05abd618e50db4861a61cc940112786474cf",
				ShippingCountry:   "IRELAND",
				ShippingCode:      "56|987",
				TimeStamp:         &timestamp,
				ProductID:         "ProductID",
				VariableReference: "VariableRef",
				PayerExists:       "0",
				SupplementaryData: map[string]interface{}{
					"UNKNOWN_1": "Unknown value 1",
					"UNKNOWN_2": "Unknown value 2",
					"UNKNOWN_3": "Unknown value 3",
					"UNKNOWN_4": "Unknown value 4",
				},
			},

			readSampleRequest("unknown-data"),
			nil,
		},
		{
			"Given the request can be marshalled",
			Request{hpp: &hpp, SupplementaryData: map[string]interface{}{"test": func() {}}},

			nil,
			fmt.Errorf("json: error calling MarshalJSON for type *hpp.Request: json: unsupported type: func()"),
		},
	}

	for _, test := range tests {
		// Subject
		js, err := json.Marshal(&test.request)

		// Assertions
		if err != nil {
			if assert.NotNil(t, test.err, test.description) {
				assert.EqualError(t, err, test.err.Error(), test.description)
			}
		} else {
			assert.Nil(t, err, test.description)
			assert.JSONEq(t, string(test.json), string(js), test.description)
		}
	}
}

func TestGenerateDefaults(t *testing.T) {
	req := Request{}
	req.GenerateDefaults()

	assert.NotNil(t, req.TimeStamp)

	assert.NotNil(t, req.OrderID)
	assert.Len(t, req.OrderID, 36)
}

func TestValidate(t *testing.T) {
	var tests = []struct {
		//given
		description string
		request     Request

		//expected
		err error
	}{
		{
			"Given the required attributes are missing",
			Request{},

			fmt.Errorf("AMOUNT: is required; MERCHANT_ID: is required"),
		},
		{
			"Given the request attributes are too long",
			Request{
				Amount:            1,
				MerchantID:        randomString(51),
				Account:           randomString(31),
				OrderID:           randomString(51),
				Hash:              randomString(41),
				CommentOne:        randomString(256),
				CommentTwo:        randomString(256),
				ShippingCode:      randomString(31),
				ShippingCountry:   randomString(51),
				BillingCode:       randomString(61),
				BillingCountry:    randomString(51),
				CustomerNumber:    randomString(51),
				VariableReference: randomString(51),
				ProductID:         randomString(51),
				CardPaymentButton: randomString(26),
				PayerReference:    randomString(51),
				PaymentReference:  randomString(51),
				PayerExists:       randomString(2),
			},

			fmt.Errorf(
				"ACCOUNT: %s; BILLING_CO: %s; BILLING_CODE: %s; CARD_PAYMENT_BUTTON: %s; "+
					"COMMENT1: %s; COMMENT2: %s; CUST_NUM: %s; MERCHANT_ID: %s; ORDER_ID: %s; "+
					"PAYER_EXIST: %s; PAYER_REF: %s; PMT_REF: %s; PROD_ID: %s; SHA1HASH: %s; "+
					"SHIPPING_CO: %s; SHIPPING_CODE: %s; VAR_REF: %s",
				accountSize,
				billingCountrySize,
				billingCodeSize,
				cardPaymentButtonTextSize,
				commentSize,
				commentSize,
				customerNumberSize,
				merchantIDSize,
				orderIDSize,
				payerExistsSize,
				payerReferenceSize,
				paymentReferenceSize,
				productIDSize,
				hashSize,
				shippingCountrySize,
				shippingCodeSize,
				variableReferenceSize,
			),
		},
		{
			"Given attributes that do not match their regexp patterns",
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
		r := test.request
		r.BuildHash("mysecret")

		// Assertions
		assert.Equal(t, r.Hash, test.hash, test.description)
	}
}

func TestMarshalJSONEncoded(t *testing.T) {
	type TestStruct struct {
		Test bool `json:"TEST"`
	}

	testReq := testRequest(true, true, true)

	var tests = []struct {
		//given
		description string
		request     interface{}
		encoded     bool

		//expected
		err error
	}{
		{
			"Given valid request",
			&testReq,
			true,

			nil,
		},
		{
			"Given a structure that cannot be marshalled",
			func() {},
			true,

			fmt.Errorf("failed to marshal HPP request: json: unsupported type: func()"),
		},
		{
			"Given a type that cannot be encoded",
			TestStruct{},
			true,

			fmt.Errorf("failed to unmarshal HPP request json: json: cannot unmarshal bool into Go value of type string"),
		},
		{
			"Given valid request which should not be encoded",
			&testReq,
			false,

			nil,
		},
	}

	for _, test := range tests {
		_, err := MarshalJSONEncoded(test.request, test.encoded)

		if err != nil && assert.NotNil(t, test.err, test.description) {
			assert.EqualError(t, err, test.err.Error())
		} else {
			assert.Nil(t, test.err)
		}
	}
}

func testRequest(cardStorage, selectStoredCard, fraudFilterMode bool) Request {
	hpp := New("mysecret")
	timestamp := time.Date(2013, 8, 14, 12, 22, 39, 0, time.UTC)
	t := JSONTime(timestamp)

	r := Request{
		hpp:        &hpp,
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

func readSampleRequest(file string) []byte {
	js, err := ioutil.ReadFile("./sample-json/hpp-request-" + file + ".json")
	if err != nil {
		log.Fatal("Could not read " + file)
	}

	return js
}
