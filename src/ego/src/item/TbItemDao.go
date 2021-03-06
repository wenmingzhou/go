package item

import (
	"database/sql"
	"ego/src/commons"
	"fmt"
)

/*
rows:每页显示的条数
page:当前第几页
*/
func selByPageDao(page, rows int) []Tbitem {
	r, err := commons.Dql("select * from tb_item limit ?,?", rows*(page-1), rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	ts := make([]Tbitem, 0)
	for r.Next() {
		var t Tbitem
		//由于barcode数据库中为Null会导致填充错误
		var s sql.NullString
		r.Scan(&t.Id, &t.Title, &t.SellPoint, &t.Price,
			&t.Num, &s, &t.Image, &t.Cid, &t.Status,
			&t.Created, &t.Updated)
		t.Barcode = s.String
		ts = append(ts, t)
	}
	commons.CloseConn()
	return ts
}

//查询总条数
func selTotalDao() (total int) {
	rows, err := commons.Dql("select count(*) from tb_item")
	if err != nil {
		fmt.Println(err)
		return -1
	}
	rows.Next()
	rows.Scan(&total)
	commons.CloseConn()
	return
}

func updateStatusByIdsDao(ids []string, status int) int {
	if len(ids) <= 0 {
		return -1
	}
	sql := "update tb_item set status= ? where "
	for i := 0; i < len(ids); i++ {
		sql += " id=" + ids[i]
		if i < len(ids)-1 {
			sql += " or "
		}
	}
	count, err := commons.Dml(sql, status)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return int(count)
}

//商品新增
func insertItemDao(t Tbitem) int {
	count, err := commons.Dml("insert into tb_item values (?,?,?,?,?,?,?,?,?,?,?) ",
		t.Id, t.Title, t.SellPoint, t.Price, t.Num, t.Barcode, t.Image, t.Cid, t.Status, t.Created, t.Updated)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return int(count)
}

//根据商品id删除数据
func delById(id int) int {
	count, err := commons.Dml("delete from tb_item where id=?", id)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return int(count)
}

//根据id查询商品信息
func SelByIdDao(id int) *Tbitem {
	rows, err := commons.Dql("select * from tb_item where id =?", id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if rows.Next() {
		t := new(Tbitem)
		var s sql.NullString
		rows.Scan(&t.Id, &t.Title, &t.SellPoint, &t.Price,
			&t.Num, &s, &t.Image, &t.Cid, &t.Status,
			&t.Created, &t.Updated)
		t.Barcode = s.String
		return t
	} else {
		return nil
	}
}

//带有事务的更新商品表数据
func updateByIdWithTx(t Tbitem) int {
	return commons.PrepareWithTx("update tb_item set title=?,sell_point=?,price=?,num=?,image=?,cid=?,updated=? where id=?",
		t.Title, t.SellPoint, t.Price, t.Num, t.Image, t.Cid, t.Updated, t.Id)
}
