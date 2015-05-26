package marsweather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestGetLatest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w,
			`{
				"report": {
			    	"terrestrial_date": "2013-05-01",
					"sol": 261,
				    "ls": 310.5,
			        "min_temp": -69.75,
					"min_temp_fahrenheit": -93.55,
				    "max_temp": -4.48,
			        "max_temp_fahrenheit": 23.94,
					"pressure": 868.05,
				    "pressure_string": "Higher",
			        "abs_humidity": null,
					"wind_speed": null,
				    "wind_direction": "--",
			        "atmo_opacity": "Sunny",
					"season": "Month 11",
				    "sunrise": "2013-05-01T11:00:00Z",
			        "sunset": "2013-05-01T22:00:00Z"
				}
			}`,
		)
	}))
	defer ts.Close()

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(ts.URL)
		},
	}
	httpClient := &http.Client{Transport: transport}
	maas := NewMaas(httpClient)
	actual, err := maas.GetLatest()
	if err != nil {
		t.Fatalf("Error calling GetLatest. %v", err)
	}

	expectedJson := `{
				    	"report": {
			        		"terrestrial_date": "2013-05-01",
					        "sol": 261,
				        	"ls": 310.5,
			        		"min_temp": -69.75,
					        "min_temp_fahrenheit": -93.55,
				        	"max_temp": -4.48,
			        		"max_temp_fahrenheit": 23.94,
					        "pressure": 868.05,
				        	"pressure_string": "Higher",
			        		"abs_humidity": null,
					        "wind_speed": null,
				        	"wind_direction": "--",
			        		"atmo_opacity": "Sunny",
					        "season": "Month 11",
				        	"sunrise": "2013-05-01T11:00:00Z",
			        		"sunset": "2013-05-01T22:00:00Z"
					    }
					}`

	expected := MaasReport{}
	err = json.Unmarshal([]byte(expectedJson), &expected)
	if err != nil {
		t.Fatalf("Error creating expected response object. %v", err)
	}

	if actual.Report.TerrestrialDate != expected.Report.TerrestrialDate {
		t.Errorf("Wrong TerrestrialDate.")
	}
	if actual.Report.CuriositySolNumber != expected.Report.CuriositySolNumber {
		t.Errorf("Wrong CuriositySolNumber.")
	}
	if actual.Report.SeasonalDate != expected.Report.SeasonalDate {
		t.Errorf("Wrong SeasonalDate.")
	}
	if actual.Report.MinTemp != expected.Report.MinTemp {
		t.Errorf("Wrong MinTemp.")
	}
	if actual.Report.MinTempFahrenheit != expected.Report.MinTempFahrenheit {
		t.Errorf("Wrong MinTempFahrenheit.")
	}
	if actual.Report.MaxTemp != expected.Report.MaxTemp {
		t.Errorf("Wrong MaxTemp.")
	}
	if actual.Report.MaxTempFahrenheit != expected.Report.MaxTempFahrenheit {
		t.Errorf("Wrong MaxTempFahrenheit.")
	}
	if actual.Report.Pressure != expected.Report.Pressure {
		t.Errorf("Wrong Pressure.")
	}
	if actual.Report.PressureString != expected.Report.PressureString {
		t.Errorf("Wrong PressureString.")
	}
	if actual.Report.AbsHumidity != expected.Report.AbsHumidity {
		t.Errorf("Wrong AbsHumidity.")
	}
	if actual.Report.WindSpeed != expected.Report.WindSpeed {
		t.Errorf("Wrong WindSpeed.")
	}
	if actual.Report.WindDirection != expected.Report.WindDirection {
		t.Errorf("Wrong WindDirection.")
	}
	if actual.Report.AtmosphericOpacity != expected.Report.AtmosphericOpacity {
		t.Errorf("Wrong AtmosphericOpacity.")
	}
	if actual.Report.Season != expected.Report.Season {
		t.Errorf("Wrong Season.")
	}
	if actual.Report.Sunrise != expected.Report.Sunrise {
		t.Errorf("Wrong Sunrise.")
	}
	if actual.Report.Sunset != expected.Report.Sunset {
		t.Errorf("Wrong Sunset.")
	}
}
