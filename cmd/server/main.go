
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Thurmanlarke server is starting...")

	// Load the port provided by the hosting service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Homepage
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Println("Homepage error:", err)
			http.Error(
				w,
				"Unable to load homepage",
				http.StatusInternalServerError,
			)
			return
		}

		if err := tmpl.Execute(w, nil); err != nil {
			log.Println("Homepage rendering error:", err)
		}
	})

	// Privacy Policy page
	http.HandleFunc("/privacy", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/privacy" {
			http.NotFound(w, r)
			return
		}

		tmpl, err := template.ParseFiles("templates/privacy.html")
		if err != nil {
			log.Println("Privacy page error:", err)
			http.Error(
				w,
				"Unable to load Privacy Notice",
				http.StatusInternalServerError,
			)
			return
		}

		if err := tmpl.Execute(w, nil); err != nil {
			log.Println("Privacy page rendering error:", err)
		}
	})

	// Serve CSS, JavaScript, and images.
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle(
		"/static/",
		http.StripPrefix("/static/", fileServer),
	)

	// Start the server.
	address := "0.0.0.0:" + port
	fmt.Println("Thurmanlarke server listening on port", port)

	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatal("Server error:", err)
	}
}
