package api

type Response struct {
	Errors []string    `json:"error"`
	Result interface{} `json:"result"`
}

type ServerTime struct {
	UnixTime int64 `json:"unixtime"`
}

type Asset struct {
	AltName string `json:"altname"`
	Class string `json:"aclass"`
	Decimals int `json:"decimals"`
	DisplayDecimals int `json:"display_decimals"`
}

type Assets map[string]Asset

