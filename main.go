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
	config, err := loadConfig("config.toml")
	if err != nil {
		log.Fatal("Error loading config:", err)
	}
	IMAPOperation(config)
}

func IMAPOperation(config Config) {
	for _, imap := range config.LoginDetails {
		success := verifyIMAPConnection(imap.Server, imap.Port, imap.Username, imap.Password)
		if success {
			log.Println("IMAP connection successful for", imap.Username)
		} else {
			log.Println("Program wont continue if one or more IMAP returns an error")
			log.Fatalln("IMAP Connection failed for", imap.Username)
		}
	}

}

func verifyIMAPConnection(server string, port int16, Username, Password string) bool {
	conn, err := imapclient.DialTLS(fmt.Sprintf("%s:%d\n", server, port), nil)
	if err != nil {
		log.Printf("failed to dial IMAP Server: %v\n", err)
		return false
	}
	defer conn.Close()

	if err := conn.Login(Username, Password).Wait(); err != nil {
		log.Printf("Failed to login: %v\n", err)
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
