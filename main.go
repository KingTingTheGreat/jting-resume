package main

import (
	"fmt"
	"net/http"

	"github.com/kingtingthegreat/jting-resume/router"
)

func main() {
	router := router.Router()

	fmt.Println("http://localhost:8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("server error:", err)
	}
}
