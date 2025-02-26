package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/viper"
)

/*
"explanation": "Do you see the ring?  If you look very closely at the center of the featured galaxy NGC 6505, a ring becomes evident. It is the gravity of NGC 6505, the nearby (z = 0.042) elliptical galaxy that you can easily see, that is magnifying and distorting the image of a distant galaxy into a complete circle. To create a complete Einstein ring there must be perfect alignment of the nearby galaxy's center and part of the background galaxy. Analysis of this ring and the multiple images of the background galaxy help to determine the mass and fraction of dark matter in NGC 6505's center, as well as uncover previously unseen details in the distorted galaxy.  The featured image was captured by ESA's Earth-orbiting Euclid telescope in 2023 and released earlier this month.",
"hdurl": "https://apod.nasa.gov/apod/image/2502/ClusterRing_Euclid_2665.jpg",
"media_type": "image",
"service_version": "v1",
"title": "Einstein Ring Surrounds Nearby Galaxy Center",
"url": "https://apod.nasa.gov/apod/image/2502/ClusterRing_Euclid_960.jpg"
*/
type ApiResponse struct {
}

func main() {
	viper.SetConfigName("configuration")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err)
	}

	api_key := viper.GetString("api_key")
	fmt.Println("API Key:", api_key)

	apod_url := viper.GetString("apod_url")
	fmt.Println("APOD URL:", apod_url)

	resp, err := http.Get(apod_url + "?api_key=" + api_key)

	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Response Data")
	fmt.Println("Response Body:", string(body))
}
