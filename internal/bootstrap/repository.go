package bootstrap

import (
	"github.com/higansama/loan-apps/config"
	loanrepo "github.com/higansama/loan-apps/module/loan/repository"
)

type Repositories struct {
	LoanRepo loanrepo.LoanRepository
}

func InitRepositories(
	mongoDB *config.MongoConnections,
) *Repositories {
	return &Repositories{
		LoanRepo: loanrepo.NewLoanRepository(mongoDB.MainDB.Client()),
	}
}
