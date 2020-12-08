package kraken

import (
	"net/url"
	"strings"
)

// A Response represents the json structure that the Kraken API outputs. All API outputs are wrapped.
type Response struct {
	Errors []string    `json:"error"`
	Result interface{} `json:"result"`
}

// Kraken: /0/public/Time result
type ServerTime struct {
	UnixTime int64 `json:"unixtime"`
}

// The asset class for currencies
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

// An Asset inside of Kraken's API
type Asset struct {
	// Alternative name for the asset
	AltName string `json:"altname"`
	// The asset class
	Class string `json:"aclass"`
	// Number of decimals used for record keeping
	Decimals int `json:"decimals"`
	// Number of decimals used for output display
	DisplayDecimals int `json:"display_decimals"`
}

// A map of asset names to their data
type Assets map[string]Asset
