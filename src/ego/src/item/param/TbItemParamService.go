package param

import (
	"ego/src/commons"
	"ego/src/item/cat"
	"fmt"
)

//显示规则参数
func showParamService(page, rows int) (e commons.Datagrid) {

	t := selByPageDao(page, rows)
	e.Count = selCountDao()
	fmt.Println(e.Count)
	cats := make([]TbItemParamCat, 0)
	for i := 0; i < len(t); i++ {
		var catItem TbItemParamCat
		catItem.Id = t[i].Id
		catItem.Updated = t[i].Updated
		catItem.Created = t[i].Created
		catItem.ParamData = t[i].ParamData
		catItem.ItemCatId = t[i].ItemCatId
		catItem.CatName = cat.ShowCatByIdService(t[i].ItemCatId).Name
		cats = append(cats, catItem)
	}
	e.Data = cats
	return
}
