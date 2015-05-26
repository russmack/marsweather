package marsweather

import (
	"encoding/json"
	"fmt"
	"log"
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
		log.Println("Error calling GetLatest.", err)
		t.Fail()
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
		log.Println("Error creating expected response object.", err)
		t.Fail()
	}

	if actual.Report.CuriositySolNumber != expected.Report.CuriositySolNumber ||
		actual.Report.SeasonalDate != expected.Report.SeasonalDate ||
		actual.Report.MinTemp != expected.Report.MinTemp {
		t.Fail()
	}
}
