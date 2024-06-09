package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/emersion/go-imap/v2"
	"github.com/emersion/go-imap/v2/imapclient"
	tele "gopkg.in/telebot.v3"
	"log"
	"os"
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
		return c.Send(" WELCOME PLEASE USE THE MENU ")
	})

	b.Start()
}

func LoadConfig(filename string) (Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}
	var config Config
	err = toml.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func EstablishIMAPconn(server string, port int16, Username, Password string) (*imapclient.Client, error) {
	addr := fmt.Sprintf("%s:%d", server, port)
	client, err := imapclient.DialTLS(addr, nil)
	if err != nil {
		return nil, fmt.Errorf("%v: failed to dial IMAP Server: %v", Username, err)
	}

	if err := client.Login(Username, Password).Wait(); err != nil {
		return nil, fmt.Errorf("%v: failed to login: %v", Username, err)
	}

	log.Println("IMAP connection and login successful for", Username)
	return client, nil
}

func IMAPOperation(config Config) (string, error) {
	var msg string
	msg += "UNREAD EMAILS IN ALL YOUR MAILBOX\n\n"
	connections := make(map[string]*imapclient.Client)
	for _, ImapInfo := range config.LoginDetails {
		client, err := EstablishIMAPconn(ImapInfo.Server, ImapInfo.Port, ImapInfo.Username, ImapInfo.Password)
		if err != nil {
			return "", err
		}
		connections[ImapInfo.Username] = client
		unreadCount, err := CheckEmail(client)
		if err != nil {
			return "", err
		}
		msg += fmt.Sprintf("%-30s %4d\n", ImapInfo.Username, *unreadCount)
		if err := client.Logout().Wait(); err != nil {
			return "", fmt.Errorf("failed to logout %v: %v", ImapInfo.Username, err)
		}
	}
	return msg, nil
}

func CheckEmail(client *imapclient.Client) (*uint32, error) {
	options := imap.StatusOptions{NumUnseen: true}
	data, err := client.Status("INBOX", &options).Wait()
	if err != nil {
		return nil, fmt.Errorf("STATUS command failed: %v", err)
	}
	return data.NumUnseen, nil
}
