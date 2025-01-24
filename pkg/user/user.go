package user

import (
	types "s21/surdo/pkg/model"
	"s21/surdo/pkg/store"
)

func FindUserByTelegramID(tgID int64) (types.TUser, error) {
 var user types.TUser
 err := store.DB.First(&user, "telegram_id = ?", tgID).Error

 return user, err
}
func CreateUserByTelegramID(user types.TUser) error {
 usr := &user
 return store.DB.Create(usr).Error
}
