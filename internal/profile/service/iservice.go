package service

import (
	"context"

	"dating-api/internal/profile/domain"
	"dating-api/pkg/errs"

	"github.com/google/uuid"
)

type Service interface {
	CreateProfilePreferences(ctx context.Context, model *domain.ProfilePreferences) errs.Error
	Login(ctx context.Context, email, password string) (*domain.Profile, errs.Error)
	GetProfileData(ctx context.Context, Id uuid.UUID) (*domain.ProfileData, errs.Error)
	CheckVerified(ctx context.Context, Id uuid.UUID) (*bool, errs.Error)
	CreateProfile(ctx context.Context, model *domain.Profile) errs.Error
	CreateVerification(ctx context.Context, model *domain.Verification) errs.Error
	UpdateVerification(ctx context.Context, Id uuid.UUID) errs.Error
}
