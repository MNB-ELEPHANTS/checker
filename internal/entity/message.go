package entity

import (
	"fmt"
	"time"
)

type Message struct {
	URL          string    `json:"address"`
	StatusCode   uint      `json:"status_code"`
	ExpectedCode uint      `json:"expected_code"`
	Body         string    `json:"body"`
	Time         time.Time `json:"time"`
}

func (m *Message) ToText() string {
	str := fmt.Sprintf(
		"Саня, тут что-то не так...\n\nURL: %s\nСтатус: %d\nОжидаемый статус: %d\nВремя: %s\nТело: %s",
		m.URL,
		m.StatusCode,
		m.ExpectedCode,
		m.Time.Format("15:04:05 02.01.2006"),
		m.Body,
	)

	return str
}
