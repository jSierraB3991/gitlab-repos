package epaycoservice

import (
	epaycoclient "gitlab.com/eliotandelon/epayco-pays/infrastructure/epayco-client"
)

type EpaycoService struct {
	UrlBase          string
	EpaycoPublicKey  string
	EpaycoPrivateKey string
	Client           epaycoclient.HttpClient
}

func NewEpaycoService(UrlBase, EpaycoPublicKey, EpaycoPrivateKey string) *EpaycoService {
	return &EpaycoService{
		UrlBase:          UrlBase,
		EpaycoPublicKey:  EpaycoPublicKey,
		EpaycoPrivateKey: EpaycoPrivateKey,
		Client:           *epaycoclient.NewHttpClient(),
	}
}
