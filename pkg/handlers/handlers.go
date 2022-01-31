package handlers

import (
	"net/http"

	"github.com/Shruty-Khullar/bookings/pkg/config"
	"github.com/Shruty-Khullar/bookings/pkg/models"
	"github.com/Shruty-Khullar/bookings/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

//make new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//sets the repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) { //Allow to access  the Respository struct
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(),"remote_ip",remoteIp)
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringmap := make(map[string]string)
	stringmap["test"] = "Hello WOrld"
	remoteIp:= m.App.Session.GetString(r.Context(),"remote_ip")
	stringmap["remote_ip"]=remoteIp
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringmap,
	})
}
