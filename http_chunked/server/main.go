package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/chunked", func(w http.ResponseWriter, r *http.Request) {
		flusher := w.(http.Flusher)
		for i := 0; i < 2; i++ {
			fmt.Fprint(w, "打gopher\n")
			flusher.Flush()
			<-time.Tick(1 * time.Second)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
