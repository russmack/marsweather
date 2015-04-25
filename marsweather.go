// Package marsweather is a library for retrieving Mars weather data.
package marsweather

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	EndpointLatest = "http://marsweather.ingenology.com/v1/latest/"
)

// Maas is the primary type for Ingenology's MAAS REST api.
type Maas struct{}
type MaasReport struct {
	Report struct {
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
	} `json:"report"`
}

// NewMaas returns a new instance of a Maas.
func NewMaas() *Maas {
	return &Maas{}
}

// GetLatest retrives the latest MAAS data.  Returns a MaasReport.
func (m *Maas) GetLatest() (MaasReport, error) {
	data, err := m.getLatestData()
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

// getLatestData is the MAAS HTTP client.  Returns the HTTP body.
func (m *Maas) getLatestData() ([]byte, error) {
	resp, err := http.Get(EndpointLatest)
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
