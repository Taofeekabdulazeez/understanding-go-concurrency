package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		println("Hello, world")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		w.Header().Set("Content-Type", "text/event-stream")

		tokens := []string{"My", "name", "is", "Abdulazeez", "Taofeek", "Taiwo", "Olanrewaju"}

		for _, token := range tokens {
			content := fmt.Sprintf("data: %s\n\n", string(token))
			w.Write([]byte(content))
			w.(http.Flusher).Flush()

			time.Sleep(time.Millisecond * 420)
		}
	})

	http.ListenAndServe(":8081", nil)

}
