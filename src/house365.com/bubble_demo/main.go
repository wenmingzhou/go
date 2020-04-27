package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

//Todo model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

var (
	DB *gorm.DB
)

func initMySQL() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = DB.DB().Ping()
	return

}
func main() {
	//连接数据库
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	defer DB.Close()
	DB.AutoMigrate(&Todo{})
	r := gin.Default()
	//告诉gin框架静态文件去哪里找
	r.Static("/static", "static")
	//告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	//v1
	v1Group := r.Group("v1")
	{
		//待办事项
		//添加
		v1Group.POST("/todo", func(c *gin.Context) {
			//1 从请求中把数据拿出来
			var todo Todo
			c.BindJSON(&todo)
			//2 存入数据库
			err = DB.Create(&todo).Error
			//3 返回响应体
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}

		})
		//查看
		v1Group.GET("/todo", func(c *gin.Context) {
			//查询todo这个表里面的所有数据
			var todoList []Todo
			err = DB.Find(&todoList).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todoList)
			}
		})
		//查看具体的某一个待办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		//修改某一个待办事项
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id, _ := c.Params.Get("id")
			var todo Todo
			if err = DB.Where("id=?", id).Find(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": "id不存在"})
				return
			} else {
				c.JSON(http.StatusOK, todo)
				return
			}
			c.BindJSON(&todo)
			if err = DB.Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error()})
				return
			} else {
				c.JSON(http.StatusOK, todo)
				return
			}
		})
		//删除某一个待办事项
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, _ := c.Params.Get("id")
			if err = DB.Where("id=?", id).Delete(Todo{}).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error()})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{
					"success": "1"})
				return
			}
		})
	}

	r.Run(":9090")
}
