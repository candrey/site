package main

import (
	"database/sql"
	"net/http"
	//	"encoding/json"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type menu struct {
	items []menuItems
}

type menuItems struct {
	MainMenuItems menuItem
	SubMenuItems  []menuItem
}

type menuItem struct {
	ID        uint
	Serial    uint
	Name      string
	SubMenuID uint
	Href      string
	Weight    uint
	Enable    bool
}

var router *gin.Engine

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

func parsSubMenu(selMenu *sql.Stmt, primMenuID uint) []menuItem {
	var subMenuItem menuItem
	var subMenu []menuItem

	subMenuItems, err := selMenu.Query(primMenuID)
	checkErr(err)

	for subMenuItems.Next() {
		err = subMenuItems.Scan(&subMenuItem.ID, &subMenuItem.Serial, &subMenuItem.Name, &subMenuItem.SubMenuID, &subMenuItem.Href, &subMenuItem.Weight, &subMenuItem.Enable)
		subMenu = append(subMenu, subMenuItem)
	}

	return subMenu
}

func main() {

	db, err := sql.Open("sqlite3", "./site.db")
	checkErr(err)
	defer db.Close()

	var primMenuItem menuItem
	var mItems menuItems
	var m menu

	selMenu, err := db.Prepare("SELECT * FROM menu WHERE subMenuID = ? and enable != 0 ORDER BY weight")
	checkErr(err)
	defer selMenu.Close()

	primMenuItems, err := selMenu.Query(0)
	checkErr(err)

	for primMenuItems.Next() {
		err = primMenuItems.Scan(&primMenuItem.ID, &primMenuItem.Serial, &primMenuItem.Name, &primMenuItem.SubMenuID, &primMenuItem.Href, &primMenuItem.Weight, &primMenuItem.Enable)

		mItems.MainMenuItems = primMenuItem
		mItems.SubMenuItems = parsSubMenu(selMenu, primMenuItem.Serial)

		m.items = append(m.items, mItems)
	}

	router = gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	initializeRoutes(m.items)
	router.Run()

}
