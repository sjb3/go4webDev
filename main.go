package main

import (
  "net/http"
  "fmt"
  "html/template"
  _ "github.com/mattn/go-sqlite3"
  "database/sql"
  "encoding/json"
  "net/url"
  "io/ioutil"
  "encoding/xml"
)

type Page struct{
  Name string
  DBStatus bool
}

type SearchResult struct {
  Title string `xml:"title,attr"`
  Author string `xml:"author,attr"`
  Year string `xml:"hyr,attr"`
  Id string `xml:"owi,attr"`
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
    // results := []SearchResult{
    //   SearchResult{"Mody Dick", "Herman Melville", "1851", "11111"},
    //   SearchResult{"Huckleberry Finn", "Mark Twain", "1884", "22222"},
    //   SearchResult{"The Catcher in the Rye", "JD Salinger", "1951", "33"},
    // }
    var results []SearchResult
    var err error

    if results, err = search(r.FormValue("Search")); err!=nil{
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }

    encoder := json.NewEncoder(w)
    if err := encoder.Encode(results); err != nil{
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  })

  http.HandleFunc("/books/add", func(w http.ResponseWriter, r*http.Request){
    var book ClassifyBookResponse
    var err error

    if book, err = find(r.FormValue("id")); err != nil{
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    if err = db.Ping(); err != nil{
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    _, err = db.Exec("insert into books(pk, title, author, id, classification) value(?,?,?,?,?)",
                      nil, book.BookData.Title, book.BookData.Author, book.BookData.ID, book.Classification.MostPopular)

    if err != nil{
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  })

  fmt.Println(http.ListenAndServe(":8000", nil))
}

type ClassifySearchResponse struct{
  Results []SearchResult `xml:"works>work"`
}

type ClassifyBookResponse struct{
  BookData struct{
    Title string `xml: "title,attr"`
    Author string `xml: "author,attr"`
    ID string `xml:"owi,attr"`
  } `xml:"work"`
  Classification struct{
    MostPopular string `xml: "sfa,attr"`
  }`xml: "recommendations>ddC>mostPopular"`
}

func find(id string)(ClassifyBookResponse, error){
  var c ClassifyBookResponse
  body, err := classifyAPI("http://classify.oclc.org/classify2/Classify?&summary=true&owi=" + url.QueryEscape(id))

  if err !=nil{
    return ClassifyBookResponse{}, err
  }
  err = xml.Unmarshal(body, &c)
  return c, err
}

func search(query string)([]SearchResult, error){

  var c ClassifySearchResponse
  body, err := classifyAPI("http://classify.oclc.org/classify2/Classify?&summary=true&title=" + url.QueryEscape(query))
  if err != nil{
    return []SearchResult{}, err
  }

  err = xml.Unmarshal(body, &c)
  return c.Results, err
}

func classifyAPI(url string)([]byte, error){
  var resp *http.Response
  var err error

  // if resp, err = http.Get("http://classify.oclc.org/classify2/Classify?&summary=true&title=" + url.QueryEscape(query)); err!=nil{
  //   return []SearchResult{}, err
  // }
  if resp, err = http.Get(url); err !=nil{
    return []byte{}, err
  }

  defer resp.Body.Close()
  // var body []byte
  // if body, err = ; err!=nil{
  //   return []byte{}, err
  // }
  return ioutil.ReadAll(resp.Body)

}
