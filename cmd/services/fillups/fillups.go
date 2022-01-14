package fillups

import (
	"AnthonyNixon/car-mileage-tracker/cmd/services/storage"
	"AnthonyNixon/car-mileage-tracker/cmd/utls/httperr"
	"fmt"
	"math"
	"time"
)

type FillUp struct {
	CarId          string    `json:"car_id"`
	Id             int       `json:"id"`
	Gallons        float64   `json:"gallons"`
	PricePerGallon float64   `json:"price_per_gallon"`
	GasStation     string    `json:"gas_station"`
	TotalCost      float64   `json:"total_cost"`
	MilesPerGallon float64   `json:"miles_per_gallon"`
	Odometer       int       `json:"odometer"`
	Timestamp      time.Time `json:"timestamp"`
}

func (fillup FillUp) print() {
	fmt.Printf("FillUp:\n\tcar_id:\t\t%s\n\tid:\t\t%d\n\tOdometer:\t%d\n\tGallons:\t%f\n\t$/gallon:\t%f\n\tCost:\t\t%f\n\tmpg:\t\t%f\n\tStation:\t%s\n\tTimestamp:\t%v\n",
		fillup.CarId, fillup.Id, fillup.Odometer, fillup.Gallons, fillup.PricePerGallon, fillup.TotalCost, fillup.MilesPerGallon, fillup.GasStation, fillup.Timestamp)
}

func NewFillup(fillup FillUp) (error httperr.HttpErr) {
	data, error := storage.GetStorageFileByID(fillup.CarId)
	if error != nil {
		return
	}

	fillup.Id = data.Last.Id + 1

	milesDriven := fillup.Odometer - data.Last.Odometer

	milesPerGallon := float64(milesDriven) / fillup.Gallons
	fillup.MilesPerGallon = milesPerGallon

	cost := fillup.Gallons * fillup.PricePerGallon
	fillup.TotalCost = math.Ceil(cost*100) / 100

	fillup.Timestamp = time.Now()

	fillup.print()

	data.Total.AddValues(milesDriven, fillup.Gallons, fillup.TotalCost)

	return nil
}
