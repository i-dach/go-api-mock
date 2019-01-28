package api

import (
	//	"encoding/json"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Book is DB table structure.
// this tables into book infomation.
type Book struct {
	ID    string `id:"split(label,2).split(3,title).insertday(dd)"`
	Title string `title:"book title"`
	Label string `label:"category"`
}

/*********************************************
*                 DB Config                  *
*********************************************/

func dbConnect(engine string, dbn string, user string, pwd string) *sqlx.DB {
	db, err := sqlx.Connect(engine, user+":"+pwd+"@tcp(db:3306)/"+dbn+"?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}

	return db
}

/**********************************************
*                API METHOD                   *
**********************************************/

func getBookList(c *gin.Context) {
	db := dbConnect("mysql", "api", "api", "password")

	books := []Book{}
	db.Select(&books, "select * from book")

	c.String(200, "getBookList : %v", books)
}

func getBookByTitle(c *gin.Context) {
	books := []Book{}

	title := c.Param("title")
	m := map[string]interface{}{"title": "%" + title + "%"}

	db := dbConnect("mysql", "api", "api", "password")
	nstmt, err := db.PrepareNamed("select * from book where title like :title")
	err = nstmt.Select(&books, m)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"books": books,
	})
}

func updateBookInfo(c *gin.Context) {
	// http request param get
	id := c.Param("id")

	// update target info format is JSON
	var target Book
	if err := c.ShouldBind(&target); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	log.Println(target)

	// get target book info
	db := dbConnect("mysql", "api", "api", "password")

	// check target data existent
	var book Book
	if err := db.Get(&book, "select * from book where id=?", id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": err})
		return
	}

	log.Println(book)

	book.Label = target.Label

	// tranxaction start
	tx := db.MustBegin()
	tx.NamedExec("update book set label = :label where id = :id", book)
	tx.Commit()

	// Query the database, storing results in book
	db.Get(&book, "select * from book where id=?", id)

	log.Println(book)

	c.JSON(http.StatusOK, gin.H{"satus": "updateBookInfo"})

}

// addBookInfo is book table insert method.
// this func used URI (title) & http request body params.
func addBookInfo(c *gin.Context) {
	// post data get
	title := c.PostForm("title")
	label := c.PostForm("label")

	log.Println(title + ":" + label)

	// db insert format set
	d := time.Date(2000, 02, 01, 12, 30, 0, 0, time.UTC)
	day := strconv.Itoa(d.Day())

	book := &Book{
		string([]rune(label)[0:2]) + "." + string([]rune(title)[0:2]) + "." + day,
		title,
		label,
	}

	// db connection start
	db := dbConnect("mysql", "api", "api", "password")

	// tranxaction start
	tx := db.MustBegin()
	tx.NamedExec("insert into book (id, title, label) values (:id, :title, :label)", book)
	tx.Commit()

	// Query the database, storing results in []Book
	books := []Book{}
	db.Select(&books, "select * from book order by id desc")

	c.JSON(200, gin.H{
		"method": "POST",
		"list":   books,
	})
}

func delBookInfo(c *gin.Context) {
	// delete method data get
	id := c.Param("id")
	m := map[string]interface{}{"id": id}

	// db connection & transaction start
	db := dbConnect("mysql", "api", "api", "password")
	tx := db.MustBegin()

	// check target data existent
	row := db.QueryRow("select count(*) as cnt from book where id=?", id)
	var cnt int
	err := row.Scan(&cnt)

	log.Println(cnt)

	// If Target Book is not found then return code 403 and msg send
	if err != nil || cnt < 1 {
		c.JSON(403, gin.H{
			"msg": "Not Found ID",
		})
		tx.Rollback()
		return
	}

	// delete target book
	tx.NamedExec("delete from book where id = :id", m)
	tx.Commit()

	// Query the database, storing results in []Book
	books := []Book{}
	db.Select(&books, "select * from book order by id desc")

	c.JSON(200, gin.H{
		"method": "DELETE",
		"list":   books,
	})

}
