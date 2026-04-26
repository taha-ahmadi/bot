package bot

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/taha/deep-bot/internal/content"
	"github.com/taha/deep-bot/internal/ticket"
)

func (b *Bot) handleCommand(m *tgbotapi.Message) {
	switch m.Command() {
	case "start", "menu":
		b.store.ResetDraft(m.From.ID)
		b.send(m.Chat.ID, welcomeText(displayName(m.From)), mainMenu(b.cfg))
	case "plans", "pricing":
		b.send(m.Chat.ID, plansListText(), plansMenu())
	case "faq":
		b.send(m.Chat.ID, faqListText(), faqMenu())
	case "ticket", "support":
		b.send(m.Chat.ID, ticketsHomeText(), ticketsMenu())
	case "tickets":
		items := b.store.ListByUser(m.From.ID)
		b.send(m.Chat.ID, ticketListText(items), ticketsMenu())
	case "about":
		b.send(m.Chat.ID, aboutText(), backToMenu())
	case "cancel":
		b.store.ResetDraft(m.From.ID)
		b.send(m.Chat.ID, "✖ Draft cancelled.", mainMenu(b.cfg))
	default:
		b.send(m.Chat.ID, "Unknown command. Try /start.", mainMenu(b.cfg))
	}
}

func (b *Bot) handleText(m *tgbotapi.Message) {
	d := b.store.Draft(m.From.ID)
	switch d.Step {
	case ticket.StepSubject:
		text := strings.TrimSpace(m.Text)
		if len(text) == 0 || len(text) > 80 {
			b.send(m.Chat.ID, "⚠️ Subject must be 1–80 characters. Try again.", ticketCancelOnly())
			return
		}
		d.Subject = text
		d.Username = displayName(m.From)
		d.Step = ticket.StepDescription
		b.send(m.Chat.ID, ticketDescriptionText(d.Category, d.Subject), ticketCancelOnly())
	case ticket.StepDescription:
		text := strings.TrimSpace(m.Text)
		if len(text) < 10 {
			b.send(m.Chat.ID, "⚠️ Description is too short — please give us at least 10 characters.", ticketCancelOnly())
			return
		}
		if len(text) > 4000 {
			text = text[:4000]
		}
		d.Description = text
		d.Step = ticket.StepConfirm
		b.send(m.Chat.ID, ticketConfirmText(d), ticketConfirm())
	default:
		b.send(m.Chat.ID, "Tap a button or use /start to open the menu.", mainMenu(b.cfg))
	}
}

func (b *Bot) handleCallback(cb *tgbotapi.CallbackQuery) {
	data := cb.Data
	chatID := cb.Message.Chat.ID
	msgID := cb.Message.MessageID

	switch {
	case data == "menu":
		b.store.ResetDraft(cb.From.ID)
		b.ack(cb, "")
		b.edit(chatID, msgID, welcomeText(displayName(cb.From)), mainMenu(b.cfg))
	case data == "plans":
		b.ack(cb, "")
		b.edit(chatID, msgID, plansListText(), plansMenu())
	case strings.HasPrefix(data, "plan:"):
		id := strings.TrimPrefix(data, "plan:")
		p, ok := content.PlanByID(id, b.cfg.PricingURL)
		if !ok {
			b.ack(cb, "Plan not found")
			return
		}
		b.ack(cb, "")
		b.edit(chatID, msgID, planDetailText(p), planDetail(p))
	case data == "faq":
		b.ack(cb, "")
		b.edit(chatID, msgID, faqListText(), faqMenu())
	case strings.HasPrefix(data, "faq:"):
		id := strings.TrimPrefix(data, "faq:")
		f, ok := content.FAQByID(id)
		if !ok {
			b.ack(cb, "Question not found")
			return
		}
		b.ack(cb, "")
		b.edit(chatID, msgID, faqDetailText(f), faqDetail())
	case data == "about":
		b.ack(cb, "")
		b.edit(chatID, msgID, aboutText(), backToMenu())
	case data == "tickets":
		b.ack(cb, "")
		b.edit(chatID, msgID, ticketsHomeText(), ticketsMenu())
	case data == "ticket:new":
		d := b.store.Draft(cb.From.ID)
		d.Step = ticket.StepCategory
		b.ack(cb, "")
		b.edit(chatID, msgID, ticketCategoryText(), ticketCategoryMenu())
	case strings.HasPrefix(data, "ticket:cat:"):
		cat := ticket.Category(strings.TrimPrefix(data, "ticket:cat:"))
		d := b.store.Draft(cb.From.ID)
		d.Category = cat
		d.Step = ticket.StepSubject
		b.ack(cb, "Category set")
		b.edit(chatID, msgID, ticketSubjectText(cat), ticketCancelOnly())
	case data == "ticket:list":
		items := b.store.ListByUser(cb.From.ID)
		b.ack(cb, "")
		b.edit(chatID, msgID, ticketListText(items), ticketsMenu())
	case data == "ticket:cancel":
		b.store.ResetDraft(cb.From.ID)
		b.ack(cb, "Draft cancelled")
		b.edit(chatID, msgID, "✖ Ticket draft cancelled.", ticketsMenu())
	case data == "ticket:submit":
		d := b.store.Draft(cb.From.ID)
		if d.Step != ticket.StepConfirm {
			b.ack(cb, "Nothing to submit")
			return
		}
		t, err := b.store.Create(ticket.Ticket{
			UserID:      cb.From.ID,
			Username:    displayName(cb.From),
			Category:    d.Category,
			Subject:     d.Subject,
			Description: d.Description,
		})
		if err != nil {
			b.ack(cb, "Could not save")
			b.edit(chatID, msgID, "⚠️ Sorry, we couldn't save your ticket. Please try again.", ticketsMenu())
			return
		}
		b.store.ResetDraft(cb.From.ID)
		b.ack(cb, "Submitted ✓")
		b.edit(chatID, msgID, ticketCreatedText(t), ticketDoneMenu())
		b.notifyAdmin(t)
	default:
		b.ack(cb, "")
	}
}

func (b *Bot) notifyAdmin(t ticket.Ticket) {
	if b.cfg.AdminChatID == 0 {
		return
	}
	msg := fmt.Sprintf(
		"🆕 <b>New ticket</b> <code>%s</code>\n"+
			"From: %s (id %d)\n"+
			"%s\n"+
			"<i>%s</i>\n\n%s",
		t.ID, escapeHTML(t.Username), t.UserID, t.Category.Label(),
		escapeHTML(t.Subject), escapeHTML(t.Description),
	)
	b.send(b.cfg.AdminChatID, msg, tgbotapi.InlineKeyboardMarkup{})
}

func displayName(u *tgbotapi.User) string {
	if u == nil {
		return ""
	}
	if u.UserName != "" {
		return "@" + u.UserName
	}
	name := u.FirstName
	if u.LastName != "" {
		name += " " + u.LastName
	}
	return name
}
