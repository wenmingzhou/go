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

func ArticlePage(pi int, ps int) ([]Article, error) {
	mods := make([]Article, 0, 10)
	err := DB.Unsafe().Select(&mods, "select * from Article order by id desc limit ?,?", (pi-1)*ps, ps)
	return mods, err
}

func ArticlePageCount() int {
	count := 0
	DB.Get(&count, "select count(*) as count from article")
	return count
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

func ArticleEdit(mod *Article) error {
	_, err := DB.Exec("update article set title=?,author=?,content=?,hits=? where id=? ",
		mod.Title, mod.Author, mod.Content, mod.Hits, mod.Id)
	return err
}
