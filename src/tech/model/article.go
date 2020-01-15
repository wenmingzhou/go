package model

import "time"

type Article struct {
	Id      int64     `json:"id"`
	Title   string    `json:"title"`
	Author  string    `json:"author"`
	Content string    `json:"content"`
	Hits    int       `json:"hits"`
	Utime   time.Time `json:"utime"`
}

func ArticleGet(id int64) (Article, error) {
	mod := Article{}
	err := DB.Unsafe().Get(&mod, "select * from Article where id =? limit 1", id)
	return mod, err
}

func ArticleList() ([]Article, error) {
	mods := make([]Article, 0, 10)
	err := DB.Unsafe().Select(&mods, "select * from Article order by id desc limit 10")
	return mods, err
}

func ArticleDel(id int64) bool {
	res, _ := DB.Exec("delete from Article where id= ?", id)
	if res == nil {
		return false
	}
	rows, _ := res.RowsAffected()
	if rows >= 1 {
		return true
	} else {
		return false
	}

}

func ArticleAdd(mod *Article) error {
	_, err := DB.Exec("insert into article (title,author,content,hits,utime) values (?,?,?,?,?) ",
		mod.Title, mod.Author, mod.Content, mod.Hits, mod.Utime)
	return err
}
