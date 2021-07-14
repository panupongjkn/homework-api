package models

type News struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func NewNews(id int, title, body string) News {
	news := News{}
	news.ID = id
	news.Title = title
	news.Body = body
	return news
}
