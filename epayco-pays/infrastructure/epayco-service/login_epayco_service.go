package epaycoservice

import (
	"encoding/base64"

	epaycoconstants "gitlab.com/eliotandelon/epayco-pays/domain/epayco-constants"
	epaycorequest "gitlab.com/eliotandelon/epayco-pays/infrastructure/epayco-request"
	epaycoresponse "gitlab.com/eliotandelon/epayco-pays/infrastructure/epayco-response"
)

func (service *EpaycoService) LoginEpayco() (*epaycoresponse.LoginResponse, error) {

	loginData := service.EpaycoPublicKey + ":" + service.EpaycoPrivateKey

	encoded := base64.StdEncoding.EncodeToString([]byte(loginData))
	var loginResponse epaycoresponse.LoginResponse
	headers := []epaycorequest.HeaderRequest{
		{Key: epaycoconstants.HEADER_AUTORIZATION, Value: epaycoconstants.HEADER_KEY_BASIC + encoded},
	}
	err := service.Client.Post(service.UrlBase, epaycoconstants.URL_EPAYCO_LOGIN, nil, &loginResponse, headers)
	if err != nil {
		return nil, err
	}
	return &loginResponse, nil
}
