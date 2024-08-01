package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/vicdevcode/exam/internal/entity"
	"github.com/vicdevcode/exam/internal/usecase"
)

type cardRoute struct{ uc usecase.Card }

func newCard(handler *gin.RouterGroup, uc usecase.Card) {
	r := cardRoute{uc}
	h := handler.Group("/card")
	{
		h.GET("/all", r.getAll)
		h.POST("/create", r.create)
		h.POST("/update", r.update)
		h.DELETE("/delete", r.delete)
	}
}

func (r *cardRoute) getAll(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		badRequest(c, err.Error())
		return
	}
	categories, err := r.uc.FindAllBySubCategoryID(c.Request.Context(), uint(id))
	if err != nil {
		internalServerError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"cards": categories})
}

type createCardRequest struct {
	Question   string `json:"question"`
	Answer     string `json:"answer"`
	CategoryID uint   `json:"sub_category_id"`
}

func (r *cardRoute) create(c *gin.Context) {
	var body createCardRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		badRequest(c, err.Error())
		return
	}

	Card, err := r.uc.Create(
		c.Request.Context(),
		entity.Card{SubCategoryID: body.CategoryID, Question: body.Question, Answer: body.Answer},
	)
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"card": Card})
}

type updateCardRequest struct {
	ID         uint   `json:"id"`
	Question   string `json:"question"`
	Answer     string `json:"answer"`
	CategoryID uint   `json:"sub_category_id"`
}

func (r *cardRoute) update(c *gin.Context) {
	var body updateCardRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		badRequest(c, err.Error())
		return
	}

	Card, err := r.uc.Update(
		c.Request.Context(),
		entity.Card{
			ID:            body.ID,
			SubCategoryID: body.CategoryID,
			Question:      body.Question,
			Answer:        body.Answer,
		},
	)
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"card": Card})
}

func (r *cardRoute) delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		badRequest(c, err.Error())
		return
	}

	if err := r.uc.Delete(c.Request.Context(), uint(id)); err != nil {
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "card was deleted"})
}
