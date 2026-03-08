# C++ Judge

A backend system for running and evaluating C++ code submissions against programming problems.

This project is being built as a learning project to explore backend architecture, database design, and building an online judge system.

⚠️ This project is still under active development.

## Features

- Create programming problems
- Retrieve problems by slug
- List problems
- PostgreSQL database with migrations
- Clean architecture (Handler → Service → Repository)
- Docker-based development environment

## Tech Stack

- Go
- Gin (HTTP framework)
- PostgreSQL
- pgx (Postgres driver)
- Docker
- golang-migrate (database migrations)

## Project Structure
```bash
cmd/
    server/           # application entry point

internal/
    server/           # server setup and routes
    problems/         # problems module
        repository.go
        service.go
        handler.go
    models/           # shared models
    database/         # postgres connection
    config/           # configuration

migrations/           # database migrations
```

## Running the Project

Start the services:

```bash
docker compose up
```

## Future Work

- Submissions system
- Code execution worker
- Testcase runner
- Judge queue system
- User authentication
- Problem testcases storage
