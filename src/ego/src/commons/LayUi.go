package commons

type Datagrid struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Count int         `json:"count"`
}
