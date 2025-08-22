package bootstrap

import (
	"github.com/gin-gonic/gin"
	loandelivery "github.com/higansama/loan-apps/module/loan/delivery"
)

type Handlers struct {
	loan *loandelivery.LoanHandlers
}

func InitHandlers(e *gin.Engine, useCases *UseCases) *Handlers {
	return &Handlers{
		loan: loandelivery.LoadLoanHandler(e, useCases.Loan),
	}
}
