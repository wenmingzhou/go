package param

import (
	"ego/src/commons"
	"fmt"
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
