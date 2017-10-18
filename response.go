package hpp

import "strings"

type Response struct {
	hpp *HPP

	// This is the merchant id that Realex Payments assign to you.
	MerchantID JSONString `json:"MERCHANT_ID"`

	// The sub-account used in the transaction.
	Account JSONString `json:"ACCOUNT"`

	// The unique order id that you sent to us.
	OrderID JSONString `json:"ORDER_ID"`

	// The amount that was authorised. Returned in the lowest unit of the currency.
	Amount JSONInt `json:"AMOUNT,string"`

	// Will contain a valid authcode if the transaction was successful. Will be empty otherwise.
	AuthCode JSONString `json:"AUTHCODE"`

	// The date and time of the transaction.
	TimeStamp *JSONTime `json:"TIMESTAMP"`

	// A SHA-1 digital signature created using the HPP response fields and your shared secret.
	Hash JSONString `json:"SHA1HASH"`

	// The outcome of the transaction. Will contain "00" if the transaction was a success or another value (depending on the error) if not.
	Result JSONString `json:"RESULT"`

	// Will contain a text message that describes the result code.
	Message JSONString `json:"MESSAGE"`

	// The result of the Card Verification check (if enabled):
	//
	// M: CVV Matched.
	// N: CVV Not Matched.
	// I: CVV Not checked due to circumstances.
	// U: CVV Not checked - issuer not certified.
	// P: CVV Not Processed.
	CvnResult JSONString `json:"CVNRESULT"`

	// A unique reference that Realex Payments assign to your transaction.
	PasRef JSONString `json:"PASREF"`

	// This is the Realex Payments batch that this transaction will be in.
	// (This is equal to "-1" if the transaction was sent in with the autosettle flag off.
	// After you settle it (either manually or programmatically) the response to that transaction will contain the batch id.)
	BatchID JSONString `json:"BATCHID"`

	// This is the ecommerce indicator (this will only be returned for 3DSecure transactions).
	ECI JSONString `json:"ECI"`

	// Cardholder Authentication Verification Value (this will only be returned for 3DSecure transactions).
	CAVV JSONString `json:"CAVV"`

	// Exchange Identifier (this will only be returned for 3DSecure transactions).
	XID JSONString `json:"XID"`

	// Whatever data you have sent in the request will be returned to you.
	CommentOne JSONString `json:"COMMENT1"`

	// Whatever data you have sent in the request will be returned to you.
	CommentTwo JSONString `json:"COMMENT2"`

	// The Transaction Suitability Score for the transaction. The RealScore is comprised of various distinct tests.
	// Using the RealControl application you can request that Realex Payments return certain individual scores to you.
	// These are identified by numbers - thus TSS_1032 would be the result of the check with id 1032.
	// You can then use these specific checks in conjunction with RealScore score to ascertain whether or not you wish to continue with the settlement.
	TSS map[string]string `json:"TSS"`

	// Anything else you sent to us in the request will be returned to you in supplementary data.
	SupplementaryData map[string]string `json:"-"`
}

// BuildHash creates the security hash from a number of fields and the shared secret.
func (r *Response) BuildHash(secret string) string {
	ts := ""
	if r.TimeStamp != nil {
		ts = r.TimeStamp.String()
	}

	s := []string{
		ts,
		r.MerchantID.String(),
		r.OrderID.String(),
		r.Result.String(),
		r.Message.String(),
		r.PasRef.String(),
		r.AuthCode.String(),
	}

	return GenerateHash(strings.Join(s, Separator), secret)
}
