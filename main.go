package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pewe21/newbelajar/config"
	"github.com/pewe21/newbelajar/database"
	"github.com/pewe21/newbelajar/product"
	"log"
)

func main() {
	conf := config.InitializedLoader()
	db := database.InitDB(conf.Database, false)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "http://localhost:5173",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	repositoryProduct := product.NewRepository(db)
	serviceProduct := product.NewService(repositoryProduct)
	handlerProduct := product.NewHandler(serviceProduct)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "hello puah",
		})
	})

	app.Get("/product", handlerProduct.Get)
	app.Get("/product/:id", handlerProduct.GetById)
	app.Post("/product", handlerProduct.Create)
	app.Put("/product/:id", handlerProduct.Update)
	app.Delete("/product/:id", handlerProduct.Delete)
	err := app.Listen(conf.Server.Host + ":" + conf.Server.Port)

	log.Println(err)
}
