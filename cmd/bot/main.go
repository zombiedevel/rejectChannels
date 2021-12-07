package main

import (
	"github.com/zombiedevel/rejectChannels/internal/bot"
	"log"
)

func main() {
	b, err := bot.New("bot.yml")
	if err != nil {
		log.Fatal(err)
	}

	b.Start()
}
