package intercom

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
) //import

const (
	MESSAGE_POST_API_ENDPOINT string = "https://api.intercom.io/messages"
) //const

func (this *Intercom_t) PostMessage(cv InitiateConversation_t) (err error) {
	var (
		req    *http.Request
		resp   *http.Response
		buffer = new(bytes.Buffer)
		client = new(http.Client)
	) //var

	// Encode tag struct into JSON
	if err = json.NewEncoder(buffer).Encode(cv); err != nil {
		return err
	} //if

	// Create new POST request
	if req, err = http.NewRequest("POST", MESSAGE_POST_API_ENDPOINT, buffer); err != nil {
		return err
	} //if

	// Set authentication and headers
	req.SetBasicAuth(this.AppId, this.ApiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	// Perform POST request
	if resp, err = client.Do(req); err != nil {
		return err
	} //if
	defer resp.Body.Close()

	// Check reponse code and report any errors
	// Intercom sends back a 200 for valid requests
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	} //if

	return err
} //PostMessage
