package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const baseUrl = "https://api.kraken.com"

type Kraken struct {
	client *http.Client
}

func New() *Kraken {
	return &Kraken{
		client: http.DefaultClient,
	}
}

func (k *Kraken) ServerTime() (time.Time, error) {
	var serverTime ServerTime
	err := k.get("/0/public/Time", &serverTime)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(serverTime.UnixTime, 0), nil
}

func (k *Kraken) Assets() (assets Assets, err error) {
	err = k.get("/0/public/Assets", &assets)
	return
}

func (k *Kraken) get(path string, retValue interface{}) error {
	resp, err := k.client.Get(baseUrl + path)
	if err != nil {
		return err
	}

	var response Response
	response.Result = retValue
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return err
	}

	if len(response.Errors) > 0 {
		return errors.New(fmt.Sprintf("got error: %v", response.Errors))
	}

	if response.Result == nil {
		return errors.New("unable to parse response")
	}
	return nil
}