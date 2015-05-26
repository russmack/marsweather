package main

import (
	"fmt"
	"github.com/russmack/marsweather"
	"os"
	"time"
)

func main() {
	maas := marsweather.NewMaas(nil)
	maasReport, err := maas.GetLatest()
	if err != nil {
		fmt.Println("Well, we got the data from Mars to Earth, but...", err)
		os.Exit(1)
	}
	fmt.Println("Latest:")
	fmt.Println("Curiosity Sol number:", maasReport.Report.CuriositySolNumber)
	fmt.Println("Atmospheric opacity:", maasReport.Report.AtmosphericOpacity)
	fmt.Println("Max temp (C):", maasReport.Report.MaxTemp)
	fmt.Println("Min temp (C):", maasReport.Report.MinTemp)
	fmt.Println("Wind speed:", maasReport.Report.WindSpeed)
	fmt.Println("Pressure:", maasReport.Report.Pressure)
	fmt.Println("Pressure string:", maasReport.Report.PressureString)
	fmt.Println("Absolute humidity:", maasReport.Report.AbsHumidity)
	fmt.Println("Terrestrial date:", maasReport.Report.TerrestrialDate)
	fmt.Println("Seasonal date:", maasReport.Report.SeasonalDate)

	page := 2
	p, err := maas.GetArchivePage(page)
	if err != nil {
		fmt.Println("Error getting archive data.", err)
		os.Exit(1)
	}
	fmt.Println("")
	fmt.Printf("Archive Data Page %v\n", page)
	fmt.Println("Count:", p.Count)
	fmt.Println("Next:", p.Next)
	fmt.Println("Previous:", p.Previous)
	fmt.Println("Results:")
	for _, j := range p.Results {
		fmt.Println("Terrestrial data:", j.TerrestrialDate)
		fmt.Println("Max Temp:", j.MaxTemp)
		fmt.Println("Min Temp:", j.MinTemp)
		fmt.Println("")
	}

	fromDate := time.Date(2015, time.April, 19, 0, 0, 0, 0, time.UTC)
	toDate := time.Date(2015, time.April, 22, 0, 0, 0, 0, time.UTC)
	d, err := maas.GetArchiveDateRange(fromDate, toDate)
	if err != nil {
		fmt.Println("Error getting archive date range.", err)
		os.Exit(1)
	}
	fmt.Println("Archive date range:")
	for _, j := range d.Results {
		fmt.Println(j)
	}
}
