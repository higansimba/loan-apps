package usecase

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/higansama/loan-apps/internal/entity"
	"github.com/higansama/loan-apps/internal/logger"
	log "github.com/higansama/loan-apps/internal/logger"
	"github.com/higansama/loan-apps/module/loan/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *LoanUsecaseImpl) PaymentInstallment(ctx context.Context, req dto.RequestPaymentInstallment) error {
	loan, err := u.repo.GetLoanByID(ctx, req.LoanID)
	if err != nil || loan == nil {
		return err
	}

	// cek apakah pinjaman menunggak
	paymentHistories, err := u.repo.GetOverduePayments(ctx, req.LoanID)
	if err != nil {
		log.Info("error find GetOverduePayments", err)
		return err
	}

	if len(paymentHistories) >= 1 && req.Amount != int(paymentHistories[0].Amount) {
		return errors.New("payment amount is not equal with loan amount")
	}

	paymentToUpdate := make([]primitive.ObjectID, 0)
	minimumPayment := 0
	for _, v := range paymentHistories {
		minimumPayment += int(v.Amount)
		paymentToUpdate = append(paymentToUpdate, v.ID)
	}

	if len(paymentHistories) > 2 || req.Amount != minimumPayment {
		return errors.New("loan deliquent , minimum payment is " + strconv.Itoa(minimumPayment))
	}

	return u.TxManager.WithTransaction(ctx, func(sessCtx context.Context) error {
		// update loan status
		loan.RemainingInstallments = loan.RemainingInstallments - minimumPayment
		err := u.repo.UpdateLoan(sessCtx, loan)
		if err != nil {
			return err
		}

		// update payment history
		dataToUpdate := entity.PaymentHistory{
			LoanID:        loan.ID,
			Amount:        float64(req.Amount),
			Paid:          true,
			PaidAt:        time.Now(),
			PaidBy:        req.PaidBy,
			PaymentMethod: req.PaymentMethod,
			Status:        entity.PaymentHistoryStatusPaid,
		}
		logger.Info("payment history updated ", paymentToUpdate)
		err = u.repo.UpdatePaymentHistory(sessCtx, paymentToUpdate, dataToUpdate)
		if err != nil {
			return err
		}

		return nil
	})

}
