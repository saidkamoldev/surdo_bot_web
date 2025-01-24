package types

import (
	"s21/surdo/pkg/i18n/lng"

	"github.com/google/uuid"
	tele "gopkg.in/telebot.v4"
	// "surdo-bot/internal/i18n/lng"
)



type TUser struct {
	ID                  uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	FirstName           string       `gorm:"type:varchar(100);not null"`
	LastName            string       `gorm:"type:varchar(100);not null"`
	TelegramID          int64        `gorm:"not null;unique"`
	PhoneNumber         string       `gorm:"type:varchar(15);not null;unique"` // Telefon raqam (majburiy, unique)
	Email               *string      `gorm:"type:varchar(100);unique"` 
	Language            lng.Language `gorm:"type:varchar(10)"`        // Email (ixtiyoriy)
	License             bool         `gorm:"not null;default:false"`
	LanguageSet         bool         `gorm:"not null;default:false"`
	Money               int64        `gorm:"not null;default:0"`
	FrozenMoney         int64        `gorm:"not null;default:0"`
	IsBlocked           bool         `gorm:"not null;default:false"`
	IsDocumentsApproved bool         `gorm:"not null;default:false"`
}


func NewUserTg(user *tele.User) TUser {
	return TUser{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		TelegramID:  user.ID,
		// PhoneNumber: phoneNumber, // Telefon raqam majburiy
		License:     false,
		LanguageSet: false,
	}
}


// newUUID := uuid.New()
// type TUser struct {
// 	ID                  uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
// 	FirstName           string       `gorm:"type:varchar(100);not null"`
// 	LastName            string       `gorm:"type:varchar(100);not null"`
// 	TelegramID          int64        `gorm:"not null;unique"`
// 	// Language            lng.Language `gorm:"type:varchar(10)"`
// 	License             bool         `gorm:"not null;default:false"`
// 	LanguageSet         bool         `gorm:"not null;default:false"`
// 	// Orders              []TOrder     `gorm:"foreignKey:UserID"`
// 	Money               int64        `gorm:"not null;default:0"`
// 	FrozenMoney         int64        `gorm:"not null;default:0"`
// 	IsBlocked           bool         `gorm:"not null;default:false"`
// 	IsDocumentsApproved bool         `gorm:"not null;default:false"`
// }

// func NewUserTg(user *tele.User) TUser {
// 	return TUser{
// 		FirstName:   user.FirstName,
// 		LastName:    user.LastName,
// 		TelegramID:  user.ID,
// 		// Language:    "---",
// 		License:     false,
// 		LanguageSet: false,
// 	}
// }

func IsUserComplete(user TUser) bool {
	if !user.License {
		return false
	}
	if !user.LanguageSet {
		return false
	}
	// if user.Language == "---" {
	// 	return false
	// }

	return true
}
