package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

type Config struct {
	LoginDetails []struct {
		server   string
		port     int16
		username string
		password string
	}
}

func main() {
	fmt.Println("hello world")
}

func verifyIMAP() {

}

func checkMail() {}

func readTOML() []byte {
	data, err := os.ReadFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func decodeTOML() Config {
	var config Config
	data := readTOML()
	_, err := toml.Decode(string(data), &config)
	if err != nil {
		log.Fatal("Error decoding TOML:", err)
	}
	return config
}
