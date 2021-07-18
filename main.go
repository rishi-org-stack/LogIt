package main

import (
	"context"
	"logit/v1/package/api"
	"logit/v1/util/config"
	"logit/v1/util/db"
	"logit/v1/util/server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	env := config.Init()
	client := db.Connect(context.Background(), env)
	s := server.Init(env)
	e := s.Start()
	ap := api.Init(client)
	ap.Route(e)
	e.Logger.Fatal(e.Start(s.Port))
}
