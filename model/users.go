package model

type Users struct {
	Username string `gorm:"type:varchar(255)"`
	Password string `gorm:"type:varchar(255)"`
}
