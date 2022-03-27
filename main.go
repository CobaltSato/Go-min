package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	post "github.com/CobaltSato/Go-min/domain"
	"github.com/CobaltSato/Go-min/mysql_min"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Server
type Server struct {
	DB mysql_min.DB
}

func (s *Server) NewDatabase() {
	s.DB = &mysql_min.Mysql{Database: NewDatabase()}
}

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	server := Server{}
	server.NewDatabase()

	db := server.DB.Get()
	db.AutoMigrate(&post.Post{}) // TODO: 運用ツールからの実行

	defer server.DB.Close()

	router := gin.Default() // TODO: dbを埋め込んだserverのルーティングをする
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context) {
		server.NewDatabase()
		var posts []post.Post

		postRepo := mysql_min.NewPost(server.DB.Get())
		postRepo.GetPosts(&posts)

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
		db.Create(&post.Post{Name: name, Message: message}) // TODO: interfaceで実行
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
		var post post.Post
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

func NewDatabase() (db *gorm.DB) {
	DBHOST := viper.GetString(`database.host`)
	USER := viper.GetString(`database.user`)
	PASS := viper.GetString(`database.pass`)
	PROTOCOL := viper.GetString(`database.protocol`)
	DBNAME := viper.GetString(`database.name`)

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	fmt.Println(CONNECT)

	count := 0
	db, err := gorm.Open(DBHOST, CONNECT)
	if err != nil {
		for {
			if err == nil {
				fmt.Println("")
				break
			}
			fmt.Print(".")
			time.Sleep(time.Second)
			count++
			if count > 180 {
				fmt.Println("")
				panic(err)
			}
			db, err = gorm.Open(DBHOST, CONNECT)
		}
	}

	return db
}
