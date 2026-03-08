package problems

import (
	"context"
	"strings"

	"github.com/Aneeshie/cpp-judge/internal/models"
)

type Service struct{
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateProblem(ctx context.Context, req models.CreateProblemRequest) (*models.Problem, error) {
	slug := slugify(req.Title)

	problem := &models.Problem{
		Slug: slug,
		Title: req.Title,
		Description: req.Description,
		Difficulty: req.Difficulty,
		TimeLimitMs: 1000,
		MemoryLimitMb: 256,
	}

	problem, err := s.repo.CreateProblem(ctx,problem)
	if err != nil {
		return nil, err
	}

	return problem,err
}

func slugify(title string) string {
	slug := strings.ToLower(title)
	slug = strings.ReplaceAll(slug," ", "-")
	return slug
}
