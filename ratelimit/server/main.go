package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"sync"
	"golang.org/x/time/rate"
)

type Response struct {
	Message string `json:"message"`
}

func getIP(r *http.Request) string {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Printf("Error parsing IP: %v", err)
		return ""
	}
	log.Println("Host: ", host)
	return host
}

var ipLimiterMap sync.Map

func rateLimiterMiddleware(next http.Handler, limit rate.Limit, burst int) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := getIP(r)

		// Checks if the IP already has a limiter
		limiterAny, _ := ipLimiterMap.LoadOrStore(ip, rate.NewLimiter(limit, burst))
		limiter := limiterAny.(*rate.Limiter)

		// Check if the limit has been reached
		if !limiter.Allow() {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(map[string]string{"error": "Too many requests"})
			return
		}

		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{Message: "Hello, World!"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", greetHandler)

	handler := rateLimiterMiddleware(mux, rate.Limit(2), 10)

	log.Println("Server started on :8080")
	if err := http.ListenAndServe("0.0.0.0:8080", handler); err != nil {
		log.Fatal(err)
	}
}