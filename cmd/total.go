package cmd

import (
	"fmt"
	"os"

	"github.com/devOpifex/cranlogs/color"
	"github.com/devOpifex/cranlogs/data"
	"github.com/spf13/cobra"
)

var totalPeriod string
var totalPkg string

var totalCmd = &cobra.Command{
	Use:   "total",
	Short: "Get total downloads of a package",
	Run: func(cmd *cobra.Command, args []string) {
		total, err := data.GetTotal(totalPeriod, totalPkg)

		if err != nil {
			color.PrintError(err.Error())
			os.Exit(0)
		}

		fmt.Printf("%v %v %v:%v\n", color.Yellow, total.Package, color.Reset, total.Downloads)
	},
}

func init() {
	totalCmd.Flags().StringVarP(&totalPeriod, "period", "p", "last-week", "time period")
	totalCmd.Flags().StringVarP(&totalPkg, "package", "a", "dplyr", "package name")
	rootCmd.AddCommand(totalCmd)
}
