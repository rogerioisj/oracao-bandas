package entities

type User struct {
	BaseEntity
	Name     string `gorm:"type:varchar(150); NOT_NULL"`
	Login    string `gorm:"type:varchar(150); unique; NOT_NULL"`
	Password string `gorm:"type:varchar(150); NOT_NULL"`
}
