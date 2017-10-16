package hpp

import (
	"strconv"
	"strings"

	"github.com/go-ozzo/ozzo-validation"
)

type Request struct {
	// The merchant ID supplied by Realex Payments – note this is not the merchant number supplied by your bank.
	MerchantID string `json:"MERCHANT_ID"`

	// The sub-account to use for this transaction. If not present, the default sub-account will be used.
	Account string `json:"ACCOUNT"`

	// A unique alphanumeric id that’s used to identify the transaction. No spaces are allowed.
	OrderID string `json:"ORDER_ID"`

	// Total amount to authorise in the lowest unit of the currency – i.e. 100 euro would be entered as 10000.
	// If there is no decimal in the currency (e.g. JPY Yen) then contact Realex Payments. No decimal points are allowed.
	// Amount should be set to 0 for OTB transactions (i.e. where validate card only is set to 1).
	Amount int `json:"AMOUNT,string"`

	// A three-letter currency code (Eg. EUR, GBP). A list of currency codes can be provided by your account manager.
	Currency string `json:"CURRENCY"`

	// Date and time of the transaction. Entered in the following format: YYYYMMDDHHMMSS. Must be within 24 hours of the current time.
	TimeStamp *JSONTime `json:"TIMESTAMP"`

	// A digital signature generated using the SHA-1 algorithm.
	Hash string `json:"SHA1HASH"`

	// Used to signify whether or not you wish the transaction to be captured in the next batch.
	// If set to "1" and assuming the transaction is authorised then it will automatically be settled in the next batch.
	// If set to "0" then the merchant must use the RealControl application to manually settle the transaction.
	// This option can be used if a merchant wishes to delay the payment until after the goods have been shipped.
	// Transactions can be settled for up to 115% of the original amount and must be settled within a certain period of time agreed with your issuing bank.
	AutoSettleFlag ConvertibleBoolean `json:"AUTO_SETTLE_FLAG"`

	// A freeform comment to describe the transaction.
	CommentOne string `json:"COMMENT1"`

	// A freeform comment to describe the transaction.
	CommentTwo string `json:"COMMENT2"`

	// Used to signify whether or not you want a Transaction Suitability Score for this transaction.
	// Can be "0" for no and "1" for yes.
	ReturnTSS string `json:"RETURN_TSS"`

	// The postcode or ZIP of the shipping address.
	ShippingCode string `json:"SHIPPING_CODE"`

	// The country of the shipping address.
	ShippingCountry string `json:"SHIPPING_CO"`

	// The postcode or ZIP of the billing address.
	BillingCode string `json:"BILLING_CODE"`

	// The country of the billing address.
	BillingCountry string `json:"BILLING_CO"`

	// The customer number of the customer. You can send in any additional information about the transaction in this field,
	// which will be visible under the transaction in the RealControl application.
	CustomerNumber string `json:"CUST_NUM"`

	// A variable reference also associated with this customer. You can send in any additional information about the transaction in this field,
	// which will be visible under the transaction in the RealControl application.
	VariableReference string `json:"VAR_REF"`

	// A product id associated with this product. You can send in any additional information about the transaction in this field,
	// which will be visible under the transaction in the RealControl application.
	ProductID string `json:"PROD_ID"`

	// Used to set what language HPP is displayed in. Currently HPP is available in English, Spanish and German, with other languages to follow.
	// If the field is not sent in, the default language is the language that is set in your account configuration. This can be set by your account manager.
	Language string `json:"HPP_LANG"`

	// Used to set what text is displayed on the payment button for card transactions. If this field is not sent in, "Pay Now" is displayed on the button by default.
	CardPaymentButton string `json:"CARD_PAYMENT_BUTTON"`

	// Enable card storage.
	EnableCardStorage ConvertibleBoolean `json:"CARD_STORAGE_ENABLE"`

	// Offer to save the card.
	OfferSaveCard ConvertibleBoolean `json:"OFFER_SAVE_CARD"`

	// The payer reference.
	PayerReference string `json:"PAYER_REF"`

	// The payment reference.
	PaymentReference string `json:"PMT_REF"`

	// Payer exists.
	PayerExists string `json:"PAYER_EXIST"`

	// Used to identify an OTB transaction.
	ValidCardOnly ConvertibleBoolean `json:"VALIDATE_CARD_ONLY"`

	// Transaction level configuration to enable/disable a DCC request. (Only if the merchant is configured).
	DCCEnable ConvertibleBoolean `json:"DCC_ENABLE"`

	// Override merchant configuration for fraud. (Only if the merchant is configured for fraud).
	FraudFilterMode string `json:"HPP_FRAUDFILTER_MODE"`

	// The HPP Version. To use HPP Card Management select HPP_VERSION = 2.
	Version string `json:"HPP_VERSION"`

	// The payer reference. If this flag is received, HPP will retrieve a list of the payment methods saved for that payer.
	SelectStoredCard string `json:"HPP_SELECT_STORED_CARD"`

	// Anything else you sent to us in the request.
	SupplementaryData map[string]string `json:"-"`
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
		validateAutoSettleFlag(&r.AutoSettleFlag),
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
		validateEnableCardStorage(&r.EnableCardStorage),
		validateOfferSaveCard(&r.OfferSaveCard),
		validatePayerReference(&r.PayerReference),
		validatePaymentReference(&r.PaymentReference),
		validatePayerExists(&r.PayerExists),
		validateValidCardOnly(&r.ValidCardOnly),
		validateDCCEnable(&r.DCCEnable),
	)
}

// BuildHash creates the security hash from a number of fields and the shared secret.
func (r *Request) BuildHash(secret string) string {
	s := r.buildHashString()
	return GenerateHash(s, secret)
}

func (r *Request) buildHashString() string {
	s := r.basicHash()

	if r.canStoreCard() {
		s = append(s, []string{r.PayerReference, r.PaymentReference}...)
	}

	if r.FraudFilterMode != "" {
		s = append(s, r.FraudFilterMode)
	}

	return strings.Join(s, Separator)
}

func (r *Request) basicHash() []string {
	ts := ""
	if r.TimeStamp != nil {
		ts = r.TimeStamp.String()
	}
	amount := strconv.Itoa(r.Amount)

	return []string{ts, r.MerchantID, r.OrderID, amount, r.Currency}
}

func (r *Request) canStoreCard() bool {
	return bool(r.EnableCardStorage) || r.SelectStoredCard != ""
}
