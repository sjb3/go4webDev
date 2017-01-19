package main

import (
  "net/http"
  "fmt"
  "html/template"
  "github.com/mattn/go-sqlite3"
  "database/sql"
)

type Page struct{
  Name string

}

func main() {
  templates := template.Must(template.ParseFiles("templates/index.html"))

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    p := Page{Name: "Gopher"}

    if name := r.FormValue("name"); name !=""{
      p.Name = name
    }
    if err := templates.ExecuteTemplate(w, "index.html", p);
    err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    fmt.Fprintf(w, "<h1>Hello HandleFunc</h1>")
  })

  fmt.Println(http.ListenAndServe(":8000", nil))
}
