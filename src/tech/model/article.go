package model

import "time"

type Article struct {
	Id      int64     `json:"id"`
	Title   string    `json:"title"`
	Author  string    `json:"author"`
	Content string    `json:"content"`
	Hits    string    `json:"hits"`
	Utime   time.Time `json:"utime"`
}

func ArticleGet(id int64) (Article, error) {
	mod := Article{}
	err := DB.Unsafe().Get(&mod, "select * from Article where id =? limit 1", id)
	return mod, err
}
