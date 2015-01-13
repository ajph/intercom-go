package intercom

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
) //import

const (
	TAG_POST_API_ENDPOINT string = "https://api.intercom.io/tags"
) //const

/*
func (this *Intercom_t) DeleteUser(user *User_t) (err error) {
	var (
		intercomUrl string = USERS_DELETE_API_ENDPOINT
		req         *http.Request
		resp        *http.Response
		client      = new(http.Client)
	) //var

	if user.UserId != "" {
		intercomUrl += "?user_id=" + user.UserId
	} else if user.Email != "" {
		intercomUrl += "?email=" + url.QueryEscape(user.Email)
	} //else if

	// Create new GET request
	if req, err = http.NewRequest("DELETE", intercomUrl, nil); err != nil {
		return err
	} //if

	// Set authentication and headers
	req.SetBasicAuth(this.AppId, this.ApiKey)
	req.Header.Set("Accept", "application/json")

	// Perform DELETE request
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
	// Full User objecs are returned on DELETE
	if err = json.NewDecoder(resp.Body).Decode(user); err != nil {
		return err
	} //if

	return nil
} //DeleteUser

func (this *Intercom_t) GetUser(user *User_t) (err error) {
	var (
		intercomUrl string = USERS_GET_API_ENDPOINT
		req         *http.Request
		resp        *http.Response
		client      = new(http.Client)
	) //var

	if user.UserId != "" {
		intercomUrl += "?user_id=" + user.UserId
	} else if user.Email != "" {
		intercomUrl += "?email=" + url.QueryEscape(user.Email)
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
	if err = json.NewDecoder(resp.Body).Decode(user); err != nil {
		return err
	} //if

	return nil
} //GetUser
*/

func (this *Intercom_t) PostTag(tag Tag_t) (err error) {
	var (
		req    *http.Request
		resp   *http.Response
		buffer = new(bytes.Buffer)
		client = new(http.Client)
	) //var

	// Encode tag struct into JSON
	if err = json.NewEncoder(buffer).Encode(tag); err != nil {
		return err
	} //if

	// Create new POST request
	if req, err = http.NewRequest("POST", TAG_POST_API_ENDPOINT, buffer); err != nil {
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
	// Intercom sends back a 202 for valid requests
	if resp.StatusCode != 202 {
		return errors.New(resp.Status)
	} //if

	return err
} //PostTag
