package router

import (
	"context"
	"logit/v1/package/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Http struct {
	serv auth.Service
}

func Route(ser auth.Service, g *echo.Group, m ...echo.MiddlewareFunc) {
	h := &Http{
		serv: ser,
	}
	grpAuth := g.Group("/auth", m...)
	grpAuth.POST("/", h.ok)
}
func (h *Http) ok(c echo.Context) error {
	ctx := context.WithValue(context.Background(), "mgClient", c.Get("mgClient"))
	res, err := h.serv.HandleAuth(ctx)
	if err != nil {
		return c.JSON(http.StatusBadGateway, err.Error())
	}

	// env := &config.Env{
	// 	Algo:        "HS256",
	// 	Key:         "RishiStack!1709",
	// 	JWTDurtaion: 60,
	// }
	// jwtService, err := auth.Init(env)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, err)
	// }
	// tk, err := jwtService.GenrateToken("ok")
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, err)
	// }
	// log.Init("AUTH", "HTTP").Warn()
	return c.JSON(http.StatusAccepted, res)
}
