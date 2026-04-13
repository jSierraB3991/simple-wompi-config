package wompiservice

import (
	eliotlibs "github.com/jSierraB3991/jsierra-libs"
	wompilibs "github.com/jSierraB3991/simple-wompi-config/wompi_domain/wompi_libs"
	wompiresponse "github.com/jSierraB3991/simple-wompi-config/wompi_infrastructure/wompi_response"
)

func (w *WompiService) getAcceptanceToken() (string, string, error) {
	var result wompiresponse.MerchantResponse
	err := eliotlibs.Get(w.urlApi, wompilibs.MERCHANT_URL+w.woompiPublicKey, &result, nil)
	if err != nil {
		return "", "", err
	}
	return result.Data.PresignedAcceptance.AcceptanceToken, result.Data.PresignedAcceptance.Permalink, nil
}
