package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentStatus string

const (
	LoanStatusActive  PaymentStatus = "active"
	LoanStatusPaid    PaymentStatus = "paid"
	LoanStatusOverdue PaymentStatus = "overdue"
)

type PaymentHistoryStatus string

func (p *PaymentHistoryStatus) String() string {
	return string(*p)
}

const (
	PaymentHistoryStatusActive  PaymentHistoryStatus = "waiting"
	PaymentHistoryStatusPaid    PaymentHistoryStatus = "paid"
	PaymentHistoryStatusOverdue PaymentHistoryStatus = "overdue"
)

type PaymentHistory struct {
	ID            primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	LoanID        primitive.ObjectID   `json:"loan_id" bson:"loan_id"`
	Amount        float64              `json:"amount" bson:"amount"`
	DueDate       time.Time            `json:"due_date" bson:"due_date,omitempty"`
	Paid          bool                 `json:"paid" bson:"paid"`
	PaidAt        time.Time            `json:"paid_at" bson:"paid_at,omitempty"`
	PaidBy        string               `json:"paid_by" bson:"paid_by,omitempty"`
	PaymentMethod string               `json:"payment_method" bson:"payment_method,omitempty"`
	Rates         float64              `json:"rates" bson:"rates,omitempty"`
	Status        PaymentHistoryStatus `json:"status" bson:"status,omitempty"`
}

func (p *PaymentHistory) CollectionName() string {
	return CollectionPayments.String()
}
