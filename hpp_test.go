package hpp

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHPPNew(t *testing.T) {
	hpp := New("mysecret")

	assert.Equal(t, hpp, HPP{Secret: "mysecret"}, "builds a new HPP")
}

func TestToJSON(t *testing.T) {
	hpp := New("mysecret")
	timestamp := JSONTime(time.Date(2099, 1, 1, 12, 0, 0, 0, time.UTC))

	req := Request{
		hpp:               &hpp,
		Account:           "myAccount",
		Currency:          "EUR",
		TimeStamp:         &timestamp,
		MerchantID:        "MerchantID",
		OrderID:           "OrderID",
		Amount:            100,
		CommentOne:        `a-z A-Z 0-9 ' ", + “” ._ - & \ / @ ! ? % ( )* : £ $ & € # [ ] | = ;ÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏÐÑÒÓÔÕÖ×ØÙÚÛÜÝÞßàáâãäåæçèéêëìíîïðñòóôõö÷ø¤ùúûüýþÿŒŽšœžŸ¥`,
		CommentTwo:        `Comment Two`,
		ReturnTSS:         false,
		ShippingCode:      "56|987",
		ShippingCountry:   "IRELAND",
		BillingCode:       "123|56",
		BillingCountry:    "IRELAND",
		CustomerNumber:    "123456",
		VariableReference: "VariableRef",
		ProductID:         "ProductID",
		Language:          "EN",
		CardPaymentButton: "Submit Payment",
		AutoSettleFlag:    "1",
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

			readSampleRequest("encoded-valid"),
			nil,
		},
		{
			"Given an invalid request",
			Request{hpp: &hpp, Amount: 100, MerchantID: "test", OrderID: "test%"},

			nil,
			fmt.Errorf("failed to validate HPP request: ORDER_ID: Order ID must only contain alphanumeric characters, dash and underscore."),
		},
	}

	for _, test := range tests {
		// Subject
		r := test.request
		j, err := hpp.ToJSON(r)

		// Assertions
		if err != nil {
			if assert.NotNil(t, test.err, test.description) {
				assert.EqualError(t, err, test.err.Error(), test.description)
			}
		} else {
			assert.Nil(t, err, test.description)
			assert.JSONEq(t, string(test.json), string(j), test.description)
		}
	}
}

func TestFromJSON(t *testing.T) {
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
			"Given valid response data",
			readSampleResponse("encoded-valid"),

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
				SupplementaryData: map[string]interface{}{
					"UNKNOWN_1": "Unknown value 1",
					"UNKNOWN_2": "Unknown value 2",
					"UNKNOWN_3": "Unknown value 3",
					"UNKNOWN_4": "Unknown value 4",
				},
				TSS: map[string]string{
					"TSS_2": "TSS_2_VALUE",
					"TSS_1": "TSS_1_VALUE",
				},
			},
			nil,
		},
		{
			"Given invalid json",
			[]byte(`invalid`),

			Response{},
			fmt.Errorf("unable to build response from json: unable to unmarshal response from json: failed to unmarshal HPP response json: invalid character 'i' looking for beginning of value"),
		},
		{
			"Given valid but incomplete json",
			[]byte(`{"ACCOUNT": "test", "SHA1HASH": "VEVTVA=="}`),

			Response{},
			fmt.Errorf("unable to build response from json: secret does not match expected: expected hash aca0089a38f647d3dae1c1fae9fa0a1c642151f0 received TEST"),
		},
	}

	for _, test := range tests {
		// Subject
		resp, err := hpp.FromJSON(test.json)

		// Assertions
		if err != nil {
			if assert.NotNil(t, test.err, test.description) {
				assert.EqualError(t, err, test.err.Error(), test.description)
			}
		} else {
			assert.Nil(t, err, test.description)
			assert.Equal(t, &test.response, resp, test.description)
		}
	}
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
