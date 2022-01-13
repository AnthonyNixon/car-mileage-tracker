package events

import (
	"AnthonyNixon/car-mileage-tracker/cmd/utls/httperr"
	"fmt"
	"log"
	"net/http"
)

func NewEvent(event Event) (error httperr.HttpErr) {
	switch event.Metadata.Type {
	case "FillUp":
		log.Println("Event type: FillUp")
		fillup := new(FillUp)
		fillup.ToStruct(event.Details)
		log.Printf("FillUp details:\n%v", fillup)
	case "Maintenance":
		log.Println("Event type: Maintenance")
		maintenance := new(Maintenance)
		maintenance.ToStruct(event.Details)
		log.Printf("Maintenance details:\n%v", maintenance)
	case "Note":
		log.Println("Event type: Note")
		note := new(Note)
		note.ToStruct(event.Details)
		log.Printf("Note details:\n%v", note)
	default:
		return httperr.New(
			http.StatusBadRequest,
			"Invalid event type",
			fmt.Sprintf("'%s' is not in ['FillUp, 'Maintenance', 'Note']", event.Metadata.Type),
		)
	}

	return nil
}
