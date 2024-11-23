package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func PrettyPrint(p any) string {
	// Marshal the struct into indented JSON
	prettyJSON, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	return string(prettyJSON)
}

func mustNot(err error) {
	if err != nil {
		panic(err)
	}
}

type ErrorResponse struct {
    Error error
    Message string
    Code int
}

func ReturnErrorResponse(w http.ResponseWriter, err error, message string, code int) {
	errRes, _ := json.MarshalIndent(ErrorResponse{
		Error: err,
		Message: message,
		Code: code,
	}, "", "  ")

	w.Write(errRes)
} 