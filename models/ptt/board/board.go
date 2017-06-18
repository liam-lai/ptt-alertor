package board

import (
	"reflect"

	"github.com/liam-lai/ptt-alertor/models/ptt/article"
	"github.com/liam-lai/ptt-alertor/rss"
)

type BoardNotExistError struct {
	Suggestion string
}

func (e BoardNotExistError) Error() string {
	return "board is not exist"
}

type Board struct {
	Name        string
	Articles    article.Articles
	NewArticles article.Articles
}

type BoardAction interface {
	OnlineArticles() article.Articles
	GetArticles() article.Articles
	WithArticles()
	Create() error
}

func (bd Board) OnlineArticles() article.Articles {
	bd.Articles = rss.BuildArticles(bd.Name)
	return bd.Articles
}

func NewArticles(bd BoardAction) article.Articles {
	savedArticles := bd.GetArticles()
	onlineArticles := bd.OnlineArticles()
	newArticles := make(article.Articles, 0)
	for _, onlineArticle := range onlineArticles {
		for index, savedArticle := range savedArticles {
			if reflect.DeepEqual(onlineArticle, savedArticle) {
				break
			}
			if index == len(savedArticles)-1 {
				newArticles = append(newArticles, onlineArticle)
			}
		}
	}
	return newArticles
}
