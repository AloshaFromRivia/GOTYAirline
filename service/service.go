package service

import "UltimateDeluxeSuperAirlineManagerGOTY/repository"

type Document interface {
}

type Passenger interface {
}

type Ticket interface {
}

type Service struct {
	Document
	Passenger
	Ticket
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
