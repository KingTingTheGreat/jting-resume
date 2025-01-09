package main

import (
	"fmt"
	"github.com/kingtingthegreat/jting-resume/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.Handler)
	http.HandleFunc("/styles.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "styles.css")
	})

	fmt.Println("http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("server error:", err)
	}
}
