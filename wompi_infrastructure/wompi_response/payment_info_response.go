package wompiresponse

type WompiInfoTransactionDataResponse struct {
	ID                string `json:"id"`
	PaymentMethodType string `json:"payment_method_type"`
	Status            string `json:"status"`
}

type WompiInfoTransactionResponse struct {
	Data WompiInfoTransactionDataResponse `json:"data"`
}
