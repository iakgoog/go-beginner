package webserver

import (
	"log"
	"net/http"
	"time"
)

// RunWebServer function
func RunWebServer() {
	runMiddleWare()
}

func hardWay() {
	err := http.ListenAndServe(":8080", &indexHandler{})
	log.Println(err)
}

type indexHandler struct{}

func (*indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Gopher"))
}

/*================================ MULTIPLEXER (MUX) ================================*/

func easyWay() {
	// h := http.HandlerFunc(anotherIndexHandler)
	h := http.HandlerFunc(mux)
	err := http.ListenAndServe(":8080", h)
	log.Println(err)
}

func anotherIndexHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		w.Write([]byte("Index Page"))
	case "/about":
		w.Write([]byte("About Page"))
	default:
		w.Write([]byte("404 Page Not Found"))
	}
}

// create MULTIPLEXER (MUX)
func mux(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/about":
		aboutPageHandler(w, r)
	case "/login":
		loginPageHandler(w, r)
	case "/":
		indexPageHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}

func indexPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Index Page"))
}

func aboutPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page"))
}

func loginPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login Page"))
}

/*================================ EASY MUX ================================*/

// easier way to create MUX
func easyMux() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexPageHandler)
	mux.HandleFunc("/about", aboutPageHandler)
	mux.HandleFunc("/login", loginPageHandler)

	err := http.ListenAndServe(":8080", mux)
	log.Println(err)
}

/*================================ MIDDLEWARE ================================*/

func runMiddleWare() {
	h := logger(http.HandlerFunc(index2Handler))
	err := http.ListenAndServe(":8080", h)
	log.Println(err)
}

func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("requestURI: %s, path: %s\n", r.RequestURI, r.URL.Query())
		t := time.Now()
		h.ServeHTTP(w, r)
		diff := time.Now().Sub(t)
		log.Printf("path: %s, time: %dns", r.URL.Path, diff/time.Nanosecond)
	})
}

func index2Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index Page"))
}
