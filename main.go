package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to Go Web Server!</h1>")
	fmt.Fprintf(w, "<p>Current time: %s</p>", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Fprintf(w, "<p>Your request method: %s</p>", r.Method)
	fmt.Fprintf(w, "<p>Your request URL: %s</p>", r.URL.Path)

}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About Page</h1>")
	fmt.Fprintf(w, "<p>This is a simple web server built with Go!</p>")
	fmt.Fprintf(w, "<a href='/'>Go back to home</a>")
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := `{
"message":"HEllo from hawkins!",
"timestamp":"` + time.Now().Format(time.RFC3339) + `",
"status":"Success"
}`

	fmt.Fprint(w, response)
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("started %s %s", r.Method, r.URL.Path)

		next(w, r)

		log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
	}
}

func main() {
	//create hhtp router
	mux := http.NewServeMux()

	mux.HandleFunc("/", loggingMiddleware(homeHandler))
	mux.HandleFunc("/about", loggingMiddleware(aboutHandler))
	mux.HandleFunc("/api", loggingMiddleware(apiHandler))

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("🚀 Server starting on http://localhost:8080")
	fmt.Println("Available routes:")
	fmt.Println("  - http://localhost:8080/")
	fmt.Println("  - http://localhost:8080/about")
	fmt.Println("  - http://localhost:8080/api")
	fmt.Println("Press Ctrl+C to stop the server")

	// Start the server
	log.Fatal(server.ListenAndServe())

}
