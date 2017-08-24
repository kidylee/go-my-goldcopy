package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Debug("Categories' key: ", vars["key"])
	val, err := redisClient.Get(vars["key"]).Result()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Category: %v\n", err)

		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", val)
}

func createCategory(w http.ResponseWriter, r *http.Request) {

	t := trade{ID: id}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
