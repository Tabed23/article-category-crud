package routes

import (
	"github.com/Tabed23/article-category-crud/app/database"
	"github.com/Tabed23/article-category-crud/app/types"
	"github.com/gin-gonic/gin"
)

var (
	r = gin.Default()
)

const (
	port = ":8080"
)

func StartApp() {
	db := database.ConnectDB()
	types.Migration(db)
	UrlMaps()
	r.Run(port)

}
