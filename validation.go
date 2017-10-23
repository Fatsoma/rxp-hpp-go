package hpp

import (
	"regexp"

	"github.com/go-ozzo/ozzo-validation"
)

var (
	merchantIDRegexp        = regexp.MustCompile(`^[a-zA-Z0-9\.]*$`)
	accountRegexp           = regexp.MustCompile(`^[a-zA-Z0-9\s]*$`)
	orderIDRegexp           = regexp.MustCompile(`^[a-zA-Z0-9_\-]*$`)
	numericRegexp           = regexp.MustCompile(`^[0-9]*$`)
	alphaRegexp             = regexp.MustCompile(`^[a-zA-Z]*$`)
	hexadecimalRegexp       = regexp.MustCompile(`^[0-9a-fA-F]+$`)
	autoSettleFlagRegexp    = regexp.MustCompile(`(?i)^on*|^off$|^*$|^multi$|^1$|^0$`)
	commentRegexp           = regexp.MustCompile(`^[\s \x{0020}-\x{003B} \x{003D} \x{003F}-\x{007E} \x{00A1}-\x{00FF}\x{20AC}\x{201A}\x{0192}\x{201E}\x{2026}\x{2020}\x{2021}\x{02C6}\x{2030}\x{0160}\x{2039}\x{0152}\x{017D}\x{2018}\x{2019}\x{201C}\x{201D}\x{2022}\x{2013}\x{2014}\x{02DC}\x{2122}\x{0161}\x{203A}\x{0153}\x{017E}\x{0178}]*$`)
	boolRegexp              = regexp.MustCompile(`^[01]*$`)
	payerExistsRegexp       = regexp.MustCompile(`^[012]*$`)
	shippingCodeRegexp      = regexp.MustCompile(`^[A-Za-z0-9\,\.\-\/\\| ]*$`)
	countryRegexp           = regexp.MustCompile(`^[A-Za-z0-9\,\.\- ]*$`)
	billingCodeRegexp       = regexp.MustCompile(`^[A-Za-z0-9\,\.\-\/\|\* ]*$`)
	referenceRegexp         = regexp.MustCompile(`^[a-zA-Z0-9\.\_\-\,\+\@ \s]*$`)
	languageRegexp          = regexp.MustCompile(`^[a-zA-Z]{2}$|^[a-zA-Z]{0}$`)
	payerRegexp             = regexp.MustCompile(`^[A-Za-z0-9\_\-\\ ]*$`)
	payRefRegexp            = regexp.MustCompile(`^[A-Za-z0-9\_\-]*$`)
	cardPaymentButtonRegexp = regexp.MustCompile(`^[\x{00C0}\x{00C1}\x{00C2}\x{00C3}\x{00C4}\x{00C5}\x{00C6}\x{00C7}\x{00C8}\x{00C9}\x{00CA}\x{00CB}\x{00CC}\x{00CD}\x{00CE}\x{00CF}\x{00D0}\x{00D1}\x{00D2}\x{00D3}\x{00D4}\x{00D5}\x{00D6}\x{00D7}\x{00D8}\x{00D9}\x{00DA}\x{00DB}\x{00DC}\x{00DD}\x{00DE}\x{00DF}\x{00E0}\x{00E1}\x{00E2}\x{00E3}\x{00E4}\x{00E5}\x{00E6}\x{00E7}\x{00E8}\x{00E9}\x{00EA}\x{00EB}\x{00EC}\x{00ED}\x{00EE}\x{00EF}\x{00F0}\x{00F1}\x{00F2}\x{00F3}\x{00F4}\x{00F5}\x{00F6}\x{00F7}\x{00F8}\x{00A4}\x{00F9}\x{00FA}\x{00FB}\x{00FC}\x{00FD}\x{00FE}\x{00FF}\x{0152}\x{017D}\x{0161}\x{0153}\x{017E}\x{0178}\x{00A5}a-zA-Z0-9\'\,\+\x{0022}\.\_\-\&\/\@\!\?\%\()\*\:\x{00A3}\$\&\x{20AC}\#\[\]\|\=\\\x{201C}\x{201D}\x{201C} ]*$`)
)

const (
	merchantIDSize    = "Merchant ID is required and must be between 1 and 50 characters"
	merchantIDPattern = "Merchant ID must only contain alphanumeric characters"

	accountSize    = "Account must be 30 characters or less"
	accountPattern = "Account must only contain alphanumeric characters"

	orderIDSize    = "Order ID must be less than 50 characters in length"
	orderIDPattern = "Order ID must only contain alphanumeric characters, dash and underscore"

	amountSize = "Amount is required and must be 11 characters or less"
	amountOTB  = "Amount must be 0 for OTB transactions (where validate card only set to 1)"

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
		validation.Required.Error("is required"),
		validation.Min(1).Error(amountSize),
		validation.Max(999999999).Error(amountSize),
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

func validateAutoSettleFlag(autoSettle *string) *validation.FieldRules {
	return validation.Field(
		autoSettle,
		validation.Match(autoSettleFlagRegexp).Error(autoSettleFlagPattern),
	)
}

func validateComment(comment *string) *validation.FieldRules {
	return validation.Field(
		comment,
		validation.Length(0, 255).Error(commentSize),
		validation.Match(commentRegexp).Error(commentPattern),
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
		validation.Match(cardPaymentButtonRegexp).Error(cardPaymentButtonTextPattern),
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

func validatePayerExists(payerExists *string) *validation.FieldRules {
	return validation.Field(
		payerExists,
		validation.Length(1, 1).Error(payerExistsSize),
		validation.Match(payerExistsRegexp).Error(payerExistsPattern),
	)
}
