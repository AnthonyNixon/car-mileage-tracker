package fillups

import (
	"AnthonyNixon/car-mileage-tracker/cmd/models/events"
	"AnthonyNixon/car-mileage-tracker/cmd/services/storage"
	"AnthonyNixon/car-mileage-tracker/cmd/utls/httperr"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
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

func (fillup FillUp) Save() (error httperr.HttpErr) {
	fileName := fmt.Sprintf("%s/fillup/%d.json", fillup.CarId, fillup.Id)

	content, err := json.Marshal(fillup)
	if err != nil {
		return httperr.New(http.StatusInternalServerError, "Failed to marshal json", err.Error())
	}

	error = storage.SaveFile(content, fileName)
	if error != nil {
		return
	}

	return nil
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
	event := events.Event{
		Id:        fillup.Id,
		Type:      "fillup",
		Timestamp: fillup.Timestamp,
		Odometer:  fillup.Odometer,
	}

	data.Last.Id = fillup.Id
	data.Last.Odometer = fillup.Odometer
	data.Last.MilesPerGallon = fillup.MilesPerGallon
	data.Last.FillUpId = fillup.Id
	data.Last.Timestamp = fillup.Timestamp

	data.Events = append(data.Events, event)
	fmt.Printf("%v", data)

	error = fillup.Save()
	if error != nil {
		return
	}

	error = storage.SaveStorageFile(data)
	if error != nil {
		return
	}

	return nil
}
