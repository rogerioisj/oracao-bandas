package entities

type Auth struct {
	BaseEntity
	Name     string `gorm:"type:varchar(150); NOT_NULL"`
	Email    string `gorm:"type:varchar(150); unique; NOT_NULL"`
	Password string `gorm:"type:varchar(100); NOT_NULL"`
}
