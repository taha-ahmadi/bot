package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/taha/deep-bot/internal/bot"
	"github.com/taha/deep-bot/internal/ticket"
)

func main() {
	loadDotEnv(".env")

	cfg := bot.Config{
		Token:       os.Getenv("TELEGRAM_BOT_TOKEN"),
		WebsiteURL:  envOr("DEEPCHARTS_URL", "https://www.deepcharts.com/"),
		PricingURL:  envOr("DEEPCHARTS_PRICING_URL", "https://www.deepcharts.com/pricing"),
		YouTubeURL:  envOr("DEEPCHARTS_YOUTUBE_URL", "https://www.youtube.com/@deepcharts-official"),
		AdminChatID: envInt64("ADMIN_CHAT_ID"),
	}
	if cfg.Token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is required")
	}

	store, err := ticket.NewStore(envOr("TICKETS_FILE", "./data/tickets.json"))
	if err != nil {
		log.Fatalf("ticket store: %v", err)
	}

	b, err := bot.New(cfg, store)
	if err != nil {
		log.Fatalf("bot init: %v", err)
	}
	b.Run()
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func envInt64(key string) int64 {
	v := os.Getenv(key)
	if v == "" {
		return 0
	}
	n, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0
	}
	return n
}

func loadDotEnv(path string) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		k, v, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		k = strings.TrimSpace(k)
		v = strings.Trim(strings.TrimSpace(v), `"'`)
		if _, exists := os.LookupEnv(k); !exists {
			os.Setenv(k, v)
		}
	}
}
