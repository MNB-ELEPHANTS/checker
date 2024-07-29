package writer

import (
	"log"

	"github.com/go-telegram/bot"
)

type Writer struct {
	ChatID int64
	Bot    *bot.Bot
}

func New(token string, chat_id int64) *Writer {
	w := &Writer{
		ChatID: chat_id,
	}

	opts := []bot.Option{}

	b, err := bot.New(token, opts...)

	if err != nil {
		log.Fatal(err)
	}

	w.Bot = b

	return w
}
