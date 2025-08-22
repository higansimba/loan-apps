package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/higansama/loan-apps/internal/logger"
	"github.com/higansama/loan-apps/internal/response"
	internalSession "github.com/higansama/loan-apps/internal/session"
	"github.com/higansama/loan-apps/module/loan/dto"
)

func (handler *LoanHandlers) CreateTokenSessionHandler(c *gin.Context) {
	var req dto.RequestToken
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Info("error binding json", err.Error())
		c.JSON(http.StatusBadRequest, response.JsonResponse(http.StatusBadRequest, "error binding json", nil, err))
		return
	}

	if req.Action == "" {
		logger.Info("error action is empty")
		c.JSON(http.StatusBadRequest, response.JsonResponse(http.StatusBadRequest, "error action is empty", nil, nil))
		return
	}

	err := handler.usecase.CheckUserid(c.Request.Context(), req.UserID)
	if err != nil {
		logger.Info("error check user id", err.Error())
		c.JSON(http.StatusInternalServerError, response.JsonResponse(http.StatusInternalServerError, "error check user id", nil, err))
		return
	}

	token := internalSession.CreateActionSession(c, internalSession.ActionType(req.Action), req.UserID)
	c.JSON(http.StatusOK, response.JsonResponse(http.StatusOK, "success create token session", gin.H{
		"token": token,
	}, nil))
}
