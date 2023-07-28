package repository

import (
	"context"

	"dating-api/internal/primary/domain"
	"dating-api/pkg/errs"
)

type Repository interface {
	GetQuiz(ctx context.Context, quizId int) (*domain.GetQuizData, errs.Error)
	GetQuizUser(ctx context.Context, quizId int, userId int) (*domain.GetQuizUserData, errs.Error)
}
