package authentication

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/arifinhermawan/blib/log"
	"github.com/arifinhermawan/probi/internal/handler"
	"github.com/arifinhermawan/probi/internal/lib/auth"
	"github.com/arifinhermawan/probi/internal/lib/errors"
	"github.com/arifinhermawan/probi/internal/usecase/authentication"
)

func (h *Handler) LogInHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req logInReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Error(ctx, nil, err, "[LogInHandler] json.NewDecoder().Decode() got error")
		handler.SendJSONResponse(w, http.StatusBadRequest, nil, "failed to log in", err)
		return
	}

	err = req.validate()
	if err != nil {
		log.Error(ctx, nil, err, "[LogInHandler] req.validate() got error")
		handler.SendJSONResponse(w, http.StatusBadRequest, nil, "failed to log in", err)
		return
	}

	req.Input = strings.ToLower(req.Input)
	username := req.Input
	email := ""

	if strings.Contains(req.Input, "@") {
		username = ""
		email = req.Input
	}

	userID, jwt, err := h.auth.LogIn(ctx, authentication.LogInReq{
		Email:    email,
		Username: username,
		Password: req.Password,
	})
	if err != nil {
		log.Error(ctx, nil, err, "[LogInHandler] h.auth.LogIn() got error")
		if err == errors.ErrUserNotFound {
			handler.SendJSONResponse(w, http.StatusBadRequest, nil, "user not exists", err)
			return
		} else if err == errors.ErrPassswordNotMatch {
			handler.SendJSONResponse(w, http.StatusBadRequest, nil, "password not match", err)
			return
		}

		handler.SendJSONResponse(w, http.StatusInternalServerError, nil, "failed to log in", err)
		return
	}

	handler.SendJSONResponse(w, http.StatusOK, logInResponse{
		UserID: userID,
		Token:  jwt,
	}, "success!", nil)
}

func (h *Handler) LogOutHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID := ctx.Value(auth.ContextKeyUserID).(int64)
	err := h.auth.LogOut(ctx, userID)
	if err != nil {
		log.Error(ctx, map[string]interface{}{
			"user_id": userID,
		}, err, "[LogOutHandler] h.auth.LogOut() got error")
		handler.SendJSONResponse(w, http.StatusBadRequest, nil, "failed to log out", err)

		return
	}

	handler.SendJSONResponse(w, http.StatusOK, nil, "success!", nil)
}
