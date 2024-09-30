package epaycoservice

import (
	"encoding/json"

	constants "gitlab.com/eliotandelon/epayco-pays/domain/epayco-constants"
	epaycorequest "gitlab.com/eliotandelon/epayco-pays/infrastructure/epayco-request"
	epaycoresponse "gitlab.com/eliotandelon/epayco-pays/infrastructure/epayco-response"
)

// Send void req, for get all transactions
func (service *EpaycoService) TransactionList(req epaycorequest.TransactionListRequest, token string) ([]epaycoresponse.DataTransactionListResponse, error) {

	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	headers := []epaycorequest.HeaderRequest{
		{Key: constants.HEADER_AUTORIZATION, Value: constants.HEADER_KEY_BEARER + token},
	}
	var transactionListResponse epaycoresponse.EpaycoTransactionResponse
	err = service.Client.Post(service.UrlBase, constants.URL_EPAYCO_TRANSACTION_LIST, jsonData, &transactionListResponse, headers)
	if err != nil {
		return nil, err
	}
	return transactionListResponse.Data.Data, nil

}
