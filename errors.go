package hpp

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

	comment1Size    = "Comment 1 must be less than 255 characters in length"
	comment1Pattern = "Comment 1 must only contain the characters a-z A-Z 0-9 ' \", + \u201C\u201D ._ - & \\ / @ ! ? % ( ) * : £ $ & \u20AC # [ ] | = ; ÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏÐÑÒÓÔÕÖ×ØÙÚÛÜÝÞßàáâãäåæçèéêëìíîïðñòóôõö÷ø¤ùúûüýþÿ\u0152\u017D\u0161\u0153\u017E\u0178¥"

	comment2Size    = "Comment 2 must be less than 255 characters in length"
	comment2Pattern = "Comment 2 must only contain the characters a-z A-Z 0-9 ' \", + \u201C\u201D ._ - & \\ / @ ! ? % ( ) * : £ $ & \u20AC # [ ] | = ; ÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏÐÑÒÓÔÕÖ×ØÙÚÛÜÝÞßàáâãäåæçèéêëìíîïðñòóôõö÷ø¤ùúûüýþÿ\u0152\u017D\u0161\u0153\u017E\u0178¥"

	returnTssSize    = "Return TSS flag must not be more than 1 character in length"
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
