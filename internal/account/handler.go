package account

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/krotesq/strowger/internal/util"
	"github.com/krotesq/strowger/internal/response"
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
		util.NewResponse(400, err.Error(), nil).Send(w)
		return
	}
	
	account, err := handler.service.create(r.Context(), createDTO.Username, createDTO.Password)
	if err != nil {
		util.NewResponse(400, err.Error(), nil).Send(w)
		return
	}
	
	util.NewResponse(201, "account created", toAccountDTO(account)).Send(w)
}

func (handler *handler) findByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	
	account, err := handler.service.findByID(r.Context(), id)
	if err != nil {
		util.NewResponse(404, "Account not found.", nil).Send(w)
		return
	}
	
	util.NewResponse(200, "Account found.", toAccountDTO(account)).Send(w)
}

func (handler *handler) deleteByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	
	account, err := handler.service.deleteByID(r.Context(), id)
	if err != nil {
		util.NewResponse(404, "Account not found.", nil).Send(w)
		return
	}
	
	util.NewResponse(200, "Account deleted.", toAccountDTO(account)).Send(w)
}

func (handler *handler) deactivateByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	
	account, err := handler.service.deactivateByID(r.Context(), id)
	if err != nil {
		util.NewResponse(404, "Account not found.", nil).Send(w)
		return
	}
	
	util.NewResponse(200, "Account deactivated.", toAccountDTO(account)).Send(w)
}

func (handler *handler) activateByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	
	account, err := handler.service.activateByID(r.Context(), id)
	if err != nil {
		util.NewResponse(404, "Account not found.", nil).Send(w)
		return
	}
	
	util.NewResponse(200, "Account activated.", toAccountDTO(account)).Send(w)
}


func (handler *handler) login(w http.ResponseWriter, r *http.Request) {
	var loginDTO loginDTO
	if err := util.ParseBody(r.Body, &loginDTO); err != nil {
		util.NewResponse(400, err.Error(), nil).Send(w)
		return
	}
	
	account, s, err := handler.service.login(r.Context(), loginDTO.Username, loginDTO.Password)
	if err != nil {
		util.NewResponse(401, err.Error(), nil).Send(w)
		return
	}

	res := response.NewBuilder(w, nil)
	res.SetCookie("access_token", s)
	res.SetBody(response.Body{
		Message: "Account logged in",
		Data: toAccountDTO(account),
	})
	res.SetHeader("random", "wow")
	res.Send()
}