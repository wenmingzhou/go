package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"address"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	//0 参数的校验
	//0.1传进来的data参数必须是指针类型(因为需要在函数中对其赋值)
	t := reflect.TypeOf(data)
	//fmt.Println(t,t.Kind())
	if t.Kind() != reflect.Ptr {
		err = errors.New("data should be a pointer")
		return
	}
	//0.2传进来的data参数必须是结构体类型指针(因为配置文件中各种键值对需要赋值给结构体的字段)
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data should be a struct")
		return
	}
	//1 读文件得到自己类型
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("read file failed,err:%v\n", err)
		return
	}
	lineSlice := strings.Split(string(b), "\r\n")
	//fmt.Printf("%#v\n",lineSlice)
	//2 一行一行的读
	var structName string
	for idx, line := range lineSlice {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		//  2.1  如果是注释就跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		//  2.2  如果是[开头的就表示是节
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			//把这行收尾的[]去掉，渠道中间的内容把空格去掉，拿到内容
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			//根据字符串sectionName去data里面根据反射找到对应的结构体
			//v :=reflect.ValueOf(data)
			//fmt.Println(t.NumField());return

			for i := 0; i < t.Elem().NumField(); i++ { //Config字段个数 2个
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					//说明找到了对应的嵌套结构体，把字段名记下来
					structName = field.Name
					fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, structName)
				}
			}

		} else {
			//  2.3  如果不是[开通的就是=分割的键值对
			//    1.以等号分割这一行,等号左边是key，等号右边是value
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			//    2.根据structName 去data里面把对应的嵌套结构体给取出来
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName) // 拿到嵌套结构体的值信息
			sType := sValue.Type()                     //拿到嵌套结构体的类型信息

			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("data中的%s字段应该是一个结构体", structName)
				return
			}
			var fieldName string
			var fieldType reflect.StructField
			//    3.遍历嵌套结构体的每一个字段，判断tag是不是等于key
			for i := 0; i < sValue.NumField(); i++ {
				field := sType.Field(i) //tag信息存储在类型信息中
				fieldType = field
				if field.Tag.Get("ini") == key {
					//找到对应的字段
					fieldName = field.Name
					break
				}
			}
			//    4.如果key =tag，给这个字段赋值
			//    4.1根据filedName 取出这个字段
			fileObj := sValue.FieldByName(fieldName)
			//    4.2对其赋值
			switch fieldType.Type.Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line:%d syntax error", idx+1)
					return
				}
				fileObj.SetInt(valueInt)
			}
		}
	}

	return
}
func main() {
	var cfg Config

	err := loadIni("D:/go/src/house365.com/studyGo/06day/06ini_demo/conf.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini failed,err:%v\n", err)
		return
	}
}
