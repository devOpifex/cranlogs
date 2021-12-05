package cmd

import (
	"fmt"
	"os"

	"github.com/devOpifex/cranlogs/color"
	"github.com/devOpifex/cranlogs/data"
	"github.com/spf13/cobra"
)

var dailyPeriod string
var dailyPkg string

var dailyCmd = &cobra.Command{
	Use:   "daily",
	Short: "Get daily downloads of a package",
	Run: func(cmd *cobra.Command, args []string) {
		daily, err := data.GetDaily(dailyPeriod, dailyPkg)

		if err != nil {
			color.PrintError(err.Error())
			os.Exit(0)
		}

		fmt.Printf("Package: %v %v %v\n", color.Yellow, daily.Package, color.Reset)
		for _, v := range daily.Downloads {
			fmt.Printf("%v%v%v: %v\n", color.Cyan, v.Day, color.Reset, v.Downloads)
		}
	},
}

func init() {
	dailyCmd.Flags().StringVarP(&dailyPeriod, "period", "p", "last-week", "time period")
	dailyCmd.Flags().StringVarP(&dailyPkg, "package", "a", "dplyr", "package name")
	rootCmd.AddCommand(dailyCmd)
}
