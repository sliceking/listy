package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sliceking/listy/models"
)

func main() {
	db, err := gorm.Open("sqlite3", "./db/app.db")
	if err != nil {
		fmt.Errorf("opening db connection", err)
	}
	defer db.Close()

	db.AutoMigrate(&models.Card{}, &models.Task{})

	newCard := models.Card{
		Title:       "first",
		Description: "i am the",
		Tasks: []models.Task{
			{
				Description: "do it",
			},
			{
				Description: "what eva",
			},
		},
	}

	db.Create(&newCard)

	var card models.Card
	var tasks models.Task
	db.Model(&card).Related(&tasks)

	fmt.Println(tasks)

	// server := gin.Default()
	// server.LoadHTMLGlob("views/*")

	// // controllers can return funcs that are gin.HandlerFuncs: func(ctx *gin.Context) {}
	// server.GET("/index", func(ctx *gin.Context) {
	// 	ctx.HTML(http.StatusOK, "index.html", gin.H{
	// 		"message": "ok!",
	// 		"ok":      "no",
	// 	})
	// })

	// server.Run(":8080")
}
