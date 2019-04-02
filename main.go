package main

import (
	"fmt"
	"html/template"
	"net/http"
	//"github.com/gorilla/mux"
)

var homeTemplate *template.Template

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}

}

func main() {

	var err error
	homeTemplate, err = template.ParseFiles(
		"views/index.gohtml",
		"views/layouts/header.gohtml",
		"views/layouts/navbar.gohtml",
		"views/layouts/footer.gohtml")
	if err != nil {
		panic(err)
	}
	fmt.Print("Listening port :3000")

	//r := mux.NewRouter()
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":3000", nil)
}
