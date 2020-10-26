package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

type order struct {
	OrderID    int    `json:"orderID"`
	CustomerID string `json:"customerID"`
	Quantity   int    `json:"quantity"`
	Amount     int    `json:"amount"`
}

type response struct {
	Message     string `json:"message"`
	StatusError int    `json:"statusError"`
}

func main() {
	PORT := ":9000"
	logger.Info("get all orders api started on port " + PORT)

	r := mux.NewRouter()
	r.Use(CORS)

	r.HandleFunc("/orders", getAllOrders)
	r.Handle("/health", healthCheck())
	log.Fatal(http.ListenAndServe(PORT, r))
}

func getAllOrders(w http.ResponseWriter, r *http.Request) {
	orders := []order{{4, "3456", 4, 7}, {5, "3678", 1, 1}, {6, "3567", 2, 3}}
	response, _ := json.Marshal(orders)

	logger.Info("response get all orders")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func healthCheck() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		JSONResponse(w, http.StatusOK, "get all is up")
	})
}

// JSONResponse returns the handler json response
func JSONResponse(w http.ResponseWriter, statusCode int, message interface{}) {
	response, _ := json.Marshal(message)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

// CORS Middleware
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Set headers
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Next
		next.ServeHTTP(w, r)
		return
	})
}
