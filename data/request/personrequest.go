package request

import "go-jwt-eg/entities"

type CreatePersonAdd struct {
	Name         string
	PhoneNumber  string
	DivisionID   uint
	PlaceOfBirth *entities.PlaceOfBirth
	Addresses    []string
	Groups       []uint
}

type UpdatePersonAdd struct {
	ID           uint
	Name         string
	PhoneNumber  string
	DivisionID   uint
	PlaceOfBirth *entities.PlaceOfBirth
	Addresses    []string
	Groups       []uint
}
