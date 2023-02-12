package routes

import (
	"net/http"

	"github.com/Tabed23/article-category-crud/app/controller"
	"github.com/gin-gonic/gin"
)

func UrlMaps() {
	r.GET("/api/v1/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "API is Alive")
	})
	art := controller.ArticleServer{}
	cat := controller.CategoryServer{}
	v1 := r.Group("/v1/article")
	{
		v1.GET("/articles", art.GetArticles)
		v1.GET("/articles/:id", art.GetArticle)
		v1.DELETE("/articles/:id", art.DeleteArticle)
		v1.PUT("/articles/:id", art.UpdateArticle)
		v1.POST("/articles", art.CreatArticle)
	}

	v2 := r.Group("/v2/category")
	{
		v2.GET("/category", cat.GetCategories)
		v2.GET("/category/:id", cat.GetCategory)
		v2.DELETE("/category/:id", cat.DeleteCategory)
		v2.PUT("/category/:id", cat.UpdateCategory)
		v2.POST("/category", cat.CreatCategory)
	}
}
