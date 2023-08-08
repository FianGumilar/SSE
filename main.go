package main

import (
	"context"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/stream-price", func(w http.ResponseWriter, r *http.Request) {

	})
	http.ListenAndServe(":9000", nil)
}

func streamUpdatePrice(ctx context.Context, channel chan<- string) {
	ticker := time.NewTicker(time.Second)

	stocks := []string{"BBCA", "INDO", "USG, UNVR", "BBCA"}

	for {
		select {
		case <-ticker.C:
			price := rand.Intn(1000000)
			stockIndex := rand.Intn(len(stocks))
		}
	}
}
