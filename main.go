package main

import (
  "log"
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/a-h/templ"
  components "bradswhite.com/src"
)

func main() {
  app := fiber.New()

  app.Static("/dist", "./dist")
  
  app.Static("/public", "./public")
	
  component := components.HomePage()
  app.Get("/", adaptor.HTTPHandler(templ.Handler(component)))
  
  component = components.AboutPage()
	app.Get("/about", adaptor.HTTPHandler(templ.Handler(component)))

	component = components.WorkPage()
	app.Get("/work", adaptor.HTTPHandler(templ.Handler(component)))

	component = components.BlogPage()
	app.Get("/blog", adaptor.HTTPHandler(templ.Handler(component)))

	component = components.ContactPage()
	app.Get("/contact", adaptor.HTTPHandler(templ.Handler(component)))

	log.Fatal(app.Listen(":3000"))
}
