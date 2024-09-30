package epaycoclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	epaycorequest "gitlab.com/eliotandelon/epayco-pays/infrastructure/epayco-request"
	epaycoresponse "gitlab.com/eliotandelon/epayco-pays/infrastructure/epayco-response"
)

type HttpClient struct {
}

func NewHttpClient() *HttpClient {
	return &HttpClient{}
}

func (HttpClient) Post(urlBase, uri string, jsonData []byte, result interface{}, headers []epaycorequest.HeaderRequest) error {
	// Crear una nueva solicitud HTTP
	var bodyIoReader io.Reader
	if jsonData != nil {
		bodyIoReader = bytes.NewBuffer(jsonData)
	}
	req, err := http.NewRequest("POST", urlBase+uri, bodyIoReader)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	if headers != nil {
		for _, v := range headers {
			req.Header.Add(v.Key, v.Value)
		}
	}

	// Enviar la solicitud
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 || resp.StatusCode < 300 {
		//body, _ := io.ReadAll(resp.Body)

		//fmt.Println(resp)
		//fmt.Println(string(body))
		return json.NewDecoder(resp.Body).Decode(&result)
	} else {
		var responseEror epaycoresponse.InvalidRequestPayResponse
		err = json.NewDecoder(resp.Body).Decode(&responseEror)
		if err != nil {
			return err
		}
		log.Println(responseEror)
		return errors.New(responseEror.MessageResponse)
	}

}
