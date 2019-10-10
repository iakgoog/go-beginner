package webserver

import (
	"log"
	"net/http"
	"time"
)

// RunWebServer function
func RunWebServer() {
	runConfigMiddleware()
}

/*================================ SIMPLE WEB SERVER ================================*/

func hardWay() {
	// SIGNATURE: func ListenAndServe(addr string, handler Handler) error {}
	/*
		type Handler interface {
			ServeHTTP(ResponseWriter, *Request)
		}
	*/
	err := http.ListenAndServe(":8080", &indexHandler{})
	log.Println(err)
}

type indexHandler struct{}

func (*indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Gopher"))
}

/*================================ MULTIPLEXER (MUX) ================================*/

func easyWay() {
	// type HandlerFunc func(ResponseWriter, *Request)
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
// THERE IS A LIB FOR MIDDLEWARE CHAINING AT "github.com/achoshift/middleware"

func runMiddleWare() {
	// demonstrate single middleware
	// h := logger(http.HandlerFunc(index2Handler))

	// demonstrate multiple middleware
	// h := m1(m2(m3(http.HandlerFunc(index2Handler))))

	// demonstrate chaining middleware
	// m := chain([]middleware{m1, m2, m3})
	// the difference between chain func and anotherChain func is
	// anotherChain func accept arguments like below
	// m := anotherChain(m1, m2, m3)
	// h := m(http.HandlerFunc(index2Handler))
	// or you it can be writen this way
	h := anotherChain(
		m1,
		m2,
		m3,
	)(http.HandlerFunc(index2Handler))

	err := http.ListenAndServe(":8080", h)
	log.Println(err)
}

type middleware func(http.Handler) http.Handler

// middleware chaining concept
func chain(hs []middleware) middleware {
	return func(h http.Handler) http.Handler {
		for i := len(hs); i > 0; i-- {
			h = hs[i-1](h)
		}
		return h
	}
}

// look for the difference from the above function
func anotherChain(hs ...middleware) middleware {
	return func(h http.Handler) http.Handler {
		for i := len(hs); i > 0; i-- {
			h = hs[i-1](h)
		}
		return h
	}
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

func m1(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("m1")
		h.ServeHTTP(w, r)
	})
}

func m2(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("m2")
		h.ServeHTTP(w, r)
	})
}

func m3(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("m3")
		h.ServeHTTP(w, r)
	})
}

/*================================ CONFIG MIDDLEWARE ================================*/

func allowRoleAdmin(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ***** coding the "Happy Path" way (easier to READ)
		// we get rid of what we don't want
		reqRole := r.Header.Get("Role")
		// This is what we DON'T want
		if reqRole != "admin" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		// This is what we WANT
		h.ServeHTTP(w, r)
	})
}

func allowRoleStaff(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqRole := r.Header.Get("Role")
		if reqRole != "staff" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		h.ServeHTTP(w, r)
	})
}

// allowRoleAdmin and allowRoleStaff are duplicated code,
// Let's DRY it.

func allowRole(role string) middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqRole := r.Header.Get("Role")
			if reqRole != role {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

// has performance issues with O(n)
func allowRoles(roles ...string) middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqRole := r.Header.Get("Role")
			// brute force way to check
			for _, role := range roles {
				if reqRole == role {
					h.ServeHTTP(w, r)
					return
				}
			}
			http.Error(w, "Forbidden", http.StatusForbidden)
		})
	}
}

// improve performance with O(1)
func allowRoles2(roles ...string) middleware {
	allow := make(map[string]bool)
	for _, role := range roles {
		allow[role] = true
	}
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqRole := r.Header.Get("Role")
			if !allow[reqRole] {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

// ***** further improve memory management
func allowRoles3(roles ...string) middleware {
	// ***** empty struct takes 0 byte, boolean takes memory
	allow := make(map[string]struct{})
	for _, role := range roles {
		// create empty struct
		allow[role] = struct{}{}
	}
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqRole := r.Header.Get("Role")
			if _, ok := allow[reqRole]; !ok {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

func runConfigMiddleware() {
	// create default Mux
	http.HandleFunc("/", index3Handler)
	// http.Handle("/admin", allowRoleAdmin(http.HandlerFunc(adminHandler)))
	// http.Handle("/staff", allowRoleStaff(http.HandlerFunc(staffHandler)))

	// Let's DRY it
	// allowRoleAdmin := allowRole("admin")
	// allowRoleStaff := allowRole("staff")
	// http.Handle("/admin", allowRoleAdmin(http.HandlerFunc(adminHandler)))
	// http.Handle("/staff", allowRoleStaff(http.HandlerFunc(staffHandler)))

	// Let's use allowRoles function
	http.Handle("/admin", allowRoles3("admin")(http.HandlerFunc(adminHandler)))
	http.Handle("/staff", allowRoles3("staff")(http.HandlerFunc(staffHandler)))
	http.Handle("/admin-staff", allowRoles3("admin", "staff")(http.HandlerFunc(adminStaffHandler)))

	err := http.ListenAndServe(":8080", nil)
	log.Println(err)
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello. Admin."))
}

func index3Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func staffHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello. Staff."))
}

func adminStaffHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello. Admin and Staff."))
}
