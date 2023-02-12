package service

import (

	"time"

	"github.com/Tabed23/article-category-crud/app/database"
	"github.com/Tabed23/article-category-crud/app/types"
)

var (
	db = database.DB
)

type ArticleService struct {
}

func (ArticleService) CreatArticle(article *types.Article) (int64, error) {
	art := &types.Article{
		Author:      article.Author,
		Title:       article.Title,
		Content:     article.Content,
		Category:    article.Category,
		PublishedAt: time.Now(),
	}
	var id int64
	tx := db.Table("articles").Exec("INSERT INTO articles (title, content, category, author, published_at) VALUES ($1, $2, $3,$4,$5) RETURNING id", &art.Title, &art.Content, &art.Category, &art.Author, &art.PublishedAt)

	return id, tx.Error
}

func (ArticleService) FindById(id string) (*types.Article, error) {
	art := types.Article{}

	err := db.Table("articles").Where("id = ?", id).First(&art).Error
	return &art, err
}

func (ArticleService) DeleteArticle(id string) (bool, error) {
	art := types.Article{}
	err := db.Table("articles").Where("id = ?", id).First(&art).Error
	if err != nil {
		return false, err
	}
	err = db.Delete(art, id).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (ArticleService) UpdateArticle(id string, artUpdate types.Article) (*types.Article, error) {
	art := types.Article{}
	err := db.Table("articles").Where("id = ?", id).First(&art).Error
	if err != nil {
		return nil, err
	}

	art.Author = artUpdate.Author
	art.Category = artUpdate.Category
	art.Content = artUpdate.Content
	art.PublishedAt = artUpdate.PublishedAt

	err = db.Save(&art).Error
	if err != nil {
		return nil, err
	}
	return &art, nil
}

func (ArticleService) GetArticle() (*[]types.Article, error) {
	art := make([]types.Article, 0)
	err := db.Table("articles").Find(&art).Error
	if err != nil {
		return nil, err
	}
	return &art, nil
}
