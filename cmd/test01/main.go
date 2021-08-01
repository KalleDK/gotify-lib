package main

import "github.com/KalleDK/gotify-lib/gotify"

func main() {
	client := gotify.New(nil, "https://gotify.example.com", "TOKEN")
	client.Notify(gotify.Message{
		Title:    "Title",
		Message:  "Message",
		Priority: 5,
	})
}
