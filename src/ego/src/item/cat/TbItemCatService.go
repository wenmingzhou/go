package cat

//根据id查询类目
func ShowCatByIdService(id int) *TbItemCat {
	return selByIdDao(id)
}
