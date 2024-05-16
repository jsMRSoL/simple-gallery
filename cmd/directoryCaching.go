package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

type DirEntry struct {
	DirName   string
	DirTarget string
	Media     []Item
}

type Item struct {
	Filename        string
	FileRelativeUrl string
	IsImage         bool
	HasThumbnail    bool
	Thumbnail       string
}

func cacheDirectories(mediaRoot string) map[string]DirEntry {

	_, err := os.Stat(mediaRoot)
	if err != nil {
		log.Fatalf("Could not access directory %s! Exiting..", mediaRoot)
	}

	dirs, err := os.ReadDir(mediaRoot)
	if err != nil {
		log.Fatalf("Could not read directory %s! Exiting..", mediaRoot)
	}

	cd := make(map[string]DirEntry)
	for _, e := range dirs {
		if e.IsDir() {
			dirName := e.Name()
			dirEntry := makeDirEntry(dirName, mediaRoot)
			cd[dirName] = dirEntry
		}
	}
	return cd
}

func makeDirEntry(dirName, mediaRoot string) DirEntry {
	dirTarget := "/directories/" + dirName
	media := []Item{}
	items, err := os.ReadDir(mediaRoot + dirName)
	if err != nil {
		log.Fatalf(
			"Could not access directory %s! Exiting..",
			mediaRoot+dirName,
		)
	}
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		filename := item.Name()
		fileExt := strings.ToLower(filepath.Ext(filename))
		if fileExt == ".zip" {
			continue
		}
		hasThumbnail := true
		thumbnail := filename + "-thumb.jpg"
		_, err := os.Stat(mediaRoot + dirName + "/thumbs/" + thumbnail)
		if err != nil {
			hasThumbnail = false
		}
		media = append(media, Item{
			Filename:        filename,
			FileRelativeUrl: "/assets/" + dirName + "/" + filename,
			IsImage:         isImage(fileExt),
			HasThumbnail:    hasThumbnail,
			Thumbnail:       "/assets/" + dirName + "/thumbs/" + thumbnail,
		})
	}
	return DirEntry{
		DirName:   dirName,
		DirTarget: dirTarget,
		Media:     media,
	}
}

func isImage(fileExt string) bool {
	imageExts := []string{
		".jpg",
		".jpeg",
		".gif",
		".png",
		".webp",
	}
	for _, imageExt := range imageExts {
		if fileExt == imageExt {
			return true
		}
	}

	return false
}
