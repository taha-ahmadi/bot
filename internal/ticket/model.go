package ticket

import "time"

type Status string

const (
	StatusOpen     Status = "open"
	StatusPending  Status = "pending"
	StatusResolved Status = "resolved"
	StatusClosed   Status = "closed"
)

type Category string

const (
	CategoryBilling   Category = "billing"
	CategoryTechnical Category = "technical"
	CategoryGeneral   Category = "general"
	CategoryFeature   Category = "feature"
)

func (c Category) Label() string {
	switch c {
	case CategoryBilling:
		return "💳 Billing & Subscription"
	case CategoryTechnical:
		return "🛠 Technical Issue"
	case CategoryFeature:
		return "💡 Feature Request"
	default:
		return "💬 General Question"
	}
}

type Ticket struct {
	ID          string    `json:"id"`
	UserID      int64     `json:"user_id"`
	Username    string    `json:"username"`
	Category    Category  `json:"category"`
	Subject     string    `json:"subject"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type DraftStep string

const (
	StepNone        DraftStep = ""
	StepCategory    DraftStep = "category"
	StepSubject     DraftStep = "subject"
	StepDescription DraftStep = "description"
	StepConfirm     DraftStep = "confirm"
)

type Draft struct {
	UserID      int64
	Username    string
	Step        DraftStep
	Category    Category
	Subject     string
	Description string
}
