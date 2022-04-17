package main

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
	"github.com/wisdom-yzh/youtube_playlist/parser"
)

type spaHandler struct {
	indexPath  string
	staticPath string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	execPath, err := os.Executable()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	basePath := filepath.Join(filepath.Dir(execPath), h.staticPath)
	path := filepath.Join(basePath, r.URL.Path)

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, filepath.Join(basePath, h.indexPath))
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.FileServer(http.Dir(basePath)).ServeHTTP(w, r)
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
	router := mux.NewRouter()
	router.Use(accessControlMiddleware)

	router.Path("/api/health").HandlerFunc(healthHandler).Methods(http.MethodGet)
	router.Path("/api/list/{list}").HandlerFunc(parser.PlaylistHandler).Methods(http.MethodGet)
	router.Path("/api/video/{video}").HandlerFunc(parser.VideoUrlHandler).Methods(http.MethodGet)

	spa := spaHandler{indexPath: "index.html", staticPath: "html"}
	router.PathPrefix("/").Handler(spa)

	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()
}
