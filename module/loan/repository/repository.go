package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/higansama/loan-apps/internal/entity"
	"github.com/higansama/loan-apps/internal/logger"
	"github.com/higansama/loan-apps/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LoanRepository interface {
	Seeder(ctx context.Context, user []entity.User) error
	GetUserByID(ctx context.Context, userID string) (*entity.User, error)
	GetUserPaymentHistory(ctx context.Context, idParam string, isUserID bool) ([]entity.PaymentHistory, error)
	CreateLoan(ctx context.Context, loan *entity.Loan) error
	GetLoanByID(ctx context.Context, id string) (loan *entity.Loan, err error)
	UpdateLoan(ctx context.Context, loan *entity.Loan) error
	CreatePaymentHistory(ctx context.Context, installments *entity.PaymentHistory) error
	GetOverduePayments(ctx context.Context, loanID string) ([]entity.PaymentHistory, error)
	UpdatePaymentHistory(ctx context.Context, paymentToUpdate []primitive.ObjectID, dataToUpdate entity.PaymentHistory) error
}
type LoanRepositoryImpl struct {
	db *mongo.Client
}

// UpdatePaymentHistory implements LoanRepository.
func (l *LoanRepositoryImpl) UpdatePaymentHistory(ctx context.Context, paymentToUpdate []primitive.ObjectID, dataToUpdate entity.PaymentHistory) error {
	collection := l.db.Database("loan").Collection(entity.CollectionPayments.String())
	_, err := collection.UpdateMany(ctx, bson.M{"_id": bson.M{"$in": paymentToUpdate}}, bson.M{"$set": dataToUpdate})
	if err != nil {
		return err
	}
	return nil
}

func (l *LoanRepositoryImpl) GetOverduePayments(
	ctx context.Context,
	loanID string,
) ([]entity.PaymentHistory, error) {
	// Tentukan waktu sekarang (UTC atau sesuai zona waktu aplikasi)

	// Filter: LoanID cocok, DueDate < hari ini, dan belum dibayar
	filter := bson.M{
		"loan_id":  utils.StringToObjectID(loanID),
		"due_date": bson.M{"$lt": time.Now()}, // jatuh tempo sebelum hari ini
		"$or": []bson.M{
			{"paid": false},
			{"is_paid": false},
		},
	}

	// Eksekusi query
	cursor, err := l.db.Database("loan").Collection(entity.CollectionPayments.String()).Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("gagal menjalankan query: %v", err)
	}
	defer cursor.Close(ctx)

	// Kumpulkan hasil
	var histories []entity.PaymentHistory
	if err = cursor.All(ctx, &histories); err != nil {
		return nil, fmt.Errorf("gagal decode hasil: %v", err)
	}

	return histories, nil
}

// CreatePaymentHistory implements LoanRepository.
func (l *LoanRepositoryImpl) CreatePaymentHistory(ctx context.Context, installments *entity.PaymentHistory) error {
	collection := l.db.Database("loan").Collection(entity.CollectionPayments.String())
	_, err := collection.InsertOne(ctx, installments)
	if err != nil {
		return err
	}
	return nil
}

// UpdateLoan implements LoanRepository.
func (l *LoanRepositoryImpl) UpdateLoan(ctx context.Context, loan *entity.Loan) error {
	collection := l.db.Database("loan").Collection(entity.CollectionLoans.String())
	_, err := collection.UpdateOne(ctx, bson.M{"_id": loan.ID}, bson.M{"$set": loan})
	if err != nil {
		return err
	}
	logger.Info("loan updated ", loan.ID.Hex())
	return nil
}

// GetLoanByID implements LoanRepository.
func (l *LoanRepositoryImpl) GetLoanByID(ctx context.Context, id string) (loan *entity.Loan, err error) {
	collection := l.db.Database("loan").Collection(entity.CollectionLoans.String())
	err = collection.FindOne(ctx, bson.M{"_id": utils.StringToObjectID(id)}).Decode(&loan)
	if err != nil && err != mongo.ErrNoDocuments {
		logger.Warn("something wrong find loan by id", err)
		return nil, err
	}
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("loan not found")
	}
	return loan, nil
}

// CreateLoan implements LoanRepository.
func (l *LoanRepositoryImpl) CreateLoan(ctx context.Context, loan *entity.Loan) error {
	collection := l.db.Database("loan").Collection(entity.CollectionLoans.String())
	_, err := collection.InsertOne(ctx, loan)
	if err != nil {
		return err
	}
	return nil
}

// GetUserPaymentHistory implements LoanRepository.
func (l *LoanRepositoryImpl) GetUserPaymentHistory(ctx context.Context, idParam string, isUserID bool) ([]entity.PaymentHistory, error) {
	filter := bson.M{"loan_id": utils.StringToObjectID(idParam)}
	if isUserID {
		filter = bson.M{"user_id": utils.StringToObjectID(idParam)}
	}
	collection := l.db.Database("loan").Collection(entity.CollectionPayments.String())
	r, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var paymentHistories []entity.PaymentHistory
	err = r.All(ctx, &paymentHistories)
	if err != nil {
		return nil, err
	}

	return paymentHistories, nil
}

// Test implements LoanRepository.
func (l *LoanRepositoryImpl) Seeder(ctx context.Context, users []entity.User) error {
	collection := l.db.Database("loan").Collection(entity.CollectionUsers.String())
	for _, user := range users {
		_, err := collection.InsertOne(ctx, user)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewLoanRepository(db *mongo.Client) LoanRepository {
	return &LoanRepositoryImpl{
		db: db,
	}
}
