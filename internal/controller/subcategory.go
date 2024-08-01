package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/vicdevcode/exam/internal/entity"
	"github.com/vicdevcode/exam/internal/usecase"
)

type subCategoryRoute struct {
	uc usecase.SubCategory
}

func newSubCategory(handler *gin.RouterGroup, uc usecase.SubCategory) {
	r := subCategoryRoute{uc}
	h := handler.Group("/sub-category")
	{
		h.GET("/all", r.getAll)
		h.POST("/create", r.create)
		h.POST("/update", r.update)
		h.DELETE("/delete", r.delete)
	}
}

func (r *subCategoryRoute) getAll(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		badRequest(c, err.Error())
		return
	}
	categories, err := r.uc.FindAllByCategoryID(c.Request.Context(), uint(id))
	if err != nil {
		internalServerError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"sub_categories": categories})
}

type createSubCategoryRequest struct {
	Name       string `json:"name"`
	CategoryID uint   `json:"category_id"`
}

func (r *subCategoryRoute) create(c *gin.Context) {
	var body createSubCategoryRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		badRequest(c, err.Error())
		return
	}

	SubCategory, err := r.uc.Create(
		c.Request.Context(),
		entity.SubCategory{Name: body.Name, CategoryID: body.CategoryID},
	)
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"sub_category": SubCategory})
}

type updateSubCategoryRequest struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	CategoryID uint   `json:"category_id"`
}

func (r *subCategoryRoute) update(c *gin.Context) {
	var body updateSubCategoryRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		badRequest(c, err.Error())
		return
	}

	SubCategory, err := r.uc.Update(
		c.Request.Context(),
		entity.SubCategory{ID: body.ID, Name: body.Name, CategoryID: body.CategoryID},
	)
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"sub_category": SubCategory})
}

func (r *subCategoryRoute) delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		badRequest(c, err.Error())
		return
	}

	if err := r.uc.Delete(c.Request.Context(), uint(id)); err != nil {
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "sub category was deleted"})
}
