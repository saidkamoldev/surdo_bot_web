package buttons

import (
	
	"log"
	"s21/surdo/internal/bot/config"
	"s21/surdo/pkg/middleware"
	"fmt"
	// "s21/surdo/pkg/middleware"

	// "strconv"

	// types "s21/surdo/pkg/model"

	tele "gopkg.in/telebot.v4"
)

// var (
// 	// Mini-app URL'sini belgilash
// 	OrderOnlineChooseTimeApp = tele.WebApp{
// 		URL: "https://static.okibiteam.ru/srd/index.html", // Bu yerga mini-app URL'sini qo'yishingiz kerak
// 	}
// )
// testtoviy dataBase
var userData = make(map[int64]map[string]string)

var Telegram_id int64
var Language string
// HelloButton funktsiyasi
func HelloButton() {
    // Botni yaratish
    b, err := config.StartBot()
    if err != nil {
        log.Fatal(err)
        return
    }
	fmt.Println("ZAYBALLLLLLLLLLLLLLLLLLLLLLLLL", Telegram_id)
    b.Handle("/start", startHandler)
    languageButtons := CreateInlineButtons()

    b.Handle(tele.OnCallback, func(c tele.Context) error {
        return callbackHandler(c, languageButtons)
    })

    b.Handle(tele.OnText, showMenu)

    log.Println("Bot ishga tushirildi...")

    // Botni alohida gorutina orqali ishga tushirish
    go func() {
        b.Start()
    }()
}



func startHandler(c tele.Context) error {
    if IdHandler(c) {
        // Agar user baza mavjud bo'lsa showMenu funksiyasini ishga tushiring
        return showMenu(c)
    }

    // Agar user baza mavjud bo'lmasa
    languageButtons := CreateInlineButtons()
    if err := c.Send("Iltimos, bitta tilni tanlang:", languageButtons); err != nil {
        return err
    }
    log.Printf("Foydalanuvchi ID: %d", Telegram_id)

    return nil
}
func IdHandler(c tele.Context) bool {
    Telegram_id = c.Sender().ID
    log.Printf("Foydalanuvchi ID: %d", Telegram_id)

    // Bu yerda userni bazadan tekshirishni qo'shish
    // Agar user baza mavjud bo'lsa true qaytaradi, bo'lmasa false qaytaradi
    userExists, err := middleware.CheckTelegramIDExists(Telegram_id)
    if err !=nil {
        log.Println(userExists)
    }
    // fmt.Println(userExists)
    return userExists
}

// Bu funksiya userni bazadan tekshirish uchun misol

// ABRABOTKA callback
func callbackHandler(c tele.Context, languageButtons *tele.ReplyMarkup) error {
	data := c.Callback().Data
	userID := c.Sender().ID
	
	// userID := c.Message().Sender.ID

	// // TUser modelini yaratish
	// user := types.TUser{
	// 	TelegramID: userID,
	// }
    // Foydalanuvchi ID sini konsolga chiqarish
	// log.Println("Foydalanuvchi ID:", user.TelegramID)

	log.Printf("Callback received: %s", data) // Callbackni loglash
	if data == "uzb" {
        Language = "uzb"
			log.Println("Mini-appni ochishga urinish...")
		err := c.Send("Buyurtma berish uchun tugmani bosing:", CreateOrderButton())
		if err != nil {
			log.Printf("Xatolik: %v", err) // Xatolikni logga yozing
			return err
		}
		log.Println("Mini-app muvaffaqiyatli yuborildi!")
		return nil
		// userData[userID] = make(map[string]string)
		// userData[userID]["language"] = "O'zbek tili"
		// return c.Send("Iltimos, ismingizni kiriting:")
	} else if data == "rus" {
		userData[userID] = make(map[string]string)
		userData[userID]["language"] = "Русский язык"
		return c.Send("Пожалуйста, введите ваше имя:")
	}

	// Menu tugmasi
	if data == "menu" {
		return showMenu(c)
	}

	// Buyurtma tugmasi
	if data == "order" {
		Orders(c)
	}
    
	// Online buyurtma tugmasi
	// "order_online" tugmasi bosilganida mini-appni ochish
// Callbackga ishlov berish
if data == "order_online" {
	log.Println("Mini-appni ochishga urinish...")
	err := c.Send("Registiratsiyadan o'tish:", CreateOrderButton())
	if err != nil {
		log.Printf("Xatolik: %v", err) // Xatolikni logga yozing
		return err
	}
	log.Println("Mini-app muvaffaqiyatli yuborildi!")
	return nil
}

	// Profile tugmasi
	if data == "profile" {
			log.Println("Mini-appni ochishga urinish...")
		err := c.Send("Buyurtma berish uchun tugmani bosing:", infoProfile())
		if err != nil {
			log.Printf("Xatolik: %v", err) // Xatolikni logga yozing
			return err
		}
		log.Println("Mini-app muvaffaqiyatli yuborildi!")
		return nil
	}

	return nil
	
}


func CreateOrderButton() *tele.ReplyMarkup {
    inlineKeys := &tele.ReplyMarkup{}

    // Online buyurtma tugmasi
    // orderButton := inlineKeys.Data( "order_online", "")
    // orderButton.WebApp = &tele.WebApp{
    //     URL: "https://static.okibiteam.ru/srd/index.html",
    // }

    // Mini App tugmasi
    webAppButton := inlineKeys.WebApp("Registratsiya", &tele.WebApp{
        URL: "https://kosmostj.uz/",
    })

    inlineKeys.Inline(
        // inlineKeys.Row(orderButton),
        inlineKeys.Row(webAppButton),
    )
    return inlineKeys
}

// data collection


// language inline buttons
func CreateInlineButtons() *tele.ReplyMarkup {
    inlineKeys := &tele.ReplyMarkup{}

    btn1 := tele.Btn{
        Text: "O'zbek tili",
        Data: "uzb", 
    }
    btn2 := tele.Btn{
        Text: "Русский язык",
        Data: "rus", // Callback data
    }

    inlineKeys.Inline(
        inlineKeys.Row(btn1, btn2), 
    )

    return inlineKeys
}


// Main menu
func showMenu(c tele.Context) error {

    menuButtons := &tele.ReplyMarkup{}

  
    orderBtn := tele.Btn{
        Text: "Tarjimon buyurtma berish",
        Data: "order", // Callback data
    }
    helpBtn := tele.Btn{
        Text: "SOS",
        Data: "help", 
    }
    supBtn := tele.Btn{
        Text: "Texnik yordam",
        Data: "support", 
    }

	profileBtn := tele.Btn{
        Text: "Profile",
        Data: "profile", 
    }


    menuButtons.Inline(
        menuButtons.Row(orderBtn, profileBtn,),
        menuButtons.Row(supBtn,helpBtn),
    )

    return c.Send("Menyu:", menuButtons)
}

func Orders(c tele.Context) error {
	orders := &tele.ReplyMarkup{}

    onlineBtn := tele.Btn{
        Text: "online",
        Data: "order_online", // Callback data
    }
    offlineBtn := tele.Btn{
        Text: "offline",
        Data: "order_offline", // Callback data
    }


  
    orders.Inline(
        orders.Row(onlineBtn, offlineBtn,),
      
    )

    return c.Send("adsadas", orders)
}



// input user_info
func infoProfile()*tele.ReplyMarkup {
    inlineKeys := &tele.ReplyMarkup{}

    // Online buyurtma tugmasi
    // orderButton := inlineKeys.Data( "order_online", "")
    // orderButton.WebApp = &tele.WebApp{
    //     URL: "https://static.okibiteam.ru/srd/index.html",
    // }

    // Mini App tugmasi
    webAppButton := inlineKeys.WebApp("Registratsiya", &tele.WebApp{
        URL: "https://kosmostj.uz/profile/index.html",
    })

    inlineKeys.Inline(
        // inlineKeys.Row(orderButton),
        inlineKeys.Row(webAppButton),
    )
    return inlineKeys
}

