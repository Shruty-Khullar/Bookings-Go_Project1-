package main
import (
	"net/http"

	"github.com/Shruty-Khullar/bookings/pkg/config"
	"github.com/Shruty-Khullar/bookings/pkg/handlers"

	//"github.com/bmizerany/pat"
	//"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//The name mux stands for “HTTP request multiplexer”. Like the standard http. ServeMux, mux. Router matches incoming requests against a list of registered routes and calls a handler for the route that matches the URL or other conditions.
//Pat knows what methods are allowed given a pattern and a URI. ... If the NotFound handler is set, then it is used whenever the pattern doesn't match the request path for the current method (and the Allow header is not altered).
//Basically we will route in separate file using various mux which will give us routes according to entered url and handlers
//We will use 2 packages for creating these mux : 1.pat 2.chi
//Middlewares: Middleware is a type of computer software that provides services to software applications beyond those available from the operating system. It can be described as "software glue"
//Chi include some useful set of middlewares
func Routes(app *config.AppConfig) http.Handler {
	/*mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))               //here mux is formed and it will ans the given url by calling the given handler funct
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	return mux*/

	mux := chi.NewRouter() //It adds its middleware
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home)) //here mux is formed and it will ans the given url by calling the given handler funct
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
  
	return mux

}
