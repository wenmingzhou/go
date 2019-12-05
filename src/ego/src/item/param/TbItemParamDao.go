package param

import (
	"ego/src/commons"
	"fmt"
	"strconv"
)

func selByPageDao(page, rows int) []TbItemParam {
	r, err := commons.Dql("select * from tb_item_param limit ?,?", rows*(page-1), rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	ts := make([]TbItemParam, 0)

	for r.Next() {
		var t TbItemParam
		r.Scan(&t.Id, &t.ItemCatId, &t.ParamData, &t.Created, &t.Updated)
		ts = append(ts, t)
	}
	commons.CloseConn()
	return ts

}

//查询总个数
func selCountDao() int {
	rows, err := commons.Dql("select count(*) from tb_item_param")
	if err != nil {
		fmt.Println(err)
		return -1
	}
	if rows.Next() {
		var count int
		rows.Scan(&count)
		return count
	}

	return -1
}

func delByIdDao(ids []int) int {
	sql := "delete from tb_item_param where id in ("
	for i := 0; i < len(ids); i++ {
		sql += strconv.Itoa(ids[i])
		if i < len(ids)-1 {
			sql += ","
		}
	}
	sql += ")"
	fmt.Println(sql)
	count, err := commons.Dml(sql)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return int(count)
}

//根据类目id查询规格参数
func selByCatIdDao(catid int) *TbItemParam {
	r, err := commons.Dql("select * from tb_item_param where item_cat_id=?", catid)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if r.Next() {
		param := new(TbItemParam)
		r.Scan(&param.Id, &param.ItemCatId, &param.ParamData, &param.Created, &param.Updated)
		return param
	}
	return nil
}
