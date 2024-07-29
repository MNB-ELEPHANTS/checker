package writer

import (
	"context"
	"mnb/checker/internal/entity"

	"github.com/go-telegram/bot"
)

func (w *Writer) SendMessage(ctx context.Context, msg *entity.Message) {
	w.Bot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: w.ChatID,
		Text:   msg.ToText(),
	})
}
