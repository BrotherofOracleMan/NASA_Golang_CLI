package cmd

import (
	"fmt"

	"github.com/BrotherofOracleMan/NASA_GOLANG_CLI/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rootcmd",
	Short: "NASA CLI application for getting pictures of the Stars and planets",
	Run: func(cmd *cobra.Command, args []string) {
		// Call the function to get the data from the API{
		config.Loadconfig()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
