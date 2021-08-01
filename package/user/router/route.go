package router

import (
	"context"
	"encoding/json"
	user "logit/v1/package/user"
	"net/http"

	"github.com/labstack/echo/v4"
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
	grpAuth.PUT("/", h.updateById)
}
func (h *Http) getById(c echo.Context) error {
	ctx := context.WithValue(context.Background(), "mgClient", c.Get("mgClient"))
	id := c.Param("id")
	user, err := h.uSer.GetUser(ctx, id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusAccepted, user)
	// return json.NewEncoder(c.Response().Writer).Encode(user)
}

func (h *Http) updateById(c echo.Context) error {
	ctx := context.WithValue(context.Background(), "mgClient", c.Get("mgClient"))
	US := &user.User{}
	if err := c.Bind(US); err != nil {
		return handleError(err, c)
	}
	user, err := h.uSer.UpdateUser(ctx, US)

	if err != nil {
		return handleError(err, c)
	}
	return json.NewEncoder(c.Response().Writer).Encode(user)
}

func handleError(e error, c echo.Context) error {
	return c.JSON(http.StatusInternalServerError, map[string]string{"we have an error:->": e.Error()})
}
