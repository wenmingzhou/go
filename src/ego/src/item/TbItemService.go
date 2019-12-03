package item

import (
	"ego/src/commons"
	"ego/src/item/cat"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
)

func showItemService(page, rows int) (e *commons.Datagrid) {

	ts := selByPageDao(page, rows)
	total := selTotalDao()
	if ts != nil {
		itemChildren := make([]TbItemChild, 0)
		for i := 0; i < len(ts); i++ {
			var itemChild TbItemChild
			itemChild.Id = ts[i].Id
			itemChild.Updated = ts[i].Updated
			itemChild.Created = ts[i].Created
			itemChild.Status = ts[i].Status
			itemChild.Barcode = ts[i].Barcode
			//itemChild.Cid = ts[i].Cid
			//itemChild.Image = ts[i].Image
			itemChild.Price = ts[i].Price
			itemChild.Num = ts[i].Num
			itemChild.SellPoint = ts[i].SellPoint
			itemChild.Title = ts[i].Title
			itemChild.Title = ts[i].Title
			itemChild.CategoryName = cat.ShowCatByIdService(ts[i].Cid).Name
			itemChildren = append(itemChildren, itemChild)
		}
		e = new(commons.Datagrid)
		e.Data = itemChildren
		e.Count = total
		return
	}
	return nil
}

//删除商品
func delByIdsService(ids string) (e commons.EgoResult) {
	count := updateStatusByIdsDao(strings.Split(ids, ","), 3)
	if count > 0 {
		e.Status = 200
	}
	return
}

//上架商品
func inStockByIdsService(ids string) (e commons.EgoResult) {
	count := updateStatusByIdsDao(strings.Split(ids, ","), 1)
	if count > 0 {
		e.Status = 200
	}
	return
}

//下架商品
func outStockByIdsService(ids string) (e commons.EgoResult) {
	count := updateStatusByIdsDao(strings.Split(ids, ","), 2)
	if count > 0 {
		e.Status = 200
	}
	return
}

func imageUploadService(f multipart.File, h *multipart.FileHeader) map[string]interface{} {
	m := make(map[string]interface{})
	b, err := ioutil.ReadAll(f)
	if err != nil {
		m["error"] = 1
		m["message"] = "上传失败,服务器错误"
		return m
	}
	//纳秒时间戳+随机数+扩展名
	rand.Seed(time.Now().UnixNano())
	fileName := "static/images/" + strconv.Itoa(int(time.Now().UnixNano())) + strconv.Itoa(rand.Intn(1000)) + h.Filename[strings.LastIndex(h.Filename, "."):]
	err = ioutil.WriteFile(fileName, b, 0777)
	if err != nil {
		m["error"] = 1
		m["message"] = "上传失败,服务器错误"
		return m
	}
	m["error"] = 0
	m["url"] = commons.CurrentPath + fileName
	return m
}
