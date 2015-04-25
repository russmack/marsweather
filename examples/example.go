package main

import (
	"fmt"
	"github.com/russmack/marsweather"
	"os"
)

func main() {
	maas := marsweather.NewMaas()
	maasReport, err := maas.GetLatest()
	if err != nil {
		fmt.Println("Well, we got the data from Mars to Earth, but...", err)
		os.Exit(1)
	}
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
}
