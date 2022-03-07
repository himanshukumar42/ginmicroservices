package main

import (
	"testing"
)

func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	if len(alist) != len(articleList) {
		t.Errorf("Lengths are not equal")
	}

	for i, v := range alist {
		if v.Content != articleList[i].Content || v.ID != articleList[i].ID || v.Title != articleList[i].Title {
			t.Errorf("Difference of articles")
		}
	}
}

func TestGetArticleByID(t *testing.T) {
	article, err := getArticleByID(1)
	if err != nil || article.ID != 1 || article.Title != "Article 1" || article.Content != "Article 1 body" {
		t.Fail()
	}
}
