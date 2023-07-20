package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/luigieai/go-study/pkg/config"
	"github.com/luigieai/go-study/pkg/handlers"
	"github.com/luigieai/go-study/pkg/render"
)

var appConfig config.AppConfig
var session *scs.SessionManager

func main() {

	appConfig.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = appConfig.InProduction
	appConfig.Session = session

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	appConfig.TemplateCache = templateCache
	repo := handlers.NewRepo(&appConfig)
	appConfig.UseCache = false

	handlers.NewHandlers(repo)
	render.NewTemplates(&appConfig)

	fmt.Println("started")
	srv := &http.Server{
		Addr:    ":8080",
		Handler: Routes(&appConfig),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
