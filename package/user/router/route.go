package router

import (
	"context"
	"encoding/json"
	user "logit/v1/package/user"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Http struct {
	uSer user.Service
}
type Res struct {
	Data interface{}
	Msg  string
}

func Route(g *echo.Group, userService *user.UserService, m ...echo.MiddlewareFunc) {
	h := &Http{
		uSer: userService,
	}
	grpAuth := g.Group("/user", m...)
	grpAuth.GET("/:id", h.getById)
	grpAuth.PUT("/:id", h.updateById)
}
func (h *Http) getById(c echo.Context) error {
	ctx := context.WithValue(context.Background(), "mgClient", c.Get("mgClient"))
	id := c.Param("id")
	user, err := h.uSer.GetUser(ctx, id)

	if err != nil {

	}
	return json.NewEncoder(c.Response().Writer).Encode(user)
}

func (h *Http) updateById(c echo.Context) error {
	ctx := context.WithValue(context.Background(), "mgClient", c.Get("mgClient"))
	id := c.Param("id")
	Id, _ := primitive.ObjectIDFromHex(id)
	US := &user.User{
		ID: Id,
	}
	user, err := h.uSer.UpdateUser(ctx, US)

	if err != nil {

	}
	return json.NewEncoder(c.Response().Writer).Encode(user)
}
