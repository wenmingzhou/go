package item

import (
	"ego/src/commons"
	"encoding/json"
	"net/http"
	"strconv"
)

func ItemHandler() {
	commons.Router.HandleFunc("/showItem", ShowItemController)
	commons.Router.HandleFunc("/item/delete", delByIdsController)
	commons.Router.HandleFunc("/item/inStock", inStockByIdsController)
	commons.Router.HandleFunc("/item/outStock", outStockByIdsController)
	commons.Router.HandleFunc("/item/imageUpload", imagesUploadController)
	commons.Router.HandleFunc("/item/add", insertController)
	commons.Router.HandleFunc("/item/showItemById", showItemDescCatController)

}

//显示修改页面信息
func showItemDescCatController(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))

	er := showItemDescCatService(id)

	b, _ := json.Marshal(er)
	w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
	w.Write(b)
}

//商品新增
func insertController(w http.ResponseWriter, r *http.Request) {
	//需要先进行解析
	r.ParseForm()
	er := insertService(r.Form)
	b, _ := json.Marshal(er)
	w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
	w.Write(b)
}

//图片上传
func imagesUploadController(w http.ResponseWriter, r *http.Request) {

	file, fileHeader, err := r.FormFile("uploadFile")
	if err != nil {
		m := make(map[string]interface{})
		m["error"] = 1
		m["message"] = "接受图片失败"
		b, _ := json.Marshal(m)
		w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
		w.Write(b)
		return
	}
	m := imageUploadService(file, fileHeader)
	b, _ := json.Marshal(m)
	w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
	w.Write(b)

}
func ShowItemController(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.FormValue("page"))
	rows, _ := strconv.Atoi(r.FormValue("limit"))

	datagrid := showItemService(page, rows)
	b, _ := json.Marshal(datagrid)
	w.Header().Set("Content-type", "application/json;charset=utf-8")
	w.Write(b)
}

//删除
func delByIdsController(w http.ResponseWriter, r *http.Request) {
	ids := r.FormValue("id")
	er := delByIdsService(ids)
	b, _ := json.Marshal(er)
	w.Header().Set("Content-type", "application/json;charset=utf-8")
	w.Write(b)
}

//上架
func inStockByIdsController(w http.ResponseWriter, r *http.Request) {
	ids := r.FormValue("id")

	er := inStockByIdsService(ids)
	b, _ := json.Marshal(er)
	w.Header().Set("Content-type", "application/json;charset=utf-8")
	w.Write(b)
}

//下架
func outStockByIdsController(w http.ResponseWriter, r *http.Request) {
	ids := r.FormValue("id")

	er := outStockByIdsService(ids)
	b, _ := json.Marshal(er)
	w.Header().Set("Content-type", "application/json;charset=utf-8")
	w.Write(b)
}
