package service

import (
	"context"

	"dating-api/internal/profile/domain"
	"dating-api/pkg/errs"
)

type Service interface {
	GetProfileData(ctx context.Context, quizId int) (*domain.Profile, errs.Error)
}
