package routes

import (
	"net/http"

	"Thurmanlark/handlers"
	"Thurmanlark/repository"
)

func SetupRoutes(userRepo repository.UserRepository) {

	// =========================
	// AUTHENTICATION
	// =========================

	authHandler := handlers.AuthHandler{
		UserRepo: userRepo,
	}

	// =========================
	// DASHBOARD
	// =========================

	dashboardHandler := handlers.DashboardHandler{
		UserRepo: userRepo,
	}

	// =========================
	// CONSULTATIONS
	// =========================

	consultationRepo := repository.ConsultationRepository{
		DB: userRepo.DB,
	}

	consultationHandler := handlers.ConsultationHandler{
		ConsultationRepo: consultationRepo,
	}

	// =========================
	// REGISTER
	// =========================

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			authHandler.RegisterPage(w, r)
			return
		}

		if r.Method == http.MethodPost {
			authHandler.Register(w, r)
			return
		}

		http.Error(
			w,
			"Method not allowed",
			http.StatusMethodNotAllowed,
		)
	})

	// =========================
	// LOGIN
	// =========================

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			authHandler.LoginPage(w, r)
			return
		}

		if r.Method == http.MethodPost {
			authHandler.Login(w, r)
			return
		}

		http.Error(
			w,
			"Method not allowed",
			http.StatusMethodNotAllowed,
		)
	})

	// =========================
	// DASHBOARD
	// =========================

	http.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			http.Error(
				w,
				"Method not allowed",
				http.StatusMethodNotAllowed,
			)
			return
		}

		dashboardHandler.Dashboard(w, r)
	})

	// =========================
	// LOGOUT
	// =========================

	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			http.Error(
				w,
				"Method not allowed",
				http.StatusMethodNotAllowed,
			)
			return
		}

		dashboardHandler.Logout(w, r)
	})

	// =========================
	// CONSULTING SERVICES
	// =========================

	http.HandleFunc("/services", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			http.Error(
				w,
				"Method not allowed",
				http.StatusMethodNotAllowed,
			)
			return
		}

		http.ServeFile(
			w,
			r,
			"templates/services.html",
		)
	})

	// =========================
	// REQUEST CONSULTATION
	// =========================

	http.HandleFunc("/consultation", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			consultationHandler.ConsultationPage(w, r)
			return
		}

		if r.Method == http.MethodPost {
			consultationHandler.CreateConsultation(w, r)
			return
		}

		http.Error(
			w,
			"Method not allowed",
			http.StatusMethodNotAllowed,
		)
	})

	// =========================
	// STATIC FILES
	// =========================

	http.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)
}