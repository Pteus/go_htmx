package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler1)
	http.HandleFunc("/add-film", handleAddFilm)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handleAddFilm(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	title := r.PostFormValue("title") // name of the input element
	director := r.PostFormValue("director")

	//htmlStr := fmt.Sprintf(" <li class='list-group-item bg-primary text-white'>\n %s - %s\n </li>", title, director)
	//tmpl, _ := template.New("t").Parse(htmlStr)

	//tmpl.Execute(w, nil)

	templ := template.Must(template.ParseFiles("index.html"))
	templ.ExecuteTemplate(w, "film-list-block", Film{Title: title, Director: director})
}

func handler1(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("index.html"))

	films := map[string][]Film{
		"Films": {
			{Title: "The godfather", Director: "Francis Ford Coppola"},
			{Title: "The moon", Director: "Moons Ford Coppol"},
		},
	}

	templ.Execute(w, films)
}

type Film struct {
	Title    string
	Director string
}
