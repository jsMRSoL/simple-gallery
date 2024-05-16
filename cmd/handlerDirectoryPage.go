package main

import (
	"net/http"
)

type DirData struct {
	PageTitle string
	Media     []Item
}

func (cfg *appConfig) handlerDirectory(
	w http.ResponseWriter,
	r *http.Request,
) {

	path := r.PathValue("dir")
	tpl := cfg.templates["directory_page.html"]
	media := cfg.directoriesCache[path].Media

	data := DirData{
		PageTitle: path,
		Media:     media,
	}

	tpl.Execute(w, data)
}
