package storage_file

import (
	"AnthonyNixon/car-mileage-tracker/cmd/models/events"
	"AnthonyNixon/car-mileage-tracker/cmd/models/last_data"
	"AnthonyNixon/car-mileage-tracker/cmd/models/total_data"
)

type StorageFile struct {
	CarId  string               `json:"car_id"`
	Last   last_data.LastData   `json:"last"`
	Total  total_data.TotalData `json:"total"`
	Events []events.Event       `json:"events"`
}
