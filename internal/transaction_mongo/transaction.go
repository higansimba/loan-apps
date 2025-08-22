package transactionmongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type TransactionManager struct {
	Client *mongo.Client
}

func NewTransactionManager(client *mongo.Client) *TransactionManager {
	return &TransactionManager{
		Client: client,
	}
}

func (tm *TransactionManager) WithTransaction(ctx context.Context, fn func(sessCtx context.Context) error) error {
	if err := tm.Client.Ping(ctx, readpref.Primary()); err != nil {
		return fmt.Errorf("gagal terhubung ke MongoDB: %v", err)
	}

	session, err := tm.Client.StartSession()
	if err != nil {
		return fmt.Errorf("gagal membuat session: %v", err)
	}
	defer session.EndSession(ctx)

	_, err = session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		return nil, fn(sessCtx)
	})

	return err
}
