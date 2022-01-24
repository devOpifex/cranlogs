package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/devOpifex/cranlogs/internal/color"
	"github.com/devOpifex/cranlogs/internal/data"
	"github.com/spf13/cobra"
)

var dailyPeriod string

func daily(_ *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Fatal("no package specified")
	}
	p, err := data.NewPeriod(dailyPeriod)
	if err != nil {
		log.Fatal(err)
	}
	dailyDls, err := data.GetDaily(p, args)
	if err != nil {
		log.Fatalf("error getting daily downloads: %s", err)
	}

	if printJson {
		var out bytes.Buffer
		dailyBytes, err := json.Marshal(dailyDls)
		if err != nil {
			log.Fatalf("error marshalling daily downloads: %s", err)
		}
		json.Indent(&out, dailyBytes, "", "  ")
		fmt.Println(out.String())
	} else {
		for _, d := range dailyDls {
			fmt.Printf("Package: %v %v %v\n", color.Yellow, d.Package, color.Reset)
			for _, v := range d.Downloads {
				fmt.Printf("%v%v%v: %v\n", color.Cyan, v.Day, color.Reset, v.Downloads)
			}
		}
	}
}

func newDailyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "daily",
		Short: "Get daily downloads packages",
		Long: `
			cranlogs daily dplyr rlang
			// can also specify the custom period
			cranlogs daily --period=last-month dplyr rlang 
		`,
		Run: daily,
	}
	cmd.Flags().StringVarP(&dailyPeriod, "period", "p", "last-week", "time period")
	return cmd
}
