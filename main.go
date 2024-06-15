package main

import (
	"flag"
	"fmt"
	tele "gopkg.in/telebot.v3"
	"log"
	"time"
)

type Config struct {
	LoginDetails []struct {
		Server   string `toml:"server"`
		Port     int16  `toml:"Port"`
		Username string `toml:"username"`
		Password string `toml:"password"`
	}

	TelegramBotToken string `toml:"TELEGRAM_BOT_TOKEN"`
	ChatID           int64  `toml:"CHAT_ID"`
	Interval         int64  `toml:"INTERVAL"`
}

func main() {
	configFile := flag.String("config", "config.toml", "")
	flag.Parse()
	config, err := LoadConfig(*configFile)
	if err != nil {
		log.Fatalln(err)
	}

	pref := tele.Settings{
		Token:  config.TelegramBotToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	fmt.Println("STARTING BOT.....")
	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/check", func(c tele.Context) error {
		msg, err := IMAPOperation(config)
		if err != nil {
			return c.Send(fmt.Sprintf("Error: %v", err))
		}
		return c.Send(msg)
	})

	b.Handle("/start", func(c tele.Context) error {
		return c.Send(" WELCOME, PLEASE USE THE MENU ")
	})
	go autoTrigger(config, b)
	b.Start()
}

func autoTrigger(config Config, b *tele.Bot) {
	for {
		msg, err := IMAPOperation(config)
		if err != nil {
			log.Printf("Error in fetching mail: %v", err)
		} else {
			chatID := config.ChatID
			_, err := b.Send(&tele.Chat{ID: chatID}, msg)
			if err != nil {
				log.Printf("Error Sending Notification: %v", err)
			}
		}
		time.Sleep(time.Duration(config.Interval) * time.Hour)
	}

}
