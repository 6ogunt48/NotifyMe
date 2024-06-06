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

func EstablishIMAPconn(server string, port int16, Username, Password string) (*imapclient.Client, error) {
	addr := fmt.Sprintf("%s:%d", server, port)
	client, err := imapclient.DialTLS(addr, nil)
	if err != nil {
		log.Printf("%v: failed to dial IMAP Server: %v", Username, err)
		return nil, err
	}

	if err := client.Login(Username, Password).Wait(); err != nil {
		log.Printf("%v:Failed to login: %v", Username, err)
		client.Logout()
		return nil, nil
	} else {
		log.Println("IMAP connection and login successful for", Username)
	}

	return client, nil
}

func IMAPOperation(config Config) {
	count := 0
	connections := make(map[string]*imapclient.Client)
	for _, imap := range config.LoginDetails {
		client, _ := EstablishIMAPconn(imap.Server, imap.Port, imap.Username, imap.Password)
		if client != nil {
			connections[imap.Username] = client
			count++
		}
	}
	if count != len(config.LoginDetails) {
		log.Fatalln("all login details in config has to be successful  before operation")
	}
}
