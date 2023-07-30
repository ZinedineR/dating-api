package repository

import (
	"context"

	"dating-api/internal/profile/domain"
	baseModel "dating-api/pkg/db"
	"dating-api/pkg/errs"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repo struct {
	db   *gorm.DB
	base *baseModel.PostgreSQLClientRepository
}

func (r repo) GetProfileData(ctx context.Context, Id uuid.UUID) (*domain.Profile, errs.Error) {
	var (
		models *domain.Profile
	)
	if err := r.db.WithContext(ctx).
		Model(&domain.Profile{}).
		Take(&models, Id).
		Error; err != nil {
		return nil, errs.Wrap(err)
	}
	return models, nil

}

func (r repo) CreateProfilePreferences(ctx context.Context, model *domain.ProfilePreferences) errs.Error {
	if err := r.db.WithContext(ctx).
		Create(&model).
		Error; err != nil {
		return errs.Wrap(err)
	}
	return nil

}

func (r repo) CreateProfile(ctx context.Context, model *domain.Profile) errs.Error {
	if err := r.db.WithContext(ctx).
		Create(&model).
		Error; err != nil {
		return errs.Wrap(err)
	}
	return nil

}

func (r repo) CreateVerification(ctx context.Context, model *domain.Verification) errs.Error {
	if err := r.db.WithContext(ctx).
		Create(&model).
		Error; err != nil {
		return errs.Wrap(err)
	}
	return nil

}

func (r repo) UpdateVerification(ctx context.Context, Id uuid.UUID) errs.Error {
	if err := r.db.Debug().WithContext(ctx).
		Model(&domain.Verification{}).
		Where("people_id = ?", Id).
		Update("verified", true).Error; err != nil {
		return errs.Wrap(err)
	}
	return nil

}

func NewRepository(db *gorm.DB, base *baseModel.PostgreSQLClientRepository) Repository {
	return &repo{db: db, base: base}
}
