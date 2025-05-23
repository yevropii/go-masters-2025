package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func fibo(n int) int {
	if n < 2 {
		return n
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}

func fiboHandler(w http.ResponseWriter, r *http.Request) {
	nStr := r.URL.Query().Get("N")
	n, err := strconv.Atoi(nStr)
	if err != nil || n < 0 {
		http.Error(w, "Bad N", http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, fibo(n))
}

func main() {
	http.HandleFunc("/fibo", fiboHandler)
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
