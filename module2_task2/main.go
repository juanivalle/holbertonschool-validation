package main

import (
  "fmt"
  "io"
  "log"
  "net/http"
  "os"

  "github.com/gorilla/mux"
)

func main() {
  httpAddr := "0.0.0.0:9999"
  if port := os.Getenv("PORT"); port != "" {
    httpAddr = "0.0.0.0:" + port
  }
  fmt.Println("HTTP Server listening on", httpAddr)


  log.Fatal(http.ListenAndServe(httpAddr, setupRouter()))
}

func setupRouter() *mux.Router {

  r := mux.NewRouter()


  r.HandleFunc("/health", HealthCheckHandler).Methods("GET")


  r.HandleFunc("/hello", HelloHandler).Methods("GET")

  return r
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {

  fmt.Println("HIT: healthcheck")


  _, _ = io.WriteString(w, "ALIVE")


}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
  // Extract the query parameters from the GET request
  queryParams := r.URL.Query()

  // Retrieve the query parameters with the key "name"
  nameParams := queryParams["name"]

  var name string
  switch len(nameParams) {
     case 0:
       // Set the name variable to there when there is no parameter "name" in the request
       name = "there"
     default:
       // Set the name variable to the first parameter "name" in the request
       name = nameParams[0]
  }

  // Return status 400 if name is empty
  if name == "" {
    w.WriteHeader(400)
    return
  }

  // Write the string "Hello <name>" into the response's body
  _, _ = io.WriteString(w, fmt.Sprintf("Hello %s!", name))

  // Print a line in the ACCESS log
  fmt.Printf("HIT: hello handler with name %s \n", name)
}