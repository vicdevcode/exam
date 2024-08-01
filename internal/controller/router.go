package controller

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/vicdevcode/exam/internal/app/config"
	"github.com/vicdevcode/exam/internal/usecase"
)

func NewRouter(handler *gin.Engine, cfg *config.Config, uc usecase.UseCases) {
	h := handler.Group("/api")

	h.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))

	{
		newCategory(h, uc.CategoryUseCase)
		newSubCategory(h, uc.SubCategoryUseCase)
		newCard(h, uc.CardUseCase)
	}
}
