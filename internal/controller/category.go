package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/vicdevcode/exam/internal/entity"
	"github.com/vicdevcode/exam/internal/usecase"
)

type categoryRoute struct {
	uc usecase.Category
}

func newCategory(handler *gin.RouterGroup, uc usecase.Category) {
	r := categoryRoute{uc}
	h := handler.Group("/category")
	{
		h.GET("/all", r.getAll)
		h.POST("/create", r.create)
		h.POST("/update", r.update)
		h.DELETE("/delete", r.delete)
	}
}

func (r *categoryRoute) getAll(c *gin.Context) {
	categories, err := r.uc.FindAll(c.Request.Context())
	if err != nil {
		internalServerError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

type createCategoryRequest struct {
	Name string `json:"name"`
}

func (r *categoryRoute) create(c *gin.Context) {
	var body createCategoryRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		badRequest(c, err.Error())
		return
	}

	category, err := r.uc.Create(c.Request.Context(), entity.Category{Name: body.Name})
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"category": category})
}

type updateCategoryRequest struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (r *categoryRoute) update(c *gin.Context) {
	var body updateCategoryRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		badRequest(c, err.Error())
		return
	}

	SubCategory, err := r.uc.Update(
		c.Request.Context(),
		entity.Category{ID: body.ID, Name: body.Name},
	)
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"category": SubCategory})
}

func (r *categoryRoute) delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		badRequest(c, err.Error())
		return
	}

	if err := r.uc.Delete(c.Request.Context(), uint(id)); err != nil {
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "category was deleted"})
}
