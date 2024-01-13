package entities

import "time"

type BaseEntity struct {
	Id         string
	CreateDate time.Time
	LastUpdate time.Time
}
