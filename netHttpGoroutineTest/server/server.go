package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("req: ping")
		fmt.Printf("Current goroutines: %d\n", runtime.NumGoroutine())
		w.Write([]byte("Pong"))
	})

	fmt.Println("Server starting on :3000")
	http.ListenAndServe(":3000", nil)
}
