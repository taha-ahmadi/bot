# Deepcharts Telegram Bot

A Telegram companion bot for [Deepcharts™](https://www.deepcharts.com/) — orderflow trading software for advanced volume analysis.

## Features

- 💎 **Plans & Pricing** — Starter / Pro Trader / Elite Annual tiers with buy buttons
- 📺 **YouTube** — quick link to the Deepcharts channel
- ❓ **FAQ** — interactive question/answer flow
- 🎫 **Support Tickets** — 3-step intake (category → subject → description), persisted to JSON
- 🌐 **Visit Site** — direct link to deepcharts.com

## Run locally

```bash
cp .env.example .env   # paste your TELEGRAM_BOT_TOKEN
go run .
```

## Run on Replit

1. Import this repo into Replit.
2. Add `TELEGRAM_BOT_TOKEN` in **Secrets**.
3. Click **Run**.

## Project layout

```
main.go
internal/
├── bot/         # Telegram bot wiring, handlers, keyboards, messages
├── content/     # Plans + FAQ data
└── ticket/      # Ticket model + JSON store
```
