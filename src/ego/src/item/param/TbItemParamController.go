package param

import (
	"ego/src/commons"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func ParamHandler() {
	commons.Router.HandleFunc("/item/param/show", showParamController)
	commons.Router.HandleFunc("/item/param/delete", delParamByIdController)
	commons.Router.HandleFunc("/item/param/iscat", isCatController)
}

//删除规格参数
func isCatController(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	fmt.Println(id)
}

//删除规格参数
func delParamByIdController(w http.ResponseWriter, r *http.Request) {
	er := delByIdsService(r.FormValue("id"))
	b, _ := json.Marshal(er)
	w.Header().Set("Content-type", "application/json;charset=utf-8")
	w.Write(b)
}

//显示规格参数
func showParamController(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.FormValue("page"))
	rows, _ := strconv.Atoi(r.FormValue("limit"))
	datagrid := showParamService(page, rows)
	b, _ := json.Marshal(datagrid)
	w.Header().Set("Content-type", "application/json;charset=utf-8")
	w.Write(b)

}
