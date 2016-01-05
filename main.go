package main

import (
	"flag"
	"net/http"
	"text/template"
	"log"
)

var addr = flag.String("addr", ":8080", "http service address")
var homeTempl = template.Must(template.ParseFiles("/html/home.html"))


func (h *hub)serveHome(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	homeTempl.Execute(w, r.Host)
}

//Minimum one ws0 socket to function.
func main() {
	//handle assets and allow router to access.
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	flag.Parse()

	//Add 3 hubs for testing
	contr.addHub()
	contr.addHub()
	contr.addHub()

  contr.run()

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

//There is a new socket.  default to last one in slice of hubs for everything though.  Why
//TODO singleton pattern for controller.
//TODO map of socket integers to name of chat channels hubs whatever
