package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
)

var navigationBarHTML string
var homepageTpl *template.Template

var templates = template.Must(template.ParseGlob("templates/*"))

// Page Data feed
type Page struct {
	Name  string
	Input string
}

// Input used to exploit
type Input struct {
	Payload string
	Request *http.Request
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	var FLAG = "F#{G0L4NG_net/http_T3mpl4tE_1nJecTION}"
	r.ParseForm() // parse arguments, you have to call this by yourself
	r.Header.Add("FLAG", FLAG)
	name := r.FormValue("name")
	if name == "" {
		name = "Woof"
	}

	var out bytes.Buffer
	var response string
	t, err := template.New("foo").Parse("Hello " + name + "!")
	if err == nil {
		t.Execute(&out, Input{Payload: name, Request: r})
		response = out.String()
	} else {
		response = "Hello " + name + "!"
	}

	p := Page{Name: response, Input: name}
	err = templates.ExecuteTemplate(w, "index.html", p)
	if err != nil {
		log.Fatal("Cannot Get View ", err)
	}
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", sayhelloName) // set router
	log.Print("Server running...")
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
