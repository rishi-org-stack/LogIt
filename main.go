package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	godotenv.Load()
	http.HandleFunc("/", greet)
	http.ListenAndServe(os.Getenv("PORT"), nil)
}
