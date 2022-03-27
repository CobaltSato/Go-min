package mysql_min

import (
	"fmt"
	"time"

	post "github.com/CobaltSato/Go-min/domain"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type Mysql struct {
	database *gorm.DB
}

type DB interface {
	Get() *gorm.DB
	Close()
	GetPosts(posts *[]post.Post)
}

func (m Mysql) Get() *gorm.DB {
	return m.database
}

func (m Mysql) GetPosts(posts *[]post.Post) {
	m.database.Order("created_at asc").Find(posts)
}

func (m Mysql) Close() {
	m.database.Close()
}

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

//
func NewPost(db *gorm.DB) post.PostRepository {
	return &Mysql{
		database: db,
	}
}

func NewDatabase() (mysql *Mysql) {
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

	return &Mysql{database: db}
}
