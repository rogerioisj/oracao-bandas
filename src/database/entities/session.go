package entities

type Session struct {
	BaseEntity
	Login string `gorm:"type:varchar(150); unique; NOT_NULL"`
}
