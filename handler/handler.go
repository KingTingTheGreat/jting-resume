package handler

import (
	"github.com/kingtingthegreat/jting-resume/data"
	"github.com/kingtingthegreat/jting-resume/render"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	render.RenderPage(w, data.ResumeData)
}
