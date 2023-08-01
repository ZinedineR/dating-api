package service

import (
	"context"

	"dating-api/internal/profile/domain"
	"dating-api/pkg/errs"

	"github.com/google/uuid"
)

type Service interface {
	CreateProfilePreferences(ctx context.Context, model *domain.ProfilePreferences) errs.Error
	Login(ctx context.Context, email string) (*domain.Profile, errs.Error)
	StoreJWT(ctx context.Context, jwt string, Id uuid.UUID) errs.Error
	CheckJWT(ctx context.Context, Id uuid.UUID) (*domain.ProfilePreferences, errs.Error)
	GetProfileData(ctx context.Context, Id uuid.UUID, sex, orientation, list string) (*domain.ProfileData, errs.Error)
	CheckVerified(ctx context.Context, Id uuid.UUID) (*bool, errs.Error)
	CreateProfile(ctx context.Context, model *domain.Profile) errs.Error
	CheckLikePass(ctx context.Context, Id uuid.UUID) (*string, errs.Error)
	UpdateLikePass(ctx context.Context, Id, target uuid.UUID, parameter string) errs.Error
	CheckView(ctx context.Context, Id uuid.UUID) (*int, errs.Error)
	UpdateView(ctx context.Context, Id uuid.UUID) errs.Error
	ResetView(ctx context.Context, Id uuid.UUID) errs.Error
	CreateVerification(ctx context.Context, model *domain.Verification) errs.Error
	UpdateVerification(ctx context.Context, Id uuid.UUID) errs.Error
}
