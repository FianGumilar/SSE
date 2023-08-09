package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/stream-price", func(w http.ResponseWriter, r *http.Request) {
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "error streaming not supported", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		stockChannel := make(chan string)
		defer close(stockChannel) // Ensure the channel is closed when the handler exits.
		go streamUpdatePrice(r.Context(), stockChannel)

		for stock := range stockChannel {
			event := fmt.Sprintf("event: %s\n"+
				"data: %s\n\n", "price-changed", stock)
			_, _ = fmt.Fprintf(w, event)
			flusher.Flush()
		}
	})
	http.ListenAndServe(":9000", nil)
}

func streamUpdatePrice(ctx context.Context, channel chan<- string) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop() // Stop the ticker when the function exits.

	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	stocks := []string{"BBCA", "INDO", "USG, UNVR", "BBCA"}

labelBreak:
	for {
		select {
		case <-ctx.Done():
			break labelBreak
		case <-ticker.C:
			price := r.Intn(1000000)
			stockIndex := r.Intn(len(stocks))
			channel <- fmt.Sprintf("%s-%d", stocks[stockIndex], price)
		}
	}
}
