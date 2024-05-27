package services

import (
	"context"
	"go-jwt-eg/configs"
	"go-jwt-eg/data/request"
	"go-jwt-eg/entities"
)

func FindAll(ctx context.Context) ([]entities.Person, error) {
	var persons []entities.Person
	err := configs.DB.WithContext(ctx).
		Joins("Division").
		Joins("PlaceOfBirth").
		Preload("Addresses").
		Preload("Groups").
		Order("persons.name asc").
		Find(&persons)
	if err.Error != nil {
		return nil, err.Error
	}

	return persons, nil
}

func FindById(ctx context.Context, id uint) (*entities.Person, error) {
	var person entities.Person
	err := configs.DB.WithContext(ctx).
		Where("persons.id = ?", id).
		Joins("Division").
		Joins("PlaceOfBirth").
		Preload("Addresses").
		Preload("Groups").
		First(&person)
	if err.Error != nil {
		return nil, err.Error
	}

	return &person, nil
}

func Create(reqPerson request.CreatePersonAdd) error {
	var addresses []entities.Address
	for _, value := range reqPerson.Addresses {
		var address entities.Address
		address.Address = value
		addresses = append(addresses, address)
	}

	var groups []entities.Group
	for _, value := range reqPerson.Groups {
		var group entities.Group
		group.ID = value
		groups = append(groups, group)
	}

	person := entities.Person{
		Name:         reqPerson.Name,
		PhoneNumber:  reqPerson.PhoneNumber,
		DivisionID:   reqPerson.DivisionID,
		PlaceOfBirth: reqPerson.PlaceOfBirth,
		Addresses:    addresses,
		Groups:       groups,
	}

	if err := configs.DB.Create(&person).Error; err != nil {
		return err
	}

	return nil
}

func Update(reqPerson request.UpdatePersonAdd) string {
	var addresses []entities.Address
	configs.DB.Where("person_id=?", reqPerson.ID).Delete(&addresses)

	var personGroups []entities.PersonGroups
	configs.DB.Where("person_id=?", reqPerson.ID).Delete(&personGroups)

	for _, value := range reqPerson.Addresses {
		var address entities.Address
		address.Address = value
		addresses = append(addresses, address)
	}
	var groups []entities.Group

	for _, value := range reqPerson.Groups {
		var group entities.Group
		group.ID = value
		groups = append(groups, group)
	}

	configs.DB.Where("person_id=?", reqPerson.ID).Updates(&reqPerson.PlaceOfBirth)

	var person = entities.Person{
		ID:          reqPerson.ID,
		Name:        reqPerson.Name,
		PhoneNumber: reqPerson.PhoneNumber,
		DivisionID:  reqPerson.DivisionID,
		Addresses:   addresses,
		Groups:      groups,
	}
	configs.DB.Where("id=?", reqPerson.ID).Updates(&person)

	return "data sucessfully updated"
}

func Delete(id uint) string {
	var addresses []entities.Address
	configs.DB.Where("person_id=?", id).Delete(&addresses)

	var personGroups []entities.PersonGroups
	configs.DB.Where("person_id=?", id).Delete(&personGroups)

	var placeOfBirth entities.PlaceOfBirth
	configs.DB.Where("person_id=?", id).Delete(&placeOfBirth)

	var person entities.Person
	configs.DB.Where("id=?", id).Delete(&person)

	return "data sucessfully deleted"
}
