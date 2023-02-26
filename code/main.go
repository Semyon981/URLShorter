package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var app application
	app.db.openDB()
	defer app.db.closeDB()

	app.mux = http.NewServeMux()
	app.mux.HandleFunc("/", app.home)
	app.mux.HandleFunc("/proc", app.proc)

	surls := app.db.GetAll()

	for i := 0; i < len(surls); i++ {
		app.mux.HandleFunc("/"+surls[i], app.red)
	}

	fileServer := http.FileServer(http.Dir("./static/"))
	app.mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	err := http.ListenAndServe("127.0.0.1:4000", app.mux)
	log.Fatal(err)

}
