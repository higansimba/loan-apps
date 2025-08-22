package session

import (
	"encoding/gob"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type ValidationPayment struct {
	Token     string
	IsDone    bool
	ExpiredAt time.Time
}

type ActionType string

const Payment ActionType = "payment"
const Submission ActionType = "submission"
const Loan ActionType = "loan"

func (a ActionType) String() string {
	return string(a)
}

func CreateActionSession(ctx *gin.Context, action ActionType, userID string) string {
	key := action.String() + "-session-" + userID
	data := map[string]ValidationPayment{
		key: {
			Token:     strconv.Itoa(int(time.Now().UnixMicro())),
			IsDone:    false,
			ExpiredAt: time.Now().Add(2 * time.Minute),
		},
	}
	session := sessions.Default(ctx)
	session.Set(string(userID), data)
	session.Save()
	return data[key].Token
}

func CekValidActionIsValid(ctx *gin.Context, action ActionType, userID string, token string) bool {
	key := action.String() + "-session-" + userID
	session := sessions.Default(ctx)
	data := session.Get(string(userID))
	if data == nil {
		return false
	}

	paymentData, ok := data.(map[string]ValidationPayment)
	if !ok {
		return false
	}

	if token != paymentData[key].Token {
		return false
	}

	payment, ok := paymentData[key]
	if !ok {
		return false
	}

	if paymentData[key].ExpiredAt.Before(time.Now()) {
		return false
	}

	return payment.IsDone
}

func DestroySession(ctx *gin.Context, action ActionType, userID string) {
	key := action.String() + "-session-" + userID
	session := sessions.Default(ctx)
	data := session.Get(string(userID))
	if data == nil {
		return
	}

	session.Delete(key)
	session.Save()
}

func InitSession(ginApp *gin.Engine) *gin.Engine {
	// 1. Setup session store (menggunakan cookie)
	store := cookie.NewStore([]byte("aofjaoejfojdfks93434nmerp04-34")) // Harus minimal 32 byte
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   86400, // 24 jam
		HttpOnly: true,
	})

	gob.Register(map[string]ValidationPayment{})
	ginApp.Use(sessions.Sessions("mysession", store))
	return ginApp
}
