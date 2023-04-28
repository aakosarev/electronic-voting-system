package http

import (
	"encoding/json"
	"fmt"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-app/internal/model"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-app/internal/storage"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	userStorage *storage.UserStorage
}

func NewHandler(userStorage *storage.UserStorage) *Handler {
	return &Handler{
		userStorage: userStorage,
	}
}

type session struct {
	username  int32
	expiresAt time.Time
}

func (s session) isExpired() bool {
	return s.expiresAt.Before(time.Now())
}

var sessions = map[string]session{}

func (h *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, "/login", h.Login)
	router.HandlerFunc(http.MethodPost, "/initial_login", h.InitialLogin)
	router.HandlerFunc(http.MethodGet, "/logout", h.Logout)
	router.HandlerFunc(http.MethodPost, "/refresh", h.Refresh)
	router.GET("/force_enter_details/:user_id", h.ForceEnterDetails)
}

func (h *Handler) ForceEnterDetails(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value

	userSession, exists := sessions[sessionToken]
	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if userSession.isExpired() {
		delete(sessions, sessionToken)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	userIDstr := ps.ByName("user_id")

	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	forceEnterDetails, err := h.userStorage.GetForceEnterDetailsByUsername(r.Context(), int32(userID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"force_enter_details": %t}`, forceEnterDetails)))

}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var loginUser model.LoginUser
	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	passwordHash, err := h.userStorage.GetPasswordHashByUsername(r.Context(), loginUser.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(loginUser.Password)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(1 * time.Hour)

	sessions[sessionToken] = session{
		username:  loginUser.Username,
		expiresAt: expiresAt,
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: expiresAt,
	})
}

func (h *Handler) InitialLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value

	userSession, exists := sessions[sessionToken]
	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if userSession.isExpired() {
		delete(sessions, sessionToken)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var ilu model.InitialLoginUser
	err = json.NewDecoder(r.Body).Decode(&ilu)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if ilu.NewPassword != ilu.ReenteredPassword {
		w.WriteHeader(http.StatusBadRequest)
		// TODO w.Write error
		return
	}

	passwordHash, err := h.userStorage.GetPasswordHashByUsername(r.Context(), userSession.username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(ilu.OldPassword)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		// TODO w.Write error
		return
	}

	newPasswordHash, err := bcrypt.GenerateFromPassword([]byte(ilu.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.userStorage.AddDetails(r.Context(), userSession.username, string(newPasswordHash), ilu.Email, ilu.FirstName, ilu.SecondName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sessionToken := c.Value
	delete(sessions, sessionToken)

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Now(),
	})
}

func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sessionToken := c.Value

	userSession, exists := sessions[sessionToken]
	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if userSession.isExpired() {
		delete(sessions, sessionToken)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	newSessionToken := uuid.NewString()
	expiresAt := time.Now().Add(1 * time.Hour)

	sessions[newSessionToken] = session{
		username:  userSession.username,
		expiresAt: expiresAt,
	}

	delete(sessions, sessionToken)

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   newSessionToken,
		Expires: expiresAt,
	})
}
