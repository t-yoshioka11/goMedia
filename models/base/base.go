package base

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//　article Struct
type Article struct {
	Id           int
	Title        string
	Body         string
	CreateUserId int
	CreatedAt    time.Time `gorm:"column:created_at"`
}

// t_user Struct
type User struct {
	UserId    int
	Email     string
	Password  string
	CreatedAt time.Time `gorm:"column:created_at"`
}

// t_sesstions
type Sessions struct {
	Id        int
	Email     string
	UserId    int
	CreatedAt time.Time `gorm:"column:created_at"`
}

//var DBConnection *sql.DB

// 初期処理
func init() {
	db, err := gormConnect()
	if err != nil {
		log.Fatal("DBconnection error", err.Error())
	}
	defer db.Close()

	// t_article生成
	// cmd := `CREATE TABLE IF NOT EXISTS t_article (
	// 	id int NOT NULL AUTO_INCREMENT,
	// 	title varchar(100) NOT NULL,
	// 	body_of_letter varchar(10000) NOT NULL,
	// 	add_timestamp timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	// 	PRIMARY KEY (id))`

	// _, err = DBConnection.Exec(cmd)

	// TODO maiglationツールに置き換え
	//db.CreateTable(&T_article{})
	//db.CreateTable(&T_User{})
}

// gormConnect DB接続
func gormConnect() (*gorm.DB, error) {
	dbConnection, err := gorm.Open("mysql", "root:@/GO_MEDIA_DATABASE?parseTime=true")
	if err != nil {
		log.Println("DBconnection error", err.Error())
		return nil, err
	}
	dbConnection.LogMode(true)
	return dbConnection, nil
}
