package repository

type Document interface {
}

type Passenger interface {
}

type Ticket interface {
}

type Repository struct {
	Document
	Passenger
	Ticket
}

func NewRepository() *Repository {
	return &Repository{}
}
