package main

import (
	"context"
	"fmt"

	"logit/v1/package/api"
	"logit/v1/util/auth"
	"logit/v1/util/config"
	"logit/v1/util/db"
	mid "logit/v1/util/middleware"
	"logit/v1/util/server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	env := config.Init()
	client := db.Connect(context.Background(), env)
	s := server.Init(env)
	e := s.Start()
	jwtService, err := auth.Init(env)
	handleError(err)
	ap := api.Init(client, jwtService, mid.JwtAuth(jwtService))
	ap.Route(e)
	e.Logger.Fatal(e.Start(s.Port))
}

func handleError(e error) {

	if e != nil {
		fmt.Println(e.Error())
	}

}
