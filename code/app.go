package main

import (
	"html/template"
	"net/http"
)

type application struct {
	db  DB
	mux *http.ServeMux
}

func (app *application) red(w http.ResponseWriter, r *http.Request) {
	surl := r.URL.Path[1:len(r.URL.Path)]

	lurl := app.db.GetLurl(surl)

	exam := lurl[0:4]
	if exam == "http" {
		http.Redirect(w, r, lurl, http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "https://"+lurl, http.StatusSeeOther)
	}
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	r.ParseForm()
	url := r.FormValue("url")

	files := []string{
		"./html/home_page.html",
	}

	ts, _ := template.ParseFiles(files...)
	ts.Execute(w, url)
}

func (app *application) proc(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	url := r.FormValue("url")

	surl, qwe := app.db.Insert(url)
	if qwe {
		app.mux.HandleFunc("/"+surl, app.red)
	}

	http.Redirect(w, r, "/?url="+surl, http.StatusSeeOther)

}
