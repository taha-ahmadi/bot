package content

type FAQ struct {
	ID       string
	Question string
	Answer   string
}

func FAQs() []FAQ {
	return []FAQ{
		{
			ID:       "what",
			Question: "What is Deepcharts?",
			Answer: "Deepcharts™ is the fastest-growing orderflow analysis and trading software. " +
				"It gives retail traders institutional-grade tools to follow the footprints of smart " +
				"money — footprint charts, volume profiles, DOM trading, replay, and more — all in one platform.",
		},
		{
			ID:       "who",
			Question: "Who is it for?",
			Answer: "Active and aspiring futures, forex, and options traders who want precision " +
				"entries/exits based on real institutional flow rather than lagging indicators. " +
				"From scalpers to swing traders.",
		},
		{
			ID:       "platform",
			Question: "Which platforms are supported?",
			Answer: "Currently Windows-only. macOS / Linux support is on the roadmap.",
		},
		{
			ID:       "data",
			Question: "Is the data real-time?",
			Answer: "The Full Plan includes a 15-minute delayed feed. Real-time data is available " +
				"by connecting your own data feed via the Data Feeds section on deepcharts.com.",
		},
		{
			ID:       "trial",
			Question: "Is there a free trial?",
			Answer: "There is no traditional free trial, but the Elite Annual promo gives you up to " +
				"4 months free — effectively a long-form trial at a steep discount.",
		},
		{
			ID:       "cancel",
			Question: "Can I cancel anytime?",
			Answer: "Yes. The Starter monthly plan can be cancelled at any time from your account dashboard at my.deepcharts.com.",
		},
		{
			ID:       "diff",
			Question: "How is it different from TradingView?",
			Answer: "Deepcharts is built specifically for orderflow: footprint charts, real DOM, " +
				"trade copier, IVB Model, and 0.015s refresh — features general-purpose charting " +
				"platforms don't offer.",
		},
		{
			ID:       "ticket",
			Question: "How do I get help?",
			Answer: "Use the 🎫 Open Ticket button in this bot — pick a category (Billing / Technical / General), " +
				"describe your issue, and our team will get back to you. You can also visit my.deepcharts.com.",
		},
	}
}

func FAQByID(id string) (FAQ, bool) {
	for _, f := range FAQs() {
		if f.ID == id {
			return f, true
		}
	}
	return FAQ{}, false
}
