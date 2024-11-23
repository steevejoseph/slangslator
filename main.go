package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)


func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", HandleHome)
	router.HandleFunc("/quick-define", HandleQuickDefine)
	router.HandleFunc("/quick-define/{term:.*}", HandleQuickDefine)

	PORT := ":8080"
	fmt.Printf("Server runnin on %s at %s\n", PORT, time.Now().Format("Monday, January 2, 2006 at 3:04 PM"))
	http.ListenAndServe(PORT, router)
}
