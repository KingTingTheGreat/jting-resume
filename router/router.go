package router

import (
	"net/http"

	"github.com/kingtingthegreat/jting-resume/handlers"
)

func Router() *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	router.HandleFunc("/", handlers.ResumeHandler)

	return router
}
