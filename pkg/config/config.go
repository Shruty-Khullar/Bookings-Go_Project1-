package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

//holds application config
type AppConfig struct {
	UseCache bool
	TemplateCache map[string]*template.Template
    InProduction bool //Used by all the files
	Session *scs.SessionManager  //as this will be mostly used in handlers which is in diff package so defining it in config is better
}