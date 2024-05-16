package main

import (
	"log"
	"net/http"
	"os"
)

func (cfg *appConfig) handlerScanNewDirectories(
	w http.ResponseWriter,
	r *http.Request,
) {

	log.Print("Scanning for new directories...")
	dirs, err := os.ReadDir(cfg.mediaRoot)
	if err != nil {
		log.Printf("Could not read directory %s! Exiting..", cfg.mediaRoot)
		return
	}

	for _, e := range dirs {
		if e.IsDir() {
			dirName := e.Name()
			_, ok := cfg.directoriesCache[dirName]
			if !ok {
				log.Print("Found new directory: ", dirName)
				dirEntry := makeDirEntry(dirName, cfg.mediaRoot)
				cfg.directoriesCache[dirName] = dirEntry
			}
		}
	}

	cfg.handlerHome(w, r)
}
