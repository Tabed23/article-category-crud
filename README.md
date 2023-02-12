# Golang Gin With PostgreSQL and Redis

## How to Run

1. sudo docker run -d --rm -v ${PWD}/postges:/var/lib/postgresql/data -e POSTGRES_DB=dev -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin123 postgres:10.4
2. endpoints
```
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
```

# To Access  the PostgreSQL database
1.sudo docker exec -it postgres  psql -U admin -d websays
