package middleware

import (
	types "s21/surdo/pkg/model"
	"s21/surdo/pkg/store"

	"gorm.io/gorm"
)

// Функция для проверки, существует ли пользователь с данным Telegram ID в базе данных
func CheckTelegramIDExists(tgID int64) (bool, error) {
    var user types.TUser

    err := store.DB.Where("telegram_id = ?", tgID).First(&user).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return false, nil
        }
        return false, err
    }

    return true, nil
}

// package middleware

// import (
// 	//  "s21/surdo/internal/types"
// 	//  "s21/surdo/internal/types/user"

// 	types "s21/surdo/pkg/model"
// 	"s21/surdo/pkg/user"

// 	tele "gopkg.in/telebot.v4"
// )

// func UserMiddleware() tele.MiddlewareFunc {
//  return func(next tele.HandlerFunc) tele.HandlerFunc {
//   return func(c tele.Context) error {
//    if c.Payment() != nil || c.PreCheckoutQuery() != nil {
//     return next(c)
//    }
//    usr, err := user.FindUserByTelegramID(c.Chat().ID)
//    if err != nil {
//     if user.CreateUserByTelegramID(types.NewUserTg(c.Sender())) != nil {
//      return err
//     }
//    } else {
//     if usr.IsBlocked {
//      return c.Send("Вы были заблокированы за нарушение условий использования")
//     }
//    }
//    return next(c)
//   }
//  }

// }

