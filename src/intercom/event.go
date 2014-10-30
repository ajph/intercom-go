package intercom

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
) //import

const (
	API_ENDPOINT string = "https://api.intercom.io/events"
) //const

func (this *Event_t) SubmitEvent(appId, apiKey string) (err error) {
	var (
		req    *http.Request
		resp   *http.Response
		buffer = new(bytes.Buffer)
		client = new(http.Client)
	) //var

	// Encode event struct into JSON
	if err = json.NewEncoder(buffer).Encode(this); err != nil {
		return err
	} //if

	// Create new POST request
	if req, err = http.NewRequest("POST", API_ENDPOINT, buffer); err != nil {
		return err
	} //if

	// Set authentication and headers
	req.SetBasicAuth(appId, apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Perform POST request
	if resp, err = client.Do(req); err != nil {
		return err
	} //if

	// Check reponse code and report any errors
	if resp.StatusCode != 202 {
		return errors.New(resp.Status)
	} //if

	return err
} //SubmitEvent
