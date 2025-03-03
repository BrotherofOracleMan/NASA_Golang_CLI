package main

import "github.com/BrotherofOracleMan/NASA_GOLANG_CLI/cmd"

func main() {
	// Call the function to get the data from the API
	cmd.Execute()
}

/*
"explanation": "Do you see the ring?  If you look very closely at the center of the featured galaxy NGC 6505, a ring becomes evident. It is the gravity of NGC 6505, the nearby (z = 0.042) elliptical galaxy that you can easily see, that is magnifying and distorting the image of a distant galaxy into a complete circle. To create a complete Einstein ring there must be perfect alignment of the nearby galaxy's center and part of the background galaxy. Analysis of this ring and the multiple images of the background galaxy help to determine the mass and fraction of dark matter in NGC 6505's center, as well as uncover previously unseen details in the distorted galaxy.  The featured image was captured by ESA's Earth-orbiting Euclid telescope in 2023 and released earlier this month.",
"hdurl": "https://apod.nasa.gov/apod/image/2502/ClusterRing_Euclid_2665.jpg",
"media_type": "image",
"service_version": "v1",
"title": "Einstein Ring Surrounds Nearby Galaxy Center",
"url": "https://apod.nasa.gov/apod/image/2502/ClusterRing_Euclid_960.jpg"

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

func main() {

	resp, err := http.Get(apod_url + "?api_key=" + api_key)

	if err != nil {
		fmt.Println("Error Getting data", err)
	}
	defer resp.Body.Close()

	var apiResp ApiResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResp)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}
	fmt.Println(apiResp)

	fmt.Printf("Explanation : %s\n", apiResp.Explanation)
	fmt.Printf("Title : %s\n", apiResp.Title)
}
*/
