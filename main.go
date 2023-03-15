package main

import (
	"github.com/gin-gonic/gin"
	"github.com/harinugroho/notion"
)

func main() {
	r := gin.Default()
	r.GET("/databases", getDatabase)
	r.POST("/databases/query", getList)
	err := r.Run()
	if err != nil {
		panic(err)
	}
}

func getDatabase(c *gin.Context) {
	integrationToken := c.Query("integration_token")
	databaseUrl := c.Query("database_url")
	client, err := notion.NewClient(integrationToken).SetDatabaseIdByUrl(databaseUrl).GetDatabase()
	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"data": client.GetObject(),
	})
}

func getList(c *gin.Context) {
	integrationToken := c.Query("integration_token")
	databaseUrl := c.Query("database_url")
	client, err := notion.NewClient(integrationToken).SetDatabaseIdByUrl(databaseUrl).GetList()
	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"data": client.GetObject(),
	})
}
