package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"logit/v1/package/api"
	"logit/v1/util/auth"
	"logit/v1/util/config"
	"logit/v1/util/db"
	mid "logit/v1/util/middleware"
	"logit/v1/util/server"
)

func main() {
	godotenv.Load()
	env := config.Init()
	client := db.Connect(context.Background(), env)
	s := server.Init(env)
	e := s.Start()
	jwtService, err := auth.Init(env)
	handleError(err)
	ap := api.Init(client, mid.JwtAuth(jwtService))
	ap.Route(e)
	e.Logger.Fatal(e.Start(s.Port))
}

// type Err struct{}

// func (err Err) Error() string {
// 	return "ok"
// }
// func main() {
// 	var e Err
// 	handleError(e)
// }
func handleError(e error) {

	if e != nil {
		fmt.Println(e.Error())
	}

}
