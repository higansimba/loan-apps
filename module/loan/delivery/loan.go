package delivery

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/higansama/loan-apps/internal/logger"
	"github.com/higansama/loan-apps/internal/response"
	"github.com/higansama/loan-apps/internal/utils"
	"github.com/higansama/loan-apps/module/loan/dto"
)

func (handler *LoanHandlers) SubmissionHandler(c *gin.Context) {
	var req dto.RequestSubmission
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Info("error binding json", err.Error())
		c.JSON(http.StatusBadRequest, response.JsonResponse(http.StatusBadRequest, "error binding json", nil, err))
		return
	}

	if err := req.Validate(); err != nil {
		logger.Info("error validate request", err.Error())
		c.JSON(http.StatusBadRequest, response.JsonResponse(http.StatusBadRequest, "error validate request", nil, err))
		return
	}

	err := handler.usecase.SubmissionHandler(c.Request.Context(), req)
	if err != nil {
		logger.Info("error create loan", err.Error())
		c.JSON(http.StatusInternalServerError, response.JsonResponse(http.StatusInternalServerError, "error create loan", nil, err))
		return
	}

	c.JSON(http.StatusOK, response.JsonResponse(http.StatusOK, "success create loan", nil, nil))
}

func (handler *LoanHandlers) ApprovalSubmissionHandler(c *gin.Context) {
	loanId := c.Param("loanid")
	var req dto.RequestApprovalSubmission
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Info("error binding json", err.Error())
		c.JSON(http.StatusBadRequest, response.JsonResponse(http.StatusBadRequest, "error binding json", nil, err))
		return
	}

	req.LoanID = loanId

	if err := req.Validate(); err != nil {
		logger.Info("error validate request", err.Error())
		c.JSON(http.StatusBadRequest, response.JsonResponse(http.StatusBadRequest, "error validate request", nil, err))
		return
	}

	err := handler.usecase.ApprovalSubmissionHandler(c.Request.Context(), req)
	if err != nil {
		logger.Info("error approval loan", err.Error())
		c.JSON(http.StatusInternalServerError, response.JsonResponse(http.StatusInternalServerError, "error approval loan", nil, err))
		return
	}

	c.JSON(http.StatusOK, response.JsonResponse(http.StatusOK, "success approval loan", nil, nil))
}

func (handler *LoanHandlers) PaymentInstallmentHandler(c *gin.Context) {
	loanId := c.Param("loanid")
	var req dto.RequestPaymentInstallment
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Info("error binding json", err.Error())
		c.JSON(http.StatusBadRequest, response.JsonResponse(http.StatusBadRequest, "error binding json", nil, err))
		return
	}

	req.LoanID = loanId

	if err := req.Validate(); err != nil {
		logger.Info("error validate request", err.Error())
		c.JSON(http.StatusBadRequest, response.JsonResponse(http.StatusBadRequest, "error validate request", nil, err))
		return
	}

	err := handler.usecase.PaymentInstallment(c.Request.Context(), req)
	if err != nil {
		logger.Info("error payment installment", err.Error())
		c.JSON(http.StatusInternalServerError, response.JsonResponse(http.StatusInternalServerError, "error payment installment", nil, err))
		return
	}

	c.JSON(http.StatusOK, response.JsonResponse(http.StatusOK, "success payment installment", nil, nil))
}

func (handler *LoanHandlers) DeliquentInstallmentHandler(c *gin.Context) {
	loanID := c.Param("loanid")

	if !utils.IsValidObjectID(loanID) {
		logger.Info("error validate loan id", "loan id is not valid")
		c.JSON(http.StatusBadRequest, response.JsonResponse(http.StatusBadRequest, "error validate loan id", nil, errors.New("loan id is not valid")))
		return
	}

	r, err := handler.usecase.CheckDeliquentInstallment(c.Request.Context(), loanID)
	if err != nil {
		logger.Info("error check deliquent installment", err.Error())
		c.JSON(http.StatusInternalServerError, response.JsonResponse(http.StatusBadRequest, "error validate loan id", nil, errors.New("loan id is not valid")))
		return
	}

	c.JSON(http.StatusOK, response.JsonResponse(http.StatusOK, "success check deliquent installment", gin.H{"deliquent_total": r}, nil))

}
