package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Loadconfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("NO configuration file found")
	}
	fmt.Println("Config file used:", viper.ConfigFileUsed())
}

func GetAPIKey() string {
	return viper.GetString("api_key")
}

func GetAPODURL() string {
	return viper.GetString("apod_url")
}

func GetEarthDateURL() string {
	return viper.GetString("earth_date_url")
}

func GetMarsRoverURL() string {
	return viper.GetString("mars_rover_url")
}

func GetApodDefaultDownloadDirectory() string {
	return viper.GetString("default_apod_download_folder")
}

func GetApodDefaultImageName() string {
	return viper.GetString("default_apod_download_file_name")
}

func GetEarthSatelliteUrl() string {
	return viper.GetString("earth_satellite_url")
}
