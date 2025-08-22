package bootstrap

import (
	transactionmongo "github.com/higansama/loan-apps/internal/transaction_mongo"
	loanusecase "github.com/higansama/loan-apps/module/loan/usecase"
)

type UseCases struct {
	Loan loanusecase.LoanUsecase
}

func InitUseCases(repos *Repositories, txManager transactionmongo.TransactionManager) *UseCases {
	return &UseCases{
		Loan: loanusecase.NewLoanUsecase(repos.LoanRepo, txManager),
	}
}
