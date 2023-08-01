package repository

import (
	"context"
	"fmt"
	"strings"

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

func (r repo) GetProfileData(ctx context.Context, Id uuid.UUID, sex, orientation, list string) (*domain.ProfileData, errs.Error) {
	var (
		models *domain.ProfileData
	)

	query := r.db.Debug().WithContext(ctx).Model(&domain.Profile{})
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
	listArray := strings.Split(list, ",")
	if list == "," {
		query.Where("id <> ?", Id)
	} else {
		query.Where("id <> ?", Id).
			Where("id NOT IN ?", listArray)
	}
	if err := query.
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

func (r repo) CheckLikePass(ctx context.Context, Id uuid.UUID) (*string, errs.Error) {
	var (
		like   *string
		pass   *string
		result string
	)
	if err := r.db.WithContext(ctx).
		Model(&domain.ProfilePreferences{}).
		// Select(`REPLACE(concat("like", pass), '{', '') AS list`).
		Select(`array_to_string("like", ',') AS like`).
		Where("people_id = ?", Id).
		Scan(&like).
		Error; err != nil {
		return nil, errs.Wrap(err)
	}

	if err := r.db.WithContext(ctx).
		Model(&domain.ProfilePreferences{}).
		// Select(`REPLACE(concat("like", pass), '{', '') AS list`).
		Select(`array_to_string("pass", ',') AS pass`).
		Where("people_id = ?", Id).
		Scan(&pass).
		Error; err != nil {
		return nil, errs.Wrap(err)
	}
	// If "like" is nil or empty, set it to an empty string
	if like == nil || *like == "" {
		emptyLike := ""
		like = &emptyLike
		result = *pass
		return &result, nil
	}

	// If "pass" is nil or empty, set it to an empty string
	if pass == nil || *pass == "" {
		emptyPass := ""
		pass = &emptyPass
		result = *like
		return &result, nil
	}

	result = *like + "," + *pass
	return &result, nil

}

func (r repo) UpdateLikePass(ctx context.Context, Id, target uuid.UUID, parameter string) errs.Error {
	query := r.db.Debug().WithContext(ctx).
		Model(&domain.ProfilePreferences{}).
		Where("people_id = ?", Id)

	if parameter == "pass" {
		query.Update(parameter, gorm.Expr(parameter+" || ?", fmt.Sprintf("{%s}", target)))
	} else if parameter == "like" {
		query.Update("like", gorm.Expr(`"like" || ?`, fmt.Sprintf("{%s}", target)))
	}
	if err := query.
		Error; err != nil {
		return errs.Wrap(err)
	}
	return nil

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
