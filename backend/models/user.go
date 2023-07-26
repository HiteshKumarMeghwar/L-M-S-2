package models

type User struct {
	ID             uint    `gorm:"primaryKey" json:"id"`
	FirstName      string  `json:"first_name"`
	LastName       string  `json:"last_name"`
	Email          string  `gorm:"unique" json:"email"`
	Password       string  `json:"-"`
	Phone          string  `json:"phone"`
	ProfilePicture string  `json:"profile_picture"` // New field for the profile picture
	Roles          []*Role `gorm:"many2many:user_roles;" json:"roles"`
}
