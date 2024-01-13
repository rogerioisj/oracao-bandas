package entities

import (
	"oracao-bandas.com/src/database/resources/common/entities"
)

type Band struct {
	entities.BaseEntity
	Title   string
	Type    string
	Country string
	UF      string
	Style   string
}
