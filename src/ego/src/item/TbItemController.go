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
