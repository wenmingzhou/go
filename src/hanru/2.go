package main

import "fmt"

type Book struct {
	bookName string
	price    float64
}

type Student struct {
	name string
	age int
	books []*Book
}



func main() {
	b1 :=Book{"孙子兵法",129.1}
	b2 :=Book{"易筋经",69.9}
	b3 :=Book{"西游记",88.9}

	bb1 :=make([]*Book,0,10)
	bb1  =append(bb1,&b1,&b2,&b3)
	fmt.Println(bb1)
	s1 :=Student{"王二狗",19,bb1}
	fmt.Println(s1)
	for i:=0;i<len(s1.books);i++{
		p :=s1.books[i]
		fmt.Println(p.bookName)
	}
}
