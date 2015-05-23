// Package marsweather is a library for retrieving Mars weather data.
// Ref: http://marsweather.ingenology.com/
package marsweather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	EndpointLatest      = "http://marsweather.ingenology.com/v1/latest/"
	EndpointArchivePage = "http://marsweather.ingenology.com/v1/archive/?page=%d"
)

// Maas is the primary type for Ingenology's MAAS REST api.
type Maas struct{}
type MaasReport struct {
	Report Report `json:"report"`
}
type MaasArchivePage struct {
	Count    int      `json:"count"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Results  []Report `json:"results"`
}
type Report struct {
	TerrestrialDate    string  `json:"terrestrial_date"`
	CuriositySolNumber float32 `json:"sol"`
	SeasonalDate       float32 `json:"ls"`
	MinTemp            float32 `json:"min_temp"`
	MinTempFahrenheit  float32 `json:"min_temp_fahrenheit"`
	MaxTemp            float32 `json:"max_temp"`
	MaxTempFahrenheit  float32 `json:"max_temp_fahrenheit"`
	Pressure           float32 `json:"pressure"`
	PressureString     string  `json:"pressure_string"`
	AbsHumidity        float32 `json:"abs_humidity"`
	WindSpeed          float32 `json:"wind_speed"`
	WindDirection      string  `json:"wind_direction"`
	AtmosphericOpacity string  `json:"atmo_opacity"`
	Season             string  `json:"season"`
	Sunrise            string  `json:"sunrise"`
	Sunset             string  `json:"sunset"`
}

// NewMaas returns a new instance of a Maas.
func NewMaas() *Maas {
	return &Maas{}
}

// GetLatest retrives the latest MAAS data.  Returns a MaasReport.
func (m *Maas) GetLatest() (MaasReport, error) {
	data, err := m.getData(EndpointLatest)
	if err != nil {
		return MaasReport{}, err
	}
	var report MaasReport
	err = json.Unmarshal(data, &report)
	if err != nil {
		return MaasReport{}, err
	}
	return report, nil
}

// getArchivePage gets a specified page from the archive data.
func (m *Maas) GetArchivePage(page int) (MaasArchivePage, error) {
	data, err := m.getData(fmt.Sprintf(EndpointArchivePage, page))
	if err != nil {
		return MaasArchivePage{}, err
	}
	var p MaasArchivePage
	err = json.Unmarshal(data, &p)
	if err != nil {
		return MaasArchivePage{}, err
	}
	return p, nil
}

// getData is the MAAS HTTP client.  Returns the HTTP body.
func (m *Maas) getData(endpoint string) ([]byte, error) {
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
