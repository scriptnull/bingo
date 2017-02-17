package users

import "github.com/jinzhu/gorm"

// User type denotes a user in bingo game
type User struct {
	gorm.Model
	ID        uint   `gorm:"AUTO_INCREMENT; primary_key" json:"id"`
	EmailID   string `gorm:"not null;unique" json:"emailId"`
	AvatarURL string `json:"avatarUrl"`
	Token     string `gorm:"not null;unique" json:"token"`
}
