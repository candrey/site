package main

import (
	"database/sql"
	"fmt"
	"net/http"
	//	"encoding/json"

	//	"net/http"

	//	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

/*
//Site is a...
type Pages struct {
	Pages []Page `json:"pages"`
}

//Page is a...
type Page struct {
	ID          uint     `json:"id"`
	Title       string   `json:"title"`
	SubMenuID   uint     `json:"subMenuID"`
	Urls        []string `json:"urls"`
	Description string   `json:"description"`
}
*/
type menu struct {
	items []menuItems
}

type menuItems struct {
	ID        uint
	Serial    uint
	Name      string
	SubMenuID uint
	weight    uint
	enable    bool
}

var router *gin.Engine

//var mjson = `{"name": "Прогулки по Москве", "subMenu": ["от Марксисткой", "от Третьяковской", "от Тверской"], "name": "Обзор книг", "subMenu":[], "name": "Handmade", "subMenu":[]}`

func initializeRoutes() {

	router.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK, "index.html", gin.H{"title": "Home Page"},
		)
	})
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	/*
		// Open our jsonFile
		jsonFile, err := os.Open("site.json")
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Successfully Opened site.json")
		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)
	*/
	db, err := sql.Open("sqlite3", "./site.db")
	checkErr(err)
	defer db.Close()

	var mi menuItems
	var m menu
	rows, err := db.Query("SELECT * FROM menu WHERE subMenuID = 0 and enable != 0")
	checkErr(err)
	for rows.Next() {
		err = rows.Scan(&mi.ID, &mi.Serial, &mi.Name, &mi.SubMenuID, &mi.weight, &mi.enable)
		fmt.Println(mi.Name, mi.enable)
		m.items = append(m.items, mi)
	}
	for i := 0; i < len(m.items); i++ {

		res, err := sq.Exec(i)
		checkErr(err)
		fmt.Println(res)
	}

	/*
		var s Pages
		json.Unmarshal(byteValue, &s)
		fmt.Println(s.Pages[1].Title)
		router = gin.Default()
		router.Static("/static", "./static")
		router.LoadHTMLGlob("templates/*")

		initializeRoutes()
		router.Run()*/
}
