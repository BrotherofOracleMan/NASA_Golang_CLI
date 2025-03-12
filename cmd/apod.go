package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"

	"github.com/BrotherofOracleMan/NASA_GOLANG_CLI/config"
	cobra "github.com/spf13/cobra"
)

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
var showImage bool
var image_uri fyne.URI

var apod_cmd = &cobra.Command{
	Use:   "apod",
	Short: "Get the Astronomy Picture of the Day from NASA's astronomy picture of the day API",
	Run: func(cmd *cobra.Command, args []string) {
		flag.Parse()
		var apiResp ApiResponse

		apiURL, err := build_api_request(datetime)
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

		if download {
			file_name := config.GetApodDefaultImageName()
			folder_name := config.GetApodDefaultDownloadDirectory()
			if len(args) > 0 {
				folder_name = args[0]
			}
			if len(args) > 1 {
				file_name = args[1]
			}
			err = downloadImage(apiResp.Url, folder_name, file_name)
			if err != nil {
				fmt.Printf("Download Image Called. See Error %v\n", err)
				return
			}
		}

		if showImage {
			a := app.New()
			w := a.NewWindow(apiResp.Title)

			if image_uri, err = storage.ParseURI(apiResp.Url); err != nil {
				w.SetContent(widget.NewLabel(apiResp.Url))
				w.ShowAndRun()
				return
			}

			image := canvas.NewImageFromURI(image_uri)
			image.FillMode = canvas.ImageFillOriginal

			explanation := widget.NewLabel(apiResp.Explanation)
			explanation.Wrapping = fyne.TextWrapWord
			explanation.Alignment = fyne.TextAlignLeading

			content := container.New(layout.NewVBoxLayout(), image, explanation)
			w.SetContent(content)
			w.ShowAndRun()
		}
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
		formatted_date, err := parse_date(date)
		if err != nil {
			return "", fmt.Errorf("an error occured while parsing the date %w", err)
		}
		queryParams.Set("date", formatted_date)
	}
	queryParams.Set("api_key", config.GetAPIKey())
	parsedURL.RawQuery = queryParams.Encode()
	return parsedURL.String(), nil
}

func downloadImage(url, folder, fileName string) error {
	// Perform the HTTP request
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download image: %w", err)
	}
	defer resp.Body.Close()

	// Check if the HTTP status code is OK
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get OK HTTP response: %s", resp.Status)
	}

	// Ensure the directory exists (create if it doesn't)
	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", folder, err)
	}

	// Create the file in the specified directory
	filePath := path.Join(folder, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filePath, err)
	}
	defer file.Close()

	// Copy the image content from the response body to the file
	if _, err := io.Copy(file, resp.Body); err != nil {
		return fmt.Errorf("failed to save image to %s: %w", filePath, err)
	}

	return nil
}

func parse_date(date string) (string, error) {
	var parsedDate time.Time
	var err error
	const desired_format = "2006-01-02"

	known_formats := []string{
		"2006-01-02",
		"2006/01/02",
		"2006.01.02",
		"01-02-2006",
		"01/02/2006",
		"01.02.2006",
		"March 2, 2006",
		"2006 March 2",
	}

	for _, format := range known_formats {
		parsedDate, err = time.Parse(format, datetime)
		if err == nil {
			break
		}
	}
	if err != nil {
		fmt.Println(datetime)
		return "", fmt.Errorf("an Error occured when parsing the date in to YYYY-MM-DD: %w", err)
	}
	return parsedDate.Format(desired_format), nil
}

func init() {
	apod_cmd.Flags().StringVarP(&datetime, "date", "d", "", "Date of the picture you want to see")
	apod_cmd.Flags().BoolVarP(&download, "download", "D", false, "Download the picture of the day")
	apod_cmd.Flags().BoolVarP(&showImage, "show", "s", true, "Show the image of the day")
	rootCmd.AddCommand(apod_cmd)
	config.Loadconfig()
}
