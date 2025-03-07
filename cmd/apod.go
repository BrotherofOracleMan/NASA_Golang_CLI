package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/BrotherofOracleMan/NASA_GOLANG_CLI/config"
	cobra "github.com/spf13/cobra"
)

//only flag that we need to provide is date

const format = "2006-01-02"

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
var download bool

var apod_cmd = &cobra.Command{
	Use:   "apod",
	Short: "Get the Astronomy Picture of the Day from NASA's astronomy picture of the day API",
	Run: func(cmd *cobra.Command, args []string) {
		flag.Parse()

		var apiResp ApiResponse
		apiURL, err := build_api_request(datetime)
		fmt.Println(apiURL)

		if err != nil {
			fmt.Println("Failed to build Api request. See error %w", err)
			return
		}
		resp, err := http.Get(apiURL)

		if err != nil {
			fmt.Println("Failed to Get API response. See ErrorL %w", err)
			return
		}
		defer resp.Body.Close()

		err = json.NewDecoder(resp.Body).Decode(&apiResp)

		if err != nil {
			fmt.Println("Failed to Decode API response. See Error %w", err)
			return
		}

		fmt.Printf("Explanation : %s\n", apiResp.Explanation)
		fmt.Printf("Title : %s\n", apiResp.Title)
		fmt.Printf("URL : %s\n", apiResp.Url)
		fmt.Printf("HD URL : %s\n", apiResp.Hdurl)
	},
}

func build_api_request(date string) (string, error) {
	baseUrl := config.GetAPODURL()

	parsedURL, err := url.Parse(baseUrl)

	if err != nil {
		return "", fmt.Errorf("an error occured while parsing the URL %w", err)
	}

	queryParams := parsedURL.Query()

	if date != "" {
		//YYYY-MM-DD is the required format
		parsed_date, err := time.Parse(format, datetime)
		if err != nil {
			return "", fmt.Errorf("an Error occured when parsing the date in to YYYY-MM-DD: %w", err)
		}
		formatted_date := parsed_date.Format(format)
		queryParams.Set("date", formatted_date)
	}
	queryParams.Set("api_key", config.GetAPIKey())
	parsedURL.RawQuery = queryParams.Encode()
	return parsedURL.String(), nil

}

func init() {
	apod_cmd.Flags().StringVarP(&datetime, "date", "d", "", "Date of the picture you want to see")
	apod_cmd.Flags().BoolVarP(&download, "download", "D", false, "Download the picture of the day")
	rootCmd.AddCommand(apod_cmd)
	config.Loadconfig()
}
