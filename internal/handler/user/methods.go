package user

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/arifinhermawan/blib/log"
	"github.com/arifinhermawan/probi/internal/handler"
	"github.com/arifinhermawan/probi/internal/usecase/user"
)

func (h *Handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req createUserReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Error(ctx, nil, err, "[CreateUserHandler] json.NewDecoder().Decode() got error")
		handler.SendJSONResponse(w, http.StatusBadRequest, nil, "failed to create user", err)
		return
	}

	err = req.validate()
	if err != nil {
		log.Error(ctx, nil, err, "[CreateUserHandler] req.validate() got error")
		handler.SendJSONResponse(w, http.StatusBadRequest, nil, "failed to create user", err)
		return
	}

	req.Username = strings.ToLower(req.Username)
	req.Email = strings.ToLower(req.Email)

	err = h.user.CreateUser(ctx, user.CreateUserReq(req))
	if err != nil {
		log.Error(ctx, nil, err, "[CreateUserHandler] h.user.CreateUser() got error")
		handler.SendJSONResponse(w, http.StatusInternalServerError, nil, "failed to create user", err)
		return
	}

	handler.SendJSONResponse(w, http.StatusCreated, nil, "success!", nil)
}
