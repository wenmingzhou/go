package control

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type Upload struct {
	Original string `json:"original"`
	State    string `json:"state"`
	Title    string `json:"title"`
	Url      string `json:"url"`
}

func ApiUpload(w http.ResponseWriter, r *http.Request) {
	//把上传的文件存储在内存和临时文件中
	r.ParseMultipartForm(1 << 20 * 20) //传入的参数是最大内存，单位是byte  20MB
	//获取文件句柄，然后对文件进行存储等处理
	f, h, err := r.FormFile("upfile")
	if err != nil {
		Fail(w, "上传失败", err.Error())
	}
	//log.Println(f, h, err)
	os.MkdirAll("static/upload/", 0666)
	name := "static/upload/" + h.Filename
	dst, _ := os.Create(name)
	io.Copy(dst, f)
	f.Close()
	dst.Close()
	w.Header().Set("Content-Type", "application/json")

	mod := Upload{
		Original: h.Filename,
		State:    "SUCCESS",
		Title:    h.Filename,
		Url:      name,
	}

	buf, _ := json.Marshal(mod)
	w.Write(buf)
}
