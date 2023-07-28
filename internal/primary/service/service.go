package service

import (
	"context"

	"dating-api/pkg/httpclient"

	"dating-api/internal/primary/domain"
	"dating-api/internal/primary/repository"
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

func (s service) GetQuiz(ctx context.Context, quizId int) (*domain.GetQuizData, errs.Error) {
	result, err := s.authRepo.GetQuiz(ctx, quizId)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s service) GetQuizUser(ctx context.Context, quizId int, userId int) (*domain.GetQuizUserData, errs.Error) {
	result, err := s.authRepo.GetQuizUser(ctx, quizId, userId)
	if err != nil {
		return nil, err
	}
	return result, nil
}
