package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoanType string

const (
	LoanTypePersonal LoanType = "personal"
	LoanTypeBusiness LoanType = "business"
)

type LoanStatus string

const (
	LoanStatusPending  LoanStatus = "pending"
	LoanStatusApproved LoanStatus = "approved"
	LoanStatusRejected LoanStatus = "rejected"
)

type Loan struct {
	ID                    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Nasabah               primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	Amount                int                `json:"amount" bson:"amount,omitempty"`
	Tenor                 int                `json:"tenor" bson:"tenor,omitempty"`
	Status                LoanStatus         `json:"status" bson:"status,omitempty"`
	RemainingInstallments int                `json:"remaining_installments" bson:"remaining_installments,omitempty"`
	CreatedAt             time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt             time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}

func ConvertStringToLoanStatus(status string) LoanStatus {
	switch status {
	case "pending":
		return LoanStatusPending
	case "approved":
		return LoanStatusApproved
	case "rejected":
		return LoanStatusRejected
	default:
		return LoanStatusPending
	}
}

func (l *Loan) CollectionName() string {
	return CollectionLoans.String()
}

func (l *Loan) IsPending() bool {
	return l.Status == LoanStatusPending
}

func (l *Loan) IsApproved() bool {
	return l.Status == LoanStatusApproved
}

func (l *Loan) IsRejected() bool {
	return l.Status == LoanStatusRejected
}
