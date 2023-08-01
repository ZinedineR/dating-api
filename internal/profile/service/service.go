package service

import (
	"context"
	"time"

	"dating-api/internal/profile/domain"
	"dating-api/internal/profile/repository"
	"dating-api/pkg/errs"

	"github.com/google/uuid"
)

// NewService creates new user service
func NewService(repo repository.Repository) Service {
	return &service{authRepo: repo}
}

type service struct {
	authRepo repository.Repository
}

func (s service) GetProfileData(ctx context.Context, Id uuid.UUID, sex, orientation, list string) (*domain.ProfileData, errs.Error) {
	result, err := s.authRepo.GetProfileData(ctx, Id, sex, orientation, list)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s service) Login(ctx context.Context, email string) (*domain.Profile, errs.Error) {
	result, err := s.authRepo.GetProfileFullData(ctx, email)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s service) StoreJWT(ctx context.Context, jwt string, Id uuid.UUID) errs.Error {
	if err := s.authRepo.StoreJWT(ctx, jwt, Id); err != nil {
		return err
	}
	return nil
}

func (s service) CheckJWT(ctx context.Context, Id uuid.UUID) (*domain.ProfilePreferences, errs.Error) {
	result, err := s.authRepo.CheckJWT(ctx, Id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s service) CheckVerified(ctx context.Context, Id uuid.UUID) (*bool, errs.Error) {
	result, err := s.authRepo.CheckVerified(ctx, Id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s service) CheckLikePass(ctx context.Context, Id uuid.UUID) (*string, errs.Error) {
	result, err := s.authRepo.CheckLikePass(ctx, Id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s service) UpdateLikePass(ctx context.Context, Id, target uuid.UUID, parameter string) errs.Error {
	if err := s.authRepo.UpdateLikePass(ctx, Id, target, parameter); err != nil {
		return errs.Wrap(err)
	}
	return nil
}

func (s service) CheckView(ctx context.Context, Id uuid.UUID) (*int, errs.Error) {
	result, err := s.authRepo.CheckView(ctx, Id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s service) UpdateView(ctx context.Context, Id uuid.UUID) errs.Error {
	if err := s.authRepo.UpdateView(ctx, Id); err != nil {
		return errs.Wrap(err)
	}
	return nil
}

func (s service) ResetView(ctx context.Context, Id uuid.UUID) errs.Error {
	if err := s.authRepo.ResetView(ctx, Id); err != nil {
		return errs.Wrap(err)
	}
	return nil
}

func (s service) CreateProfilePreferences(ctx context.Context, model *domain.ProfilePreferences) errs.Error {
	if model.Id == uuid.Nil {
		model.Id = uuid.New()
	}
	if err := s.authRepo.CreateProfilePreferences(ctx, model); err != nil {
		return errs.Wrap(err)
	}
	return nil
}

func (s service) CreateProfile(ctx context.Context, model *domain.Profile) errs.Error {
	if model.Id == uuid.Nil {
		model.Id = uuid.New()
		model.Account = "free"
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
