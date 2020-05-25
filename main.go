package main

import (
	"fmt"

	"github.com/sliceking/listy/controllers"
	"github.com/sliceking/listy/database"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	db, err := gorm.Open("sqlite3", "./database/app.db")
	if err != nil {
		fmt.Errorf("opening db connection", err)
	}
	defer db.Close()

	database.Seeder(db)

	cardController := controllers.NewCardsController(db)

	server := gin.Default()
	server.LoadHTMLGlob("views/*")
	server.Static("/assets", "./assets")

	// controllers can return funcs that are gin.HandlerFuncs: func(ctx *gin.Context) {}
	server.GET("/index", cardController.Index())

	server.Run(":8080")
}
