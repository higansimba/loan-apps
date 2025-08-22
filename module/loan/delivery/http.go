package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/higansama/loan-apps/internal/logger"
	"github.com/higansama/loan-apps/internal/response"
	"github.com/higansama/loan-apps/module/loan/usecase"
)

type LoanHandlers struct {
	usecase usecase.LoanUsecase
}

func LoadLoanHandler(e *gin.Engine, useCase usecase.LoanUsecase) *LoanHandlers {
	handler := &LoanHandlers{
		usecase: useCase,
	}
	e.GET("/seed", handler.SeedHandler)
	e.POST("/token", handler.CreateTokenSessionHandler)
	group := e.Group("/loan")
	group.POST("/installment", handler.SubmissionHandler)
	group.POST("/payment/installment/:loanid", handler.PaymentInstallmentHandler)
	group.PUT("/approval/installment/:loanid", handler.ApprovalSubmissionHandler)
	group.GET("/deliquent/installment/cek/:loanid", handler.DeliquentInstallmentHandler)
	return handler
}

func (handler *LoanHandlers) TestHandler(c *gin.Context) {
	handler.usecase.TestUsecase()
}

func (handler *LoanHandlers) SeedHandler(c *gin.Context) {
	err := handler.usecase.SeedUsecase(c.Request.Context())
	if err != nil {
		logger.Info("error seeding", err.Error())
		c.JSON(http.StatusInternalServerError, response.JsonResponse(http.StatusInternalServerError, "error seeding", nil, err))
		return
	}
	c.JSON(http.StatusOK, response.JsonResponse(http.StatusOK, "success seeding", nil, nil))
}
