package main
import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "привет")
	})

	fmt.Printf("порт сервера: %s\n", port)
	
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
