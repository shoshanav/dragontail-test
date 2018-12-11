package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
)


const (
	StatusOk = "OK"
	API_KEY = "AIzaSyBw8kNatPCD1U8V3hkDp6OcqI2SbzS7pTk"
)

type Response struct {
	Status  string   `json:"status"`
	Results []Result `json:"results"`
}

type Result struct {
	FormattedAddress  string	`json:"formatted_address"`
}

func GetAddress(lat string, len string) (string, error) {
	var address string
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?sensor=false&latlng=%s,%s&key=%s", lat, len, API_KEY)
	req, err := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return address, err
	}

	defer resp.Body.Close()

	response := Response{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return address, err
	}

	if response.Status != StatusOk {
		return address, fmt.Errorf(response.Status)
	}


	return response.Results[0].FormattedAddress, nil
}
