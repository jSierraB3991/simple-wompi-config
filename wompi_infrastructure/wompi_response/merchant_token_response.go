package wompiresponse

type MerchantDataPreSignedResponse struct {
	AcceptanceToken string `json:"acceptance_token"`
	Permalink       string `json:"permalink"`
	Type            string `json:"type"`
}

type MerchantDataResponse struct {
	ID                  int                           `json:"id"`
	PresignedAcceptance MerchantDataPreSignedResponse `json:"presigned_acceptance"`
}

type MerchantResponse struct {
	Data MerchantDataResponse `json:"data"`
}
