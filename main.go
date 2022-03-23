package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type Post struct {
  gorm.Model
  Name string
  Message string
}

func init() {
  viper.SetConfigFile("config.json")
    err := viper.ReadInConfig()
    if err != nil {
        panic(err)
     }
}


func main() {
  db := sqlConnect()
  db.AutoMigrate(&Post{})
  defer db.Close()

  router := gin.Default()
  router.LoadHTMLGlob("templates/*.html")

  router.GET("/", func(ctx *gin.Context){
    db := sqlConnect()
    var posts []Post
    db.Order("created_at asc").Find(&posts)
    defer db.Close()

    ctx.HTML(200, "index.html", gin.H{
      "posts": posts,
    })
  })

  router.POST("/new", func(ctx *gin.Context) {
    db := sqlConnect()
    name := ctx.PostForm("name")
    message := ctx.PostForm("message")
    fmt.Println("create user " + name + " and message" + message)
    db.Create(&Post{Name: name, Message: message})
    defer db.Close()

    ctx.Redirect(302, "/")
  })

  router.POST("/delete/:id", func(ctx *gin.Context) {
    db := sqlConnect()
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

  router.Run()
}

func sqlConnect() (database *gorm.DB) {
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
