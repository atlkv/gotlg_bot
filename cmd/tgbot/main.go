package main

import (	
	//"tgbot/internal/config"
	"tgbot/internal/bot"
	"log"
)

func main() {
	// config, err := config.NewConfig("local")
	
	// if err != nil {
	// 	log.Panicf("Config error: %s", err)
	// }

	// token, err := config.GetTelegramToken()

	// if err != nil {
	// 	log.Panicf("Can't get telegram token: %s", err)
	// }

	 bot, err := bot.NewBot("5047366335:AAFOoTRzS6CZmDbJ65_PIO5NTe_SKw9tFyI")

	log.Printf("Success auth by: %s", bot.UserName())

	 if err != nil {
		log.Panicf("Error initializing Telegram Bot: %s", err)
	 }

	 bot.Run(60)
}