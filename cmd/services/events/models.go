package events

import "time"

type Event struct {
	Metadata Metadata               `json:"metadata"`
	Details  map[string]interface{} `json:"details"`
}

type Metadata struct {
	Id        string    `json:"id"`
	Type      string    `json:"type"`
	Odometer  int       `json:"odometer"`
	Timestamp time.Time `json:"timestamp"`
}

type FillUp struct {
	Gallons        float64 `json:"gallons"`
	PricePerGallon float64 `json:"price_per_gallon"`
	GasStation     string  `json:"gas_station"`
	TotalCost      float64 `json:"total_cost"`
	MilesPerGallon float64 `json:"miles_per_gallon"`
}

func (fillup *FillUp) ToStruct(details map[string]interface{}) {
	fillup.Gallons = details["gallons"].(float64)
	fillup.PricePerGallon = details["price_per_gallon"].(float64)
	fillup.GasStation = details["gas_station"].(string)
}

type Maintenance struct {
	Type            string  `json:"type"`
	Notes           string  `json:"notes"`
	Price           float64 `json:"price"`
	ServiceLocation string  `json:"service_location"`
}

func (maintenance *Maintenance) ToStruct(details map[string]interface{}) {
	maintenance.Type = details["type"].(string)
	maintenance.Notes = details["notes"].(string)
	maintenance.Price = details["price"].(float64)
	maintenance.ServiceLocation = details["service_location"].(string)
}

type Note struct {
	Body string `json:"body"`
}

func (note *Note) ToStruct(details map[string]interface{}) {
	note.Body = details["body"].(string)
}
