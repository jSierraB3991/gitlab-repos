package brevoservice

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	brevorequest "gitlab.com/eliotandelon/brevo-email/infrastructure/brevo_request"
)

func (BrevoMailClient) get(urlBase, uri string, result interface{}, headers []brevorequest.HeaderRequest) error {

	req, err := http.NewRequest("GET", urlBase+uri, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")

	for _, v := range headers {
		req.Header.Add(v.Key, v.Value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 || resp.StatusCode == 201 {
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New(resp.Status)
}

func (BrevoMailClient) post(urlBase, uri string, jsonData []byte, result interface{}, headers []brevorequest.HeaderRequest) error {
	var bodyIoReader io.Reader
	if jsonData != nil {
		bodyIoReader = bytes.NewBuffer(jsonData)
	}
	req, err := http.NewRequest("POST", urlBase+uri, bodyIoReader)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")

	for _, v := range headers {
		req.Header.Add(v.Key, v.Value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 || resp.StatusCode == 201 {
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New(resp.Status)
}
