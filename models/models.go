package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Login          string
	FullName       string
	Role           string
	BirthDate      string
	PhoneNumber    string
	ActivityStatus string
	PasswordHash   string
}

type Queue struct {
	gorm.Model
	Name        string
	Description string
	Features    string
	Cost        float64
}

type Portfolio struct {
	gorm.Model
	Name            string
	Cost            float64
	PlannedDeadline string
	Status          string
	QueueID         uint
	AssignedToID    uint
}
