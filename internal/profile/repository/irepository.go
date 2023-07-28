package repository

import (
	"context"

	"dating-api/internal/profile/domain"
	"dating-api/pkg/errs"
)

type Repository interface {
	GetProfileData(ctx context.Context, Id int) (*domain.Profile, errs.Error)
}
