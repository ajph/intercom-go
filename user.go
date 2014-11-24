package intercom

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
) //import

const (
	USERS_GET_API_ENDPOINT    string = "https://api.intercom.io/users/"
	USERS_POST_API_ENDPOINT   string = "https://api.intercom.io/users"
	USERS_DELETE_API_ENDPOINT string = "https://api.intercom.io/users?"
) //const

func NewAvatar() *Avatar_t {
	return &Avatar_t{
		"IntercomType": "avatar",
	} //Avatar_t
} //NewAavatar

func NewLocation() *Location_t {
	return &Location_t{
		"IntercomType": "location_data",
	} //Location_t
} //NewLocation

func NewUser() *User_t {
	return &User_t{
		"IntercomType": "user",
		"UpdatedAt":    time.Now().Unix(),
	} //User_t
} //NewUser

func NewUserList() *UserList_t {
	return &UserList_t{
		"IntercomType": "user.list",
	} //UserList_t
} //NewUserList

func (this *Intercom_t) PostUser(event Event_t) (err error) {
	var (
		req    *http.Request
		resp   *http.Response
		buffer = new(bytes.Buffer)
		client = new(http.Client)
	) //var

	// Encode event struct into JSON
	if err = json.NewEncoder(buffer).Encode(event); err != nil {
		return err
	} //if

	// Create new POST request
	if req, err = http.NewRequest("POST", USERS_POST_API_ENDPOINT, buffer); err != nil {
		return err
	} //if

	// Set authentication and headers
	req.SetBasicAuth(this.AppId, this.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	// Perform POST request
	if resp, err = client.Do(req); err != nil {
		return err
	} //if

	// Check reponse code and report any errors
	// Intercom sends back a 202 for valid requests
	if resp.StatusCode != 202 {
		return errors.New(resp.Status)
	} //if

	return err
} //PostUser
