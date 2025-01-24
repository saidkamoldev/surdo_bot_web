package types

// import (
// 	"github.com/google/uuid"
// )

// type UserTransaction struct {
// 	ID               uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
// 	UserID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
// 	DeltaMoney       int64     `gorm:"not null;default:0"`
// 	DeltaMoneyFrozen int64     `gorm:"not null;default:0"`
// 	Type             int       `gorm:"not null"`
// 	Comment          string    `gorm:"type:text"`
// }
// type TranslatorTransaction struct {
// 	ID           uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
// 	TranslatorID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
// 	DeltaMoney   int64     `gorm:"not null;default:0"`
// 	Type         int       `gorm:"not null"`
// 	Comment      string    `gorm:"type:text"`
// }

// //
// //func WriteUserTransaction(tx UserTransaction) error {
// //	return db.ORM.Create(&tx).Error
// //}

// const (
// 	TXU_ORDER_FREEZE    = 0
// 	TXU_ORDER_CANCELLED = 1
// 	TXU_ORDER_PAID      = 2
// 	TXU_PAYMENT         = 3
// 	TXU_REFUND          = 4
// )
