package hpp

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/pkg/errors"
)

// Response represents a response from HPP
type Response struct {
	hpp *HPP

	// This is the merchant id that Realex Payments assign to you.
	MerchantID string `json:"MERCHANT_ID"`

	// The sub-account used in the transaction.
	Account string `json:"ACCOUNT"`

	// The unique order id that you sent to us.
	OrderID string `json:"ORDER_ID"`

	// The amount that was authorised. Returned in the lowest unit of the currency.
	Amount int `json:"AMOUNT,string"`

	// Will contain a valid authcode if the transaction was successful. Will be empty otherwise.
	AuthCode string `json:"AUTHCODE"`

	// The date and time of the transaction.
	TimeStamp *JSONTime `json:"TIMESTAMP"`

	// A SHA-1 digital signature created using the HPP response fields and your shared secret.
	Hash string `json:"SHA1HASH"`

	// The outcome of the transaction. Will contain "00" if the transaction was a success or another value (depending on the error) if not.
	Result string `json:"RESULT"`

	// Will contain a text message that describes the result code.
	Message string `json:"MESSAGE"`

	// The result of the Card Verification check (if enabled):
	//
	// M: CVV Matched.
	// N: CVV Not Matched.
	// I: CVV Not checked due to circumstances.
	// U: CVV Not checked - issuer not certified.
	// P: CVV Not Processed.
	CvnResult string `json:"CVNRESULT"`

	// A unique reference that Realex Payments assign to your transaction.
	PasRef string `json:"PASREF"`

	// This is the Realex Payments batch that this transaction will be in.
	// (This is equal to "-1" if the transaction was sent in with the autosettle flag off.
	// After you settle it (either manually or programmatically) the response to that transaction will contain the batch id.)
	BatchID string `json:"BATCHID"`

	// This is the ecommerce indicator (this will only be returned for 3DSecure transactions).
	ECI string `json:"ECI"`

	// Cardholder Authentication Verification Value (this will only be returned for 3DSecure transactions).
	CAVV string `json:"CAVV"`

	// Exchange Identifier (this will only be returned for 3DSecure transactions).
	XID string `json:"XID"`

	// Whatever data you have sent in the request will be returned to you.
	CommentOne string `json:"COMMENT1"`

	// Whatever data you have sent in the request will be returned to you.
	CommentTwo string `json:"COMMENT2"`

	// The Transaction Suitability Score for the transaction. The RealScore is comprised of various distinct tests.
	// Using the RealControl application you can request that Realex Payments return certain individual scores to you.
	// These are identified by numbers - thus TSS_1032 would be the result of the check with id 1032.
	// You can then use these specific checks in conjunction with RealScore score to ascertain whether or not you wish to continue with the settlement.
	TSS map[string]string `json:"TSS"`

	// Anything else you sent to us in the request will be returned to you in supplementary data.
	SupplementaryData map[string]interface{} `json:"-"`
}

// FromJSON converts valid JSON into the Response
func (r *Response) FromJSON(data []byte, encoded bool) error {
	fmt.Println("Converting JSON to HppResponse.")

	if encoded {
		err := UnmarshalJSONEncoded(r, data)
		if err != nil {
			return errors.Wrap(err, "unable to unmarshal encoded response from json")
		}
	} else {
		err := json.Unmarshal(data, r)
		if err != nil {
			return errors.Wrap(err, "unable to unmarshal response from json")
		}
	}

	fmt.Println("Validating response hash.")
	err := r.ValidateHash(r.hpp.Secret)
	if err != nil {
		return errors.Wrap(err, "secret does not match expected")
	}

	return nil
}

// ValidateHash ensure the HPP response hash is what we expect it to be
func (r *Response) ValidateHash(secret string) error {
	expected := r.BuildHash(secret)
	if expected != r.Hash {
		return fmt.Errorf("expected hash %s received %s", expected, r.Hash)
	}

	return nil
}

// UnmarshalJSON override the standard JSON unmarshaller to include the supplementary data
func (r *Response) UnmarshalJSON(data []byte) error {
	type Alias Response
	ra := (*Alias)(r)
	err := json.Unmarshal(data, &ra)
	if err != nil {
		return errors.Wrap(err, "unable to unmarshal response")
	}

	// Add the supplementary data from the response
	extra := map[string]interface{}{}
	err = json.Unmarshal(data, &extra)
	if err != nil {
		return errors.Wrap(err, "unable to unmarshal response to map")
	}

	// delete any keys that are already in the response struct fields
	for _, k := range r.jsonfields() {
		delete(extra, k)
	}

	r.SupplementaryData = extra

	return nil
}

// BuildHash creates the security hash from a number of fields and the shared secret.
func (r *Response) BuildHash(secret string) string {
	ts := ""
	if r.TimeStamp != nil {
		ts = r.TimeStamp.String()
	}

	s := []string{ts, r.MerchantID, r.OrderID, r.Result, r.Message, r.PasRef, r.AuthCode}

	return GenerateHash(strings.Join(s, Separator), secret)
}

func (r Response) jsonfields() (names []string) {
	val := reflect.ValueOf(r)
	for i := 0; i < val.Type().NumField(); i++ {
		jt := val.Type().Field(i).Tag.Get("json")
		jt = strings.TrimRight(jt, ",string")
		names = append(names, jt)
	}
	return
}

// UnmarshalJSONEncoded Base64 decodes the values and unmarshals the response
func UnmarshalJSONEncoded(resp interface{}, data []byte) error {
	decoded := map[string]interface{}{}
	err := json.Unmarshal(data, &decoded)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal HPP response json")
	}

	for k, v := range decoded {
		switch val := v.(type) {
		case string:
			s, derr := base64.StdEncoding.DecodeString(val)
			if derr != nil {
				return errors.Wrap(derr, "failed to decode string from json response")
			}
			decoded[k] = string(s)
		case map[string]interface{}:
			m, derr := decodeNestedMap(val)
			if derr != nil {
				return errors.Wrap(derr, "failed to decode map from json response")
			}
			decoded[k] = m
		}
	}

	dd, err := json.Marshal(decoded)
	if err != nil {
		return errors.Wrap(err, "failed to marshal decoded response")
	}

	return json.Unmarshal(dd, resp)
}

func decodeNestedMap(val map[string]interface{}) (map[string]string, error) {
	res := map[string]string{}
	for k, v := range val {
		s, err := base64.StdEncoding.DecodeString((v).(string))
		if err != nil {
			return nil, errors.Wrap(err, "cannot decode val")
		}
		res[k] = string(s)
	}

	return res, nil
}
