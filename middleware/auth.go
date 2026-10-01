package middleware

import (
	"net/http"
	"strconv"
)

const sessionCookieName = "thurmanlark_session"

// CreateSession creates a login session for the user.
func CreateSession(w http.ResponseWriter, userID int) {

	http.SetCookie(w, &http.Cookie{
		Name:     sessionCookieName,
		Value:    strconv.Itoa(userID),
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
}

// GetUserID gets the logged-in user's ID from the session.
func GetUserID(r *http.Request) (int, bool) {

	cookie, err := r.Cookie(sessionCookieName)

	if err != nil {
		return 0, false
	}

	userID, err := strconv.Atoi(cookie.Value)

	if err != nil {
		return 0, false
	}

	return userID, true
}

// ClearSession logs the user out.
func ClearSession(w http.ResponseWriter) {

	http.SetCookie(w, &http.Cookie{
		Name:     sessionCookieName,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
}
