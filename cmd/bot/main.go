package main

import (
	// "log"
	"s21/surdo/internal/bot/buttons"
	"s21/surdo/pkg/api"
	"s21/surdo/pkg/store"
	"time"

	// "your_project/api"
	// "your_project/buttons"
	// "your_project/store"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
    store.InitDB() // Подключение к базе данных
    router := gin.Default()

    // CORS middleware ni qo'shing
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://127.0.0.1:5501"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

    buttons.HelloButton()

    // Swagger
    // router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    router.GET("/users", api.GetAllUsers)   
    router.GET("profile", api.GetUserByTelegramID)            // Получить все песни
    router.POST("/users", api.CreateUser)               // Добавить новую песню
    // router.GET("/songs/:id/lyrics", api.GetSongsLyrics) // Получить текст песни по id
    // router.GET("/songs/:id", api.Getsong)               // Получить песню по id
    router.PUT("/users/", api.UpdateUser)            // Обновить данные песни
    router.DELETE("/users", api.DeleteUser)         // Удалить песню по id
    // router.DELETE("/songs/all", api.AllDeleteSongs)     // Удалить все песни

    // Запуск сервера на порту 8080
    router.Run(":8080")
}
// package main

// import (
// 	// "log"
// 	//  "s21/surdo/internal/api"

// 	"s21/surdo/internal/bot/buttons"


// 	// "s21/surdo/internal/bot/config"
// 	"s21/surdo/pkg/api"
// 	"s21/surdo/pkg/store"

// 	//  "s21/surdo/internal/db"

// 	"github.com/gin-gonic/gin"
// 	// "github.com/swaggo/files"
// 	// ginSwagger "github.com/swaggo/gin-swagger"
// )

// // @title Music Library API
// // @version 1.0
// // @description API для музыкальной библиотеки.
// // @termsOfService http://swagger.io/terms/

// // @contact.name Поддержка API
// // @contact.url http://www.swagger.io/support
// // @contact.email support@swagger.io

// // @host localhost:8080
// // @BasePath /
// func main() {
//  store.InitDB() // Подключение к базе данных
//  router := gin.Default()

//  buttons.HelloButton()




//  // Swagger
// //  router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

//  router.GET("/users", api.GetAllUsers)               // Получить все песни
//  router.POST("/users", api.CreateUser)               // Добавить новую песню
//  // router.GET("/songs/:id/lyrics", api.GetSongsLyrics) // Получить текст песни по id
//  // router.GET("/songs/:id", api.Getsong)               // Получить песню по id
//  router.PUT("/users/", api.UpdateUser)            // Обновить данные песни
//  router.DELETE("/users", api.DeleteUser)         // Удалить песню по id
//  // router.DELETE("/songs/all", api.AllDeleteSongs)     // Удалить все песни

//  // Запуск сервера на порту 8080
//  router.Run(":8080")
// }

// package main

// import (
// 	"log"
// 	// "net/http"
// 	"s21/surdo/internal/bot/buttons"
// 	"s21/surdo/pkg/api"
// 	"s21/surdo/pkg/store"

// 	"github.com/gin-contrib/cors"
// 	"github.com/gin-gonic/gin"
// 	// "github.com/rs/cors"
// 	// Для генерации документации
// )

// func main() {
// 	store.InitDB() // Baza bilan bog'lanish

// 	router := gin.Default()

// 	// `buttons.HelloButton()`ni chaqirish
// 	buttons.HelloButton()

// 	// CORS sozlamalarini qo'shish
// 	router.Use(cors.New(cors.Config{
// 		AllowOrigins:     []string{"http://127.0.0.1:5500", "https://f870-217-30-173-218.ngrok-free.app/web/"}, // Faqat shu origin-dan so'rovlar qabul qilinadi
// 		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"}, // Ruxsat berilgan metodlar
// 		AllowHeaders:     []string{"Content-Type", "Authorization"}, // Ruxsat berilgan headerlar
// 		AllowCredentials: true, // Agar kerak bo'lsa, cookie'lar bilan ishlash uchun
// 	}))

// 	// API marshrutlari
// 	router.GET("/users", api.GetAllSongs)          // Barcha foydalanuvchilarni olish
// 	router.POST("/users", api.CreateSong)          // Yangi foydalanuvchi qo'shish
// 	router.DELETE("/users", api.DeleteSong)        // Foydalanuvchini o'chirish

// 	// Swagger qo'shish
// 	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

// 		// Serverni ishga tushirish
// 		if err := router.Run(":8080"); err != nil {
// 			log.Fatal("Serverni ishga tushirishda xatolik yuz berdi:", err)
// 		}
// 	}





// package main

// import (
//     "log"
//     "net/http"
//     "s21/surdo/internal/bot/buttons"
//     // "s21/surdo/pkg/api"
//     types "s21/surdo/pkg/model"
//     "s21/surdo/pkg/store"

//     "github.com/gin-gonic/gin"
//     // "gorm.io/driver/sqlite"
//     "gorm.io/gorm"
// )

// var db *gorm.DB

// // func initDB() {
// //     var err error
// //     db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
// //     if err != nil {
// //         log.Fatalf("Failed to connect database: %v", err)
// //     }
// //     db.AutoMigrate(&types.TUser{})
// //     store.DB = db
// // }

// func CreateUser(c *gin.Context) {
//     var userData types.TUser
//     if err := c.ShouldBindJSON(&userData); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
//     }

//     // Yuborilgan ma'lumotni konsolga chiqarish
//     log.Printf("Qabul qilindi: %+v", userData)

//     // Ma'lumotlarni bazaga saqlash
//     result := store.DB.Create(&userData)
//     if result.Error != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
//         return
//     }

//     // Javobni yuborish
//     c.JSON(http.StatusOK, gin.H{"message": "Ma'lumotlar muvaffaqiyatli qabul qilindi"})
// }

// func GetAllUsers(c *gin.Context) {
//     var users []types.TUser
//     result := store.DB.Find(&users)
//     if result.Error != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
//         return
//     }

//     c.JSON(http.StatusOK, users)
// }

// func main() {
//     // initDB()
// 	store.InitDB()
//     r := gin.Default()
//     buttons.HelloButton()
//     // CORS middleware
//     r.Use(func(c *gin.Context) {
//         c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
//         c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
//         c.Writer.Header().Set("Access-Control-Allow-Headers",
//             "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, "+
//                 "Authorization, accept, origin, Cache-Control, X-Requested-With")
//         c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
//         if c.Request.Method == "OPTIONS" {
//             c.AbortWithStatus(204)
//             return
//         }
//         c.Next()
//     })

//     // Foydalanuvchi yaratish endpointi
//     r.POST("/user", CreateUser)
//     r.GET("/users", GetAllUsers)

//     // Serverni ishga tushirish
//     r.Run(":8080")
// }
