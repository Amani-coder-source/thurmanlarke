package handlers

import (
	"html/template"
	"net/http"

	"Thurmanlark/middleware"
	"Thurmanlark/models"
	"Thurmanlark/repository"
)

type CompanyHandler struct {
	CompanyRepo repository.CompanyRepository
}

// CompanyPage displays the company registration form.
func (h CompanyHandler) CompanyPage(w http.ResponseWriter, r *http.Request) {

	_, loggedIn := middleware.GetUserID(r)

	if !loggedIn {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles(
		"templates/company.html",
	)

	if err != nil {
		http.Error(
			w,
			"Unable to load company registration page",
			http.StatusInternalServerError,
		)
		return
	}

	err = tmpl.Execute(w, nil)

	if err != nil {
		http.Error(
			w,
			"Unable to display company registration page",
			http.StatusInternalServerError,
		)
		return
	}
}

// CreateCompany handles a company registration submission.
func (h CompanyHandler) CreateCompany(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(
			w,
			"Method not allowed",
			http.StatusMethodNotAllowed,
		)
		return
	}

	userID, loggedIn := middleware.GetUserID(r)

	if !loggedIn {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
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

	companyName := r.FormValue("company_name")
	companyType := r.FormValue("company_type")
	businessActivity := r.FormValue("business_activity")
	address := r.FormValue("address")
	city := r.FormValue("city")
	country := r.FormValue("country")

	if companyName == "" ||
		companyType == "" ||
		businessActivity == "" ||
		address == "" ||
		city == "" ||
		country == "" {

		http.Error(
			w,
			"All fields are required",
			http.StatusBadRequest,
		)
		return
	}

	company := models.Company{
		UserID:           userID,
		CompanyName:      companyName,
		CompanyType:      companyType,
		BusinessActivity: businessActivity,
		Address:          address,
		City:             city,
		Country:          country,
		Status:           "pending",
	}

	err = h.CompanyRepo.CreateCompany(company)

	if err != nil {
		http.Error(
			w,
			"Unable to submit company registration",
			http.StatusInternalServerError,
		)
		return
	}

	http.Redirect(
		w,
		r,
		"/dashboard",
		http.StatusSeeOther,
	)
}