package mysql_min

import (
	post "github.com/CobaltSato/Go-min/domain"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Mysql struct {
	Database *gorm.DB
}

type DB interface {
	Get() *gorm.DB
	Close()
	GetPosts(posts *[]post.Post)
}

func (m Mysql) Get() *gorm.DB {
	return m.Database
}

func (m Mysql) GetPosts(posts *[]post.Post) {
	m.Database.Order("created_at asc").Find(posts)
}

func (m Mysql) Close() {
	m.Database.Close()
}

func NewPost(db *gorm.DB) post.PostRepository {
	return &Mysql{
		Database: db,
	}
}
