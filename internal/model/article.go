package model

import "github.com/halweg/gin-blog/pkg/app"

type Article struct {
    *Model
    Title string `json:"title"`
    Desc string `json:"desc"`
    Content string `json:"Content"`
    CoverImageUrl string `json:"CoverImageUrl"`
    State uint8 `json:"State"`
}

type articleSwagger struct {
    List []*Article
    Pager *app.Pager
}

func (a *Article) TableName() string {
    return "blog_article"
}
