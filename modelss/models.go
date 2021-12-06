package modelss

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Names    string `json:"names"`
	Username string `gorm:"unique_index;not null" json:"username"`
	Email    string `gorm:"unique_index;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}
