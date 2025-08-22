package dto

import (
	"errors"

	"github.com/higansama/loan-apps/internal/utils"
)

type RequestToken struct {
	UserID string `json:"user_id" binding:"required"`
	Action string `json:"action" binding:"required"`
}

type RequestSubmission struct {
	AccessToken string  `json:"access_token" binding:"required"`
	UserID      string  `json:"user_id" binding:"required"`
	Amount      float64 `json:"amount" binding:"required"`
	Tenor       int     `json:"tenor" binding:"required"`
}

func (r *RequestSubmission) Validate() error {
	if !utils.IsValidObjectID(r.UserID) {
		return errors.New("user id is not valid")
	}

	if r.AccessToken == "" {
		return errors.New("access token is required")
	}
	if r.UserID == "" {
		return errors.New("user id is required")
	}
	if r.Amount <= 0 {
		return errors.New("amount is required")
	}
	if r.Tenor <= 0 {
		return errors.New("tenor is required")
	}
	return nil
}

type RequestApprovalSubmission struct {
	AccessToken    string `json:"access_token" binding:"required"`
	OfficerID      string `json:"officer_id" binding:"required"`
	LoanID         string `json:"loan_id"`
	ApprovalStatus string `json:"approval_status" binding:"required"`
}

func (r *RequestApprovalSubmission) Validate() error {
	if !utils.IsValidObjectID(r.OfficerID) {
		return errors.New("user id is not valid")
	}

	if r.AccessToken == "" {
		return errors.New("access token is required")
	}
	return nil
}

type RequestPaymentInstallment struct {
	Amount        int    `json:"amount" binding:"required"`
	PaymentMethod string `json:"payment_method" binding:"required"`
	PaidBy        string `json:"paid_by" binding:"required"`
	LoanID        string `json:"loan_id"`
}

func (r *RequestPaymentInstallment) Validate() error {
	if !utils.IsValidObjectID(r.LoanID) {
		return errors.New("loan id is not valid")
	}
	if r.Amount <= 0 {
		return errors.New("amount is required")
	}
	if r.PaymentMethod == "" {
		return errors.New("payment method is required")
	}
	if r.PaidBy == "" {
		return errors.New("paid by is required")
	}
	return nil
}
