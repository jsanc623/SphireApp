package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/spf13/viper"

	"sphire/configuration"
	"sphire/geofence"
	sfxlog "sphire/log"
)

type Page struct {
	Title string
	Body  template.HTML
}

// get viper configuration pointer
var vpx *viper.Viper = configuration.Configuration("DEV")

// logrus log manager init
var _ = sfxlog.Init("json", "DEV", "/tmp/sphire.log")

func main() {
	// Start listening for all requests on "/"
	http.HandleFunc("/", router)

	sfxlog.Log(nil, "main.go:main() Attempting to listen for requests", "info")
	http.ListenAndServe(vpx.Get("application.http.listen").(string), nil)
}

func router(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")

	// build the url
	url := request.URL.Scheme + request.URL.Opaque + request.URL.Host + request.URL.Path

	// Append the query if any exists
	if request.URL.RawQuery != "" {
		url += "?" + request.URL.RawQuery
	}

	// kill the favicon
	if url == "/favicon.ico" {
		kill_favicon(writer)
		return
	}

	path := request.URL.Path[len("/"):]
	sfxlog.Log(nil, "main.go:router() " + request.Method + " " + request.Proto + url, "info")

	switch path {
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

func kill_favicon(writer http.ResponseWriter) {
	fmt.Fprintf(writer, "{'Content-Type': 'image/x-icon'}")
}

func loadPage(title string, body string) *Page {
	return &Page{Title: title, Body: template.HTML(body)}
}

func renderTemplate(writer http.ResponseWriter, templ string, page *Page) {
	t, _ := template.ParseFiles("resources/views/" + templ + ".html")
	t.Execute(writer, page)
}
