CREATE TABLE problems (
  id BIGSERIAL PRIMARY KEY,
  slug TEXT UNIQUE NOT NULL,
  title TEXT NOT NULL,
  description TEXT NOT NULL,
  difficulty TEXT,

  time_limit_ms INTEGER NOT NULL DEFAULT 1000 CHECK (time_limit_ms > 0),
  memory_limit_mb INTEGER NOT NULL DEFAULT 256 CHECK (memory_limit_mb > 0),

  created_at TIMESTAMP DEFAULT NOW()
);
