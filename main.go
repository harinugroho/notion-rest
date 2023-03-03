package main

import (
	"NotionRest/notion"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/databases/:id", getList)
	r.GET("/databases/:id/info", getInfo)
	r.GET("/databases/:id/columns", getColumns)
	err := r.Run()
	if err != nil {
		panic(err)
	}
}

func getInfo(c *gin.Context) {
	id := c.Param("id")
	secret := c.Request.Header["Integration-Token"][0]
	database := notion.NewDatabase(secret, id)
	info := database.GetInfo()

	c.JSON(200, gin.H{
		"data": info,
	})
}

func getColumns(c *gin.Context) {
	id := c.Param("id")
	secret := c.Request.Header["Integration-Token"][0]
	database := notion.NewDatabase(secret, id)
	columns := database.GetColumns()

	c.JSON(200, gin.H{
		"data": columns,
	})
}

func getList(c *gin.Context) {
	id := c.Param("id")
	secret := c.Request.Header["Integration-Token"][0]
	list := notion.NewList(secret, id)
	data, err := list.QueryDatabase()
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{
		"data": data,
	})
}
