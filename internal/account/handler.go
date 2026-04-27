package account

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/krotesq/strowger/internal/response"
	"github.com/krotesq/strowger/internal/util"
)

type handler struct {
	service *service
}

func newHandler(service *service) *handler {
	return &handler{service: service}
}

func (handler *handler) create(w http.ResponseWriter, r *http.Request) {
	var createDTO createDTO
	if err := util.ParseBody(r.Body, &createDTO); err != nil {
		response.Send(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	account, err := handler.service.create(r.Context(), createDTO.Username, createDTO.Password)
	if err != nil {
		response.Send(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response.Send(w, http.StatusOK, "Account created", toAccountDTO(account))
}

func (handler *handler) findByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	account, err := handler.service.findByID(r.Context(), id)
	if err != nil {
		response.Send(w, http.StatusNotFound, "Account not found", nil)
		return
	}

	response.Send(w, http.StatusOK, "Account found", toAccountDTO(account))
}

func (handler *handler) deleteByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	account, err := handler.service.deleteByID(r.Context(), id)
	if err != nil {
		response.Send(w, http.StatusNotFound, "Account not found", nil)
		return
	}

	response.Send(w, http.StatusOK, "Account deleted", toAccountDTO(account))
}

func (handler *handler) deactivateByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	account, err := handler.service.deactivateByID(r.Context(), id)
	if err != nil {
		response.Send(w, http.StatusNotFound, "Account not found", nil)
		return
	}

	response.Send(w, http.StatusOK, "Account deactivated", toAccountDTO(account))
}

func (handler *handler) activateByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	account, err := handler.service.activateByID(r.Context(), id)
	if err != nil {
		response.Send(w, http.StatusNotFound, "Account not found", nil)
		return
	}

	response.Send(w, http.StatusOK, "Account activated", toAccountDTO(account))
}

func (handler *handler) login(w http.ResponseWriter, r *http.Request) {
	var loginDTO loginDTO
	if err := util.ParseBody(r.Body, &loginDTO); err != nil {
		response.Send(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	account, s, err := handler.service.login(r.Context(), loginDTO.Username, loginDTO.Password)
	if err != nil {
		response.Send(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response.SendWithSimpleCookie(w, http.StatusOK, "Account logged in.", toAccountDTO(account), "access_token", s)
}

func (handler *handler) logout(w http.ResponseWriter, r *http.Request) {
	response.SendWithSimpleCookie(w, http.StatusOK, "Account logged out.", nil, "access_token", "")
}