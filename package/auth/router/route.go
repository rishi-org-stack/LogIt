package router

import (
	log "logit/v1/util/log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Route(g *echo.Group, m ...echo.MiddlewareFunc) {
	grpAuth := g.Group("/auth")
	grpAuth.GET("/", ok)
}
func ok(c echo.Context) error {
	log.Init("AUTH", "HTTP").Warn()
	return c.String(http.StatusAccepted, "ok form auth")
}
