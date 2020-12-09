package main

import (
	"fmt"

	"github.com/shaunbennett/kraken"
)

func main() {
	api := kraken.New()
	assets, err := api.Assets(kraken.AssetOptions{Assets: []string{"ADA", "BCH"}})
	if err != nil {
		fmt.Printf("Got error: %v\n", err)
		return
	}

	fmt.Printf("got assets %+v!\n", assets)

	pairs, err := api.AssetPairs(kraken.AssetPairOptions{})
	if err != nil {
		fmt.Printf("Got error: %v\n", err)
		return
	}

	fmt.Printf("got assets pairs %+v!\n", pairs)
}
