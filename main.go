package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the details of the incoming request
		log.Printf("Request: Method=%s, URL=%s, User-Agent=%s, IP=%s, Time=%s\n", 
			r.Method, r.URL, r.UserAgent(), r.RemoteAddr, time.Now().Format(time.RFC3339))

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter()
	router.Use(loggingMiddleware)

	router.HandleFunc("/", HandleHome)
	router.HandleFunc("/quick-define", HandleQuickDefine)
	router.HandleFunc("/quick-define/{term:.*}", HandleQuickDefine)

	// Use the PORT environment variable, default to 8080 if not set
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on %s at %s\n", port, time.Now().Format("Monday, January 2, 2006 at 3:04 PM"))
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
