package usecase

import (
	"context"

	"github.com/higansama/loan-apps/internal/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SeedUsecase implements LoanUsecase.
func (l *LoanUsecaseImpl) SeedUsecase(ctx context.Context) error {
	users := []entity.User{
		{
			ID:       primitive.NewObjectID(),
			Nama:     "admin",
			NoKTP:    "1234567890",
			NoHP:     "081234567890",
			Email:    "admin@example.com",
			Alamat:   "Jl. Admin",
			Password: "123456",
			Role:     entity.UserRoleAdmin,
		},
		{
			ID:       primitive.NewObjectID(),
			Nama:     "nasabah 1",
			NoKTP:    "1234567890",
			NoHP:     "081234567890",
			Email:    "nasabah1@example.com",
			Alamat:   "Jl. Nasabah",
			Password: "123456",
			Role:     entity.UserRoleNasabah,
		},
		{
			ID:       primitive.NewObjectID(),
			Nama:     "nasabah 2",
			NoKTP:    "1234567890",
			NoHP:     "081234567890",
			Email:    "nasabah2@example.com",
			Alamat:   "Jl. Nasabah",
			Password: "123456",
			Role:     entity.UserRoleNasabah,
		},
		{
			Nama:     "officer 1",
			NoKTP:    "1234567890",
			NoHP:     "081234567890",
			Email:    "officer1@example.com",
			Alamat:   "Jl. Officer",
			Password: "123456",
			Role:     entity.UserRoleOfficer,
		},
		{
			Nama:     "officer 2",
			NoKTP:    "1234567890",
			NoHP:     "081234567890",
			Email:    "officer2@example.com",
			Alamat:   "Jl. Officer",
			Password: "123456",
			Role:     entity.UserRoleOfficer,
		},
	}
	err := l.repo.Seeder(ctx, users)
	if err != nil {
	}
	return nil
}
