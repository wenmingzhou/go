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
func selByPageDao(page ,rows int) []Tbitem {
	r,err :=commons.Dql("select * from tb_item limit ?,?",rows*(page-1),rows)
	if err!=nil{
		fmt.Println(err)
		return nil
	}
	ts :=make([]Tbitem,0)
	for r.Next(){
		var t Tbitem
		//由于barcode数据库中为Null会导致填充错误
		var s sql.NullString
		r.Scan(&t.Id,&t.Title,&t.SellPoint,&t.Price,
			&t.Num,&s,&t.Image,&t.Cid,&t.Status,
			&t.Created,&t.Updated)
		t.Barcode =s.String
		ts =append(ts,t)
	}
	commons.CloseConn()
	return ts
}


//查询总条数
func selTotalDao() (total int) {
	rows,err :=commons.Dql("select count(*) from tb_item")
	if err!=nil{
		fmt.Println(err)
		return -1
	}
	rows.Next()
	rows.Scan(&total)
	commons.CloseConn()
	return
}