package last_data

import "time"

type LastData struct {
	Id             int       `json:"id"`
	Timestamp      time.Time `json:"timestamp"`
	Odometer       int       `json:"odometer"`
	MilesPerGallon float64   `json:"miles_per_gallon"`
	FillUpId       int       `json:"fill_up_id"`
	MaintenanceId  int       `json:"maintenance_id"`
	NoteId         int       `json:"note_id"`
}
