package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ry-animal/go-htmx/ent"
)

func addRoutes(app *echo.Echo, client *ent.Client) {
	// SEO
	app.GET("/robots.txt", robotsHandler)
	app.GET("/sitemap.xml", sitemapHandler)

	// Routes
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World 👋!")
	})

	app.GET("/users", userIndexHandler(client))
	app.POST("/users", userCreateHandler(client))
}
