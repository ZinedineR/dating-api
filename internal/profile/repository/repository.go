package repository

import (
	"context"

	"dating-api/internal/profile/domain"
	baseModel "dating-api/pkg/db"
	"dating-api/pkg/errs"

	"gorm.io/gorm"
)

type repo struct {
	db   *gorm.DB
	base *baseModel.PostgreSQLClientRepository
}

func (r repo) GetProfileData(ctx context.Context, quizId int) (*domain.Profile, errs.Error) {
	var (
		models *domain.Profile
	)
	return models, nil

}

func NewRepository(db *gorm.DB, base *baseModel.PostgreSQLClientRepository) Repository {
	return &repo{db: db, base: base}
}
