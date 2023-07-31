package service

import (
	"context"
	"time"

	"dating-api/pkg/httpclient"

	"dating-api/internal/profile/domain"
	"dating-api/internal/profile/repository"
	"dating-api/pkg/errs"

	"github.com/google/uuid"
)

// NewService creates new user service
func NewService(repo repository.Repository, httpClient httpclient.Client) Service {
	return &service{authRepo: repo, httpClient: httpClient}
}

type service struct {
	authRepo   repository.Repository
	httpClient httpclient.Client
}

func (s service) GetProfileData(ctx context.Context, Id uuid.UUID) (*domain.ProfileData, errs.Error) {
	result, err := s.authRepo.GetProfileData(ctx, Id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s service) Login(ctx context.Context, email, password string) (*domain.Profile, errs.Error) {
	result, err := s.authRepo.GetProfileFullData(ctx, email, password)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s service) CreateProfilePreferences(ctx context.Context, model *domain.ProfilePreferences) errs.Error {
	if model.Id == uuid.Nil {
		model.Id = uuid.New()
		model.LastLogin = time.Now()
	}
	if err := s.authRepo.CreateProfilePreferences(ctx, model); err != nil {
		return errs.Wrap(err)
	}
	return nil
}

func (s service) CreateProfile(ctx context.Context, model *domain.Profile) errs.Error {
	if model.Id == uuid.Nil {
		model.Id = uuid.New()
		model.Account = "Free"
		model.CreatedAt = time.Now()
	}

	if err := s.authRepo.CreateProfile(ctx, model); err != nil {
		return errs.Wrap(err)
	}
	return nil
}

func (s service) CreateVerification(ctx context.Context, model *domain.Verification) errs.Error {
	if model.Id == uuid.Nil {
		model.Id = uuid.New()
		model.Expiresat = time.Now().Add(time.Hour * 24)
	}

	if err := s.authRepo.CreateVerification(ctx, model); err != nil {
		return errs.Wrap(err)
	}
	return nil
}

func (s service) UpdateVerification(ctx context.Context, Id uuid.UUID) errs.Error {
	if err := s.authRepo.UpdateVerification(ctx, Id); err != nil {
		return errs.Wrap(err)
	}
	return nil
}
