package handlers

import (
	"net/http"

	"github.com/luigieai/go-study/pkg/config"
	"github.com/luigieai/go-study/pkg/models"
	"github.com/luigieai/go-study/pkg/render"
)

// Repository methods
var Repo *Repository

type Repository struct {
	AppConfig *config.AppConfig
}

func NewRepo(appConfig *config.AppConfig) *Repository {
	return &Repository{
		AppConfig: appConfig,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// Routers
func (repo *Repository) Home(writer http.ResponseWriter, request *http.Request) {
	remoteIP := request.RemoteAddr
	repo.AppConfig.Session.Put(request.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(writer, "home.page.tmpl", &models.TemplateData{})

}

func (repo *Repository) About(writer http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := repo.AppConfig.Session.GetString(request.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(writer, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
