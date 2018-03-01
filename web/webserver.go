package web

import (
	"net/http"
	"encoding/json"
	"goworkshop/model"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
)

func GetAllAuthors(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if serializedData, err := json.Marshal(model.Authors); err != nil {
		fmt.Fprintf(w, "{\"errorMessage\":\"%s\"}", err.Error())
	} else {
		fmt.Fprintf(w, string(serializedData))
	}

}

func GetAuthorByUuid(w http.ResponseWriter, req *http.Request) {
	uuid := mux.Vars(req)["uuid"]
	w.Header().Set("Content-Type", "application/json")
	for _, author := range model.Authors {
		if author.UUID == uuid {
			if serializedData, err := json.Marshal(author); err != nil {
				fmt.Fprintf(w, "{\"errorMessage\":\"%s\"}", err.Error())
				return
			} else {
				fmt.Fprintf(w, string(serializedData))
				return
			}
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func CreateAuthor(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		fmt.Fprintf(w, "{\"errorMessage\":\"%s\"}", err.Error())
		return
	} else {
		var author model.AuthorDto
		if err := json.Unmarshal(body, &author); err != nil {
			fmt.Fprintf(w, "{\"errorMessage\":\"%s\"}", err.Error())
			return
		} else {
			model.Authors = append(model.Authors, author)
			if serializedData, err := json.Marshal(author); err != nil {
				fmt.Fprintf(w, "{\"errorMessage\":\"%s\"}", err.Error())
				return
			} else {
				fmt.Fprintf(w, string(serializedData))
				return
			}
		}
	}
}

func StartWebServer() {
	gorillaMux := mux.NewRouter()
	gorillaMux.HandleFunc("/authors", GetAllAuthors).Methods("GET")
	gorillaMux.HandleFunc("/authors", CreateAuthor).Methods("POST")
	gorillaMux.HandleFunc("/authors/{uuid}", GetAuthorByUuid).Methods("GET")

	if err := http.ListenAndServe(":8080", gorillaMux); err != nil {
		panic(err)
	}
}
