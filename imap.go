package main

import (
	"fmt"
	"github.com/emersion/go-imap/v2"
	"github.com/emersion/go-imap/v2/imapclient"
	"log"
)

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
		fmt.Println("Notification Sent for ", ImapInfo.Username)
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
