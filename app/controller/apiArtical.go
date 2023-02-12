package controller

import (
	"fmt"
	"net/http"

	"github.com/Tabed23/article-category-crud/app/service"
	"github.com/Tabed23/article-category-crud/app/types"
	"github.com/gin-gonic/gin"
)

type ArticleServer struct {
	srv service.ArticleService
}

func (a *ArticleServer) CreatArticle(c *gin.Context) {

	var art types.Article
	if err := c.ShouldBindJSON(&art); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "invalid json",
		})
		return
	}

	err := a.srv.CreatArticle(&art)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Cannot insert data into database",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": fmt.Sprintf("article created successfully with id %s", art.ArticleID),
	})

}
func (a *ArticleServer) GetArticle(c *gin.Context) {
	artId := c.Param("id")

	resp, err := a.srv.FindById(artId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Cannot find the article",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": resp,
	})
}

func (a *ArticleServer) GetArticles(c *gin.Context) {
	art, err := a.srv.GetArticle()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "there are no articles available",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   &art,
	})
}

func (a *ArticleServer) UpdateArticle(c *gin.Context) {
	artId := c.Param("id")

	art := types.Article{}
	if err := c.ShouldBindJSON(&art); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "invalid json",
		})
		return
	}
	resp, err := a.srv.UpdateArticle(artId, art)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Cannot update the article",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": resp,
	})

}

func (a *ArticleServer) DeleteArticle(c *gin.Context) {
	artId := c.Param("id")

	ok, _ := a.srv.DeleteArticle(artId)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Cannot delete the article",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("article deleted %t", ok),
	})
}
