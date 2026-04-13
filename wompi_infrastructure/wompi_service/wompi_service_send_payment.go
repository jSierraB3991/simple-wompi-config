package wompiservice

import (
	"encoding/json"

	eliotlibs "github.com/jSierraB3991/jsierra-libs"
	wompilibs "github.com/jSierraB3991/simple-wompi-config/wompi_domain/wompi_libs"
	wompirequest "github.com/jSierraB3991/simple-wompi-config/wompi_infrastructure/wompi_request"
	wompiresponse "github.com/jSierraB3991/simple-wompi-config/wompi_infrastructure/wompi_response"
)

func (w *WompiService) SendPaymentLink(req wompirequest.WompiSendPaymentRequest, cellphone string) (*wompiresponse.WompiPaymentResponse, error) {
	acceptedToken, _, err := w.getAcceptanceToken()
	if err != nil {
		return nil, err
	}

	req.AcceptanceToken = acceptedToken
	req.PaymentMethod = wompirequest.PaymentMethodReq{
		Type:        wompilibs.NEQUI,
		PhoneNumber: cellphone,
	}

	signture := w.generateSignature(req.Reference, req.AmountInCents, eliotlibs.COP)
	req.Currency = eliotlibs.COP
	req.Signature = signture
	req.Metadata.SingleUse = true
	req.Metadata.CollectShipping = false

	jsonBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	header := w.getHeaderBearer()

	response := wompiresponse.WompiPaymentResponse{}
	err = eliotlibs.Post(w.urlApi, wompilibs.TRANSACTIONS, jsonBody, &response, []eliotlibs.HeaderRequest{header})
	if err != nil {
		return nil, err
	}
	return &response, nil
}
