package main

import (
	"fmt"
	"github.com/Arman92/go-tdlib"
	"log"
)

func main() {
	// create new instance of TDLib client
	client := tdlib.NewClient(tdlib.Config{
		APIID:               "23161762",
		APIHash:             "34e8d54e22cb4a23c5122a723ca38559",
		SystemLanguageCode:  "en",
		DeviceModel:         "Desktop",
		SystemVersion:       "1.0",
		ApplicationVersion:  "1.0",
		UseMessageDatabase:  true,
		UseFileDatabase:     true,
		UseChatInfoDatabase: true,
		UseTestDataCenter:   false,
	})

	// connect to the TDLib client
	err := client.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to Telegram: %v", err)
	}
	defer client.Disconnect()

	// create a new chat with yourself
	chat, err := client.CreatePrivateChat("username")
	if err != nil {
		log.Fatalf("Failed to create chat: %v", err)
	}

	// send a message to the chat
	_, err = client.SendMessage(chat.ID, &tdlib.InputMessageText{
		Text: tdlib.NewFormattedText("Hello, World!", nil),
	})
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	fmt.Println("Message sent successfully")
}
