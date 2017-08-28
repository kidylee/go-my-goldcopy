package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

type record struct {
	ID        string `json:"ssbtradeid"`
	TradeDate string `json:"tradedate"`
	Amount    string `json:"amount"`
	Status    string `json:"status"`
}

func getCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Debug("Categories' key: ", vars["key"])
	val, err := redisClient.ZRangeWithScores(vars["key"], 0, -1).Result()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Category: %v\n", err)

		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Category: %v\n", err)

		return
	}
	respondWithJSON(w, http.StatusOK, val)

}

func createCategory(w http.ResponseWriter, r *http.Request) {

	var re record
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&re); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	log.Debug(re)
	defer r.Body.Close()

	b, err := json.Marshal(re)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	var score int64
	if re.TradeDate != "" {
		t, err := time.Parse("2006/01/02", re.TradeDate)
		if err != nil {
			log.Debug(err)
			score = time.Now().Unix()
		} else {
			score = t.Unix()
		}
	}

	member := &redis.Z{
		Score:  float64(score),
		Member: string(b),
	}

	result := redisClient.ZAdd("GC:UNMATCHED", *member).Err()

	log.Debug(result)
	respondWithJSON(w, http.StatusCreated, re)
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
