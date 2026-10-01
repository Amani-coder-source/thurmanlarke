package handlers

import (
	"html/template"
	"log"
	"net/http"

	"Thurmanlark/middleware"
	"Thurmanlark/models"
	"Thurmanlark/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	UserRepo repository.UserRepository
}

// RegisterPage displays the registration page.
func (h AuthHandler) RegisterPage(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/register.html")

	if err != nil {
		http.Error(
			w,
			"Unable to load registration page",
			http.StatusInternalServerError,
		)
		return
	}

	err = tmpl.Execute(w, nil)

	if err != nil {
		http.Error(
			w,
			"Unable to display registration page",
			http.StatusInternalServerError,
		)
		return
	}
}

// Register creates a new client account.
func (h AuthHandler) Register(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(
			w,
			"Method not allowed",
			http.StatusMethodNotAllowed,
		)
		return
	}

	err := r.ParseForm()

	if err != nil {
		http.Error(
			w,
			"Invalid form",
			http.StatusBadRequest,
		)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm_password")

	if name == "" || email == "" || password == "" {
		http.Error(
			w,
			"All fields are required",
			http.StatusBadRequest,
		)
		return
	}

	if password != confirmPassword {
		http.Error(
			w,
			"Passwords do not match",
			http.StatusBadRequest,
		)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		http.Error(
			w,
			"Unable to secure password",
			http.StatusInternalServerError,
		)
		return
	}

	user := models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Role:     "client",
	}

	err = h.UserRepo.CreateUser(user)

	if err != nil {
		log.Println("CreateUser error:", err)

		http.Error(
			w,
			"Unable to create account",
			http.StatusInternalServerError,
		)
		return
	}

	log.Println("User registered successfully:", email)

	http.Redirect(
		w,
		r,
		"/login",
		http.StatusSeeOther,
	)
}

// LoginPage displays the login page.
func (h AuthHandler) LoginPage(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/login.html")

	if err != nil {
		http.Error(
			w,
			"Unable to load login page",
			http.StatusInternalServerError,
		)
		return
	}

	err = tmpl.Execute(w, nil)

	if err != nil {
		http.Error(
			w,
			"Unable to display login page",
			http.StatusInternalServerError,
		)
		return
	}
}

// Login authenticates an existing user.
func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(
			w,
			"Method not allowed",
			http.StatusMethodNotAllowed,
		)
		return
	}

	err := r.ParseForm()

	if err != nil {
		log.Println("Login form error:", err)

		http.Error(
			w,
			"Invalid form",
			http.StatusBadRequest,
		)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	log.Println("Login attempt for:", email)

	if email == "" || password == "" {
		http.Error(
			w,
			"Email and password are required",
			http.StatusBadRequest,
		)
		return
	}

	user, err := h.UserRepo.GetUserByEmail(email)

	if err != nil {
		log.Println("GetUserByEmail error:", err)

		http.Error(
			w,
			"Invalid email or password",
			http.StatusUnauthorized,
		)
		return
	}

	log.Println("User found with ID:", user.ID)

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		log.Println("Password comparison failed:", err)

		http.Error(
			w,
			"Invalid email or password",
			http.StatusUnauthorized,
		)
		return
	}

	log.Println("Login successful for user ID:", user.ID)

	// Create a session for the authenticated user.
	middleware.CreateSession(w, user.ID)

	http.Redirect(
		w,
		r,
		"/dashboard",
		http.StatusSeeOther,
	)
}