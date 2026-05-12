package osint

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
)

var InvalidIPRepresentation = fmt.Errorf("Entered IP is not a Valid IP")
var FetchInfoError = fmt.Errorf("Error fetching Info from API")
var InvalidorEmptyResponse = fmt.Errorf("Error Reading Response Body: Either invalid or Empty")
var FailedParsingIntoJSON = fmt.Errorf("Ran into Error parsing the response into JSON")

type IPLookup struct {
	Country     string  `json:"country,omitempty"`
	CountryCode string  `json:"countryCode,omitempty"`
	Region      string  `json:"region,omitempty"`
	RegionName  string  `json:"regionName,omitempty"`
	City        string  `json:"city,omitempty"`
	Zip         string  `json:"zip,omitempty"`
	Latitude    float32 `json:"latitude,omitempty"`
	Longitude   float32 `json:"longitude,omitempty"`
	Timezone    string  `json:"timezone,omitempty"`
	ISP         string  `json:"isp,omitempty"`
	ORG         string  `json:"org,omitempty"`
	AS          string  `json:"as,omitempty"`
	Status      string  `json:"status,omitempty"`
	Message     string  `json:"message,omitempty"`
}

func IPLookupRun(addr string) (*IPLookup, error) {
	if net.ParseIP(addr) == nil {
		return nil, InvalidIPRepresentation
	}

	response, err := http.Get(fmt.Sprintf("http://ip-api.com/json/%s", addr))
	if err != nil {
		return nil, FetchInfoError
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, InvalidorEmptyResponse
	}

	var ipRes IPLookup
	err = json.Unmarshal(body, &ipRes)
	if err != nil {
		// TODO: Handle error responses
		return nil, FailedParsingIntoJSON
	}

	return &ipRes, nil
}
