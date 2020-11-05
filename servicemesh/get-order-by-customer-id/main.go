package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

// Orders struct for individual order
type Orders struct {
	OrderID    int    `json:"orderID"`
	CustomerID string `json:"customerID"`
	Quantity   int    `json:"quantity"`
	Amount     int    `json:"amount"`
}

func main() {
	PORT := ":9001"
	
	r := mux.NewRouter()
	r.Use(CORS)
	
	r.HandleFunc("/order/{customerID}", getOrderByCustomerID)
	
	logger.Info("get order by customerID on port " + PORT)
	logger.Fatal(http.ListenAndServe(PORT, r))
}

func getOrderByCustomerID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["customerID"]

	customer, err := strconv.Atoi(customerID)
	if err != nil {
		logger.Error("conversion error")
		return
	}
	w.WriteHeader(http.StatusOK)

	ordersResp, err := http.Get("http://all-svc:9000/orders")
	if err != nil {
		logger.Error(err)
	}

	orders, err := ioutil.ReadAll(ordersResp.Body)
	if err != nil {
		panic(err)
	}

	var orderResult []Orders
	json.NewDecoder(strings.NewReader(string(orders))).Decode(&orderResult)
	logger.Debug(orderResult)

	for _, orderSearch := range orderResult {
		customerInt, _ := strconv.ParseInt(orderSearch.CustomerID, 10, 32)
		if int64(customer) == customerInt {
			out, _ := json.Marshal(orderSearch)

			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(out))
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"statusError": "can not find customer"}`))
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
