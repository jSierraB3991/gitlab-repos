package epaycoserviceinterface

import (
	epaycorequest "gitlab.com/eliotandelon/epayco-pays/infrastructure/epayco-request"
	epaycoresponse "gitlab.com/eliotandelon/epayco-pays/infrastructure/epayco-response"
)

type EpaycoServiceInterface interface {
	LoginEpayco() (*epaycoresponse.LoginResponse, error)
	Pay(paymentRequest epaycorequest.PaymentRequest, token string) (*epaycoresponse.PaymentResponse, error)
	TransactionList(req epaycorequest.TransactionListRequest, token string) ([]epaycoresponse.DataTransactionListResponse, error)
}
