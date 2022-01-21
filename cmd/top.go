package cmd

import (
	"fmt"
	"os"

	"github.com/devOpifex/cranlogs/internal/color"
	"github.com/devOpifex/cranlogs/internal/data"
	"github.com/spf13/cobra"
)

var topPeriod string
var topN int

var topCmd = &cobra.Command{
	Use:   "top",
	Short: "Get most downloaded packages",
	Run: func(cmd *cobra.Command, args []string) {
		top, err := data.GetTop(topPeriod, topN)

		if err != nil {
			color.PrintError(err.Error())
			os.Exit(0)
		}

		for _, v := range top.Downloads {
			fmt.Printf("%v%v%v: %v\n", color.Cyan, v.Package, color.Reset, v.Downloads)
		}
	},
}

func init() {
	topCmd.Flags().StringVarP(&topPeriod, "period", "p", "last-week", "time period")
	topCmd.Flags().IntVarP(&topN, "max", "n", 5, "number of packages")
	rootCmd.AddCommand(topCmd)
}
