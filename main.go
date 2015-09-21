package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/spf13/viper"

	"sphire/configuration"
	"sphire/geofence"
)

type Page struct {
	Title string
	Body  template.HTML
}

// viper configuration pointer
var vpx *viper.Viper

func main() {
	fmt.Println("")

	// Get a pointer to our configuration
	vpx = configuration.Configuration("DEV")

	// Start listening for requests
	http.HandleFunc("/", router)
	http.ListenAndServe(vpx.Get("application.http.listen").(string), nil)
}

func router(writer http.ResponseWriter, request *http.Request) {
	// Set headers
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")

	switch request.URL.Path[len("/"):] {
	case "geofence":
		rt_geomap(writer, request)
		return
	}

	fmt.Fprintf(writer, "Sphire API")
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
