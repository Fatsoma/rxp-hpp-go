package hpp

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

const (
	TimeLayout = "20060102150405"
	Separator  = "."
)

type HPP struct {
	// The merchant ID supplied by Realex Payments – note this is not the merchant number supplied by your bank.
	MerchantID string `json:"MERCHANT_ID" valid:"-"`

	// The sub-account to use for this transaction. If not present, the default sub-account will be used.
	Account string `json:"ACCOUNT" valid:"-"`

	// A unique alphanumeric id that’s used to identify the transaction. No spaces are allowed.
	OrderID string `json:"ORDER_ID" valid:"uuidv4"`

	// Total amount to authorise in the lowest unit of the currency – i.e. 100 euro would be entered as 10000.
	// If there is no decimal in the currency (e.g. JPY Yen) then contact Realex Payments. No decimal points are allowed.
	// Amount should be set to 0 for OTB transactions (i.e. where validate card only is set to 1).
	Amount int `json:"AMOUNT" valid:"int"`

	// A three-letter currency code (Eg. EUR, GBP). A list of currency codes can be provided by your account manager.
	Currency string `json:"CURRENCY" valid:"alpha"`

	// Date and time of the transaction. Entered in the following format: YYYYMMDDHHMMSS. Must be within 24 hours of the current time.
	TimeStamp *time.Time `json:"TIMESTAMP" valid:"-"`

	// A digital signature generated using the SHA-1 algorithm.
	Hash string `json:"SHA1HASH" valid:"-"`

	// Used to signify whether or not you wish the transaction to be captured in the next batch.
	// If set to "1" and assuming the transaction is authorised then it will automatically be settled in the next batch.
	// If set to "0" then the merchant must use the RealControl application to manually settle the transaction.
	// This option can be used if a merchant wishes to delay the payment until after the goods have been shipped.
	// Transactions can be settled for up to 115% of the original amount and must be settled within a certain period of time agreed with your issuing bank.
	AutoSettleFlag ConvertibleBoolean `json:"AUTO_SETTLE_FLAG" valid:"int"`

	// A freeform comment to describe the transaction.
	CommentOne string `json:"COMMENT1" valid:"-"`

	// A freeform comment to describe the transaction.
	CommentTwo string `json:"COMMENT2" valid:"-"`

	// Used to signify whether or not you want a Transaction Suitability Score for this transaction.
	// Can be "0" for no and "1" for yes.
	ReturnTSS string `json:"RETURN_TSS" valid:"-"`

	// The postcode or ZIP of the shipping address.
	ShippingCode string `json:"SHIPPING_CODE" valid:"-"`

	// The country of the shipping address.
	ShippingCountry string `json:"SHIPPING_CO" valid:"-"`

	// The postcode or ZIP of the billing address.
	BillingCode string `json:"BILLING_CODE" valid:"-"`

	// The country of the billing address.
	BillingCountry string `json:"BILLING_CO" valid:"-"`

	// The customer number of the customer. You can send in any additional information about the transaction in this field,
	// which will be visible under the transaction in the RealControl application.
	CustomerNumber string `json:"CUST_NUM" valid:"-"`

	// A variable reference also associated with this customer. You can send in any additional information about the transaction in this field,
	// which will be visible under the transaction in the RealControl application.
	VariableReference string `json:"VAR_REF" valid:"-"`

	// A product id associated with this product. You can send in any additional information about the transaction in this field,
	// which will be visible under the transaction in the RealControl application.
	ProductID string `json:"PROD_ID" valid:"-"`

	// Used to set what language HPP is displayed in. Currently HPP is available in English, Spanish and German, with other languages to follow.
	// If the field is not sent in, the default language is the language that is set in your account configuration. This can be set by your account manager.
	Language string `json:"HPP_LANG" valid:"-"`

	// Used to set what text is displayed on the payment button for card transactions. If this field is not sent in, "Pay Now" is displayed on the button by default.
	CardPaymentButton string `json:"CARD_PAYMENT_BUTTON" valid:"-"`

	// Enable card storage.
	EnableCardStorage ConvertibleBoolean `json:"CARD_STORAGE_ENABLE" valid:"-"`

	// Offer to save the card.
	OfferSaveCard ConvertibleBoolean `json:"OFFER_SAVE_CARD" valid:"-"`

	// The payer reference.
	PayerReference string `json:"PAYER_REF" valid:"-"`

	// The payment reference.
	PaymentReference string `json:"PMT_REF" valid:"-"`

	// Payer exists.
	PayerExists ConvertibleBoolean `json:"PAYER_EXIST" valid:"-"`

	// Used to identify an OTB transaction.
	ValidCardOnly ConvertibleBoolean `json:"VALIDATE_CARD_ONLY" valid:"-"`

	// Transaction level configuration to enable/disable a DCC request. (Only if the merchant is configured).
	DCCEnable ConvertibleBoolean `json:"DCC_ENABLE" valid:"-"`

	// Override merchant configuration for fraud. (Only if the merchant is configured for fraud).
	FraudFilterMode string `json:"HPP_FRAUDFILTER_MODE" valid:"-"`

	// The HPP Version. To use HPP Card Management select HPP_VERSION = 2.
	Version string `json:"HPP_VERSION" valid:"-"`

	// The payer reference. If this flag is received, HPP will retrieve a list of the payment methods saved for that payer.
	SelectStoredCard string `json:"HPP_SELECT_STORED_CARD" valid:"-"`
}

// ConvertibleBoolean is a boolean that represents "1" and "0" as true / false
type ConvertibleBoolean bool

// MarshalJSON converts bools to "1" / "0"
func (b *ConvertibleBoolean) MarshalJSON() ([]byte, error) {
	if *b {
		return json.Marshal("1")
	}

	return json.Marshal("0")
}

// UnmarshalJSON converts "1" / "0" to bool
func (b *ConvertibleBoolean) UnmarshalJSON(data []byte) error {
	s := string(data)
	if s == "1" || s == "true" {
		*b = true
	} else if s == "0" || s == "false" {
		*b = false
	} else {
		return fmt.Errorf("Boolean unmarshal error: invalid input %s", s)
	}
	return nil
}

// BuildHash - Each message sent to Realex should have a hash, attached. For a message using the remote
// interface this is generated using the This is generated from the TIMESTAMP, MERCHANT_ID,
// ORDER_ID, AMOUNT, and CURRENCY fields concatenated together with "." in between each field.
// This confirms the message comes
// from the client and
// Generate a hash, required for all messages sent to IPS to prove it was not tampered with.
// <p>
// Hashing takes a string as input, and produce a fixed size number (160 bits for SHA-1 which
// this implementation uses). This number is a hash of the input, and a small change in the
// input results in a substantial change in the output. The functions are thought to be secure
// in the sense that it requires an enormous amount of computing power and time to find a string
// that hashes to the same value. In others words, there's no way to decrypt a secure hash.
// Given the larger key size, this implementation uses SHA-1 which we prefer that you, but Realex
// has retained compatibilty with MD5 hashing for compatibility with older systems.
// <p>
// <p>
// To construct the hash for the remote interface follow this procedure:
//
// To construct the hash for the remote interface follow this procedure:
// Form a string by concatenating the above fields with a period ('.') in the following order
// <p>
// (TIMESTAMP.MERCHANT_ID.ORDER_ID.AMOUNT.CURRENCY)
// <p>
// Like so (where a field is empty an empty string "" is used):
// <p>
// (20120926112654.thestore.ORD453-11.29900.EUR)
// <p>
// Get the hash of this string (SHA-1 shown below).
// <p>
// (b3d51ca21db725f9c7f13f8aca9e0e2ec2f32502)
// <p>
// Create a new string by concatenating this string and your shared secret using a period.
// <p>
// (b3d51ca21db725f9c7f13f8aca9e0e2ec2f32502.mysecret )
// <p>
// Get the hash of this value. This is the value that you send to Realex Payments.
// <p>
// (3c3cac74f2b783598b99af6e43246529346d95d1)
//
// This method takes the pre-built string of concatenated fields and the secret and returns the
// SHA-1 hash to be placed in the request sent to Realex.
func (hpp *HPP) BuildHash(secret string) string {
	//first pass hashes the String of required fields
	s := hpp.buildHashString()

	firstHash := sha1.Sum([]byte(s))
	firstHashStr := fmt.Sprintf("%x", firstHash)

	//second pass takes the first hash, adds the secret and hashes again
	firstWithSecret := strings.Join([]string{firstHashStr, secret}, Separator)
	secondHash := sha1.Sum([]byte(firstWithSecret))

	return fmt.Sprintf("%x", secondHash)
}

func (hpp *HPP) buildHashString() string {
	a := hpp.basicHash()

	if hpp.canStoreCard() {
		a = append(a, []string{hpp.PayerReference, hpp.PaymentReference}...)
	}

	if hpp.FraudFilterMode != "" {
		a = append(a, hpp.FraudFilterMode)
	}

	return strings.Join(a, Separator)
}

func (hpp *HPP) basicHash() []string {
	amount := strconv.Itoa(hpp.Amount)
	return []string{hpp.timeStampStr(), hpp.MerchantID, hpp.OrderID, amount, hpp.Currency}
}

func (hpp *HPP) canStoreCard() bool {
	if hpp.EnableCardStorage || hpp.SelectStoredCard != "" {
		return true
	}
	return false
}

func (hpp *HPP) timeStampStr() string {
	ts := hpp.TimeStamp
	if ts == nil {
		return ""
	}
	return ts.Format(TimeLayout)
}
