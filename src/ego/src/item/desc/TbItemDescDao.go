package desc

import (
	"ego/src/commons"
)

func insertDescDao(t TbItemDesc) int {
	count, err := commons.Dml("insert into tb_item_desc values (?,?,?,?)", t.ItemId, t.ItemDesc, t.Created, t.Updated)
	if err != nil {
		return -1
	}
	return int(count)
}
