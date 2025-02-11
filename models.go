package airlineManager

type Document struct {
	Id     int    `json:"id"`
	Type   int    `json:"type"`
	Number string `json:"number"`
}

type Passenger struct {
	Id         string `json:"id"`
	DocumentId int    `json:"document_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
}

type Ticket struct {
	Id            string    `json:"id"`
	Passenger     Passenger `json:"passenger"`
	Document      Document  `json:"document"`
	Number        string    `json:"number"`
	Provider      int       `json:"provider"`
	Departure     string    `json:"departure"`
	Arrival       string    `json:"arrival"`
	EventDate     string    `json:"event_date"`
	DepartureDate string    `json:"departure_date"`
	ArrivalDate   string    `json:"arrival_date"`
}
