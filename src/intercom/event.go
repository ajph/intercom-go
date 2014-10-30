package intercom

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"
) //import

const (
	API_ENDPOINT string = "https://api.intercom.io/events"
) //const

func (this *Event_t) SubmitEvent(appId, apiKey string) (err error) {
	var (
		eventJSON []byte
		data      url.Values   = make(url.Values)
		client    *http.Client = &http.Client{}
		req       *http.Request
		resp      *http.Response
	) //var

	if eventJSON, err = json.Marshal(this); err != nil {
		return err
	} //if

	data.Set("-d", string(eventJSON))

	// Create new POST request
	if req, err = http.NewRequest("POST", API_ENDPOINT, strings.NewReader(data.Encode())); err != nil {
		return err
	} //if

	req.SetBasicAuth(appId, apiKey)
	req.Header.Set("Content-Type", "application/json")

	//req.PostForm = data

	//req.Header.Add("d", string(eventJSON))

	/*
		if resp, err = http.PostForm(API_ENDPOINT, data); err != nil {
			return err
		} //if
	*/

	log.Println(req)

	if resp, err = client.Do(req); err != nil {
		return err
	} //if

	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	} //if

	return err
} //SubmitEvent
