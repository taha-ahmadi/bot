package bot

import (
	"fmt"
	"strings"
	"time"

	"github.com/taha/deep-bot/internal/content"
	"github.com/taha/deep-bot/internal/ticket"
)

func welcomeText(name string) string {
	if name == "" {
		name = "trader"
	}
	return fmt.Sprintf(
		"👋 <b>Welcome, %s!</b>\n\n"+
			"You're inside the official <b>Deepcharts™</b> companion bot — your shortcut to "+
			"<b>orderflow trading software for advanced volume analysis</b>.\n\n"+
			"🔍 Follow the footprints of <b>smart money</b>\n"+
			"⚡ <b>0.015s</b> refresh rate, 80+ indicators\n"+
			"🎯 Footprint, DOM, Replay, Trade Copier — all in one\n\n"+
			"Pick an option below 👇",
		escapeHTML(name),
	)
}

func aboutText() string {
	return "ℹ️ <b>About Deepcharts™</b>\n\n" +
		"Deepcharts is the <i>fastest-growing</i> orderflow analysis platform — built so retail " +
		"traders can see what institutions see.\n\n" +
		"<b>What you get:</b>\n" +
		"• 🦴 <b>Deep Print</b> — institutional footprint charts\n" +
		"• 📊 Volume Profiles, VWAPs, Delta\n" +
		"• 🎯 IVB Model + Deep Effort proprietary models\n" +
		"• ⚙️ Trade Copier, Auto-Tracker, Replay\n" +
		"• 🎓 Free orderflow mastery course\n" +
		"• 👥 Weekly sessions with Fabervaale, Andrea Cimi & the Deep Team\n\n" +
		"<i>Trading involves substantial risk and is not appropriate for everyone.</i>"
}

func plansListText() string {
	return "💎 <b>Choose Your Plan</b>\n\n" +
		"All plans unlock the full Deepcharts platform. The longer the commitment, the bigger your savings.\n\n" +
		"🥉 <b>Starter</b>     — $99/mo, monthly\n" +
		"🥈 <b>Pro Trader</b>  — $74/mo, quarterly  ⭐\n" +
		"🥇 <b>Elite Annual</b> — up to <b>4 months FREE</b> 🔥\n\n" +
		"Tap a plan for full details 👇"
}

func planDetailText(p content.Plan) string {
	var b strings.Builder
	if p.Highlight != "" {
		b.WriteString("<i>" + p.Highlight + "</i>\n")
	}
	b.WriteString(fmt.Sprintf("%s <b>%s Plan</b>\n", p.BadgeIcon, p.Name))
	b.WriteString("<i>" + p.Tagline + "</i>\n\n")
	b.WriteString(fmt.Sprintf("💰 <b>%s</b> / month\n", p.PricePerMo))
	b.WriteString("🧾 " + p.BilledAs + "\n\n")
	b.WriteString("<b>What's included:</b>\n")
	for _, f := range p.Features {
		b.WriteString("  ✓ " + f + "\n")
	}
	b.WriteString("\n👉 Tap <b>Buy</b> below to checkout on deepcharts.com")
	return b.String()
}

func faqListText() string {
	return "❓ <b>Frequently Asked Questions</b>\n\n" +
		"Everything you wanted to know about Deepcharts, in one tap. Pick a question:"
}

func faqDetailText(f content.FAQ) string {
	return "❔ <b>" + escapeHTML(f.Question) + "</b>\n\n" + escapeHTML(f.Answer)
}

func ticketsHomeText() string {
	return "🎫 <b>Support Tickets</b>\n\n" +
		"Need help? Open a ticket and our team will get back to you. " +
		"You can also review your past tickets here."
}

func ticketCategoryText() string {
	return "🎫 <b>New Ticket — Step 1 / 3</b>\n\n" +
		"Pick a category for your ticket:"
}

func ticketSubjectText(cat ticket.Category) string {
	return "🎫 <b>New Ticket — Step 2 / 3</b>\n\n" +
		"Category: <b>" + cat.Label() + "</b>\n\n" +
		"📝 <b>Send a short subject</b> (e.g. <i>“Charts not loading on Win11”</i>).\n" +
		"Keep it under 80 characters."
}

func ticketDescriptionText(cat ticket.Category, subject string) string {
	return "🎫 <b>New Ticket — Step 3 / 3</b>\n\n" +
		"Category: <b>" + cat.Label() + "</b>\n" +
		"Subject: <i>" + escapeHTML(subject) + "</i>\n\n" +
		"📄 <b>Now describe the issue in detail.</b> Include error messages, " +
		"OS version, and steps to reproduce if possible."
}

func ticketConfirmText(d *ticket.Draft) string {
	return "🔍 <b>Review your ticket</b>\n\n" +
		"<b>Category:</b> " + d.Category.Label() + "\n" +
		"<b>Subject:</b> " + escapeHTML(d.Subject) + "\n\n" +
		"<b>Description:</b>\n" + escapeHTML(d.Description) + "\n\n" +
		"Submit it?"
}

func ticketCreatedText(t ticket.Ticket) string {
	return fmt.Sprintf(
		"✅ <b>Ticket submitted!</b>\n\n"+
			"🆔 <code>%s</code>\n"+
			"📂 %s\n"+
			"📌 <i>%s</i>\n\n"+
			"You'll be notified here as soon as our team responds. "+
			"Average response time: <b>under 24h</b>.",
		t.ID, t.Category.Label(), escapeHTML(t.Subject),
	)
}

func ticketListText(items []ticket.Ticket) string {
	if len(items) == 0 {
		return "📭 <b>No tickets yet</b>\n\nWhen you open a support ticket, it'll show up here."
	}
	var b strings.Builder
	b.WriteString("📋 <b>Your Tickets</b>\n\n")
	for _, t := range items {
		b.WriteString(fmt.Sprintf(
			"🎫 <code>%s</code>  <i>%s</i>\n"+
				"   %s — %s\n"+
				"   <i>%s</i>\n\n",
			t.ID,
			statusBadge(t.Status),
			t.Category.Label(),
			t.CreatedAt.Format(time.RFC822),
			escapeHTML(t.Subject),
		))
	}
	return b.String()
}

func statusBadge(s ticket.Status) string {
	switch s {
	case ticket.StatusOpen:
		return "🟢 open"
	case ticket.StatusPending:
		return "🟡 pending"
	case ticket.StatusResolved:
		return "🔵 resolved"
	case ticket.StatusClosed:
		return "⚫ closed"
	}
	return string(s)
}

func escapeHTML(s string) string {
	r := strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;")
	return r.Replace(s)
}
