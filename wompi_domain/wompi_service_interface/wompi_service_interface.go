package wompiserviceinterface

import (
	wompirequest "github.com/jSierraB3991/simple-wompi-config/wompi_infrastructure/wompi_request"
	wompiresponse "github.com/jSierraB3991/simple-wompi-config/wompi_infrastructure/wompi_response"
)

type WompiServiceInterface interface {
	SendPaymentLink(req wompirequest.WompiSendPaymentRequest, cellphone string) (*wompiresponse.WompiPaymentResponse, error)
	WompiInfoPay(transactionId string) (*wompiresponse.WompiInfoTransactionDataResponse, error)
}
