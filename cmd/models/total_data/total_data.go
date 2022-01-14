package total_data

import "fmt"

type TotalData struct {
	MilesPerGallon float64 `json:"miles_per_gallon"`
	Gallons        float64 `json:"gallons"`
	Miles          int     `json:"miles"`
	Cost           float64 `json:"cost"`
}

func (totalData TotalData) Print() {
	fmt.Printf("Totals:\n\tMiles:\t%d\n\tGallons:\t%f\n\tmpg:\t%f\n\tcost:\t%f\n",
		totalData.Miles, totalData.Gallons, totalData.MilesPerGallon, totalData.Cost)

}

func (totalData *TotalData) AddValues(miles int, gallons float64, cost float64) {
	totalData.Miles += miles
	totalData.Gallons += gallons
	totalData.MilesPerGallon = float64(miles) / gallons

	totalData.Cost += cost
}
