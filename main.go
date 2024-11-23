package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", HandleHome)
	router.HandleFunc("/quick-define", HandleQuickDefine)
	router.HandleFunc("/quick-define/{term:.*}", HandleQuickDefine)

	// Use the PORT environment variable, default to 8080 if not set
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on %s at %s\n", port, time.Now().Format("Monday, January 2, 2006 at 3:04 PM"))
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
