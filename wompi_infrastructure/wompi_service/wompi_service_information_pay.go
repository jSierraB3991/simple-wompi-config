package wompiservice

import (
	eliotlibs "github.com/jSierraB3991/jsierra-libs"
	wompilibs "github.com/jSierraB3991/simple-wompi-config/wompi_domain/wompi_libs"
	wompiresponse "github.com/jSierraB3991/simple-wompi-config/wompi_infrastructure/wompi_response"
)

func (w *WompiService) WompiInfoPay(transactionId string) (*wompiresponse.WompiInfoTransactionDataResponse, error) {
	var responseApi wompiresponse.WompiInfoTransactionResponse
	err := eliotlibs.Get(w.urlApi, wompilibs.TRANSACTION+transactionId, &responseApi, nil)
	if err != nil {
		return nil, err
	}
	return &responseApi.Data, nil
}
