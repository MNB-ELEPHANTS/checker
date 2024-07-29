package main

import (
	"context"
	"log"
	"mnb/checker/internal/checker"
	"mnb/checker/internal/writer"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	chat_id, err := strconv.ParseInt(os.Getenv("CHAT"), 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	w := writer.New(os.Getenv("TOKEN"), chat_id)

	c := checker.New(
		w,
		15*time.Second,
	)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	c.Run(ctx)
}
