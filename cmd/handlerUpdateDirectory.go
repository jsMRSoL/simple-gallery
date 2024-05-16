package main

import (
	"log"
	"net/http"
)

func (cfg *appConfig) handlerRescanDirectory(
	w http.ResponseWriter,
	r *http.Request,
) {
	path := r.PathValue("dir")

	cfg.rescanDirectory(path)

	tpl := cfg.templates["directory_page.html"]
	media := cfg.directoriesCache[path].Media

	data := DirData{
		PageTitle: path,
		Media:     media,
	}

	tpl.Execute(w, data)
}

func (cfg *appConfig) rescanDirectory(dir string) {
	log.Printf("Rescanning for directory %s...", dir)
	cfg.directoriesCache[dir] = makeDirEntry(dir, cfg.mediaRoot)
}
