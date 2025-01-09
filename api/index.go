package handler

import (
	"net/http"

	"github.com/kingtingthegreat/jting-resume/handlers"
)

func VercelHandler(w http.ResponseWriter, r *http.Request) {
	handlers.ResumeHandler(w, r)
}
