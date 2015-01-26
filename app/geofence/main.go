package main
 
import (
        "fmt"
        "net/http"
        "strconv"
        "encoding/json"
)
 
const HTTP_LISTEN = "104.131.54.150:7227"
 
type (
        Bounds struct {
                Lat BoundsLatLng
                Lng BoundsLatLng
        }
 
        BoundsLatLng struct {
                Min float64
                Max float64
        }
)
 
func main() {
        http.HandleFunc("/", bound)
        http.ListenAndServe(HTTP_LISTEN, nil)
}
 
func bound(writer http.ResponseWriter, request *http.Request) {
        dirty_lat := request.FormValue("lat")
        dirty_lng := request.FormValue("lng")
        dirty_miles := request.FormValue("miles")
 
        // allow cross domain AJAX requests
        writer.Header().Set("Access-Control-Allow-Origin", "*")
        writer.Header().Set("Content-Type", "application/json; charset=utf-8")
 
        if dirty_lat != "" && dirty_lng != "" && dirty_miles != "" {
                lat, err := strconv.ParseFloat(dirty_lat, 32)
                lng, err := strconv.ParseFloat(dirty_lng, 32)
                miles, err := strconv.ParseFloat(dirty_miles, 32)
 
                if err != nil {
                        fmt.Fprint(writer, "error in strconv")
                }
 
                var lat_bnd float64 = 0.0144697 * miles
                var lng_bnd float64 = 0.0144812 * miles
 
                bounds := &Bounds{
                        Lat: BoundsLatLng{Min: lat - lat_bnd, Max: lat + lat_bnd},
                        Lng: BoundsLatLng{Min: lng - lng_bnd, Max: lng + lng_bnd},
                }
 
                output, err := json.Marshal(bounds)
 
                if err != nil {
                        fmt.Fprint(writer, "error in marshal")
                }
 
                fmt.Fprint(writer, string(output))
        }
}
 
