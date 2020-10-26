package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

var ctx = context.Background()

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
	records, err := getFromRedis(ctx, "records")

	if err != nil {
		stringOrders, _ := json.Marshal(orders)
		stored := string(stringOrders)
		saveToRedis(ctx, "records", stored)

		logger.Info("response get all orders: static")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response, _ := json.Marshal(orders)
		w.Write(response)
		return
	}

	logger.Info("response get all orders: redis")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(records))
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

// redisClient creates a connection to the redis database
func redisClient(ctx context.Context) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := client.Ping(ctx).Result()

	if err != nil {
		logger.Error("can not connect to redis database")
		panic(err)
	}
	return client
}

// saveToRedis saves key-value pair to redis
func saveToRedis(ctx context.Context, key string, data string) error {
	rClient := redisClient(ctx)
	err := rClient.Set(ctx, key, data, 20*time.Second).Err()
	if err != nil {
		logger.Error("Failed to save to redis")
		panic(err)
	}
	logger.Info("saved " + key + " to redis")
	return err
}

// gets key-value pair from redis
func getFromRedis(ctx context.Context, key string) (string, error) {
	rClient := redisClient(ctx)
	result, err := rClient.Get(ctx, key).Result()
	if err != nil {
		logger.Error("Failed to read from redis")
		return result, err
	}
	logger.Info("got " + key + " from redis")
	return result, nil
}
