package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BrotherofOracleMan/NASA_GOLANG_CLI/config"
	cobra "github.com/spf13/cobra"
)

//only flag that we need to provide is date

type ApiResponse struct {
	Copyright       string `json:"copyright"`
	Date            string `json:"date"`
	Explanation     string `json:"explanation"`
	Hdurl           string `json:"hdurl"`
	Media_type      string `json:"media_type"`
	Service_version string `json:"service_version"`
	Title           string `json:"title"`
	Url             string `json:"url"`
}

var datetime string

var apod_cmd = &cobra.Command{
	Use:   "apod",
	Short: "Get the Astronomy Picture of the Day from NASA's astronomy picture of the day API",
	Run: func(cmd *cobra.Command, args []string) {

		resp, err := http.Get(config.GetAPODURL() + "?api_key=" + config.GetAPIKey())

		if err != nil {
			fmt.Println("Unable to get HTTP response", err)
		}
		defer resp.Body.Close()

		var apiResp ApiResponse
		err = json.NewDecoder(resp.Body).Decode(&apiResp)

		if err != nil {
			fmt.Println("Error decoding JSON:", err)
		}

		fmt.Printf("Explanation : %s\n", apiResp.Explanation)
		fmt.Printf("Title : %s\n", apiResp.Title)
		fmt.Printf("URL : %s\n", apiResp.Url)
		fmt.Printf("HD URL : %s\n", apiResp.Hdurl)
	},
}

func init() {
	apod_cmd.Flags().StringVarP(&datetime, "date", "d", "", "Date of the picture you want to see")
	rootCmd.AddCommand(apod_cmd)
}
