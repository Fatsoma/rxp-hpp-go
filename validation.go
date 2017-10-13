package hpp

import (
	"regexp"

	"github.com/go-ozzo/ozzo-validation"
)

var (
	merchantIDRegexp     = regexp.MustCompile(`^[a-zA-Z0-9\.]*$`)
	accountRegexp        = regexp.MustCompile(`^[a-zA-Z0-9\s]*$`)
	orderIDRegexp        = regexp.MustCompile(`^[a-zA-Z0-9_\-]*$`)
	numericRegexp        = regexp.MustCompile(`^[0-9]*$`)
	alphaRegexp          = regexp.MustCompile(`^[a-zA-Z]*$`)
	hexadecimalRegexp    = regexp.MustCompile(`^[0-9a-fA-F]+$`)
	autoSettleFlagRegexp = regexp.MustCompile(`(?i)^on*|^off$|^*$|^multi$|^1$|^0$`)
	// commentRegexp        = regexp.MustCompile(`^[\s \u0020-\u003B \u003D \u003F-\u007E \u00A1-\u00FF\u20AC\u201A\u0192\u201E\u2026\u2020\u2021\u02C6\u2030\u0160\u2039\u0152\u017D\u2018\u2019\u201C\u201D\u2022\u2013\u2014\u02DC\u2122\u0161\u203A\u0153\u017E\u0178]*$`)
	boolRegexp         = regexp.MustCompile(`^[01]*$`)
	payerExistsRegexp  = regexp.MustCompile(`^[012]*$`)
	shippingCodeRegexp = regexp.MustCompile(`^[A-Za-z0-9\,\.\-\/\\| ]*$`)
	countryRegexp      = regexp.MustCompile(`^[A-Za-z0-9\,\.\- ]*$`)
	billingCodeRegexp  = regexp.MustCompile(`^[A-Za-z0-9\,\.\-\/\|\* ]*$`)
	referenceRegexp    = regexp.MustCompile(`^[a-zA-Z0-9\.\_\-\,\+\@ \s]*$`)
	languageRegexp     = regexp.MustCompile(`^[a-zA-Z]{2}$|^[a-zA-Z]{0}$`)
	payerRegexp        = regexp.MustCompile(`^[A-Za-z0-9\_\-\\ ]*$`)
	payRefRegexp       = regexp.MustCompile(`^[A-Za-z0-9\_\-]*$`)
	// cardPaymentButtonRegexp = regexp.MustCompile(`^[ÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏÐÑÒÓÔÕÖ×ØÙÚÛÜÝÞßàáâãäåæçèéêëìíîïðñòóôõö÷ø¤ùúûüýþÿ\\u0152\\u017D\\u0161\u0153\u017E\u0178¥a-zA-Z0-9\\'\\,\"\\+\\.\\_\\-\\&\\/\\@\\!\\?\\%\\()\\*\\:\\£\\$\\&\\u20AC\\#\\[\\]\\|\\=\\\\\u201C\u201D\u201C ]*$`)
)

const (
	merchantIDSize    = "Merchant ID is required and must be between 1 and 50 characters"
	merchantIDPattern = "Merchant ID must only contain alphanumeric characters"

	accountSize    = "Account must be 30 characters or less"
	accountPattern = "Account must only contain alphanumeric characters"

	orderIDSize    = "Order ID must be less than 50 characters in length"
	orderIDPattern = "Order ID must only contain alphanumeric characters, dash and underscore"

	amountSize    = "Amount is required and must be 11 characters or less"
	amountPattern = "Amount must only include numeric characters"
	amountOTB     = "Amount must be 0 for OTB transactions (where validate card only set to 1)"

	currencySize    = "Currency is required and must be 3 characters in length"
	currencyPattern = "Currency must only consist of alphabetic characters"

	timestampSize    = "Time stamp is required and must be 14 characters in length"
	timestampPattern = "Time stamp must be in YYYYMMDDHHMMSS format"

	hashSize    = "Security hash must be 40 characters in length"
	hashPattern = "Security hash must only contain numeric and a-f characters"

	autoSettleFlagPattern = "Auto settle flag must be 0, 1, on, off or multi"

	commentSize    = "Comment must be less than 255 characters in length"
	commentPattern = "Comment must only contain the characters a-z A-Z 0-9 ' \", + \u201C\u201D ._ - & \\ / @ ! ? % ( ) * : £ $ & \u20AC # [ ] | = ; ÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏÐÑÒÓÔÕÖ×ØÙÚÛÜÝÞßàáâãäåæçèéêëìíîïðñòóôõö÷ø¤ùúûüýþÿ\u0152\u017D\u0161\u0153\u017E\u0178¥"

	returnTssPattern = "Return TSS must be 1 or 0"

	shippingCodeSize    = "Shipping code must not be more than 30 characters in length"
	shippingCodePattern = "Shipping code must be of format <digits from postcode>|<digits from address> and contain only a-z A-Z 0-9 , . - / | spaces"

	shippingCountrySize    = "Shipping country must not contain more than 50 characters"
	shippingCountryPattern = "Shipping country must only contain the characters A-Z a-z 0-9 , . -"

	billingCodeSize    = "Billing code must not be more than 60 characters in length"
	billingCodePattern = "Billing code must be of format <digits from postcode>|<digits from address> and contain only a-z A-Z 0-9 , . - / | spaces"

	billingCountrySize    = "Billing country must not contain more than 50 characters"
	billingCountryPattern = "Billing country must only contain the characters A-Z a-z 0-9 , . -"

	customerNumberSize    = "Customer number must not contain more than 50 characters"
	customerNumberPattern = "Customer number must only contain the characters a-z A-Z 0-9 - _ . , + @ spaces"

	variableReferenceSize    = "Variable reference must not contain more than 50 characters"
	variableReferencePattern = "Variable reference must only contain the characters a-z A-Z 0-9 - _ . , + @ spaces"

	productIDSize    = "Product ID must not contain more than 50 characters"
	productIDPattern = "Product ID must only contain the characters a-z A-Z 0-9 - _ . , + @ spaces"

	languagePattern = "Language must be 2 alphabetic characters only"

	cardPaymentButtonTextSize    = "Card payment button text must not contain more than 25 characters"
	cardPaymentButtonTextPattern = "Card payment button text must only contain the characters a-z A-Z 0-9 ' , + \u201C\u201D ._ - & \\ / @!? % ( ) * :£ $ & \u20AC # [] | ="

	cardStorageEnableSize    = "Card storage enable flag must not be more than 1 character in length"
	cardStorageEnablePattern = "Card storage enable flag must be 1 or 0"

	offerSaveCardSize    = "Offer to save card flag must not be more than 1 character in length"
	offerSaveCardPattern = "Offer to save card flag must be 1 or 0"

	payerReferenceSize    = "Payer reference must not be more than 50 characters in length"
	payerReferencePattern = "Payer reference must only contain the characters a-z A-Z\\ 0-9 _ spaces"

	paymentReferenceSize    = "Payment reference must not be more than 50 characters in length"
	paymentReferencePattern = "Payment reference must only contain  characters a-z A-Z 0-9 _ - spaces"

	payerExistsSize    = "Payer exists flag must not be more than 1 character in length"
	payerExistsPattern = "Payer exists flag must be 0, 1 or 2"

	validateCardOnlySize    = "Validate card only flag must not be more than 1 character in length"
	validateCardOnlyPattern = "Validate card only flag must be 1 or 0"

	dccEnableSize    = "DCC enable flag must not be more than 1 character in length"
	dccEnablePattern = "DCC enable flag must be 1 or 0"
)

func validateMerchantID(merchantID *string) *validation.FieldRules {
	return validation.Field(
		merchantID,
		validation.Required.Error("is required"),
		validation.Length(1, 50).Error(merchantIDSize),
		validation.Match(merchantIDRegexp).Error(merchantIDPattern),
	)
}

func validateAccount(account *string) *validation.FieldRules {
	return validation.Field(
		account,
		validation.Length(0, 30).Error(accountSize),
		validation.Match(accountRegexp).Error(accountPattern),
	)
}

func validateOrderID(orderID *string) *validation.FieldRules {
	return validation.Field(
		orderID,
		validation.Length(0, 50).Error(orderIDSize),
		validation.Match(orderIDRegexp).Error(orderIDPattern),
	)
}

func validateAmount(amount *int) *validation.FieldRules {
	return validation.Field(
		amount,
		validation.Length(1, 11).Error(amountSize),
		validation.Match(numericRegexp).Error(amountPattern),
	)
}

func validateCurrency(currency *string) *validation.FieldRules {
	return validation.Field(
		currency,
		validation.Length(3, 3).Error(currencySize),
		validation.Match(alphaRegexp).Error(currencyPattern),
	)
}

func validateHash(hash *string) *validation.FieldRules {
	return validation.Field(
		hash,
		validation.Length(40, 40).Error(hashSize),
		validation.Match(hexadecimalRegexp).Error(hashPattern),
	)
}

func validateAutoSettleFlag(autoSettle *ConvertibleBoolean) *validation.FieldRules {
	return validation.Field(
		autoSettle,
		validation.Match(autoSettleFlagRegexp).Error(autoSettleFlagPattern),
	)
}

func validateComment(comment *string) *validation.FieldRules {
	return validation.Field(
		comment,
		validation.Length(0, 255).Error(commentSize),
		// validation.Match().Error(commentPattern),
	)
}

func validateReturnTss(returnTss *string) *validation.FieldRules {
	return validation.Field(
		returnTss,
		validation.Match(boolRegexp).Error(returnTssPattern),
	)
}

func validateShippingCode(shippingCode *string) *validation.FieldRules {
	return validation.Field(
		shippingCode,
		validation.Length(0, 30).Error(shippingCodeSize),
		validation.Match(shippingCodeRegexp).Error(shippingCodePattern),
	)
}

func validateShippingCountry(shippingCountry *string) *validation.FieldRules {
	return validation.Field(
		shippingCountry,
		validation.Length(0, 50).Error(shippingCountrySize),
		validation.Match(countryRegexp).Error(shippingCountryPattern),
	)
}

func validateBillingCode(billingCode *string) *validation.FieldRules {
	return validation.Field(
		billingCode,
		validation.Length(0, 60).Error(billingCodeSize),
		validation.Match(billingCodeRegexp).Error(billingCodePattern),
	)
}

func validateBillingCountry(billingCountry *string) *validation.FieldRules {
	return validation.Field(
		billingCountry,
		validation.Length(0, 50).Error(billingCountrySize),
		validation.Match(countryRegexp).Error(billingCountryPattern),
	)
}

func validateCustomerNumber(customerNumber *string) *validation.FieldRules {
	return validation.Field(
		customerNumber,
		validation.Length(0, 50).Error(customerNumberSize),
		validation.Match(referenceRegexp).Error(customerNumberPattern),
	)
}

func validateVariableReference(variableReference *string) *validation.FieldRules {
	return validation.Field(
		variableReference,
		validation.Length(0, 50).Error(variableReferenceSize),
		validation.Match(referenceRegexp).Error(variableReferencePattern),
	)
}

func validateProductID(productID *string) *validation.FieldRules {
	return validation.Field(
		productID,
		validation.Length(0, 50).Error(productIDSize),
		validation.Match(referenceRegexp).Error(productIDPattern),
	)
}

func validateLanguage(language *string) *validation.FieldRules {
	return validation.Field(
		language,
		validation.Match(languageRegexp).Error(languagePattern),
	)
}

func validateCardPaymentButton(cardPaymentButton *string) *validation.FieldRules {
	return validation.Field(
		cardPaymentButton,
		validation.Length(0, 25).Error(cardPaymentButtonTextSize),
		// validation.Match(languageRegexp).Error(languagePattern),
	)
}

func validateEnableCardStorage(enableCardStorage *ConvertibleBoolean) *validation.FieldRules {
	return validation.Field(
		enableCardStorage,
		validation.Match(boolRegexp).Error(cardStorageEnablePattern),
	)
}

func validateOfferSaveCard(offerSaveCard *ConvertibleBoolean) *validation.FieldRules {
	return validation.Field(
		offerSaveCard,
		validation.Match(boolRegexp).Error(offerSaveCardPattern),
	)
}

func validatePayerReference(payerReference *string) *validation.FieldRules {
	return validation.Field(
		payerReference,
		validation.Length(0, 50).Error(payerReferenceSize),
		validation.Match(payerRegexp).Error(payerReferenceSize),
	)
}

func validatePaymentReference(paymentReference *string) *validation.FieldRules {
	return validation.Field(
		paymentReference,
		validation.Length(0, 50).Error(paymentReferenceSize),
		validation.Match(payRefRegexp).Error(paymentReferencePattern),
	)
}

func validatePayerExists(payerExists *ConvertibleBoolean) *validation.FieldRules {
	return validation.Field(
		payerExists,
		validation.Match(payerExistsRegexp).Error(payerExistsPattern),
	)
}

func validateValidCardOnly(validCardOnly *ConvertibleBoolean) *validation.FieldRules {
	return validation.Field(
		validCardOnly,
		validation.Match(boolRegexp).Error(validateCardOnlyPattern),
	)
}

func validateDCCEnable(DCCEnable *ConvertibleBoolean) *validation.FieldRules {
	return validation.Field(
		DCCEnable,
		validation.Match(boolRegexp).Error(dccEnablePattern),
	)
}
