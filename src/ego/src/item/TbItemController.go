package item

import (
	"ego/src/commons"
	"encoding/json"
	"net/http"
	"strconv"
)

func ItemHandler()  {
	commons.Router.HandleFunc("/showItem",ShowItemController)
}
func ShowItemController(w http.ResponseWriter ,r *http.Request)  {
	page,_ :=strconv.Atoi(r.FormValue("page"))
	rows,_ :=strconv.Atoi(r.FormValue("limit"))

	datagrid :=showItemService(page,rows)
	b,_ :=json.Marshal(datagrid)
	w.Header().Set("Content-type","application/json;charset=utf-8")
	w.Write(b)
}
