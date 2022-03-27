package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/CobaltSato/Go-min/mysql_min"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Post is a struct for db
// TODO: domainで定義
type Post struct {
	gorm.Model
	Name    string
	Message string
}

// Server
type Server struct {
	DB mysql_min.DB
}

func (s *Server) NewDatabase() {
	s.DB = mysql_min.NewDatabase()
}

func main() {
	server := Server{}
	server.NewDatabase()

	db := server.DB.Get()
	db.AutoMigrate(&Post{}) // TODO: 運用ツールからの実行

	defer server.DB.Close()

	router := gin.Default() // TODO: dbを埋め込んだserverのルーティングをする
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context) {
		server.NewDatabase()
		var posts []Post

		db := server.DB.Get()
		db.Order("created_at asc").Find(&posts)
		defer server.DB.Close()

		ctx.HTML(200, "index.html", gin.H{
			"posts": posts,
		})
	})

	router.POST("/new", func(ctx *gin.Context) {
		server.NewDatabase()
		name := ctx.PostForm("name")
		message := ctx.PostForm("message")
		fmt.Println("create user " + name + " and message" + message)

		db := server.DB.Get()
		db.Create(&Post{Name: name, Message: message}) // TODO: interfaceで実行
		defer server.DB.Close()

		ctx.Redirect(302, "/")
	})

	router.POST("/delete/:id", func(ctx *gin.Context) {
		server.NewDatabase()
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("id is not a number")
		}
		var post Post
		db := server.DB.Get()
		db.First(&post, id) // TODO: interfaceで実行
		db.Delete(&post)    // TODO: interfaceで実行
		defer server.DB.Close()

		ctx.Redirect(302, "/")
	})

	if err := router.Run(); err != nil {
		log.Fatalf("server can't start :%v", err)
	}
}
