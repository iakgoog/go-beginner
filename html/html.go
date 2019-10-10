package html

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

// RunHTML function
func RunHTML() {
	// html/public/css/style.css => /-/css/style.css
	// / index
	runWithTemplate()
}

func runWithDebugMiddleWare() {
	http.HandleFunc("/", simpleIndexHandler)
	http.HandleFunc("/-/", fileServerHandler)
	http.ListenAndServe(":8080", nil)
}

// Vulnerable Warning: The served path expose file
func fileServerHandler(w http.ResponseWriter, r *http.Request) {
	// log.Println(r.URL.Path) // for debug
	// r.URL.Path = "/public/css/style.css" can be changed for debug
	h := http.FileServer(http.Dir("html"))
	http.StripPrefix("/-", h).ServeHTTP(w, r)
}

// Vulnerable Warning: The served path expose file
func runWihoutDebug() {
	http.HandleFunc("/", simpleIndexHandler)
	http.Handle("/-/", http.StripPrefix("/-", http.FileServer(http.Dir("html"))))
	http.ListenAndServe(":8080", nil)
}

// Fix vulnerability issue
// Syntax like `class noDir extends http.Dir`
type noDir struct {
	http.Dir
}

// Wrap Dir.Open with custom check Open
func (d noDir) Open(name string) (http.File, error) {
	f, err := d.Dir.Open(name)
	if err != nil {
		return nil, err
	}
	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if stat.IsDir() {
		return nil, os.ErrNotExist
	}
	return f, nil
}

func runWithFixedVulnerability() {
	http.HandleFunc("/", simpleIndexHandler)
	http.Handle("/-/", http.StripPrefix("/-", http.FileServer(noDir{http.Dir("html")})))
	http.ListenAndServe(":8080", nil)
}

func simpleIndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(`
		<!doctype html>
		<title>Static Web Server</title>
		<link href=/-/public/css/style.css rel=stylesheet>
		<p class=red>
			Static web server.
		</p>
	`))
}

func runWithTemplate() {
	http.HandleFunc("/", templateIndexHandler)
	http.Handle("/-/", http.StripPrefix("/-", http.FileServer(noDir{http.Dir("html")})))
	http.ListenAndServe(":8080", nil)
}

func templateIndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles("index.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
