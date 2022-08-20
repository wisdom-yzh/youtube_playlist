package main

import (
	"embed"
	"encoding/json"
	"io"
	"io/fs"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/wisdom-yzh/youtube_playlist/parser"
)

//go:embed html/build/*
var html embed.FS

type spaHandler struct {
	indexPath  string
	staticPath string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fSys, err := fs.Sub(html, "html/build")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	f, err := fSys.Open(r.URL.Path)
	if err != nil {
		f, err = fSys.Open("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	content, err := io.ReadAll(f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func accessControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS,PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	router := mux.NewRouter()
	router.Use(accessControlMiddleware)

	router.Path("/api/health").HandlerFunc(healthHandler).Methods(http.MethodGet)
	router.Path("/api/list/{list}").HandlerFunc(parser.PlaylistHandler).Methods(http.MethodGet)
	router.Path("/api/video/{video}").HandlerFunc(parser.VideoUrlHandler).Methods(http.MethodGet)

	spa := spaHandler{indexPath: "index.html", staticPath: "html"}
	router.PathPrefix("/").Handler(http.StripPrefix("/", spa))

	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Begin to serve at %v\n", srv.Addr)
	srv.ListenAndServe()
}
