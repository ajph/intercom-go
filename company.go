package intercom

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
) //import

const (
	COMPANIES_API_ENDPOINT string = "https://api.intercom.io/companies"
) //const

func NewCompany() *Company_t {
	return &Company_t{
		CustomAttributes: make(map[string]interface{}),
	} //User_t
} //NewCompany

func (this *Intercom_t) PostCompany(com *Company_t) (err error) {
	var (
		req    *http.Request
		resp   *http.Response
		buffer = new(bytes.Buffer)
		client = new(http.Client)
	) //var

	// Encode user struct into JSON
	if err = json.NewEncoder(buffer).Encode(*com); err != nil {
		return err
	} //if

	// Create new POST request
	if req, err = http.NewRequest("POST", COMPANIES_API_ENDPOINT, buffer); err != nil {
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
} //SubmitCompany
