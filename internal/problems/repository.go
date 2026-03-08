package problems

import (
	"context"

	"github.com/Aneeshie/cpp-judge/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct{
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}


func (r *Repository) CreateProblem(ctx context.Context, p *models.Problem) (*models.Problem, error){

	query := `
		INSERT into problems (
			slug,
			title,
			description,
			difficulty,
			time_limit_ms,
			memory_limit_mb
		)
		VALUES ($1,$2,$3,$4,$5,$6)
		RETURNING id, created_at
	`

	err := r.db.QueryRow(
		ctx,
		query,
		p.Slug,
		p.Title,
		p.Description,
		p.Difficulty,
		p.TimeLimitMs,
		p.MemoryLimitMb,
	).Scan(&p.ID, &p.CreatedAt)

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *Repository) GetProblems(ctx context.Context) ([]models.Problem, error){
	query := `
	SELECT
	id,
	slug,
	title,
	difficulty,
	time_limit_ms,
	memory_limit_mb,
	created_at
	FROM problems
	ORDER BY created_at DESC
	LIMIT 20;
	`

	rows,err := r.db.Query(ctx,query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var problems []models.Problem

	for rows.Next() {
		var p models.Problem

		err := rows.Scan(
			&p.Slug,
			&p.Title,
			&p.Difficulty,
			&p.Description,
			&p.TimeLimitMs,
			&p.MemoryLimitMb,
			&p.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		problems = append(problems, p)
	}

	return problems,err
}

