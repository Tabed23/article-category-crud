package service

import (
	"errors"
	"time"

	"github.com/Tabed23/article-category-crud/app/database"
	"github.com/Tabed23/article-category-crud/app/types"
)

var (
	db = database.DB
	ar map[string]types.Article
)

type ArticleService struct {

}


func ( s *ArticleService) CreatArticle(article *types.Article) (error) {
	art := &types.Article{
		ArticleID: article.ArticleID,
		Author:      article.Author,
		Title:       article.Title,
		Content:     article.Content,
		Category:    article.Category,
		PublishedAt: time.Now(),
	}
	ar = make(map[string]types.Article, 0)
	ar[art.ArticleID] = *art
	tx := db.Table("articles").Exec("INSERT INTO articles (article_id, title, content, category, author, published_at) VALUES ($1, $2, $3,$4,$5,$6) RETURNING id",art.ArticleID, &art.Title, &art.Content, &art.Category, &art.Author, &art.PublishedAt)

	return tx.Error
}

func (s *ArticleService) FindById(id string) (*types.Article, error) {
	art := types.Article{}
	_, ok := ar[id]

	if !ok {
		return nil, errors.New("no such article")
	}else if ok {
		found := ar[id]
		return &found, nil
	}

	err := db.Table("articles").Where("article_id = ?", id).First(&art).Error
	return &art, err
}

func ( s* ArticleService) DeleteArticle(id string) (bool, error) {
	_, ok := ar[id]
	if !ok {
		return false, errors.New("no such article")
	}

	art := types.Article{}
	err := db.Table("articles").Where("article_id = ?", id).First(&art).Error
	if err != nil {
		return false, err
	}

	delete(ar, id)

	err = db.Delete(art, id).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *ArticleService) UpdateArticle(id string, artUpdate types.Article) (*types.Article, error) {
	_, ok := ar[id]

	if !ok {
		return nil, errors.New("no such article")
	}
	art := types.Article{}
	err := db.Table("articles").Where("article_id = ?", id).First(&art).Error
	if err != nil {
		return nil, err
	}

	art.Author = artUpdate.Author
	art.Category = artUpdate.Category
	art.Content = artUpdate.Content
	art.PublishedAt = artUpdate.PublishedAt

	ar[id] = art
	err = db.Save(&art).Error
	if err != nil {
		return nil, err
	}
	return &art, nil
}

func (s *ArticleService) GetArticle() (*[]types.Article, error) {

	art := make([]types.Article, 0)
	err := db.Table("articles").Find(&art).Error
	if err != nil {
		return nil, err
	}
	return &art, nil
}
