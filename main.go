package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/emersion/go-imap/v2"
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
	config, _ := LoadConfig("config.toml")
	IMAPOperation(config)
}

func LoadConfig(filename string) (Config, error) {
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
		log.Fatalf("%v:Failed to login: %v", Username, err)
	} else {
		log.Println("IMAP connection and login successful for", Username)
	}

	return client, nil
}

func IMAPOperation(config Config) {
	count := 0
	connections := make(map[string]*imapclient.Client)
	for _, ImapInfo := range config.LoginDetails {
		client, _ := EstablishIMAPconn(ImapInfo.Server, ImapInfo.Port, ImapInfo.Username, ImapInfo.Password)
		if client != nil {
			connections[ImapInfo.Username] = client
			count++
		}
		CheckEmail(connections, ImapInfo.Username)
	}
}

func CheckEmail(client map[string]*imapclient.Client, email string) *uint32 {
	var UnreadCount *uint32
	imapClient := client[email]
	options := imap.StatusOptions{NumUnseen: true}
	if data, err := imapClient.Status("INBOX", &options).Wait(); err != nil {
		log.Fatalf("STATUS command failed: %v", err)
	} else {
		UnreadCount = data.NumUnseen
	}
	return UnreadCount
}
