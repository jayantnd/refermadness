package web

import (
  "github.com/codegangsta/negroni"
  "github.com/gorilla/mux"
  "github.com/larryprice/refermadness/controllers"
  "html/template"
  "net/http"
)

type Server struct {
  *negroni.Negroni
}

func NewServer(database Database, clientID, clientSecret string, isDevelopment bool) *Server {
  s := Server{negroni.Classic()}

  router := mux.NewRouter()
  router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("views/layout.html", "views/index.html")
    t.Execute(w, nil)
  })
  router.HandleFunc("/legal", func(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("views/layout.html", "views/legal.html")
    t.Execute(w, nil)
  })
  router.HandleFunc("/account", func(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("views/layout.html", "views/account.html")
    t.Execute(w, nil)
  })
  router.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("views/layout.html", "views/search.html")
    t.Execute(w, nil)
  })
  router.HandleFunc("/service/create", func(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("views/layout.html", "views/create-service.html")
    t.Execute(w, nil)
  })
  router.HandleFunc("/service/{id}", func(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("views/layout.html", "views/service.html")
    t.Execute(w, nil)
  })
  authenticationController := controllers.NewAuthenticationController(clientID, clientSecret, isDevelopment)
  authenticationController.Register(router)

  s.UseHandler(router)
  s.Use(database.Middleware())

  return &s
}
