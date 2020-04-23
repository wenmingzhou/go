package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

//初始化ES，准备接受kafka那边发来的数据
var (
	client *elastic.Client
)

func Init(address string) (err error) {

	client, err = elastic.NewClient(elastic.SetURL("http://" + address))
	if err != nil {
		return err
	}
	fmt.Println("connect to es success")
	return
}

func SendToES(indexStr string, data interface{}) error {
	fmt.Println(indexStr, data)
	put1, err := client.Index().
		Index(indexStr).
		Type("xxx").
		BodyJson(data).
		Do(context.Background())
	if err != nil {
		// Handle error
		return err
	}

	fmt.Printf("Indexed student %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	return err
}
