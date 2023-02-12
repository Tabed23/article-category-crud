package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Tabed23/article-category-crud/app/service"
	"github.com/Tabed23/article-category-crud/app/types"
	"github.com/gin-gonic/gin"
)


type ArticleServer struct {
	srv service.ArticleService
	ar map[string]types.Article
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
	resp, err := a.srv.CreatArticle(&art)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Cannot insert data into database",
		})
		return
	}

	id := strconv.Itoa(int(resp))
	a.ar[id] = art

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": fmt.Sprintf("article created successfully with id %d", resp),
	})

}
func (a *ArticleServer) GetArticle(c *gin.Context) {
	artId := c.Param("id")
	_, ok := a.ar[artId]

	if !ok {
		c.JSON(400, gin.H{"error": "Article not found"})
		return
	}

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
	_, ok := a.ar[artId]
		if !ok {
			c.JSON(400, gin.H{"error": "Article not found"})
			return
		}

	art := types.Article{}
	if err := c.ShouldBindJSON(&art); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "invalid json",
		})
		return
	}
	a.ar[artId] = art
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

	_, ok := a.ar[artId]
		if !ok {
			c.JSON(400, gin.H{"error": "Article not found"})
			return
		}

	delete(a.ar, artId)
	ok, _ = a.srv.DeleteArticle(artId)
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
