package controller

import (
	"fmt"
	"github.com/Tabed23/article-category-crud/app/service"
	"github.com/Tabed23/article-category-crud/app/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryServer struct {
	srv service.CategoryService
}

func (cat *CategoryServer) CreatCategory(c *gin.Context) {
	var ca types.Category
	if err := c.ShouldBindJSON(&ca); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "invalid json",
		})
		return
	}
	resp, err := cat.srv.CreatCategory(&ca)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Cannot insert data into database",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": resp,
	})

}
func (cat *CategoryServer) GetCategory(c *gin.Context) {
	catId := c.Param("id")

	resp, err := cat.srv.CategoryFindById(catId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Cannot find the category",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": resp,
	})
}

func (cat *CategoryServer) GetCategories(c *gin.Context) {

	ca, err := cat.srv.GetCategory()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "there are no categories",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   &ca,
	})
}

func (cat *CategoryServer) UpdateCategory(c *gin.Context) {
	catId := c.Param("id")
	ca := types.Category{}
	if err := c.ShouldBindJSON(&ca); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "invalid json",
		})
		return
	}
	resp, err := cat.srv.UpdateCategory(catId, ca)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Cannot update the category",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": resp,
	})

}

func (cat *CategoryServer) DeleteCategory(c *gin.Context) {
	catId := c.Param("id")
	ok, _ := cat.srv.DeleteCategory(catId)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Cannot delete the category",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("category deleted %t", ok),
	})
}
