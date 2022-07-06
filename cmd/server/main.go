package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/broothie/ghx/server"
)

func main() {
	srv, err := server.New()
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	fmt.Println("serving on", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), srv.Handler()); err != nil {
		fmt.Println("error", err)
	}
}
