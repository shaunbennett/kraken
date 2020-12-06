package main

import (
	"fmt"
	"github.com/shaunbennett/kraken/pkg/api"
)

func main() {
	api := api.New()
	assets, err := api.Assets()
	if err != nil {
		fmt.Printf("Got error: %v\n", err)
	}

	fmt.Printf("got assets time %+v!\n", assets["BCH"])
}
