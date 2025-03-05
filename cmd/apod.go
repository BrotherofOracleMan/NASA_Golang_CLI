package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
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

		//check if the date flag is set
		flag.Parse()

		if datetime != "" {
			//YYYY-MM-DD is the required format
			test_parse, err := time.Parse(format, datetime)
			formatted_date := test_parse.Format(format)
			if err != nil {
				fmt.Println("An Error occured when parsing the date in to YYYY-MM-DD", err)
			}
			fmt.Printf("Date : %s\n", test_parse)
			fmt.Printf("Formatted Date : %s\n", formatted_date)
			fmt.Println(test_parse)
		}
		//check if the download flag is set
		if download {
			fmt.Println("Downloading the image")
		}
	},
}

func build_api_request(date string, download bool) string {
	return config.GetAPODURL() + "?api_key=" + config.GetAPIKey()
	/*
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
	*/
}

func init() {
	apod_cmd.Flags().StringVarP(&datetime, "date", "d", "", "Date of the picture you want to see")
	apod_cmd.Flags().BoolVarP(&download, "download", "D", false, "Download the picture of the day")
	rootCmd.AddCommand(apod_cmd)
	config.Loadconfig()
}
