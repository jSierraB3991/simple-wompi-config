package wompirequest

type PaymentMethodReq struct {
	Type        string `json:"type"`
	PhoneNumber string `json:"phone_number"`
}

type MetadataReq struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	SingleUse       bool   `json:"single_use"`
	CollectShipping bool   `json:"collect_shipping"`
	RedirectUrl     string `json:"redirect_url"`
}

type WompiSendPaymentRequest struct {
	AmountInCents   int64  `json:"amount_in_cents"`
	Currency        string `json:"currency"`
	Reference       string `json:"reference"`
	CustomerEmail   string `json:"customer_email,omitempty"`
	AcceptanceToken string `json:"acceptance_token"`

	PaymentMethod PaymentMethodReq `json:"payment_method"`
	Signature     string           `json:"signature"`

	Metadata MetadataReq `json:"metadata"`
}
