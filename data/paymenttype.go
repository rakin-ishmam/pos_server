package data

// PaymentType represents pos payment type
type PaymentType string

const (
	// CashPaymentType represents cash payment
	CashPaymentType PaymentType = "Cash"

	// CheckPaymentType represnets check payment
	CheckPaymentType PaymentType = "Check"

	// CreditCardPaymentType represents credit card payment
	CreditCardPaymentType PaymentType = "CreditCard"

	// DebitCardPaymentType represents debit card payment
	DebitCardPaymentType PaymentType = "DebitCard"
)

// Validate returns true if payment type is valid
func (p PaymentType) Validate() bool {
	switch p {
	case CashPaymentType:
		return true
	case CheckPaymentType:
		return true
	case CreditCardPaymentType:
		return true
	case DebitCardPaymentType:
		return true
	}

	return false
}
