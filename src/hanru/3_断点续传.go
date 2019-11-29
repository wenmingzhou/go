package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	srcFile  :="./src/hanru/1.wmv"
	descFile  :="./src/hanru/aa.wmv"
	tmpFile  :=descFile+"temp.txt"

	file1,_ :=os.Open(srcFile)
	file2,_ :=os.OpenFile(descFile,os.O_CREATE|os.O_WRONLY,os.ModePerm)
	file3,_ :=os.OpenFile(tmpFile,os.O_CREATE|os.O_RDWR,os.ModePerm)

	defer file1.Close()
	defer file2.Close()
	//1 读取临时文件的数据，根据seek
	file3.Seek(0,0)
	bs:=make([]byte,100,100)
	n1,err:=file3.Read(bs)
	countStr :=string(bs[:n1])
	count,_ :=strconv.ParseInt(countStr,10,64)
	fmt.Println(count)

	//2 设置读写的偏移量
	file1.Seek(count,0)
	file2.Seek(count,0)

	n2:=-1 //读入的数据量
	n3:=-1 //写入的数据量
	total :=int(count)  //读取的总量
	data:=make([]byte,1024,1024)

	for{
		//3 读取数据
		n2,err=file1.Read(data)
		if(err ==io.EOF){
			fmt.Println("文件复制完毕...")
			file3.Close()
			os.Remove(tmpFile)
			break
		}
		//讲数据些人目标文件
		n3,_=file2.Write(data[:n2])
		total +=n3
		//将复制总量，存储到临时文件中
		file3.Seek(0,0)
		file3.WriteString(strconv.Itoa(total))

		//假装断电
		/*if total>200000{
			panic("假装断电3。。。。。。。。。。。。")
		}
		
		 */
	}
}
