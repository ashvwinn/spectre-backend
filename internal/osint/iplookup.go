package osint

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
)

var InvalidIPRepresentation = errors.New("entered ip is not valid ip")

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
	Status      string
	Message     string
}

func IPLookupRun(addr string) (*IPLookup, error) {
	if net.ParseIP(addr) == nil {
		return nil, InvalidIPRepresentation
	}

	response, err := http.Get(fmt.Sprintf("http://ip-api.com/json/%s", addr))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var ipRes IPLookup
	err = json.Unmarshal(body, &ipRes)
	if err != nil {
		return nil, err
	}

	return &ipRes, nil
}
