package item

import (
	"ego/src/commons"
	"ego/src/item/cat"
	"strings"
)

func showItemService(page, rows int) (e *commons.Datagrid) {

	ts := selByPageDao(page, rows)
	total := selTotalDao()
	if ts != nil {
		itemChildren := make([]TbItemChild, 0)
		for i := 0; i < len(ts); i++ {
			var itemChild TbItemChild
			itemChild.Id = ts[i].Id
			itemChild.Updated = ts[i].Updated
			itemChild.Created = ts[i].Created
			itemChild.Status = ts[i].Status
			itemChild.Barcode = ts[i].Barcode
			//itemChild.Cid = ts[i].Cid
			//itemChild.Image = ts[i].Image
			itemChild.Price = ts[i].Price
			itemChild.Num = ts[i].Num
			itemChild.SellPoint = ts[i].SellPoint
			itemChild.Title = ts[i].Title
			itemChild.Title = ts[i].Title
			itemChild.CategoryName = cat.ShowCatByIdService(ts[i].Cid).Name
			itemChildren = append(itemChildren, itemChild)
		}
		e = new(commons.Datagrid)
		e.Data = itemChildren
		e.Count = total
		return
	}
	return nil
}

//删除商品
func delByIdsService(ids string) (e commons.EgoResult) {
	count := updateStatusByIdsDao(strings.Split(ids, ","), 3)
	if count > 0 {
		e.Status = 200
	}
	return
}
