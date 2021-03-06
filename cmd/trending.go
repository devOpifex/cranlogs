package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/devOpifex/cranlogs/internal/color"
	"github.com/devOpifex/cranlogs/internal/data"
	"github.com/spf13/cobra"
)

var round bool

var trendingCmd = &cobra.Command{
	Use:   "trending",
	Short: "Get trending packages",
	Run: func(cmd *cobra.Command, args []string) {
		trending, err := data.GetTrending()

		if err != nil {
			color.PrintError(err.Error())
			os.Exit(0)
		}

		for _, v := range trending {
			increase := v.Increase
			if round {
				increase = strings.Split(v.Increase, ".")[0]
			}
			fmt.Printf("%v: %v%v%v\n", v.Package, color.Green, increase, color.Reset)
		}
	},
}

func init() {
	trendingCmd.Flags().BoolVarP(&round, "round", "r", true, "round increase value")
	rootCmd.AddCommand(trendingCmd)
}
