package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		responseData, err := json.Marshal(map[string]string{"response": "pong"})
		if err != nil {
			fmt.Printf("Encountered Error: %v\n", err)
		}
		w.Write(responseData)
	})
	fmt.Printf("Listening on port 3000\n")
	http.ListenAndServe(":8080", nil)
}
