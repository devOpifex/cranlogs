package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/devOpifex/cranlogs/color"
	"github.com/devOpifex/cranlogs/data"
	"github.com/spf13/cobra"
)

func grandTotalPeriod() string {
	currentTime := time.Now()
	to := currentTime.Format("2006-01-02")
	return "1900-01-01:" + to
}

var grandTotalPkg string

var grandTotalCmd = &cobra.Command{
	Use:   "grand-total",
	Short: "Get total downloads of a package over its lifetime",
	Run: func(cmd *cobra.Command, args []string) {

		total, err := data.GetTotal(grandTotalPeriod(), grandTotalPkg)

		if err != nil {
			color.PrintError(err.Error())
			os.Exit(0)
		}

		fmt.Printf("%v %v %v:%v\n", color.Yellow, total.Package, color.Reset, total.Downloads)
	},
}

func init() {
	grandTotalCmd.Flags().StringVarP(&grandTotalPkg, "package", "a", "dplyr", "package name")
	rootCmd.AddCommand(grandTotalCmd)
}
