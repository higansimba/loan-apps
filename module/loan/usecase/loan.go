package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/higansama/loan-apps/internal/entity"
	log "github.com/higansama/loan-apps/internal/logger"
	"github.com/higansama/loan-apps/internal/utils"
	"github.com/higansama/loan-apps/module/loan/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SubmissionHandler implements LoanUsecase.
func (l *LoanUsecaseImpl) SubmissionHandler(ctx context.Context, req dto.RequestSubmission) error {
	user, err := l.repo.GetUserByID(ctx, req.UserID)
	if err != nil {
		log.Info("error find user by id ", err)
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}

	// cek user deliquent status
	isEligible, err := l.CheckUserDeliquentStatus(ctx, req.UserID)
	if err != nil {
		log.Info("error check user deliquent status", err)
		return err
	}
	if !isEligible {
		return errors.New("user is deliquent")
	}

	return l.TxManager.WithTransaction(ctx, func(sessCtx context.Context) error {
		// create loan
		loan := &entity.Loan{
			ID:                    primitive.NewObjectID(),
			Nasabah:               utils.StringToObjectID(req.UserID),
			Amount:                int(req.Amount),
			Tenor:                 req.Tenor,
			Status:                entity.LoanStatusPending,
			RemainingInstallments: 0,
			CreatedAt:             time.Now(),
			UpdatedAt:             time.Now(),
		}
		if err := l.repo.CreateLoan(sessCtx, loan); err != nil {
			log.Info("error create loan", err)
			return err
		}
		return nil
	})
}

func (l *LoanUsecaseImpl) CheckUserDeliquentStatus(ctx context.Context, userID string) (bool, error) {
	paymentHistories, err := l.repo.GetUserPaymentHistory(ctx, userID, true)
	if err != nil {
		log.Info("error find user loan history", err)
		return false, err
	}

	if len(paymentHistories) < 2 {
		return true, nil
	}

	isEligible := true
	unpaidHistory := 0
	for _, v := range paymentHistories {
		if !v.Paid && v.DueDate.Before(time.Now()) {
			unpaidHistory++
		}
	}

	if unpaidHistory > 1 {
		isEligible = false
	}

	return isEligible, nil
}
