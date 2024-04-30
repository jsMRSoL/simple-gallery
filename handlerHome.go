package main

import (
	"cmp"
	"math/rand"
	"net/http"
	"slices"
)

type Entry struct {
	DirName      string
	DirTarget    string
	HasThumbnail bool
	Thumbnail    string
}

type HomePageData struct {
	PageTitle string
	Entries   []Entry
}

func (cfg *appConfig) handlerHome(w http.ResponseWriter, _ *http.Request) {

	tpl := cfg.templates["landing_page.html"]
	entries := []Entry{}

	for dirName, dirEntry := range cfg.directoriesCache {
		items := dirEntry.Media
		random := rand.Intn(len(items))
		hasThumbnail := items[random].HasThumbnail
		thumbnail := items[random].Thumbnail
		entries = append(entries, Entry{
			DirName:      dirName,
			DirTarget:    dirEntry.DirTarget,
			HasThumbnail: hasThumbnail,
			Thumbnail:    thumbnail,
		})
	}
	slices.SortFunc(entries, func(a, b Entry) int {
		return cmp.Compare(a.DirName, b.DirName)
	})

	data := HomePageData{
		PageTitle: "Welcome to the gallery!",
		Entries:   entries,
	}

	tpl.Execute(w, data)
}
