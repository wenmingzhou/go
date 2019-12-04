package desc

import (
	"ego/src/commons"
	"fmt"
)

func insertDescDao(t TbItemDesc) int {
	count, err := commons.Dml("insert into tb_item_desc values (?,?,?,?)", t.ItemId, t.ItemDesc, t.Created, t.Updated)
	if err != nil {
		return -1
	}
	return int(count)
}

//根据id查询分类信息
func selByIdDao(id int) *TbItemDesc {
	rows, err := commons.Dql("select * from tb_item_desc where item_id =?", id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if rows.Next() {
		t := new(TbItemDesc)
		rows.Scan(&t.ItemId, &t.ItemDesc, &t.Created, &t.Updated)
		return t
	} else {
		return nil
	}
}
