package usecase

import (
	"context"
	"errors"
	"fmt"

	transactionmongo "github.com/higansama/loan-apps/internal/transaction_mongo"
	"github.com/higansama/loan-apps/module/loan/dto"
	"github.com/higansama/loan-apps/module/loan/repository"
)

type LoanUsecase interface {
	TestUsecase()
	SubmissionHandler(ctx context.Context, req dto.RequestSubmission) error
	SeedUsecase(ctx context.Context) error
	CheckUserid(ctx context.Context, userID string) error
	ApprovalSubmissionHandler(ctx context.Context, req dto.RequestApprovalSubmission) error
	PaymentInstallment(ctx context.Context, req dto.RequestPaymentInstallment) error
	CheckDeliquentInstallment(ctx context.Context, loanid string) (int, error)
}

type LoanUsecaseImpl struct {
	repo      repository.LoanRepository
	TxManager transactionmongo.TransactionManager
}

// CheckUserid implements LoanUsecase.
func (l *LoanUsecaseImpl) CheckUserid(ctx context.Context, userID string) error {
	fmt.Println("masuk sini gak ")

	user, err := l.repo.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}
	fmt.Println("masuk sini gak ", user)

	if user == nil {
		fmt.Println("user not found")
		return errors.New("user not found")
	}
	return nil
}

// TestUsecase implements LoanUsecase.
func (l *LoanUsecaseImpl) TestUsecase() {
	panic("unimplemented")
}

func NewLoanUsecase(repo repository.LoanRepository, txManager transactionmongo.TransactionManager) LoanUsecase {
	return &LoanUsecaseImpl{
		repo:      repo,
		TxManager: txManager,
	}
}
