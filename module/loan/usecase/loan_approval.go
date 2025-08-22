package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/higansama/loan-apps/internal/entity"
	"github.com/higansama/loan-apps/internal/logger"
	"github.com/higansama/loan-apps/module/loan/dto"
	"github.com/higansama/loan-apps/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ApprovalSubmissionHandler implements LoanUsecase.
func (l *LoanUsecaseImpl) ApprovalSubmissionHandler(ctx context.Context, req dto.RequestApprovalSubmission) error {
	// cek user should be  officer
	user, err := l.repo.GetUserByID(ctx, req.OfficerID)
	if err != nil {
		return err
	}

	if user.Role == entity.UserRoleNasabah {
		return errors.New("user should be officer")
	}

	// cek loan should be exist
	loan, err := l.repo.GetLoanByID(ctx, req.LoanID)
	if err != nil || loan == nil {
		return err
	}

	if !loan.IsPending() {
		logger.Info("loan not pending ", loan.ID.Hex())
		return errors.New("loan not pending")
	}
	return l.TxManager.WithTransaction(ctx, func(sessCtx context.Context) error {
		// update loan status
		loan.Status = entity.ConvertStringToLoanStatus(req.ApprovalStatus)
		loan.RemainingInstallments = loan.Amount
		err := l.repo.UpdateLoan(sessCtx, loan)
		if err != nil {
			return err
		}

		if loan.Status == entity.LoanStatusApproved {
			// create installments
			err = l.createInstallments(sessCtx, loan)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (l *LoanUsecaseImpl) createInstallments(ctx context.Context, loan *entity.Loan) error {
	dueDates := time.Now().AddDate(0, 0, 7)
	weeklyInstallment, rates := utils.GenerateInstallmentRate(float64(loan.Amount), float64(loan.Tenor))
	for i := 0; i < loan.Tenor; i++ {

		installments := &entity.PaymentHistory{
			ID:      primitive.NewObjectID(),
			LoanID:  loan.ID,
			Amount:  weeklyInstallment,
			DueDate: dueDates,
			Paid:    false,
			Status:  entity.PaymentHistoryStatusActive,
			Rates:   rates / 50,
		}
		err := l.repo.CreatePaymentHistory(ctx, installments)
		if err != nil {
			return err
		}
		dueDates = dueDates.AddDate(0, 0, 7)
	}

	return nil
}
