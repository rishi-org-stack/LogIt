package api

import (
	"logit/v1/package/auth"
	amdb "logit/v1/package/auth/databases/mongo"
	authR "logit/v1/package/auth/router"
	userR "logit/v1/package/user/router"
	mid "logit/v1/util/middleware"
	"os"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type api struct {
	Client      *mongo.Database
	Version     string
	MiddleWares []echo.MiddlewareFunc
}

func Init(c *mongo.Database, m ...echo.MiddlewareFunc) *api {
	return &api{
		Client:      c,
		Version:     os.Getenv("VERSION"),
		MiddleWares: m,
	}
}
func (ap *api) Route(e *echo.Echo) {
	e.Use(mid.ConnectionMDB(ap.Client))

	v1 := e.Group("/api/" + ap.Version)

	// v1.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusAccepted, "Works well\n")
	// })
	authService := auth.Init(amdb.AuthDb{})
	authR.Route(authService, v1, mid.ConnectionMDB(ap.Client))
	userR.Route(v1, ap.MiddleWares...)
}
