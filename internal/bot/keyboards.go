package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/taha/deep-bot/internal/content"
	"github.com/taha/deep-bot/internal/ticket"
)

func mainMenu(cfg Config) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("💎 Plans & Pricing", "plans"),
			tgbotapi.NewInlineKeyboardButtonData("🎫 Support Tickets", "tickets"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("❓ FAQ", "faq"),
			tgbotapi.NewInlineKeyboardButtonData("ℹ️ About", "about"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("📺 YouTube Channel", cfg.YouTubeURL),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("🌐 Visit Deepcharts.com", cfg.WebsiteURL),
		),
	)
}

func plansMenu() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🥉 Starter", "plan:starter"),
			tgbotapi.NewInlineKeyboardButtonData("🥈 Pro Trader", "plan:pro"),
			tgbotapi.NewInlineKeyboardButtonData("🥇 Elite Annual", "plan:elite"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬅ Back to Menu", "menu"),
		),
	)
}

func planDetail(p content.Plan) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("💳  Buy "+p.Name+"  →", p.BuyURL),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬅ All Plans", "plans"),
			tgbotapi.NewInlineKeyboardButtonData("🏠 Menu", "menu"),
		),
	)
}

func faqMenu() tgbotapi.InlineKeyboardMarkup {
	rows := [][]tgbotapi.InlineKeyboardButton{}
	for _, f := range content.FAQs() {
		rows = append(rows, tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("❔ "+f.Question, "faq:"+f.ID),
		))
	}
	rows = append(rows, tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("⬅ Back to Menu", "menu"),
	))
	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}

func faqDetail() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬅ All Questions", "faq"),
			tgbotapi.NewInlineKeyboardButtonData("🏠 Menu", "menu"),
		),
	)
}

func ticketsMenu() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("➕ Open New Ticket", "ticket:new"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("📋 My Tickets", "ticket:list"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬅ Back to Menu", "menu"),
		),
	)
}

func ticketCategoryMenu() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("💳 Billing", "ticket:cat:"+string(ticket.CategoryBilling)),
			tgbotapi.NewInlineKeyboardButtonData("🛠 Technical", "ticket:cat:"+string(ticket.CategoryTechnical)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("💡 Feature", "ticket:cat:"+string(ticket.CategoryFeature)),
			tgbotapi.NewInlineKeyboardButtonData("💬 General", "ticket:cat:"+string(ticket.CategoryGeneral)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("✖ Cancel", "ticket:cancel"),
		),
	)
}

func ticketCancelOnly() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("✖ Cancel Draft", "ticket:cancel"),
		),
	)
}

func ticketConfirm() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("✅ Submit Ticket", "ticket:submit"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("✖ Cancel", "ticket:cancel"),
		),
	)
}

func ticketDoneMenu() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("📋 My Tickets", "ticket:list"),
			tgbotapi.NewInlineKeyboardButtonData("🏠 Menu", "menu"),
		),
	)
}

func backToMenu() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🏠 Menu", "menu"),
		),
	)
}
