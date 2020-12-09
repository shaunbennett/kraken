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

type AssetPairOptions struct {
	Info  string
	Pairs []string
}

func (a AssetPairOptions) QueryString() string {
	v := url.Values{}
	if a.Info != "" {
		v.Add("info", a.Info)
	}
	if len(a.Pairs) > 0 {
		v.Add("pair", strings.Join(a.Pairs, ","))
	}
	return v.Encode()
}

// A tradable asset pair
type AssetPair struct {
	// Alternative pair name
	AltName string `json:"altname"`
	// WebSocket pair name (if available)
	WebSocketName string `json:"wsname"`
	// Asset class of the base component
	ClassBase string `json:"aclass_base"`
	// Asset id of the base component
	Base string `json:"base"`
	// Asset class of the quote component
	ClassQuote string `json:"aclass_quote"`
	// Asset id of the quote component
	Quote string `json:"quote"`
	// Volume lot size
	Lot string `json:"lot"`
	// Scaling decimal places for the pair
	PairDecimals int `json:"pair_decimals"`
	// Scaling decimal places for volume
	LotDecimals int `json:"lot_decimals"`
	// Amount to multiply the lot volume by to get the currency volume
	LotMultiplier int `json:"lot_multiplier"`
	// Array of possible leverage amounts when buying
	LeverageBuy []int `json:"leverage_buy"`
	// Array of possible leverage amounts when selling
	LeverageSell []int `json:"leverage_sell"`
	// Fee schedule in the form [[volume, percent fee], ...]
	Fees [][]float32 `json:"fees"`
	// Maker fee schedule in the form [[volume, percent fee], ...]
	FeesMaker [][]float32 `json:"fees_maker"`
	// Volume discount currency
	FeeVolumeCurrency string `json:"fee_volume_currency"`
	// Margin call level
	MarginCall int `json:"margin_call"`
	// Stop/liquidation margin level
	MarginStop int `json:"margin_stop"`
	// Minimum order volume for the pair
	OrderMin string `json:"ordermin"`
}

// A map of asset pair ids to their data
type AssetPairs map[string]AssetPair
