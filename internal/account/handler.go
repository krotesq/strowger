package account

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/krotesq/strowger/internal/util"
)

type handler struct {
	service *service
}

func newHandler(service *service) *handler {
	return &handler{service: service}
}

func (handler *handler) findByUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	account, err := handler.service.findByUsername(r.Context(), username)
	if err != nil {
		util.NewResponse(404, "account not found", nil).Send(w)
		return
	}
	
	util.NewResponse(200, "found account", toAccountDTO(account)).Send(w)
}


func (handler *handler) login(w http.ResponseWriter, r *http.Request) {
	var loginData loginDto
	if err := util.ParseBody(r.Body, &loginData); err != nil {
		util.NewResponse(400, err.Error(), nil).Send(w)
		return
	}
	
	s, err := handler.service.login(r.Context(), loginData.Username, loginData.Password)
	if err != nil {
		util.NewResponse(401, err.Error(), nil).Send(w)
		return
	}
	
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    s,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	util.NewResponse(200, "logged in", nil).Send(w)
}