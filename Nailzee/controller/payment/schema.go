package payment

type PaymentService interface {
	CreatePaymentIntent(userId, token string, payments PaymentIntentSchema) (string, string, error)
	CreateNewPaymentIntent(userId, token string, payments PaymentIntentSchema) (string, string, error)
	ConfirmPaymentIntent(userId, token string, payments PaymentIntentSchema) (string, error)
}
type PaymentIntentSchema struct {
	Name            string `json:"name,omitempty"`
	AddressLine     string `json:"address,omitempty"`
	Pin             string `json:"pin,omitempty"`
	State           string `json:"state,omitempty"`
	City            string `json:"city,omitempty"`
	Country         string `json:"country,omitempty"`
	Amount          int64  `json:"amount,omitempty"`
	Currency        string `json:"currency,omitempty"`
	Description     string `json:"description,omitempty"`
	ReceiptEmail    string `json:"receipt_email,omitempty"`
	UserId          string `json:"user_id,omitempty,omitempty"`
	PaymentMethodID string `json:"payment_method_id,omitempty"`
	PaymentIntentID string `json:"payment_intent_id,omitempty"`
	IsSaved         bool   `json:"is_saved,omitempty"`
}
