package entities

import "time"

type Person struct {
	ID           uint          `gorm:"primaryKey"`
	Name         string        `validate:"required"`
	PhoneNumber  string        `validate:"required"`
	DivisionID   uint          `validate:"required"`
	PlaceOfBirth *PlaceOfBirth //has one relationship
	Division     *Division     //belongs to relationship
	Addresses    []Address     `gorm:"foreignKey:PersonID"`     //has many relationship
	Groups       []Group       `gorm:"many2many:person_groups"` //many-to-many relationship
}

func (Person) TableName() string {
	return "persons"
}

type PlaceOfBirth struct {
	ID       uint `gorm:"primaryKey"`
	City     string
	PersonID uint
	Date     *time.Time
}

type Address struct {
	ID       uint `gorm:"primaryKey"`
	Address  string
	PersonID uint
}

type Division struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

type Group struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

type PersonGroups struct {
	PersonId uint
	GroupId  uint
}
