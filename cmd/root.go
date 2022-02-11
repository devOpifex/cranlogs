package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var printJson bool
var rootCmd = &cobra.Command{
	Use:   "Cranlogs",
	Short: "Access the cranlogs API",
	Long:  `Access the cranlogs API from the command line`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute(version string) {
	rootCmd.Version = version
	rootCmd.SetVersionTemplate(`{{printf "%s\n" .Version}}`)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(newDailyCmd())
	rootCmd.PersistentFlags().BoolVar(&printJson, "json", false, "output in json format")
}
