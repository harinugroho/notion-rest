package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harinugroho/notion"
	"github.com/harinugroho/notion/properties"
)

func main() {
	app := fiber.New()

	// web
	app.Static("/static", "./ui/dist/static")
	app.Get("/web/*", getView)

	// api
	app.Get("/api/v1/databases", getDatabase)
	app.Post("/api/v1/databases/query", getList)

	err := app.Listen("0.0.0.0:9000")
	if err != nil {
		panic(err)
	}
}

func getDatabase(c *fiber.Ctx) error {
	integrationToken := c.Query("integration_token")
	databaseUrl := c.Query("database_url")
	client, err := notion.NewClient(integrationToken).SetDatabaseIdByUrl(databaseUrl).GetDatabase()
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": client.GetObject(),
	})
}

type List struct {
	Filter properties.Filter `json:"filter"`
}

func getList(c *fiber.Ctx) error {
	integrationToken := c.Query("integration_token")
	databaseUrl := c.Query("database_url")
	client := notion.NewClient(integrationToken).SetDatabaseIdByUrl(databaseUrl)

	var list List
	if c.Get("Content-Type") == "application/json" {
		err := c.BodyParser(&list)
		if err != nil {
			return err
		}
		client = client.Filters(list.Filter)
	} else {
		list = List{}
	}
	result, err := client.GetList()
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":       result.GetObject(),
		"properties": result.GetProperties(),
	})
}

func getView(c *fiber.Ctx) error {
	return c.SendFile("./ui/dist/index.html")
}
