package types

// import (
// 	"errors"
// 	"github.com/google/uuid"
// 	"time"
// )

// type TOrder struct {
// 	ID           uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
// 	UserID       uuid.UUID  `gorm:"type:uuid;not null"` // Внешний ключ на TUser
// 	TranslatorID *uuid.UUID `gorm:"type:uuid"`          // Внешний ключ на TTranslator (опционально)
// 	Slot         *TSlot     `gorm:"foreignKey:OrderID"` // Связь с TSlot
// 	CreatedAt    time.Time
// 	Status       OrderStage `gorm:"not null"` // Тип int для OrderStage
// 	Type         OrderType  `gorm:"not null"` // Тип int для OrderType
// 	Cost         int64      `gorm:"not null;default:0"`
// }

// func NewOrder(user TUser, oType OrderType,
// 	slot *TSlot) (TOrder, error) {
// 	status := ORDER_STAGE_CREATED

// 	switch oType {
// 	case ORDER_TYPE_ONLINE_ASAP:
// 		break
// 	case ORDER_TYPE_ONLINE_PLANNED:
// 		break
// 	case ORDER_TYPE_OFFLINE:
// 		return TOrder{}, errors.New("not implemented")

// 	}

// 	return TOrder{
// 		ID:           uuid.New(),
// 		UserID:       user.ID,
// 		TranslatorID: nil,
// 		Slot:         slot,
// 		CreatedAt:    time.Now(),
// 		Status:       status,
// 		Type:         oType,
// 	}, nil
// }

// type OrderStage int
// type OrderType int

// const (
// 	ORDER_STAGE_NOT_APPROVED_YET  OrderStage = -1
// 	ORDER_STAGE_CREATED           OrderStage = 0
// 	ORDER_STAGE_SETTED_TRANSLATOR OrderStage = 1
// 	ORDER_STAGE_15_MIN            OrderStage = 2
// 	ORDER_STAGE_LINK_SENT         OrderStage = 3
// 	ORDER_STAGE_5_TO_END          OrderStage = 4
// 	ORDER_STAGE_FINISHED          OrderStage = 5
// 	//
// 	ORDER_STAGE_CANCELED OrderStage = 99
// )

// const (
// 	ORDER_TYPE_ONLINE_ASAP    OrderType = 0
// 	ORDER_TYPE_ONLINE_PLANNED OrderType = 1
// 	ORDER_TYPE_OFFLINE        OrderType = 2
// 	ORDER_TYPE_SOS            OrderType = 3
// )
