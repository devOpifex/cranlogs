package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/devOpifex/cranlogs/internal/color"
	"github.com/devOpifex/cranlogs/internal/data"
	"github.com/spf13/cobra"
)

var dailyPeriod string

func daily(_ *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Fatal("no package specified")
	}
	// TODO: this still has some issues in how its handled
	// as if a specific package fails, it will cause the entire process to
	// exit then. A different pattern could be to aggregate errors for
	// failed packages, then print out successful ones, then at the end
	// the error(s) for the failed ones
	for _, pkg := range args {
		daily, err := data.GetDaily(dailyPeriod, pkg)
		if err != nil {
			color.PrintError(err.Error())
			// given an error should not exit with a 0 exit code
			os.Exit(-1)
		}

		fmt.Printf("Package: %v %v %v\n", color.Yellow, daily.Package, color.Reset)
		for _, v := range daily.Downloads {
			fmt.Printf("%v%v%v: %v\n", color.Cyan, v.Day, color.Reset, v.Downloads)
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
