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

	return s.repo.CreateProblem(ctx,problem)
}

func slugify(title string) string {
	slug := strings.ToLower(title)
	slug = strings.ReplaceAll(slug," ", "-")
	return slug
}

func (s *Service) GetProblems(ctx context.Context) ([]models.Problem, error){
	return s.repo.GetProblems(ctx)
}

func(s *Service) GetProblemBySlug(ctx context.Context, slug string) (*models.Problem, error){
	return s.repo.GetProblemBySlug(ctx,slug)
}

func (s *Service) DeleteProblemBySlug(ctx context.Context, slug string) error {
	return s.repo.DeleteProblemBySlug(ctx,slug)
}
