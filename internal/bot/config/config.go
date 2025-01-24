package config

import (

	"log"
	"os"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v4"
)

func ReadEnv() string{
	
	 err := godotenv.Load(".env")
	 if err != nil{
	  log.Fatalf("Error loading .env file: %s", err)
	 }
	
	
	 greeting := os.Getenv("TOKEN")
	
	//  fmt.Println(greeting)
	return greeting
}



func StartBot() (*tele.Bot, error) {
	b, err := tele.NewBot(tele.Settings{
		Token: ReadEnv(), 
	})
	if err != nil {
		return nil, err
	}
	return b, nil
}