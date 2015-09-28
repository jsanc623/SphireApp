package main

import (
	"fmt"
	"flag"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/spf13/viper"

	"sphire/configuration"
	"sphire/geofence"
	sfxlog "sphire/log"
)

type Error struct {
	Error string `json:error`
	Values map[string]string `json:values`
}

// viper configuration pointer
var vpx *viper.Viper

// define a generic fatal error
var fatal_error = "{\"error\": \"something went wrong.\"}"

// main handles initial setup and configuration of the router, configuration, etc
func main() {
	env := flag.String("env", "DEV", "Environment (DEV, STG, PRD)")
	vpx = configuration.Configuration(*env)

	// logrus log manager init
	sfxlog.Init(vpx.Get("application.log.type").(string), vpx.Get("environment").(string), vpx.Get("application.log.file").(string))

	// Start listening for all requests on "/"
	http.HandleFunc("/", router)

	sfxlog.Log(nil, "main.go:main() Attempting to listen for requests", "info")
	http.ListenAndServe(vpx.Get("application.http.listen").(string), nil)
}

// router handles all routing of requests to appropriate handler functions
// writer http.ResponseWriter
// request *http.Request
func router(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	//writer.Header().Set("Content-Type", "text/html; charset=utf-8")
  	writer.Header().Set("Content-Type", "application/json")

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
	sfxlog.Log(nil, "main.go:router() " + request.Method + " " + request.Proto + " " + url, "info")

	switch path {
	case "geofence":
		rt_geomap(writer, request)
		return
	}

	fmt.Fprintf(writer, "Sphire API")
}

// rt_geomap handles all requests for geofence calculations
// writer http.ResponseWriter
// request *http.Request
// request.URL.Query()[miles, latitude, longitude]
func rt_geomap(writer http.ResponseWriter, request *http.Request) {
	miles, _ := strconv.ParseFloat(request.URL.Query().Get("miles"), 64)
	latitude, _ := strconv.ParseFloat(request.URL.Query().Get("latitude"), 64)
	longitude, _ := strconv.ParseFloat(request.URL.Query().Get("longitude"), 64)

	if latitude == 0 || longitude == 0 {
		error := Error{Error: "Latitude/longitude is required.", Values: map[string]string{"latitude": fmt.Sprint(latitude), "longitude": fmt.Sprint(longitude)}}
		jsonval, err := json.Marshal(error)

		if err != nil {
			sfxlog.Log(nil, "main.go:rt_geomap() " + err.Error(), "error")
			fmt.Fprintf(writer, fatal_error)
			return
		}

		sfxlog.Log(nil, "main.go:rt_geomap() " + error.Error, "error")
		fmt.Fprintf(writer, string(jsonval))
		return
	}

	var res string = geofence.BoundingBox(latitude, longitude, miles)
	fmt.Fprintf(writer, res)
}

// kill_favicon kills requests for favicon
// writer http.ResponseWriter
func kill_favicon(writer http.ResponseWriter) {
	fmt.Fprintf(writer, "{'Content-Type': 'image/x-icon'}")
}
