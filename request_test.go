package hpp

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRequestToJSON(t *testing.T) {
	hpp := New("mysecret")
	timestamp := time.Date(2099, 1, 1, 12, 0, 0, 0, time.UTC)
	jsonTime := JSONTime(timestamp)

	req := Request{
		hpp:               &hpp,
		Account:           "myAccount",
		Currency:          "EUR",
		TimeStamp:         &jsonTime,
		MerchantID:        "MerchantID",
		OrderID:           []byte("OrderID"),
		Amount:            100,
		CommentOne:        `a-z A-Z 0-9 ' ", + “” ._ - & \ / @ ! ? % ( )* : £ $ & € # [ ] | = ;ÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏÐÑÒÓÔÕÖ×ØÙÚÛÜÝÞßàáâãäåæçèéêëìíîïðñòóôõö÷ø¤ùúûüýþÿŒŽšœžŸ¥`,
		CommentTwo:        `Comment Two`,
		ReturnTSS:         "0",
		ShippingCode:      "56|987",
		ShippingCountry:   "IRELAND",
		BillingCode:       "123|56",
		BillingCountry:    "IRELAND",
		CustomerNumber:    "123456",
		VariableReference: "VariableRef",
		ProductID:         "ProductID",
		Language:          "EN",
		CardPaymentButton: "Submit Payment",
		AutoSettleFlag:    true,
		EnableCardStorage: false,
		OfferSaveCard:     false,
		PayerReference:    "PayerRef",
		PaymentReference:  "PaymentRef",
		PayerExists:       "0",
		ValidCardOnly:     false,
		DCCEnable:         false,
	}

	var tests = []struct {
		//given
		description string
		request     Request

		//expected
		json json.RawMessage
		err  error
	}{
		{
			"Given a valid request",
			req,

			[]byte(`{"MERCHANT_ID":"TWVyY2hhbnRJRA==","ACCOUNT":"bXlBY2NvdW50","ORDER_ID":"T3JkZXJJRA==","AMOUNT":"MTAw","CURRENCY":"RVVS","TIMESTAMP":"MjA5OTAxMDExMjAwMDA=","SHA1HASH":"YWFmYWIxMmRkNGY5MmUwZDZlMTNkZDhjM2ZjZTkzMjMyYWVkZjI4YQ==","AUTO_SETTLE_FLAG":"MQ==","COMMENT1":"YS16IEEtWiAwLTkgJyAiLCArIOKAnOKAnSAuXyAtICYgXCAvIEAgISA/ICUgKCApKiA6IMKjICQgJiDigqwgIyBbIF0gfCA9IDvDgMOBw4LDg8OEw4XDhsOHw4jDicOKw4vDjMONw47Dj8OQw5HDksOTw5TDlcOWw5fDmMOZw5rDm8Ocw53DnsOfw6DDocOiw6PDpMOlw6bDp8Oow6nDqsOrw6zDrcOuw6/DsMOxw7LDs8O0w7XDtsO3w7jCpMO5w7rDu8O8w73DvsO/xZLFvcWhxZPFvsW4wqU=","COMMENT2":"Q29tbWVudCBUd28=","RETURN_TSS":"MA==","SHIPPING_CODE":"NTZ8OTg3","SHIPPING_CO":"SVJFTEFORA==","BILLING_CODE":"MTIzfDU2","BILLING_CO":"SVJFTEFORA==","CUST_NUM":"MTIzNDU2","VAR_REF":"VmFyaWFibGVSZWY=","PROD_ID":"UHJvZHVjdElE","HPP_LANG":"RU4=","CARD_PAYMENT_BUTTON":"U3VibWl0IFBheW1lbnQ=","CARD_STORAGE_ENABLE":"MA==","OFFER_SAVE_CARD":"MA==","PAYER_REF":"UGF5ZXJSZWY=","PMT_REF":"UGF5bWVudFJlZg==","PAYER_EXIST":"MA==","VALIDATE_CARD_ONLY":"MA==","DCC_ENABLE":"MA=="}`),
			nil,
		},
	}

	for _, test := range tests {
		// Subject
		r := test.request
		j, err := r.ToJSON()

		// Assertions
		if err != nil {
			if assert.NotNil(t, test.err, test.description) {
				assert.EqualError(t, err, test.err.Error(), test.description)
			}
		} else {
			assert.Nil(t, err, test.description)
			assert.Equal(t, string(test.json), string(j), test.description)
		}
	}
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
			"Given the merchant ID is missing",
			Request{MerchantID: ""},

			fmt.Errorf("MERCHANT_ID: is required"),
		},
		{
			"Given the merchant ID is too long",
			Request{MerchantID: JSONString(randomString(51))},

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
		assert.Equal(t, r.Hash.String(), test.hash, test.description)
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
		OrderID:    []byte("ORD453-11"),
		Amount:     JSONInt(29900),
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
