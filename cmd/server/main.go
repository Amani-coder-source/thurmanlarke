package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	fmt.Println("Thurmanlarke server is starting...")

	// Homepage
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Only handle the homepage at "/"
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(
				w,
				"Unable to load homepage",
				http.StatusInternalServerError,
			)
			fmt.Println("Homepage error:", err)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			fmt.Println("Homepage rendering error:", err)
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
			http.Error(
				w,
				"Unable to load Privacy Notice",
				http.StatusInternalServerError,
			)
			fmt.Println("Privacy page error:", err)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			fmt.Println("Privacy page rendering error:", err)
		}
	})

	// Serve CSS, JavaScript and images
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle(
		"/static/",
		http.StripPrefix("/static/", fileServer),
	)

	fmt.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}