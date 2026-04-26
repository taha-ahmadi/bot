package content

type Plan struct {
	ID          string
	Name        string
	Tagline     string
	PricePerMo  string
	BilledAs    string
	Highlight   string
	Features    []string
	BadgeIcon   string
	BuyURL      string
}

func Plans(pricingURL string) []Plan {
	return []Plan{
		{
			ID:         "starter",
			Name:       "Starter",
			Tagline:    "Try Deepcharts month-to-month",
			PricePerMo: "$99",
			BilledAs:   "Billed monthly · cancel anytime",
			BadgeIcon:  "🥉",
			Features: []string{
				"Deep Print (footprint analysis)",
				"Volume Profiles & VWAPs",
				"Delta + Volume indicators",
				"Deep Trades + Deep V-Tracker",
				"15-min delayed data feed",
				"Free orderflow mastery course",
			},
			BuyURL: pricingURL,
		},
		{
			ID:         "pro",
			Name:       "Pro Trader",
			Tagline:    "Best balance — quarterly billing",
			PricePerMo: "$74",
			BilledAs:   "$222 every 3 months · save 25%",
			Highlight:  "⭐ MOST POPULAR",
			BadgeIcon:  "🥈",
			Features: []string{
				"Everything in Starter",
				"IVB Model + Deep Effort",
				"Automatic risk management",
				"Trade Copier (beta)",
				"Unlimited charts & templates",
				"Weekly Fabervaale & Andrea Cimi sessions",
			},
			BuyURL: pricingURL,
		},
		{
			ID:         "elite",
			Name:       "Elite Annual",
			Tagline:    "Up to 4 months FREE",
			PricePerMo: "From $66",
			BilledAs:   "Billed yearly · best value",
			Highlight:  "🔥 LIMITED PROMO",
			BadgeIcon:  "🥇",
			Features: []string{
				"Everything in Pro Trader",
				"Up to 4 months free with annual commit",
				"Priority support queue",
				"Monthly competitions entry",
				"Weekly Deep Team sessions",
				"Early access to new indicators",
			},
			BuyURL: pricingURL,
		},
	}
}

func PlanByID(id, pricingURL string) (Plan, bool) {
	for _, p := range Plans(pricingURL) {
		if p.ID == id {
			return p, true
		}
	}
	return Plan{}, false
}
