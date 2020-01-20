package model

type Class struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

//分页
func ClassPage(pi, ps int) ([]Class, error) {
	mods := make([]Class, 0, ps)
	err := DB.Select(&mods, "select * from class limit ?, ?", (pi-1)*ps, ps)
	return mods, err
}

//数据总数
func ClassCount() int {
	count := 0
	DB.Get(&count, "select count(*) as count from class")
	return count
}

//查询所有
func ClassAll() ([]Class, error) {
	mods := make([]Class, 0, 8)
	err := DB.Select(&mods, "select * from class ")
	return mods, err
}

//查询m某条数据
func ClassGet(id int64) (Class, error) {
	mod := Class{}
	err := DB.Get(&mod, "select  * from class where id= ? limit 1", id)
	return mod, err
}

//添加
//修改
//删除
