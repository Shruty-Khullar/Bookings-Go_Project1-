package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Shruty-Khullar/bookings/pkg/config"
	"github.com/Shruty-Khullar/bookings/pkg/models"
)

//funcmap maps functions and allow us to use it in template
var functions = template.FuncMap{

}

var app *config.AppConfig 
//newTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig){
	app = a
}
//Rendering of Templates:A template contains the static parts of the desired HTML output as well as some special syntax describing how dynamic content will be inserted. ... Rendering means interpolating the template with context data and returning the resulting string.
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	//This creates a new template and parses the filename that is passed as input. 
	//In optimastion time we called Renderfunct using config only once else use this ;;;;; tc, err := RenderTemplate1()          //this will return us err and map of pointers of template
    var tc map[string]*template.Template
	//Used in dev mode if made changes it will again read from disk
	if app.UseCache {
		tc = app.TemplateCache
	} else {
          tc, _ = RenderTemplate1()
	}
	
	//ok tells us if ele we are searching for is present in map or not
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get Template Cache")
	}

	buff := new(bytes.Buffer)   //intialises zero size buffer of bytes
	 _ = t.Execute(buff,td)   //Executes template t and put result in buff
	 _ ,err := buff.WriteTo(w)
	 if err!=nil {
		 fmt.Println("Error writing to the browser", err)
	 }

	 /*parsedTemplate, _ := template.ParseFiles("./Templates/" + tmpl)
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing Template : ", err) 
		return
	} else {
		fmt.Println("No Error")
	}*/

}
//The filepath package provides functions to parse and construct file paths in a way that is portable between operating systems; dir/file on Linux vs. dir\file on Windows, for example.
//Glob return all the files matching the pattern so filepath.Glob return path of all the files matching the pattern
//filepath.Base(): Go language, path package used for paths separated by forwarding slashes, such as the paths in URLs. The filepath.Base() function in Go language used to return the last element of the specified path. Here the trailing path separators are removed before extracting the last element. If the path is empty, Base returns “.”. If the path consists entirely of separators, Base returns a single separator. Moreover, this function is defined under the path package. Here, you need to import the “path/filepath” package in order to use these functions
func RenderTemplate1() (map[string]*template.Template, error){
  myCache := map[string]*template.Template{}              //another way of declaring empty map. Here *templae. Template is a specialized Template from "text/template" that produces a safe HTML document fragment.
  //this will give us about and hoem.page.html path
  pages,err := filepath.Glob("./Templates/*.page.html")
  if err!=nil {
	  fmt.Println("Error parsing template",err)
	  return nil, err
  }
  //iterate through all the pages \. range pages will return index and each page iteratively. here it will give 0,about.page.html and 1,home.page.html
  for _,page := range pages {
      name := filepath.Base(page)
      fmt.Println("Current Page: ", page)
	  ts,err := template.New(name).Funcs(functions).ParseFiles(page)
	  if err!=nil {
            fmt.Println("error")
			return myCache,err
	  }
	  //see if this templates matches any layout.html layout template funct
	  matches,err:= filepath.Glob("./Templates/*.layout.html")
	  if err!=nil {
		fmt.Println("Error")
		return myCache,err
	  }
	  if len(matches)>0 {
		  ts,err = ts.ParseGlob("./Templates/*.layout.html")   //ParseGlob parses the template def in the files which is identified by the pattern and associates the resulting template with t
	
		  if err!=nil {
			  fmt.Println("Error")
			  return myCache,err
			}
		}
	    myCache[name] = ts
	}
	return myCache, nil
}
