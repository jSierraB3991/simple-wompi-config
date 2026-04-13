package wompiresponse

type WompiDataCustomerReference struct {
	Label      string `json:"label"`
	IsRequired bool   `json:"is_required"`
}

type WompiDataCustomerData struct {
	CustomerReferences []WompiDataCustomerReference `json:"customer_references"`
}

type WoompiDataResponse struct {
	ID                string                 `json:"id"`
	CustomerEmail     string                 `json:"customer_email"`
	PaymentMethodType string                 `json:"payment_method_type"`
	Status            string                 `json:"status"`
	StatusMessage     *string                `json:"status_message"`
	Currency          string                 `json:"currency"`
	AmountInCents     int64                  `json:"amount_in_cents"`
	RedirectUrl       *string                `json:"redirect_url"`
	CustomerData      *WompiDataCustomerData `json:"customer_data"`
	CreatedAt         string                 `json:"created_at"`
	Reference         string                 `json:"reference"`
	FinalizedAt       *string                `json:"finalized_at"`
}

type WompiPaymentResponse struct {
	Data WoompiDataResponse `json:"data"`
}
