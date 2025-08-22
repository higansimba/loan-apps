package repository

import (
	"context"

	"github.com/higansama/loan-apps/internal/entity"
	log "github.com/higansama/loan-apps/internal/logger"
	"github.com/higansama/loan-apps/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CheckUserid implements LoanRepository.
func (l *LoanRepositoryImpl) GetUserByID(ctx context.Context, userID string) (*entity.User, error) {
	collection := l.db.Database("loan").Collection(entity.CollectionUsers.String())
	var user entity.User
	err := collection.FindOne(ctx, bson.M{"_id": utils.StringToObjectID(userID)}).Decode(&user)
	if err != nil && err != mongo.ErrNoDocuments {
		log.Info("error find user by id ", err)
		return nil, err
	}
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &user, nil
}
