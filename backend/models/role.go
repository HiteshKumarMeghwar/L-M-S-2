package models

type Role struct {
	ID    uint    `gorm:"primaryKey" json:"id"`
	Name  string  `json:"name"`
	Users []*User `gorm:"many2many:user_roles;" json:"users"`
}
