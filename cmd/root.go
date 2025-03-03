package cmd

import (
	"fmt"

	"github.com/BrotherofOracleMan/NASA_GOLANG_CLI/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "NASA CLI application ",
	Short: "NASA CLI application for getting pictures of the Stars and planets",
	Run: func(cmd *cobra.Command, args []string) {
		// Call the function to get the data from the API{
		config.Loadconfig()

		apod_url := config.GetAPODURL()
		earth_date_url := config.GetEarthDateURL()
		mars_rover_url := config.GetMarsRoverURL()

		fmt.Printf("APOD URL: %s\n", apod_url)
		fmt.Printf("Earth Date URL: %s\n", earth_date_url)
		fmt.Printf("Mars Rover URL: %s\n", mars_rover_url)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
