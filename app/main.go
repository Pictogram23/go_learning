package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

type MessageBoard struct {
	gorm.Model
	Username string
	Content  string
}

func gormConnect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:password@tcp(mysql:3306)/bbs_db?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func dbInit() {
	db := gormConnect()
	defer db.Close()
	db.AutoMigrate(&MessageBoard{})
}

func dbInsert(username string, content string) {
	db := gormConnect()
	defer db.Close()
	db.Create(&MessageBoard{Username: username, Content: content})
}

func dbGetAll() []MessageBoard {
	db := gormConnect()
	defer db.Close()
	var messageBoards []MessageBoard
	db.Find(&messageBoards)
	return messageBoards
}

func main() {
	log.Print("server - [URL] http://localhost:8080/")

	dbInit()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", func(c *gin.Context) {
		dbInsert("イッチ", "そうです私が百科一之進です")
		//dbInsert("吸う", "ネズミの吸うで吸う")
		messageBoards := dbGetAll()
		c.HTML(http.StatusOK, "index.html", gin.H{
			"messageBoards": messageBoards,
		})

	})
	r.Run(":8080")
}
