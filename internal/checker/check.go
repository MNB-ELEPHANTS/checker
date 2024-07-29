package checker

import (
	"context"
	"io"
	"mnb/checker/internal/entity"
	"net/http"
	"time"
)

func (c *Checker) Run(ctx context.Context) {
	for {
		c.UpdateList(ctx)

		duration := c.Duration

		for url, code := range c.WebSites {
			resp, err := http.Get(url)

			if err != nil {
				msg := &entity.Message{
					URL:          url,
					Time:         time.Now(),
					ExpectedCode: code,
					Body:         "Не удалось получить (ошибка при запросе)",
				}

				c.Writer.SendMessage(ctx, msg)
				duration = time.Hour

			} else if uint(resp.StatusCode) != code {
				msg := &entity.Message{
					URL:          url,
					StatusCode:   uint(resp.StatusCode),
					ExpectedCode: code,
					Time:         time.Now(),
				}

				body, err := io.ReadAll(resp.Body)

				if err != nil {
					msg.Body = "Не удалось получить"
				}

				msg.Body = string(body)

				c.Writer.SendMessage(ctx, msg)

				duration = time.Hour
			}
		}

		time.Sleep(duration)
	}
}
