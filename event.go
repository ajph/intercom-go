package event

import (
	"encoding/json"
	"error"
	"net/http"
	"net/url"
	"time"
) //import

const (
	API_ENDPOINT string = "https://api.intercom.io/events"
) //const

func (this *Event_t) SubmitEvent(apiKey string) (err error) {
	var (
		eventJSON []byte
		data      url.Values
		resp      *http.Response
	) //var

	if eventJSON, err = json.Marshal(this); err != nil {
		return err
	} //if

	this.EventName = "company-profile-view"
	this.CreatedAt = time.Now()
	this.UserId = "8"
	data.Set("-d", eventJSON)

	// Make request to API
	if resp, err = http.PostForm(API_ENDPOINT, data); err != nil {
		return err
	} //if

	if resp.StatusCode != 200 {
		return error.Error(resp.Status)
	} //if

	return err
} //SubmitEvent
