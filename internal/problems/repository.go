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
	description,
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
			&p.ID,
			&p.Slug,
			&p.Title,
			&p.Description,
			&p.Difficulty,
			&p.TimeLimitMs,
			&p.MemoryLimitMb,
			&p.CreatedAt,
			)
		if err != nil {
			return nil, err
		}

		problems = append(problems, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return problems, nil

}

func (r *Repository) GetProblemBySlug(ctx context.Context, slug string) (*models.Problem, error){
	query:= `
	SELECT
	id,
	slug,
	title,
	description,
	difficulty,
	time_limit_ms,
	memory_limit_mb,
	created_at
	FROM problems
	WHERE slug = $1
	`

	var problem models.Problem

	err := r.db.QueryRow(ctx, query, slug).Scan(
		&problem.ID,
		&problem.Slug,
		&problem.Title,
		&problem.Description,
		&problem.Difficulty,
		&problem.TimeLimitMs,
		&problem.MemoryLimitMb,
		&problem.CreatedAt,
		)

	if err != nil {
		return nil, err
	}

	return &problem,nil

}

func (r *Repository) DeleteProblemBySlug(ctx context.Context, slug string) (error) {
	query := `
	DELETE FROM problems
	WHERE slug = $1
	returning slug
	`

	var deleted string

	err := r.db.QueryRow(ctx,query,slug).Scan(&deleted)
	if err != nil {
		return err
	}

	return nil
}


func (r *Repository) UpdateProblem(ctx context.Context, slug string, req models.UpdateProblemRequest) (*models.Problem, error) {

	query := `
	UPDATE problems
	SET
		title = $1,
		description = $2,
		difficulty = $3,
		time_limit_ms = $4,
		memory_limit_mb = $5
	WHERE slug = $6
	RETURNING slug, title, description, difficulty, time_limit_ms, memory_limit_mb, created_at
	`

	var p models.Problem

	err := r.db.QueryRow(ctx, query,
		req.Title,
		req.Description,
		req.Difficulty,
		req.TimeLimitMs,
		req.MemoryLimitMb,
		slug,
	).Scan(
		&p.Slug,
		&p.Title,
		&p.Description,
		&p.Difficulty,
		&p.TimeLimitMs,
		&p.MemoryLimitMb,
		&p.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &p, nil
}
