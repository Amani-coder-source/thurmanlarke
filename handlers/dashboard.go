package handlers

import (
	"html/template"
	"net/http"

	"Thurmanlark/middleware"
	"Thurmanlark/repository"
)

type DashboardHandler struct {
	UserRepo repository.UserRepository
}

// Dashboard displays the logged-in client's dashboard.
func (h DashboardHandler) Dashboard(w http.ResponseWriter, r *http.Request) {

	// Check if the user is logged in.
	userID, loggedIn := middleware.GetUserID(r)

	if !loggedIn {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Get the user's information from MySQL.
	user, err := h.UserRepo.GetUserByID(userID)

	if err != nil {
		http.Error(
			w,
			"Unable to find your account",
			http.StatusInternalServerError,
		)
		return
	}

	// Load dashboard template.
	tmpl, err := template.ParseFiles("templates/dashboard.html")

	if err != nil {
		http.Error(
			w,
			"Unable to load dashboard",
			http.StatusInternalServerError,
		)
		return
	}

	// Send user information to the template.
	err = tmpl.Execute(w, user)

	if err != nil {
		http.Error(
			w,
			"Unable to display dashboard",
			http.StatusInternalServerError,
		)
		return
	}
}

// Logout logs the user out.
func (h DashboardHandler) Logout(w http.ResponseWriter, r *http.Request) {

	middleware.ClearSession(w)

	http.Redirect(
		w,
		r,
		"/login",
		http.StatusSeeOther,
	)
}