package main

import (
	"fmt"
	"log"

	"github.com/devOpifex/go-cranlogs/data"
)

func main() {
	api := data.API{
		URL: "https://cranlogs.r-pkg.org/",
	}

	top, err := api.Top("last-day", 3)

	if err != nil {
		log.Fatal(err)
	}

	for i, v := range top.Downloads {
		fmt.Printf("%v. %v: %v\n", i, v.Package, v.Downloads)
	}

	trending, err := api.Trending()

	if err != nil {
		log.Fatal(err)
	}

	for i, v := range trending {
		fmt.Printf("%v. %v: %v\n", i, v.Package, v.Increase)
	}
}
