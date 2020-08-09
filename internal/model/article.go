package model

type Article struct {
    *Model
    Title string `json:"title"`
    Desc string `json:"desc"`
    Content string `json:"Content"`
    CoverImageUrl string `json:"CoverImageUrl"`
    State uint8 `json:"State"`
}

func (a *Article) TableName() string {
    return "blog_article"
}
