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
	USERS_GET_API_ENDPOINT    string = "https://api.intercom.io/users"
	USERS_POST_API_ENDPOINT   string = "https://api.intercom.io/users"
	USERS_DELETE_API_ENDPOINT string = "https://api.intercom.io/users"
) //const

func NewUserAvatar() *UserAvatar_t {
	return &UserAvatar_t{
		IntercomType: "avatar",
	} //UserAvatar_t
} //NewUserAavatar

/*
func NewCompanies() *Companies_t {
	return &Companies_t{
		IntercomType: "company.list",
	} //Avatar_t
} //NewAavatar
*/

func NewUserLocation() *UserLocation_t {
	return &UserLocation_t{
		IntercomType: "location_data",
	} //UserLocation_t
} //NewUserLocation

func NewUser(httpMethod string) *User_t {
	httpMethod = strings.TrimSpace(strings.ToUpper(httpMethod))

	if httpMethod == "POST" {
		return &User_t{
			CustomAttributes: make(map[string]interface{}),
			Companies:        make([]UserCompanyPost_t, 0),
		} //User_t
	} //if

	return &User_t{
		CustomAttributes: make(map[string]interface{}),
		Companies:        UserCompanyGet_t{},
	} //User_t
} //NewUser

func NewUserList() *UserList_t {
	return &UserList_t{
		IntercomType: "user.list",
	} //UserList_t
} //NewUserList

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

func (this *Intercom_t) PostUser(user *User_t) (err error) {
	var (
		req    *http.Request
		resp   *http.Response
		buffer = new(bytes.Buffer)
		client = new(http.Client)
	) //var

	// Encode user struct into JSON
	if err = json.NewEncoder(buffer).Encode(*user); err != nil {
		return err
	} //if

	// Create new POST request
	if req, err = http.NewRequest("POST", USERS_POST_API_ENDPOINT, buffer); err != nil {
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
} //PostUser
