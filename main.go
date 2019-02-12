package main

import (
	"net/http"

	"./views"

	"github.com/gorilla/mux"
)

var (
	homeView    *views.View
	contactView *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)

	if err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil)

	if err != nil {
		panic(err)
	}
}

// func faq(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	fmt.Fprint(w, "<h1>FAQ</h1><p>Here is a list of questions</p>")
// }

// func notFound(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	w.WriteHeader(http.StatusNotFound)
// 	fmt.Fprint(w, "<h1>Sorry, but we couldn't find the page you were looking for</h1>")
// }

func main() {

	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")

	r := mux.NewRouter()
	// r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	// r.HandleFunc("/faq", faq)
	http.ListenAndServe(":3000", r)
}
