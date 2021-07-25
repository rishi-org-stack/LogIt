package router

import (
	"logit/v1/util/auth"
	"logit/v1/util/config"
	log "logit/v1/util/log"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type test struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
}

func Route(g *echo.Group, m ...echo.MiddlewareFunc) {
	grpAuth := g.Group("/auth", m...)
	grpAuth.GET("/", ok)
}
func ok(c echo.Context) error {
	env := &config.Env{
		Algo:        "HS256",
		Key:         "RishiStack!1709",
		JWTDurtaion: 60,
	}
	jwtService, err := auth.Init(env)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	tk, err := jwtService.GenrateToken("ok")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	log.Init("AUTH", "HTTP").Warn()
	return c.JSON(http.StatusAccepted, tk)
}
