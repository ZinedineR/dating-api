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

func (r repo) GetProfileData(ctx context.Context, Id uuid.UUID, sex, orientation string) (*domain.ProfileData, errs.Error) {
	var (
		models *domain.ProfileData
	)
	query := r.db.WithContext(ctx).Model(&domain.Profile{})
	if sex == "male" && orientation == "straight" {
		query.Where("sex = ?", "female").
			Where("orientation = ?", "straight")
	} else if sex == "female" && orientation == "straight" {
		query.Where("sex = ?", "male").
			Where("orientation = ?", "straight")
	} else if sex == "male" && orientation == "gay" {
		query.Where("sex = ?", "male").
			Where("orientation = ?", "gay")
	} else if sex == "female" && orientation == "gay" {
		query.Where("sex = ?", "female").
			Where("orientation = ?", "gay")
	} else if sex == "non-binary" {
		query.Where("sex = ?", "non-binary")
	} else if orientation == "experiment" {
		query.Where("sex IS NOT NULL")
	}

	if err := query.
		Where("id <> ?", Id).
		Order("RANDOM ()").
		First(&models).
		Error; err != nil {
		return nil, errs.Wrap(err)
	}
	return models, nil

}

func (r repo) CheckVerified(ctx context.Context, Id uuid.UUID) (*bool, errs.Error) {
	var (
		model *domain.Verification
	)
	if err := r.db.WithContext(ctx).
		Model(&domain.Verification{}).
		Where("people_id = ?", Id).
		First(&model).
		Error; err != nil {
		return nil, errs.Wrap(err)
	}
	return &model.Verified, nil

}

func (r repo) CheckView(ctx context.Context, Id uuid.UUID) (*int, errs.Error) {
	var (
		model *domain.ProfilePreferences
	)
	if err := r.db.WithContext(ctx).
		Model(&domain.ProfilePreferences{}).
		Where("people_id = ?", Id).
		First(&model).
		Error; err != nil {
		return nil, errs.Wrap(err)
	}
	return &model.View, nil

}

func (r repo) UpdateView(ctx context.Context, Id uuid.UUID) errs.Error {
	if err := r.db.WithContext(ctx).
		Model(&domain.ProfilePreferences{}).
		Where("people_id = ?", Id).
		Update("view", gorm.Expr("view + 1")).Error; err != nil {
		return errs.Wrap(err)
	}
	return nil

}

func (r repo) ResetView(ctx context.Context, Id uuid.UUID) errs.Error {
	if err := r.db.WithContext(ctx).
		Model(&domain.ProfilePreferences{}).
		Where("people_id = ?", Id).
		Update("view", 0).Error; err != nil {
		return errs.Wrap(err)
	}
	return nil

}

func (r repo) GetProfileFullData(ctx context.Context, email string) (*domain.Profile, errs.Error) {
	var (
		models *domain.Profile
	)
	if err := r.db.WithContext(ctx).
		Model(&domain.Profile{}).
		Where("email = ?", email).
		First(&models).
		Error; err != nil {
		return nil, errs.Wrap(err)
	}
	return models, nil

}

func (r repo) StoreJWT(ctx context.Context, jwt string, Id uuid.UUID) errs.Error {
	if err := r.db.WithContext(ctx).
		Model(&domain.ProfilePreferences{}).
		Where("people_id = ?", Id).
		Update("jwt", jwt).Error; err != nil {
		return errs.Wrap(err)
	}
	return nil

}

func (r repo) CheckJWT(ctx context.Context, Id uuid.UUID) (*domain.ProfilePreferences, errs.Error) {
	var (
		models *domain.ProfilePreferences
	)
	if err := r.db.WithContext(ctx).
		Model(&domain.ProfilePreferences{}).
		Where("people_id = ?", Id).
		First(&models).
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
	if err := r.db.WithContext(ctx).
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
