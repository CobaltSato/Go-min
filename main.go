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
type Post struct {
	gorm.Model
	Name    string
	Message string
}

// Server
type Server struct {
	DB mysql_min.DB
}

func main() {
	mysql := mysql_min.NewDatabase()
	server := Server{DB: mysql}

	db := server.DB.Get()
	db.AutoMigrate(&Post{}) // TODO: 運用ツールからの実行

	defer db.Close()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context) {
		server.DB = mysql_min.NewDatabase()
		db := server.DB.Get()
		var posts []Post
		db.Order("created_at asc").Find(&posts)
		defer db.Close()

		ctx.HTML(200, "index.html", gin.H{
			"posts": posts,
		})
	})

	router.POST("/new", func(ctx *gin.Context) {
		server.DB = mysql_min.NewDatabase()
		db := server.DB.Get()
		name := ctx.PostForm("name")
		message := ctx.PostForm("message")
		fmt.Println("create user " + name + " and message" + message)
		db.Create(&Post{Name: name, Message: message})
		defer db.Close()

		ctx.Redirect(302, "/")
	})

	router.POST("/delete/:id", func(ctx *gin.Context) {
		server.DB = mysql_min.NewDatabase()
		db := server.DB.Get()
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("id is not a number")
		}
		var post Post
		db.First(&post, id)
		db.Delete(&post)
		defer db.Close()

		ctx.Redirect(302, "/")
	})

	if err := router.Run(); err != nil {
		log.Fatalf("server can't start :%v", err)
	}
}
