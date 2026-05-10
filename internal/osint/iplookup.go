package osint

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
)

var InvalidIPRepresentation = fmt.Errorf("Invalid IP representation")

type IPLookup struct {
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Latitude    float32 `json:"latitude"`
	Longitude   float32 `json:"longitude"`
	Timezone    string  `json:"timezone"`
	ISP         string  `json:"isp"`
	ORG         string  `json:"org"`
	AS          string  `json:"as"`
}

func IPLookupRun(addr string) (*IPLookup, error) {
	if net.ParseIP(addr) == nil {
		return nil, InvalidIPRepresentation
	}

	response, err := http.Get(fmt.Sprintf("http://ip-api.com/json/%s", addr))
	if err != nil {
		return nil, fmt.Errorf("Failed IPLookup: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed IPLookup: %v", err)
	}

	var ipRes IPLookup
	err = json.Unmarshal(body, &ipRes)
	if err != nil {
		return nil, fmt.Errorf("Failed IPLookup: %v", err)
	}

	return &ipRes, nil
}
