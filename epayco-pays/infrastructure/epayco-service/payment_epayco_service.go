package epaycoservice

import (
	"encoding/json"

	constants "gitlab.com/eliotandelon/epayco-pays/domain/epayco-constants"
	epaycorequest "gitlab.com/eliotandelon/epayco-pays/infrastructure/epayco-request"
	epaycoresponse "gitlab.com/eliotandelon/epayco-pays/infrastructure/epayco-response"
)

func (service *EpaycoService) Pay(paymentRequest epaycorequest.PaymentRequest, token string) (*epaycoresponse.PaymentResponse, error) {

	paymentRequest.PrivateKey = service.EpaycoPrivateKey
	paymentRequest.PublicKey = service.EpaycoPublicKey
	paymentRequest.TypeSell = constants.TYPE_SHELL_NORMAL

	jsonData, err := json.Marshal(paymentRequest)
	if err != nil {
		return nil, err
	}
	var paymentDataResponse epaycoresponse.PaymentDataResponse
	headers := []epaycorequest.HeaderRequest{
		{Key: constants.HEADER_AUTORIZATION, Value: constants.HEADER_KEY_BEARER + token},
	}

	err = service.Client.Post(service.UrlBase, constants.URL_EPAYCO_CREATE_PAY, jsonData, &paymentDataResponse, headers)
	if err != nil {
		return nil, err
	}
	return &paymentDataResponse.Data, nil
}
