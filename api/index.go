package handler

import (
	"net/http"

	"github.com/kingtingthegreat/jting-resume/router"
)

func VercelHandler(w http.ResponseWriter, r *http.Request) {
	router := router.Router()

	router.ServeHTTP(w, r)
}
