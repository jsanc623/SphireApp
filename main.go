package main
 
import (
    "fmt"
    "net/http"
)
 
const (
    HTTP_PORT = "7227"
    PORT_SEPARATOR = ":"
    HTTP_IP = "63.217.249.27"
)
 
func main(){
    http.HandleFunc("/", getHandler)
    http.ListenAndServe(HTTP_IP + PORT_SEPARATOR + HTTP_PORT, nil)
}
 
func getHandler(writer http.ResponseWriter, request *http.Request){
    strFromUrl := request.URL.Path[len("/"):]
    fmt.Fprintf(writer, "String: %s\n", strFromUrl);
}
