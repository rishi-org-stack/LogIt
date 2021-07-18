package api

import (
	authR "logit/v1/package/auth/router"
	"os"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type api struct {
	Client  *mongo.Client
	Version string
}

func Init(c *mongo.Client) *api {
	return &api{
		Client:  c,
		Version: os.Getenv("VERSION"),
	}
}
func (ap *api) Route(e *echo.Echo) {
	v1 := e.Group("/api/" + ap.Version)
	// v1.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusAccepted, "Works well\n")
	// })
	authR.Route(v1)
}
