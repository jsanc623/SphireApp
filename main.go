package main
 
import (
    "fmt"
    "net/http"
)
 
const (
    HTTP_LISTEN = "0.0.0.0:7227"
)
 
func main(){
    http.HandleFunc("/", router)
    http.ListenAndServe(HTTP_LISTEN, nil)
}
 
func router(writer http.ResponseWriter, request *http.Request){
    route := request.URL.Path[len("/"):]
    fmt.Fprintf(writer, "Requested route: %s\n", route);
}
