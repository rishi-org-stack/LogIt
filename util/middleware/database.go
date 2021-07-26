package middleware

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func ConnectionMDB(client *mongo.Database) echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("mgClient", client)
			return hf(c)
		}
	}
}
