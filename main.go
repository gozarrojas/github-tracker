package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received POST request!")
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading the request:", err)
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	fmt.Println(string(body))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Request received successfully"))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", postHandler).Methods("POST")

	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
