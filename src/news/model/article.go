package model

import "time"

type Article struct {
	Id   	int64 `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
	Hits   	string   `json:"hits"`
	Utime   time.Time `json:"utime"`
}

//返回单条
func ArticleGet(id int64)(Article,error){
	mod :=Article{}
	err :=Db.Unsafe().Get(&mod,"select * from Article where id=? limit 1",id)
	return  mod,err
}

//返回多条
func ArticleList()([]Article,error){
	mods :=make([]Article,0,10)
	err :=Db.Unsafe().Select(&mods,"select * from Article limit 10")
	return  mods,err
}

//返回多条
func ArticleDel(id int64) bool{
	res,_ :=Db.Exec("delete from Article where id =?",id)
	if res ==nil{
		return false
	}
	rows,_ :=res.RowsAffected()
	if rows>=1{
		return  true
	}else{
		return false
	}
}

//添加
func ArticleAdd(mod * Article) error{
	_,err:=Db.Unsafe().Exec("insert into Article (title,author,content,hits,utime) value(?,?,?,?,?)",
		mod.Title,mod.Author,mod.Content,mod.Hits,mod.Utime)
	return err

}