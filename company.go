package intercom

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
) //import

const (
	COMPANIES_API_ENDPOINT string = "https://api.intercom.io/companies"
) //const

func NewCompany(httpMethod string) *Company_t {
	httpMethod = strings.TrimSpace(strings.ToUpper(httpMethod))

	if httpMethod == "POST" {
		return &Company_t{
			CustomAttributes: make(map[string]interface{}),
			Plan:             "",
		} //NewCompany
	} //if

	return &Company_t{
		CustomAttributes: make(map[string]interface{}),
		Plan:             CompanyPlan_t{},
	} //NewCompany
} //NewCompany

func (this *Intercom_t) GetCompany(com *Company_t) (err error) {
	var (
		intercomUrl string = USERS_GET_API_ENDPOINT
		req         *http.Request
		resp        *http.Response
		client      = new(http.Client)
	) //var

	if com.CompanyId != "" {
		intercomUrl += "?company_id=" + com.CompanyId
	} else if com.Name != "" {
		intercomUrl += "?name=" + url.QueryEscape(com.Name)
	} //else if

	// Create new GET request
	if req, err = http.NewRequest("GET", intercomUrl, nil); err != nil {
		return err
	} //if

	// Set authentication and headers
	req.SetBasicAuth(this.AppId, this.ApiKey)
	req.Header.Set("Accept", "application/json")

	// Perform GET request
	if resp, err = client.Do(req); err != nil {
		return err
	} //if
	defer resp.Body.Close()

	// Check reponse code and report any errors
	// Intercom sends back a 200 for valid requests
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	} //if

	// Decode JSON response into User_t struct
	if err = json.NewDecoder(resp.Body).Decode(com); err != nil {
		return err
	} //if

	return nil
} //GetUser

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
} //PostCompany
