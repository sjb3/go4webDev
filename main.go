package main

import (
  "net/http"
  "fmt"
  "html/template"
  _ "github.com/mattn/go-sqlite3"
  "database/sql"
  "encoding/json"
)

type Page struct{
  Name string
  DBStatus bool
}

type SearchResult struct {
  Title string
  Author string
  Year string
  Id string
}

func main() {
  templates := template.Must(template.ParseFiles("templates/index.html"))

  db, _ := sql.Open("sqlite3", "dev.db")


  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    p := Page{Name: "Gopher"}

    if name := r.FormValue("name"); name !=""{
      p.Name = name
    }
    p.DBStatus = db.Ping() == nil

    if err := templates.ExecuteTemplate(w, "index.html", p);
    err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  })
  http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request){
    results := []SearchResult{
      SearchResult{"Mody Dick", "Herman Melville", "1851", "11111"},
      SearchResult{"Huckleberry Finn", "Mark Twain", "1884", "22222"},
      SearchResult{"The Catcher in the Rye", "JD Salinger", "1951", "33"},

    }
    encoder := json.NewEncoder(w)
    if err := encoder.Encode(results); err != nil{
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  })

  fmt.Println(http.ListenAndServe(":8000", nil))
}
