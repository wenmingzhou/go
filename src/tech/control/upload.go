package control

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path"
	"time"
)

var all = "QWERTYUIOPASDFGHJKLZXCVBNM1234567890"

type editUpload struct {
	Original string `json:"original"`
	State    string `json:"state"`
	Title    string `json:"title"`
	Url      string `json:"url"`
}

//传指针8个字节     //非指针    len cap point (8+8+8)*4 96字节
func (er *editUpload) json() []byte {
	buf, _ := json.Marshal(er)
	return buf
}

//构建随机字符串
func RandStr(ln int) string {
	res := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < ln; i++ {
		res += string(all[rand.Intn(36)])
	}
	return res

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
	ext := path.Ext(h.Filename)
	name := "static/upload/" + RandStr(10) + ext
	dst, _ := os.Create(name)
	io.Copy(dst, f)
	f.Close()
	dst.Close()
	w.Header().Set("Content-Type", "application/json")

	mod := editUpload{
		Original: h.Filename,
		State:    "SUCCESS",
		Title:    h.Filename,
		Url:      name,
	}
	w.Write(mod.json())
	/*
		mod := editUpload{
			Original: h.Filename,
			State:    "SUCCESS",
			Title:    h.Filename,
			Url:      name,
		}
		buf, _ := json.Marshal(mod)
		w.Write(buf)
	*/
}
