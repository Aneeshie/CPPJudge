CREATE TYPE verdict as ENUM (
    'accepted',
    'wrong_answer',
    'tle',
    'runtime_error',
    'compile_error',
    'mle'
);

CREATE TYPE submission_status as ENUM (
    'pending',
    'running',
    'completed'
);

CREATE TABLE submissions (
    id BIGSERIAL PRIMARY KEY,

    problem_id BIGINT NOT NULL REFERENCES problems(id),

    user_id BIGINT,

    code TEXT NOT NULL,
    language TEXT NOT NULL,

    status submission_status NOT NULL DEFAULT 'pending',
    verdict verdict,


    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_submissions_problem_id
ON submissions(problem_id);