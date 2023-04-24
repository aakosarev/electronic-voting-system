package http

import (
	"encoding/json"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-app/internal/model"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-app/internal/storage"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
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
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var loginUser model.LoginUser
	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	password, err := h.userStorage.GetPassByUsername(r.Context(), loginUser.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if password != loginUser.Password {
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

	password, err := h.userStorage.GetPassByUsername(r.Context(), userSession.username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if password != ilu.OldPassword {
		w.WriteHeader(http.StatusBadRequest)
		// TODO w.Write error
		return
	}

	err = h.userStorage.Update(r.Context(), userSession.username, ilu.NewPassword, ilu.Email, ilu.FirstName, ilu.SecondName)
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
	log.Println("out", sessionToken)
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
