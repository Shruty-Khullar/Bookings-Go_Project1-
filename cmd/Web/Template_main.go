package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Shruty-Khullar/bookings/pkg/config"
	"github.com/Shruty-Khullar/bookings/pkg/handlers"
	"github.com/Shruty-Khullar/bookings/pkg/render"
	"github.com/alexedwards/scs/v2"
)

var portnumber string = ":8080"
var session *scs.SessionManager //Automatic loading and saving of session data via middleware. it is used for global def . Now middleware.go can also use this
func main() {
	var app config.AppConfig
	app.InProduction = false
	tc, err := render.RenderTemplate1()
	if err != nil {
		fmt.Println("Error occured")
	}
	//A session is started once a user logs in, and expires some time after that. Each logged in user has some reference to the session, which they send with their requests. We then use this reference to look up the user that it belongs to and return information specific to them.
	//Sessions and cookies:Cookies are client-side files on a local computer that hold user information. Sessions are server-side files that contain user data. Cookies end on the lifetime set by the user. When the user quits the browser or logs out of the programmed, the session is over.
	//We need to intialise the session used
	session = scs.New()
	session.Lifetime = 24 * time.Hour              //session should live for 24hr max
	session.Cookie.Persist = true                  //cookie in client side should persist
	session.Cookie.SameSite = http.SameSiteLaxMode //by default settings
	session.Cookie.Secure = app.InProduction       //as we are in devlopment level and we are using localhost so we don't need a encrypted session
	app.Session = session
	app.TemplateCache = tc
	app.UseCache = false
	//sends template cache as received from RenderTemplate1() to NewRepo function leading to formation of new repository
	repo := handlers.NewRepo(&app)
	//It returns new repos and it is future put to new handlers
	handlers.NewHandlers(repo)
	//call newtemplate funct
	render.NewTemplates(&app)

	//handled using routes.go //http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)
	fmt.Println("Running folder Udemy_Course")
	//err = http.ListenAndServe(":8080", nil)
	srv := &http.Server{
		Addr:    portnumber,
		Handler: Routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	/*if err == nil {
		fmt.Println("no error")
	} else {
		fmt.Println("Error123")
	}*/

}
