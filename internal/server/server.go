package server

import (
	"log"
	"net/http"

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
	router.GET("/", func (c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

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

