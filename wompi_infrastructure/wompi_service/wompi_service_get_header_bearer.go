package wompiservice

import (
	eliotlibs "github.com/jSierraB3991/jsierra-libs"
	wompilibs "github.com/jSierraB3991/simple-wompi-config/wompi_domain/wompi_libs"
)

func (w *WompiService) getHeaderBearer() eliotlibs.HeaderRequest {
	return eliotlibs.HeaderRequest{
		Key:   eliotlibs.HeaderAuthorization,
		Value: wompilibs.BEARER + w.woompiPrivateKey,
	}
}
