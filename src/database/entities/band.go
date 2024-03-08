package entities

type Band struct {
	BaseEntity
	Title   string `gorm:"type:varchar(150); unique; NOT_NULL"`
	Type    string
	Country string
	UF      string
	Style   string
}
