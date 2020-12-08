package kraken

import (
	"net/url"
	"strings"
)

type Response struct {
	Errors []string    `json:"error"`
	Result interface{} `json:"result"`
}

type ServerTime struct {
	UnixTime int64 `json:"unixtime"`
}

// TODO: Can't find documentation on other asset classes, do they exist?
const AssetCurrency = "currency"

// AssetOptions represent the options that can be used when retrieving asset data
type AssetOptions struct {
	// The info for assets to retrieve from the server. Defaults to all info.
	Info string
	// The asset class to retrieve data for. Defaults to "currency".
	AssetClass string
	// Assets to retrieve data for. Defaults to all assets if none are specified.
	Assets []string
}

func (a AssetOptions) QueryString() string {
	v := url.Values{}
	if a.Info != "" {
		v.Add("info", a.Info)
	}
	if a.AssetClass != "" {
		v.Add("aclass", a.AssetClass)
	}
	if len(a.Assets) > 0 {
		v.Add("asset", strings.Join(a.Assets, ","))
	}
	return v.Encode()
}

type Asset struct {
	AltName         string `json:"altname"`
	Class           string `json:"aclass"`
	Decimals        int    `json:"decimals"`
	DisplayDecimals int    `json:"display_decimals"`
}

type Assets map[string]Asset
