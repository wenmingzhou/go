package model

import (
	"errors"
)

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
func ClassAdd(mod *Class) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	res, err := tx.Exec("insert into class (`name`,`desc`) values (?,?)", mod.Name, mod.Desc)
	if err != nil {
		tx.Rollback()
		return err
	}
	rows, _ := res.RowsAffected()
	if rows < 1 {
		tx.Rollback()
		return errors.New("rows affected <1 ")
	}
	tx.Commit()
	return nil
}

//修改
//添加
func ClassEdit(mod *Class) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	res, err := tx.Exec("update  class set `name`=?,`desc`=? where id=?", mod.Name, mod.Desc, mod.Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	rows, _ := res.RowsAffected()
	if rows < 1 {
		tx.Rollback()
		return errors.New("rows affected <1 ")
	}
	tx.Commit()
	return nil
}

//删除

func ClassDrop(id int64) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	res, err := tx.Exec("delete from class where  id =? ", id)
	if err != nil {
		tx.Rollback()
		return err
	}
	rows, _ := res.RowsAffected()
	if rows < 1 {
		tx.Rollback()
		return errors.New("rows affected <1 ")
	}
	tx.Commit()
	return nil
}
