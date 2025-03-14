package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"

	"github.com/BrotherofOracleMan/NASA_GOLANG_CLI/config"
	"github.com/spf13/cobra"
)

type resource struct {
	Dataset string `json:"dataset"`
	Planet  string `json:"planet"`
}

// Define the e_sat_api_response struct
type e_sat_api_response struct {
	Date           string   `json:"date"`            // Timestamp with precision
	ID             string   `json:"id"`              // ID as string
	Resource       resource `json:"resource"`        // Nested resource struct
	ServiceVersion string   `json:"service_version"` // Service version as string
	URL            string   `json:"url"`             // URL as string
}

var lon float64
var lat float64
var dim float64
var e_sat_datetime string

var e_st_cmd = &cobra.Command{
	Use:   "e_sat",
	Short: "Get Earth satellite imagery from NASA's Earth Land Sat Imagery",
	Run: func(cmd *cobra.Command, args []string) {
		flag.Parse()
		var e_sat_response e_sat_api_response

		apiUrl, err := build_esat_api_request(lon, lat, e_sat_datetime)

		if err != nil {
			fmt.Println("An error occured while sending the request", err)
		}

		resp, err := http.Get(apiUrl)

		if err != nil {
			fmt.Println("An error occured while sending the request", err)
		}
		defer resp.Body.Close()

		if err = json.NewDecoder(resp.Body).Decode(&e_sat_response); err != nil {
			fmt.Println("An error occured while decoding the response", err)
		}
		fmt.Println("The URL for the Earth Satellite Imagery is: ", e_sat_response.URL)
		fmt.Println("The date of the image is: ", e_sat_response.Date)
		fmt.Println("The API called used is: ", apiUrl)
	},
}

func build_esat_api_request(lon float64, lat float64, date string) (string, error) {
	baseUrl := config.GetEarthSatelliteUrl()
	parsedURL, err := url.Parse(baseUrl)
	queryParams := parsedURL.Query()

	if err != nil {
		return "", fmt.Errorf("an error occured while parsing the URL %w", err)
	}

	if e_sat_datetime != "" {
		formatted_date, err := parse_date(date)
		if err != nil {
			return "", fmt.Errorf("an error occured while parsing the date. See error: %w", err)
		}
		queryParams.Add("date", formatted_date)
	}

	queryParams.Add("dim", fmt.Sprintf("%f", dim))
	queryParams.Add("lon", fmt.Sprintf("%f", lon))
	queryParams.Add("lat", fmt.Sprintf("%f", lat))
	queryParams.Set("api_key", config.GetAPIKey())
	parsedURL.RawQuery = queryParams.Encode()
	return parsedURL.String(), nil
}

func init() {
	e_st_cmd.Flags().Float64VarP(&lon, "long", "l", 0.0, "Longitude of the location")
	e_st_cmd.Flags().Float64VarP(&lat, "lat", "t", 0.0, "Latitude of the location")
	e_st_cmd.Flags().Float64VarP(&dim, "dim", "i", 0.025, "Dimension of the image")
	e_st_cmd.Flags().StringVarP(&e_sat_datetime, "date", "d", "", "Date of the location you want to look for")
	rootCmd.AddCommand(e_st_cmd)
	config.Loadconfig()
	e_st_cmd.MarkFlagRequired("long")
	e_st_cmd.MarkFlagRequired("lat")
}
