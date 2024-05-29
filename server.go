package main

import (
    "log"
    "net/http"
    "os"
    "path/filepath"
    "time"
)

var PORT string
PORT = "8080"

// LoggingMiddleware logs the details of each request and prints them out to cmdline
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
    })
}

// NotFoundHandler handles 404 errors
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, filepath.Join("static", "404.html"))
}

// FileServerWith404 creates a file server with custom 404 error handling
func FileServerWith404(root http.FileSystem) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        path := r.URL.Path
        file, err := root.Open(path)
        if err != nil {
            if os.IsNotExist(err) {
                NotFoundHandler(w, r)
                return
            }
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }
        defer file.Close()

        fi, err := file.Stat()
        if err != nil {
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }

        // Check if the path is a directory
        if fi.IsDir() {
            // Serve the index.html file if the directory is requested (central operation)
            index := filepath.Join(path, "index.html")
            _, err := root.Open(index)
            if err != nil {
                NotFoundHandler(w, r)
                return
            }
            path = index
        }
        
        http.ServeFile(w, r, filepath.Join("static", path))
    })
}

func main() {
    // Directory of service
    fs := http.Dir("./static")
    http.Handle("/", LoggingMiddleware(FileServerWith404(fs)))

    // Listen on specified port
    log.Println("Listening on 8080...")
    err := http.ListenAndServe(":"+PORT, nil)
    if err != nil {
        log.Fatal(err)
    }
}
