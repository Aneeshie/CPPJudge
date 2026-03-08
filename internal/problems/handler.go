package problems

import (
	"net/http"

	"github.com/Aneeshie/cpp-judge/internal/models"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateProblemHandler(c *gin.Context){
	var req models.CreateProblemRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if req.Description == "" || req.Difficulty == "" || req.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid inputs"})
		return
	}

	problem, err := h.service.CreateProblem(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create problem", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"problem": problem})
}

func (h *Handler) GetProblemsHandler(c *gin.Context){
	problems, err := h.service.GetProblems(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not fetch problems",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"problems": problems,
	})
}

func (h *Handler) GetProblemBySlugHandler(c *gin.Context){
	slug := c.Param("slug")
	problem, err:= h.service.GetProblemBySlug(c.Request.Context(), slug)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "could'nt find the problem", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, problem)
}

func (h *Handler) DeleteProblemHandler (c *gin.Context){
	slug := c.Param("slug")

	err := h.service.DeleteProblemBySlug(c.Request.Context(), slug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "problem not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "problem deleted",
	})
}
