package server

import (
	"log"

	"github.com/Aneeshie/cpp-judge/internal/problems"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct{
	router *gin.Engine
	db *pgxpool.Pool
}

func NewServer(db *pgxpool.Pool) *Server{
	router := gin.Default()

	//all the routes and stuff below
	problemRepo := problems.NewRepository(db)
	problemService := problems.NewService(problemRepo)
	problemHandler := problems.NewHandler(problemService)

	//problem routes
	router.POST("/problems",problemHandler.CreateProblemHandler)
	router.POST("/problems/:slug",problemHandler.GetProblemBySlugHandler)
	router.DELETE("/problems/:slug",problemHandler.DeleteProblemHandler)
	router.GET("/problems", problemHandler.GetProblemsHandler)

	return &Server{
		router: router,
		db: db,
	}
}

func (s *Server) Run(port string){
	if err := s.router.Run(":"+port); err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}
}

