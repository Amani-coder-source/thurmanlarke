package handlers

import (
	"html/template"
	"net/http"

	"Thurmanlark/middleware"
	"Thurmanlark/models"
	"Thurmanlark/repository"
)

type ConsultationHandler struct {
	ConsultationRepo repository.ConsultationRepository
}

// ConsultationPage displays the consultation form.
func (h ConsultationHandler) ConsultationPage(w http.ResponseWriter, r *http.Request) {

	_, loggedIn := middleware.GetUserID(r)

	if !loggedIn {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles(
		"templates/consultation.html",
	)

	if err != nil {
		http.Error(
			w,
			"Unable to load consultation page",
			http.StatusInternalServerError,
		)
		return
	}

	err = tmpl.Execute(w, nil)

	if err != nil {
		http.Error(
			w,
			"Unable to display consultation page",
			http.StatusInternalServerError,
		)
		return
	}
}

// CreateConsultation saves a consultation request.
func (h ConsultationHandler) CreateConsultation(w http.ResponseWriter, r *http.Request) {

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

	service := r.FormValue("service")
	businessName := r.FormValue("business_name")
	businessType := r.FormValue("business_type")
	description := r.FormValue("description")
	budget := r.FormValue("budget")
	preferredDate := r.FormValue("preferred_date")

	if service == "" ||
		businessName == "" ||
		businessType == "" ||
		description == "" ||
		budget == "" ||
		preferredDate == "" {

		http.Error(
			w,
			"All fields are required",
			http.StatusBadRequest,
		)
		return
	}

	consultation := models.Consultation{
		UserID:         userID,
		Service:        service,
		BusinessName:  businessName,
		BusinessType:  businessType,
		Description:   description,
		Budget:        budget,
		PreferredDate: preferredDate,
		Status:        "pending",
	}

	err = h.ConsultationRepo.CreateConsultation(
		consultation,
	)

	if err != nil {
		http.Error(
			w,
			"Unable to submit consultation request",
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