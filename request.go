package hpp

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

type Request struct {
	hpp *HPP

	// The merchant ID supplied by Realex Payments – note this is not the merchant number supplied by your bank.
	MerchantID JSONString `json:"MERCHANT_ID"`

	// The sub-account to use for this transaction. If not present, the default sub-account will be used.
	Account JSONString `json:"ACCOUNT"`

	// A unique alphanumeric id that’s used to identify the transaction. No spaces are allowed.
	OrderID []byte `json:"ORDER_ID"`

	// Total amount to authorise in the lowest unit of the currency – i.e. 100 euro would be entered as 10000.
	// If there is no decimal in the currency (e.g. JPY Yen) then contact Realex Payments. No decimal points are allowed.
	// Amount should be set to 0 for OTB transactions (i.e. where validate card only is set to 1).
	Amount JSONInt `json:"AMOUNT,string"`

	// A three-letter currency code (Eg. EUR, GBP). A list of currency codes can be provided by your account manager.
	Currency JSONString `json:"CURRENCY,omitempty"`

	// Date and time of the transaction. Entered in the following format: YYYYMMDDHHMMSS. Must be within 24 hours of the current time.
	TimeStamp *JSONTime `json:"TIMESTAMP"`

	// A digital signature generated using the SHA-1 algorithm.
	Hash JSONString `json:"SHA1HASH"`

	// Used to signify whether or not you wish the transaction to be captured in the next batch.
	// If set to "1" and assuming the transaction is authorised then it will automatically be settled in the next batch.
	// If set to "0" then the merchant must use the RealControl application to manually settle the transaction.
	// This option can be used if a merchant wishes to delay the payment until after the goods have been shipped.
	// Transactions can be settled for up to 115% of the original amount and must be settled within a certain period of time agreed with your issuing bank.
	AutoSettleFlag JSONBool `json:"AUTO_SETTLE_FLAG,omitempty"`

	// A freeform comment to describe the transaction.
	CommentOne JSONString `json:"COMMENT1,omitempty"`

	// A freeform comment to describe the transaction.
	CommentTwo JSONString `json:"COMMENT2,omitempty"`

	// Used to signify whether or not you want a Transaction Suitability Score for this transaction.
	// Can be "0" for no and "1" for yes.
	ReturnTSS JSONString `json:"RETURN_TSS,omitempty"`

	// The postcode or ZIP of the shipping address.
	ShippingCode JSONString `json:"SHIPPING_CODE,omitempty"`

	// The country of the shipping address.
	ShippingCountry JSONString `json:"SHIPPING_CO",omitempty`

	// The postcode or ZIP of the billing address.
	BillingCode JSONString `json:"BILLING_CODE,omitempty"`

	// The country of the billing address.
	BillingCountry JSONString `json:"BILLING_CO,omitempty"`

	// The customer number of the customer. You can send in any additional information about the transaction in this field,
	// which will be visible under the transaction in the RealControl application.
	CustomerNumber JSONString `json:"CUST_NUM,omitempty"`

	// A variable reference also associated with this customer. You can send in any additional information about the transaction in this field,
	// which will be visible under the transaction in the RealControl application.
	VariableReference JSONString `json:"VAR_REF,omitempty"`

	// A product id associated with this product. You can send in any additional information about the transaction in this field,
	// which will be visible under the transaction in the RealControl application.
	ProductID JSONString `json:"PROD_ID,omitempty"`

	// Used to set what language HPP is displayed in. Currently HPP is available in English, Spanish and German, with other languages to follow.
	// If the field is not sent in, the default language is the language that is set in your account configuration. This can be set by your account manager.
	Language JSONString `json:"HPP_LANG,omitempty"`

	// Used to set what text is displayed on the payment button for card transactions. If this field is not sent in, "Pay Now" is displayed on the button by default.
	CardPaymentButton JSONString `json:"CARD_PAYMENT_BUTTON,omitempty"`

	// Enable card storage.
	EnableCardStorage JSONBool `json:"CARD_STORAGE_ENABLE"`

	// Offer to save the card.
	OfferSaveCard JSONBool `json:"OFFER_SAVE_CARD"`

	// The payer reference.
	PayerReference JSONString `json:"PAYER_REF"`

	// The payment reference.
	PaymentReference JSONString `json:"PMT_REF,omitempty"`

	// Payer exists.
	PayerExists JSONString `json:"PAYER_EXIST,omitempty"`

	// Used to identify an OTB transaction.
	ValidCardOnly JSONBool `json:"VALIDATE_CARD_ONLY"`

	// Transaction level configuration to enable/disable a DCC request. (Only if the merchant is configured).
	DCCEnable JSONBool `json:"DCC_ENABLE"`

	// Override merchant configuration for fraud. (Only if the merchant is configured for fraud).
	FraudFilterMode JSONString `json:"HPP_FRAUDFILTER_MODE,omitempty"`

	// The HPP Version. To use HPP Card Management select HPP_VERSION = 2.
	Version JSONString `json:"HPP_VERSION,omitempty"`

	// The payer reference. If this flag is received, HPP will retrieve a list of the payment methods saved for that payer.
	SelectStoredCard JSONString `json:"HPP_SELECT_STORED_CARD,omitempty"`

	// Anything else you sent to us in the request.
	SupplementaryData map[string]string `json:"-"`
}

// ToJSON converts the request into valid JSON
// Validates inputs, generates security hash, order ID and time stamp (if required)
// Base64 encodes inputs, and serialises itself to JSON
func (r *Request) ToJSON() (json.RawMessage, error) {
	fmt.Println("Converting HppRequest to JSON.")

	fmt.Println("Generating defaults.")
	r.generateDefaults()

	r.BuildHash(r.hpp.Secret)

	fmt.Println("Validating request.")

	err := r.Validate()
	if err != nil {
		return nil, errors.Wrap(err, "failed to validate HPP request")
	}

	js, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal HPP request")
	}

	return js, nil
}

func (r *Request) generateDefaults() {
	if r.TimeStamp == nil {
		now := JSONTime(time.Now())
		r.TimeStamp = &now
	}

	if len(r.OrderID) < 1 {
		r.OrderID = uuid.NewV4().Bytes()
	}
}

// Validate the HPP request fields
func (r *Request) Validate() error {
	return validation.ValidateStruct(r,
		validateMerchantID(&r.MerchantID),
		validateAccount(&r.Account),
		validateOrderID(&r.OrderID),
		validateAmount(&r.Amount),
		validateCurrency(&r.Currency),
		validateHash(&r.Hash),
		validateComment(&r.CommentOne),
		validateComment(&r.CommentTwo),
		validateReturnTss(&r.ReturnTSS),
		validateShippingCode(&r.ShippingCode),
		validateShippingCountry(&r.ShippingCountry),
		validateBillingCode(&r.BillingCode),
		validateBillingCountry(&r.BillingCountry),
		validateCustomerNumber(&r.CustomerNumber),
		validateVariableReference(&r.VariableReference),
		validateProductID(&r.ProductID),
		validateLanguage(&r.Language),
		validateCardPaymentButton(&r.CardPaymentButton),
		validatePayerReference(&r.PayerReference),
		validatePaymentReference(&r.PaymentReference),
		validatePayerExists(&r.PayerExists),
	)
}

// BuildHash creates the security hash from a number of fields and the shared secret.
func (r *Request) BuildHash(secret string) {
	s := r.buildHashString()
	r.Hash = JSONString(GenerateHash(s, secret))
}

func (r *Request) buildHashString() string {
	s := r.basicHash()

	if r.canStoreCard() {
		s = append(s, []string{r.PayerReference.String(), r.PaymentReference.String()}...)
	}

	if r.FraudFilterMode != "" {
		s = append(s, r.FraudFilterMode.String())
	}

	return strings.Join(s, Separator)
}

func (r *Request) basicHash() []string {
	ts := ""
	if r.TimeStamp != nil {
		ts = r.TimeStamp.String()
	}
	amount := r.Amount.String()
	orderID := string(r.OrderID)

	return []string{ts, r.MerchantID.String(), orderID, amount, r.Currency.String()}
}

func (r *Request) canStoreCard() bool {
	return bool(r.EnableCardStorage) || r.SelectStoredCard != ""
}
