// -----------------------------------------------------------------------------
// App: SlackDataStrava
//
// Description:
//    This app provides users the capabilty to query and obtain all Strava related
//    data via a simple REST API interface. It basically provides more or less
//    the same functionality as the stock Strava UI, but can now be invoked
//    via Slack. (wwww.slack.com)
// Author:
//    Barry Winata, 2017
// Version:
//    1.0.
// -----------------------------------------------------------------------------

package main

import (
  "fmt"
  "time"
  "net/http"
  "github.com/gorilla/mux"
)


func handleMe (w http.ResponseWriter, r * http.Request) {
  fmt.Println ("Getting information about me!")
}

func handleMeBike (w http.ResponseWriter, r * http.Request) {
  fmt.Println ("Getting information about me / bike")
}

func handleMeRun (w http.ResponseWriter, r * http.Request) {
  fmt.Println ("Getting information about me / run")
}


func main () {
  // Create new router instance
  r := mux.NewRouter ()

  //
  r.HandleFunc ("/me", handleMe).Methods ("GET")
  r.HandleFunc ("/me/bike", handleMeBike).Methods ("GET")
  r.HandleFunc ("/me/run", handleMeRun).Methods ("GET")

  server := &http.Server {
    Handler      : r,
    Addr         : "0.0.0.0:8000",
    WriteTimeout : 15 * time.Second,
    ReadTimeout  : 15 * time.Second,
  }

  // Initialise server
  server.ListenAndServe ()
}
