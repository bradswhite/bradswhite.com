package main

import (
  "log"
  "io"
  "context"
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/a-h/templ"
  "github.com/russross/blackfriday/v2"
  components "bradswhite.com/src"
  fns "bradswhite.com/fns"
)

func Unsafe(html string) templ.Component {
  return templ.ComponentFunc(func (ctx context.Context, w io.Writer) (err error) {
    _, err = io.WriteString(w, html)
    return
  })
}

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

  posts, _ := fns.GetPosts()
  app.Get("/blog", adaptor.HTTPHandler(templ.Handler(components.BlogPage(posts))))
	
  app.Get("/post/:title", func (c *fiber.Ctx) error {
    front, body, err := fns.GetPost(c.Params("title") + ".md")
    if err != nil {
      return adaptor.HTTPHandler(templ.Handler(components.ErrorPage()))(c)
    }

    output := blackfriday.Run(body)

    return adaptor.HTTPHandler(templ.Handler(components.PostPage(front, Unsafe(string(output)))))(c)
  })

	component = components.ContactPage()
	app.Get("/contact", adaptor.HTTPHandler(templ.Handler(component)))

	component = components.ErrorPage()
	app.Get("/error", adaptor.HTTPHandler(templ.Handler(component)))

	log.Fatal(app.Listen(":3000"))
}
