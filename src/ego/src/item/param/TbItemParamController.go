package param

import (
	"ego/src/commons"
	"encoding/json"
	"net/http"
	"strconv"
)

func ParamHandler() {
	commons.Router.HandleFunc("/item/param/show", showParamController)
}

func showParamController(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.FormValue("page"))
	rows, _ := strconv.Atoi(r.FormValue("limit"))
	datagrid := showParamService(page, rows)
	b, _ := json.Marshal(datagrid)
	w.Header().Set("Content-type", "application/json;charset=utf-8")
	w.Write(b)

}
