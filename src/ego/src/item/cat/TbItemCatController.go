package cat

import (
	"ego/src/commons"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func ItemCatHandler() {
	fmt.Println("11")
	commons.Router.HandleFunc("/item/cat/show", ShowItemCatController)
	fmt.Println("222")
}

func ShowItemCatController(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		id = "0"
	}
	fmt.Println("222")
	idInt, _ := strconv.Atoi(id)
	t := ShowCateByPidService(idInt)
	b, _ := json.Marshal(t)
	w.Header().Set("Content-type", "application/json;charset=utf-8")
	w.Write(b)
}
