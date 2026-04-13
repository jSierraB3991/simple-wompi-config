package wompiservice

type WompiService struct {
	woompiPrivateKey string
	woompiPublicKey  string
	integrityKey     string
	urlApi           string
}

func NewWompiClient(woompiPrivateKey, woompiPublicKey, integrityKey, urlApi string) *WompiService {
	return &WompiService{
		woompiPrivateKey: woompiPrivateKey,
		woompiPublicKey:  woompiPublicKey,
		integrityKey:     integrityKey,
		urlApi:           urlApi,
	}
}
