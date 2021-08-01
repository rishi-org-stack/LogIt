package api

import (
	"logit/v1/package/auth"
	amdb "logit/v1/package/auth/databases/mgdb"
	authR "logit/v1/package/auth/router"
	"logit/v1/package/user"
	umdb "logit/v1/package/user/databases/mgdb"
	userR "logit/v1/package/user/router"
	jAuth "logit/v1/util/auth"
	mid "logit/v1/util/middleware"
	"os"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type api struct {
	Client      *mongo.Database
	Version     string
	MiddleWares []echo.MiddlewareFunc
	Jwt         *jAuth.Auth
}

func Init(c *mongo.Database, jwt *jAuth.Auth, m ...echo.MiddlewareFunc) *api {
	return &api{
		Client:      c,
		Version:     os.Getenv("VERSION"),
		MiddleWares: m,
		Jwt:         jwt,
	}
}
func (ap *api) Route(e *echo.Echo) {
	e.Use(mid.ConnectionMDB(ap.Client))

	v1 := e.Group("/api/" + ap.Version)

	// v1.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusAccepted, "Works well\n")
	// })
	authService := auth.Init(amdb.AuthDb{}, ap.Jwt)
	userService := user.Init(&umdb.UserDb{},authService)
	authR.Route(authService, v1, mid.ConnectionMDB(ap.Client))
	userR.Route(v1, userService, ap.MiddleWares...)
}
