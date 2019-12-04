package item

import (
	"ego/src/commons"
	"ego/src/item/cat"
	"ego/src/item/desc"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"net/url"
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

//商品新增
func insertService(f url.Values) (e commons.EgoResult) {
	var t Tbitem
	cid, _ := strconv.Atoi(f["cid"][0])

	t.Cid = cid
	t.Title = f["title"][0]
	t.SellPoint = f["sell_point"][0]
	price, _ := strconv.Atoi(f["price"][0])
	t.Price = price
	num, _ := strconv.Atoi(f["num"][0])
	t.Num = num
	t.Image = f["image"][0]

	t.Status = 1
	date := time.Now().Format("2006-01-02 15:04:05")
	t.Created = date
	t.Updated = date
	id := commons.GenId()
	t.Id = id
	//商品表新增
	count := insertItemDao(t)
	if count > 0 {
		var tbItemDesc desc.TbItemDesc
		tbItemDesc.ItemId = id
		tbItemDesc.Created = date
		tbItemDesc.Updated = date
		tbItemDesc.ItemDesc = f["desc"][0]
		countDesc := desc.Insert(tbItemDesc)
		if countDesc > 0 {
			e.Status = 200
		} else {
			//删除商品中的数据
			delById(id)
		}
	}
	return
}
