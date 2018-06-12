package main

import (
	"database/sql"
	"net/http"
	//	"encoding/json"

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

/*
type menuItem struct {
	mainMenuItems menuItems
	subMenuItems  []menuItems
}
*/

type menuItems struct {
	ID        uint
	Serial    uint
	Name      string
	SubMenuID uint
	Weight    uint
	Enable    bool
}

var router *gin.Engine

//var mjson = `{"name": "Прогулки по Москве", "subMenu": ["от Марксисткой", "от Третьяковской", "от Тверской"], "name": "Обзор книг", "subMenu":[], "name": "Handmade", "subMenu":[]}`

func initializeRoutes(mItems []menuItems) {

	router.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK, "index.html", gin.H{"title": "Home Page", "menuItems": mItems},
		)
	})
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	db, err := sql.Open("sqlite3", "./site.db")
	checkErr(err)
	defer db.Close()

	var mis menuItems
	//var mi menuItem
	var m menu

	selMenu, err := db.Query("SELECT * FROM menu WHERE enable != 0 ORDER BY weight")
	checkErr(err)
	//defer selMenu.Close()
	//var subMenuItems
	//mMenuItems, err := selMenu.Query(0)
	//checkErr(err)

	for selMenu.Next() {
		err = selMenu.Scan(&mis.ID, &mis.Serial, &mis.Name, &mis.SubMenuID, &mis.Weight, &mis.Enable)
		//fmt.Println(mi.Name, mi.enable)

		//		subMenuItems, err := selMenu.Query(&mis.ID)
		//		checkErr(err)

		//		for subMenuItems.Next() {
		//			err = subMenuItems.Scan(&mis.ID, &mis.Serial, &mis.Name, &mis.SubMenuID, &mis.weight, &mis.enable)
		//			//fmt.Println(mi.Name, mi.enable)
		//			mi.subMenu = append(mi.subMenu, mis)
		//		}

		m.items = append(m.items, mis)
	}
	//fmt.Println(m)

	router = gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	initializeRoutes(m.items)
	router.Run()

}
