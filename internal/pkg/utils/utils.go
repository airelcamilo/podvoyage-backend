package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ParseBody(w http.ResponseWriter, r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		err := json.Unmarshal([]byte(body), x)
		if err != nil {
			FormatResponse(w, http.StatusBadRequest, nil)
			return
		}
	}
}

func CheckErrIsNil(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func FormatResponse(w http.ResponseWriter, statusCode int, obj interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(obj)
}

func GetId(w http.ResponseWriter, r *http.Request) int {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		FormatResponse(w, http.StatusBadRequest, nil)
		return -1
	}
	return id
}