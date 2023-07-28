package service

import (
	"context"

	"dating-api/pkg/httpclient"

	"dating-api/internal/profile/domain"
	"dating-api/internal/profile/repository"
	"dating-api/pkg/errs"
)

// NewService creates new user service
func NewService(repo repository.Repository, httpClient httpclient.Client) Service {
	return &service{authRepo: repo, httpClient: httpClient}
}

type service struct {
	authRepo   repository.Repository
	httpClient httpclient.Client
}

func (s service) GetProfileData(ctx context.Context, Id int) (*domain.Profile, errs.Error) {
	result, err := s.authRepo.GetProfileData(ctx, Id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
