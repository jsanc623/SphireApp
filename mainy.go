package main

import (
	"fmt"
	"net/http"
	"sphire/geofence"
	"html/template"
	"strconv"
)

const (
	HTTP_LISTEN = "0.0.0.0:7227"
)

type Page struct {
    Title string
    Body  template.HTML
}

func main() {
	http.HandleFunc("/", router)
	http.ListenAndServe(HTTP_LISTEN, nil)
}

func router(writer http.ResponseWriter, request *http.Request) {
	// Set headers
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")

	route := request.URL.Path[len("/"):]
 	switch route {
    case "geofence":
        rt_geomap(writer, request)
		return
    case "home":
        fmt.Fprintf(writer, "home hommily")
		return
    }
    err404(writer)
}

func err404(writer http.ResponseWriter){
	writer.Header().Set("Content-Type", "text/plaintext; charset=utf-8")
	fmt.Fprintf(writer, "404 Not Found")
}

func rt_geomap(writer http.ResponseWriter, request *http.Request) {
	miles, _ := strconv.ParseFloat(request.URL.Query().Get("miles"), 64)
	var res string = geofence.BoundingBox(40.752087, -73.980190, miles)
    renderTemplate(writer, "map", loadPage("Map test", res))
}

func loadPage(title string, body string) *Page {
    return &Page{Title: title, Body: template.HTML(body)}
}

func renderTemplate(writer http.ResponseWriter, templ string, page *Page) {
    t, _ := template.ParseFiles("resources/views/" + templ + ".html")
    t.Execute(writer, page)
}
