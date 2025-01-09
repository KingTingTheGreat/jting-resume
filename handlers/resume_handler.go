package handlers

import (
	"github.com/kingtingthegreat/jting-resume/data"
	"github.com/kingtingthegreat/jting-resume/render"
	"net/http"
)

func ResumeHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderPage(w, data.ResumeData)
}
