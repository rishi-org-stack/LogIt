package router

import (
	"fmt"
	log "logit/v1/util/log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Route(g *echo.Group, m ...echo.MiddlewareFunc) {
	grpAuth := g.Group("/user", m...)
	grpAuth.GET("/", ok)
}
func ok(c echo.Context) error {
	fmt.Println(c.Get("id"))
	log.Init("AUTH", "HTTP").Warn()
	return c.String(http.StatusAccepted, "ok form auth")
}
