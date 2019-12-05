package param

type TbItemParam struct {
	Id        int    `json:"id"`
	ItemCatId int    `json:"item_cat_id"`
	ParamData string `json:"param_data"`
	Created   string `json:"created"`
	Updated   string `json:"updated"`
}

type TbItemParamCat struct {
	TbItemParam
	CatName string `json:"cat_name"`
}
