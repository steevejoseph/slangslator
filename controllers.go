package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/index.html")
}

func HandleQuickDefine(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	term := vars["term"]
	client := &http.Client{}

	if  len(term) < 1 {
		ReturnErrorResponse(w, errors.New("no term specified"),  "no term specified", http.StatusBadRequest)
		return 
	}

	url := fmt.Sprintf("https://api.urbandictionary.com/v0/define?term=%s", term)
	message := fmt.Sprintf("Failed to get definition of \"%s\" from Urban Dictionary", term)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		ReturnErrorResponse(w, err, message, http.StatusInternalServerError)
		return
	}

	res, err := client.Do(req)

	if err != nil {
		ReturnErrorResponse(w, err, message, http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			ReturnErrorResponse(w, err, message, http.StatusInternalServerError)
			return
		}

		defs := &UrbanDictionaryResponse{}
		err = json.Unmarshal(bodyBytes, defs)
		if err != nil {
			ReturnErrorResponse(w, err, message, http.StatusInternalServerError)
			return
		}
		
		retval, err := json.MarshalIndent(defs, "", "  ")
		if err != nil {
			ReturnErrorResponse(w, err, message, http.StatusInternalServerError)
			return
		} // log.Print(PrettyPrint(defs.List[0]))
		
		w.Write(retval)
		return
	}

	ReturnErrorResponse(w, err, message, http.StatusInternalServerError)
}
