package param

import (
	"ego/src/commons"
	"ego/src/item/cat"
	"strconv"
	"strings"
)

//显示规则参数
func showParamService(page, rows int) (e commons.Datagrid) {
	t := selByPageDao(page, rows)
	e.Count = selCountDao()
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

//删除规格参数
func delByIdsService(ids string) (e commons.EgoResult) {
	idsStr := strings.Split(ids, ",")
	idInt := make([]int, 0)
	for _, n := range idsStr {
		id, _ := strconv.Atoi(n)
		idInt = append(idInt, id)
	}
	count := delByIdDao(idInt)
	if count > 0 {
		e.Status = 200
	}
	return
}

//根据类目id查询规则参数是否已经添加
func catIdService(catid int) (e commons.EgoResult) {
	p := selByCatIdDao(catid)
	if p == nil {
		e.Status = 200
	}
	return

}
