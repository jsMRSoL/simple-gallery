package main

import (
	"html/template"
	"log"
	"net/http"
)

type appConfig struct {
	mediaRoot        string
	templates        map[string]*template.Template
	directoriesCache map[string]DirEntry
}

func main() {

	dbg, addr, port, mediaRoot := parseFlags()

	if dbg == true {
		log.Println("In debug mode.................")
	}

	cfg := appConfig{
		mediaRoot:        mediaRoot,
		templates:        cacheTemplates(),
		directoriesCache: cacheDirectories(mediaRoot),
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", cfg.handlerHome)
	mux.HandleFunc("GET /directories/{dir}", cfg.handlerDirectory)
	mux.HandleFunc("GET /directories/update", cfg.handlerScanNewDirectories)
	mux.HandleFunc("GET /directories/update/{dir}", cfg.handlerRescanDirectory)

	// static files for js and frontend ux resources
	mux.Handle(
		"/static/",
		http.StripPrefix("/static/", http.FileServer(
			http.Dir("static/"))))
	// assets for image directories
	mux.Handle("/assets/", http.StripPrefix("/assets/",
		http.FileServer(http.Dir(mediaRoot))))

	corsMux := middlewareCors(mux)

	var srv http.Server
	srv.Handler = corsMux
	srv.Addr = addr + ":" + port

	log.Printf(
		"Serving files from %s from %s on port: %s\n",
		mediaRoot,
		addr,
		port,
	)
	log.Fatal(srv.ListenAndServe())
}
