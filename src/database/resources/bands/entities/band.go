package entities

import (
	"oracao-bandas.com/src/database/resources/common/entities"
)

type Band struct {
	entities.BaseEntity
	Title   string `gorm:"type:varchar(150); unique; NOT_NULL"`
	Type    string
	Country string
	UF      string
	Style   string
}
