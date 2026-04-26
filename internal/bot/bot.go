package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/taha/deep-bot/internal/ticket"
)

type Config struct {
	Token       string
	WebsiteURL  string
	PricingURL  string
	YouTubeURL  string
	AdminChatID int64
}

type Bot struct {
	api   *tgbotapi.BotAPI
	cfg   Config
	store *ticket.Store
}

func New(cfg Config, store *ticket.Store) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		return nil, err
	}
	return &Bot{api: api, cfg: cfg, store: store}, nil
}

func (b *Bot) Run() {
	log.Printf("authorized as @%s", b.api.Self.UserName)

	cfg := tgbotapi.NewUpdate(0)
	cfg.Timeout = 30
	updates := b.api.GetUpdatesChan(cfg)

	for u := range updates {
		switch {
		case u.CallbackQuery != nil:
			b.handleCallback(u.CallbackQuery)
		case u.Message != nil && u.Message.IsCommand():
			b.handleCommand(u.Message)
		case u.Message != nil:
			b.handleText(u.Message)
		}
	}
}

func (b *Bot) send(chatID int64, text string, kb tgbotapi.InlineKeyboardMarkup) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = tgbotapi.ModeHTML
	msg.DisableWebPagePreview = true
	msg.ReplyMarkup = kb
	if _, err := b.api.Send(msg); err != nil {
		log.Printf("send: %v", err)
	}
}

func (b *Bot) edit(chatID int64, msgID int, text string, kb tgbotapi.InlineKeyboardMarkup) {
	edit := tgbotapi.NewEditMessageTextAndMarkup(chatID, msgID, text, kb)
	edit.ParseMode = tgbotapi.ModeHTML
	edit.DisableWebPagePreview = true
	if _, err := b.api.Send(edit); err != nil {
		// Falling back to a fresh message keeps the UX flowing if the original was deleted.
		b.send(chatID, text, kb)
	}
}

func (b *Bot) ack(cb *tgbotapi.CallbackQuery, text string) {
	c := tgbotapi.NewCallback(cb.ID, text)
	if _, err := b.api.Request(c); err != nil {
		log.Printf("ack: %v", err)
	}
}
