package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

func Perform_go_line_sdk() {

	// 1. Configuring linebot
	// 2. Completed Creating a test line dev official account
	// Filling out both channel secret & channel access token below
	client := &http.Client{}
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			// This statment controls whether bot receives message
			if event.Type == linebot.EventTypeMessage {
				// Using switch to check event type is equivalent to message
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})

}

func Perform_DTO_MongoDB() {

}

func main() {
	fmt.Println("Cinnox project:")
	fmt.Println("Step 1. Performing the configuration of LINE & MongoDB")
	Perform_go_line_sdk()
	fmt.Println("Step 2. Connecting to MongoDB & Creating DTO to save/query message to MongoDB.")
	Perform_DTO_MongoDB()
}
