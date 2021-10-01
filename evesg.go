package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Item struct {
	Name string
	Body []byte
}

func (i *Item) save() error {
	filename := i.Name + ".txt"
	return ioutil.WriteFile(filename, i.Body, 0600)
}

func loadItem(name string) (*Item, error) {
	filename := name + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Item{Name: name, Body: body}, nil
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, i *Item) {
	err := templates.ExecuteTemplate(w, tmpl+".html", i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/view/"):]
	i, err := loadItem(name)
	if err != nil {
		http.Redirect(w, r, "/edit/"+name, http.StatusFound)
		return
	}
	renderTemplate(w, "view", i)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/edit/"):]
	i, err := loadItem(name)
	if err != nil {
		i = &Item{Name: name}
	}
	renderTemplate(w, "edit", i)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	i := &Item{Name: name, Body: []byte(body)}
	err := i.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+name, http.StatusFound)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
