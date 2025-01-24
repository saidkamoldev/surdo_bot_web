package api

import (
	"log"
	// "s21/surdo/internal/bot/buttons"
	// "s21/surdo/internal/bot/buttons"
	// "s21/surdo/internal/bot/buttons"
	"s21/surdo/internal/bot/buttons"
	"s21/surdo/pkg/i18n/lng"
	types "s21/surdo/pkg/model"
	"s21/surdo/pkg/store"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	// "gopkg.in/telebot.v4"
	"gorm.io/gorm"

	// "github.com/google/uuid"
	"net/http"
)
func GetAllUsers(c *gin.Context) {
 var users []types.TUser
 db := store.GetDB()

 if err := db.Find(&users).Error; err != nil {
  c.JSON(http.StatusNotFound, gin.H{"error": "Database is empty"})
  return
 }

 c.JSON(http.StatusOK, users)
}


// func GetUserByTelegramID(c *gin.Context) {
//     var user types.TUser
//     // telegramID := c.Param("telegramID")
// 	user.TelegramID = buttons.Telegram_id
// 	user.Language = lng.Language(buttons.Language)
//     tgID, err := strconv.ParseInt(user.TelegramID, 10, 64)
//     if err != nil {
//         log.Printf("Ошибка преобразования Telegram ID: %v", err)
//         c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный Telegram ID"})
//         return
//     }

//     db := store.GetDB()
//     result := db.Where("telegram_id = ?", tgID).First(&user)
//     if result.Error != nil {
//         if result.Error == gorm.ErrRecordNotFound {
//             c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
//             return
//         }
//         log.Printf("Ошибка при поиске пользователя: %v", result.Error)
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при поиске данных пользователя"})
//         return
//     }

//     c.JSON(http.StatusOK, user)
// }
func GetUserByTelegramID(c *gin.Context) {
    telegramID := buttons.Telegram_id
    var user types.TUser

    db := store.GetDB()
    result := db.Where("telegram_id = ?", telegramID).First(&user)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
            return
        }
        log.Printf("Ошибка при поиске пользователя: %v", result.Error)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при поиске данных пользователя"})
        return
    }

    c.JSON(http.StatusOK, user)
}


func CreateUser(c *gin.Context) {
 var user types.TUser
 if err := c.ShouldBindJSON(&user); err != nil {
  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  return
 }
//  user.TelegramID = 
 user.TelegramID = buttons.Telegram_id
 user.Language = lng.Language(buttons.Language)
 db := store.GetDB()
 result := db.Create(&user)
 if result.Error != nil {
  log.Printf("Error adding user: %s", result.Error)
  c.JSON(http.StatusInternalServerError, gin.H{"error": "Error adding data"})
  return
 }

 c.JSON(http.StatusCreated, user)
}
func UpdateUser(c *gin.Context) {
 id := c.Param("id")
 var user types.TUser
 var new_user types.TUser
 db := store.GetDB()
 if err := db.First(&user, "id = ?", id).Error; err != nil {
  c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
  return
 }
 if err := c.ShouldBindJSON(&new_user); err != nil {
  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  return
 }
 db.Model(&user).Updates(new_user)
 c.JSON(http.StatusOK, user)
}
func DeleteUser(c *gin.Context) {
 var body struct {
  ID string `json:"ID"`
 }

 if err := c.ShouldBindJSON(&body); err != nil {
  c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
  return
 }

 parsedID, err := uuid.Parse(body.ID)
 if err != nil {
  c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
  return
 }

 db := store.GetDB()

 var user types.TUser
 if err := db.First(&user, "id = ?", parsedID).Error; err != nil {
  if err == gorm.ErrRecordNotFound {
   log.Printf("User with ID %s not found", body.ID)
   c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
   return
  }
  log.Printf("Error finding user: %s", err.Error())
  c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding user"})
  return
 }

 if err := db.Delete(&user).Error; err != nil {
  log.Printf("Error deleting user with ID %s: %s", body.ID, err.Error())
  c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
  return
 }

 c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}



// Получить текст песни по её ID
// func GetSongsLyrics(c *gin.Context) {
// 	id := c.Param("id")
// 	var song model.Song

// 	db := store.GetDB()
// 	if err := db.First(&song, id).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
// 		return
// 	}

// 	lyrics := strings.Split(song.Text, "\n\n")

// 	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
// 	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "3"))
// 	total := len(lyrics)
// 	start := (page - 1) * pageSize

// 	if start > total {
// 		start = total
// 	}

// 	end := start + pageSize
// 	if end > total {
// 		end = total
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"total":    total,
// 		"page":     page,
// 		"pageSize": pageSize,
// 		"lyrics":   lyrics[start:end],
// 	})

// }

// Получить песню по её ID
// func Getsong(c *gin.Context) {
// 	id := c.Param("id")
// 	var song model.Song
// 	db := store.GetDB()

// 	if err := db.First(&song, id).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Song not found!"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, song)
// }

// Обновить данные песни
// func UpdateSong(c *gin.Context) {
// 	id := c.Param("id")
// 	var song model.Song
// 	var input model.Song

// 	db := store.GetDB()

// 	if err := db.First(&song, id).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
// 		return
// 	}

// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	db.Model(&song).Updates(input)
// 	c.JSON(http.StatusOK, song)
// }

// // Удалить песню по её ID
// func DeleteSong(c *gin.Context) {
// 	// Body-dan ID olish
// 	var body struct {
// 		ID string `json:"ID"` // JSON body dan ID ni olish
// 	}

// 	if err := c.ShouldBindJSON(&body); err != nil {
// 		// Agar format noto'g'ri bo'lsa
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
// 		return
// 	}

// 	// ID ni UUID formatiga aylantirish
// 	parsedID, err := uuid.Parse(body.ID)
// 	if err != nil {
// 		// Agar ID noto'g'ri formatda bo'lsa
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
// 		return
// 	}

// 	// DBga ulanish
// 	db := store.GetDB()

// 	// Yozuvni bazada qidirish
// 	var user types.TUser
// 	if err := db.First(&user, "id = ?", parsedID).Error; err != nil {
// 		// Agar yozuv topilmasa
// 		if err == gorm.ErrRecordNotFound {
// 			// Log chiqarish
// 			log.Printf("Song with ID %s not found", body.ID)
// 			c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
// 			return
// 		}
// 		// Boshqa xato bo'lsa
// 		log.Printf("Error finding song: %s", err.Error())
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting song"})
// 		return
// 	}

// 	// Yozuvni o'chirish
// 	if err := db.Delete(&user).Error; err != nil {
// 		log.Printf("Error deleting song with ID %s: %s", body.ID, err.Error())
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting song"})
// 		return
// 	}

// 	// Muvaffaqiyatli o'chirish
// 	c.JSON(http.StatusOK, gin.H{"message": "Song deleted"})
// }



// // Удалить все песни
// func AllDeleteSongs(c *gin.Context) {
// 	db := store.GetDB()
// 	var song model.Song

// 	if err := db.First(&song).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Database is empty"})
// 		return
// 	}

// 	db.Where("1=1").Delete(&model.Song{})

// 	c.JSON(http.StatusOK, gin.H{"message": "All songs deleted"})
// }
