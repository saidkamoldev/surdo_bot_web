package types

// import (
// 	"github.com/google/uuid"
// 	tele "gopkg.in/telebot.v4"
// 	"surdo-bot/internal/i18n/lng"
// )

// type TTranslator struct {
// 	ID                  uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
// 	FirstName           string       `gorm:"type:varchar(100);not null"`
// 	LastName            string       `gorm:"type:varchar(100);not null"`
// 	TelegramID          int64        `gorm:"not null"`
// 	Availability        []TWorktime  `gorm:"foreignKey:TranslatorID"`
// 	Language            lng.Language `gorm:"type:varchar(10)"`
// 	License             bool         `gorm:"not null;default:false"`
// 	LanguageSet         bool         `gorm:"not null;default:false"`
// 	Orders              []TOrder     `gorm:"foreignKey:TranslatorID"` // Связь с заказами
// 	Money               int64        `gorm:"not null;default:0"`
// 	IsDocumentsApproved bool         `gorm:"not null;default:false"`
// 	IsBlocked           bool         `gorm:"not null;default:false"`
// }

// func NewTranslatorTg(user *tele.User) TTranslator {
// 	return TTranslator{
// 		FirstName:   user.FirstName,
// 		LastName:    user.LastName,
// 		TelegramID:  user.ID,
// 		Language:    "---",
// 		License:     false,
// 		LanguageSet: false,
// 	}
// }
