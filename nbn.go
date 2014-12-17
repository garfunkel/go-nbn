// Package nbn implements functions to analyse the Australian National Broadband Network (NBN) rollout.
package nbn

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Service URLs.
const (
	NBNStatusRefererURL = "http://www.nbnco.com.au/when-do-i-get-it/rollout-map.html"
	NBNStatusURL        = "http://www.nbnco.com.au/api/map/search.html"
)

// Info is the structure representing NBN rollout information.
type Info struct {
	RolloutPlans []struct {
		ID          string `json:"id"`
		ServiceType string `json:"serviceType"`
		Suburb      string `json:"suburb"`
		Postcode    string `json:"postcode"`
		State       string `json:"state"`
		Status      string `json:"status"`
		FirstDate   string `json:"firstDate"`
		LastDate    string `json:"lastDate"`
	} `json:"rolloutPlans"`
	ServingArea struct {
		ID                        string `json:"id"`
		IsDisconnectionDatePassed bool   `json:"isDisconnectionDatePassed"`
		IsFrustratedMduAddress    bool   `json:"isFrustratedMduAddress"`
		ServiceStatus             string `json:"serviceStatus"`
		DisconnectionDate         string `json:"disconnectionDate"`
		Description               string `json:"description"`
		IsServiceClassZeroAddress bool   `json:"isServiceClassZeroAddress"`
		CSAID                     string `json:"csaId"`
		ServiceCategory           string `json:"serviceCategory"`
		ServiceType               string `json:"serviceType"`
	} `json:"servingArea"`
	// Fibre Service Area Modules
	FSAMs []struct {
		ID          string `json:"id"`
		ServiceType string `json:"serviceType"`
		State       string `json:"state"`
		Status      string `json:"status"`
		FirstDate   string `json:"firstDate"`
		LastDate    string `json:"lastDate"`
	} `json:"fsams"`
}

// request processes HTTP requests.
func request(params string) (info *Info, err error) {
	url := fmt.Sprintf("%v?%v", NBNStatusURL, params)
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return
	}

	req.Header.Add("Referer", NBNStatusRefererURL)

	response, err := client.Do(req)

	if err != nil {
		return
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return
	}

	info = new(Info)
	err = json.Unmarshal(data, info)

	return
}

// RolloutInfo accepts latlng params and returns NBN rollout info.
func RolloutInfo(latitude, longitude float64) (info *Info, err error) {
	params := fmt.Sprintf("lat=%v&lng=%v", latitude, longitude)

	return request(params)
}
