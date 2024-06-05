package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/emersion/go-imap/v2/imapclient"
	"log"
	"os"
)

type Config struct {
	LoginDetails []struct {
		Server   string `toml:"server"`
		Port     int16  `toml:"Port"`
		Username string `toml:"username"`
		Password string `toml:"password"`
	}
}

func main() {
	config, _ := loadConfig("config.toml")
	IMAPOperation(config)
}

func IMAPOperation(config Config) {
	for _, imap := range config.LoginDetails {
		success := verifyIMAPConnection(imap.Server, imap.Port, imap.Username, imap.Password)
		if success {
			log.Println("IMAP connection successful for", imap.Username)
		}
	}

}

func verifyIMAPConnection(server string, port int16, Username, Password string) bool {
	conn, err := imapclient.DialTLS(fmt.Sprintf("%s:%d", server, port), nil)
	if err != nil {
		log.Printf("%v: failed to dial IMAP Server: %v", Username, err)
		return false
	}
	defer conn.Close()

	if err := conn.Login(Username, Password).Wait(); err != nil {
		log.Printf("%v:Failed to login: %v", Username, err)
		return false
	}
	return true
}

func checkMail() {}

func loadConfig(filename string) (Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}
	var config Config
	err = toml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config, nil
}
