package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type stationData struct {
	LastUpdated int `json:"last_updated"`
	TTL         int `json:ttl"`
	Data        struct {
		Stations []station `json:"stations"`
	} `json:"data"`
}

type station struct {
	ID                string `json:"station_id"`
	NumBikesAvailable int    `json:"num_bikes_available"`
	NumBikesDisabled  int    `json:"num_bike_disabled"`
	NumDocksAvailable int    `json:"num_docks_available"`
	NumDocksDisabled  int    `json:"num_docks_disabled"`
	IsInstalled       int    `json:"is_installed"`
	IsRenting         int    `json:"is_renting"`
	IsReturning       int    `json:"is_returning"`
	LastReported      int    `json:"last_reported"`
	HasAvailableKeys  bool   `json:"eightd_has_available_keys"`
}

func main() {
	const citiBikeURL = "https://gbfs.citibikenyc.com/gbfs/en/station_status.json"
	response, err := http.Get(citiBikeURL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var sd stationData

	if err := json.Unmarshal(body, &sd); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n\n", sd.Data.Stations[0])

}
