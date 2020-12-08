package kraken

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
	err := k.get("/0/public/Time", "", &serverTime)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(serverTime.UnixTime, 0), nil
}

func (k *Kraken) Assets(opts AssetOptions) (assets Assets, err error) {
	err = k.get("/0/public/Assets", opts.QueryString(), &assets)
	return
}

func (k *Kraken) get(path string, queryString string, retValue interface{}) error {
	req, err := http.NewRequest(http.MethodGet, baseUrl+path, nil)
	if err != nil {
		return err
	}
	if queryString != "" {
		req.URL.Query()
		req.URL.RawQuery = queryString
	}
	resp, err := k.client.Do(req)
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
		return fmt.Errorf("got error: %v", response.Errors)
	}

	if response.Result == nil {
		return errors.New("unable to parse response")
	}
	return nil
}
