package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/luigieai/go-study/pkg/config"
	"github.com/luigieai/go-study/pkg/models"
)

var appConfig *config.AppConfig

func NewTemplates(app *config.AppConfig) {
	appConfig = app
}

func addDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	// create a template cache
	var templateCache map[string]*template.Template
	if appConfig.UseCache {
		templateCache = appConfig.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	// get requested template from cache
	template, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)
	td = addDefaultData(td)
	_ = template.Execute(buf, td)

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all files that represents a page (*.page.tmpl) from templates folder
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		log.Println(err)
		return myCache, err
	}

	// get all files that represent a page
	for _, page := range pages {
		fileName := filepath.Base(page)
		templateCreated, err := template.New(fileName).ParseFiles(page)
		if err != nil {
			log.Println(err)
			return myCache, err
		}
		// let's check if we have a layout
		layouts, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			log.Println(err)
			return myCache, err
		}
		if len(layouts) > 0 {
			templateCreated, err = templateCreated.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				log.Println(err)
				return myCache, err
			}
		}
		myCache[fileName] = templateCreated
	}
	return myCache, nil
}
