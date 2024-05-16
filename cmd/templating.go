package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

func cacheTemplates() map[string]*template.Template {
	templatesDirectory := "./templates/"
	_, err := os.Stat(templatesDirectory)
	if err != nil {
		log.Fatal("Could not access templates directory! Exiting..")
	}
	files, err := os.ReadDir(templatesDirectory)
	if err != nil {
		log.Fatal(
			"Could not read in templates from templates directory! Exiting..",
		)
	}

	templates := make(map[string]*template.Template)
	for _, file := range files {
		name := file.Name()
		if strings.HasSuffix(name, ".html") {
			tmpl, err := template.ParseFiles(templatesDirectory + name)
			if err != nil {
				log.Fatalf("Could not parse template %s! Exiting..", name)
			}
			templates[name] = tmpl
		}
	}

	return templates

}
