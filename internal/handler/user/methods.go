package user

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/arifinhermawan/blib/log"
	"github.com/arifinhermawan/blib/tracer"
	"github.com/arifinhermawan/probi/internal/handler"
	"github.com/arifinhermawan/probi/internal/lib/auth"
	"github.com/arifinhermawan/probi/internal/lib/errors"
	"github.com/arifinhermawan/probi/internal/usecase/user"
	"github.com/gorilla/mux"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func (h *Handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx, txn := tracer.StartHTTPTransaction(r.Context(), tracer.Handler+"CreateUserHandler", r)
	w = txn.SetWebResponse(w)
	defer txn.End()

	var req createUserReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Error(ctx, nil, err, "[CreateUserHandler] json.NewDecoder().Decode() got error")
		handler.SendJSONResponse(w, http.StatusBadRequest, nil, "failed to create user", err)
		return
	}

	err = validate(req)
	if err != nil {
		log.Error(ctx, nil, err, "[CreateUserHandler] validate() got error")
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

	handler.SendJSONResponse(w, http.StatusCreated, nil, handler.SuccessMessage, nil)
}

func (h *Handler) GetUserDetailsHandler(w http.ResponseWriter, r *http.Request) {
	ctx, txn := tracer.StartHTTPTransaction(r.Context(), tracer.Handler+"GetUserDetailsHandler", r)
	w = txn.SetWebResponse(w)
	defer txn.End()

	vars := mux.Vars(r)
	userIDstr := vars["user_id"]
	userID, err := strconv.ParseInt(userIDstr, 10, 64)
	if err != nil {
		log.Error(ctx, nil, err, "[GetUserDetailsHandler] strconv.ParseInt() got error")
		handler.SendJSONResponse(w, http.StatusBadRequest, nil, "failed to get user", err)
		return
	}
	res, err := h.user.GetUserDetails(ctx, userID)
	if err != nil {
		log.Error(ctx, nil, err, "[GetUserDetailsHandler] h.user.GetUserDetails() got error")
		if err == errors.ErrUserNotFound {
			handler.SendJSONResponse(w, http.StatusBadRequest, nil, "user not exists", err)
			return
		}

		handler.SendJSONResponse(w, http.StatusInternalServerError, nil, "failed to get user", err)
		return
	}

	user := User{
		ID:          res.ID,
		Username:    res.Username,
		Email:       res.Email,
		DisplayName: res.DisplayName,
	}

	handler.SendJSONResponse(w, http.StatusOK, user, handler.SuccessMessage, nil)
}

func (h *Handler) UpdateUserDetailsHandler(w http.ResponseWriter, r *http.Request) {
	ctx, txn := tracer.StartHTTPTransaction(r.Context(), tracer.Handler+"UpdateUserDetailsHandler", r)
	w = txn.SetWebResponse(w)
	defer txn.End()

	var req updateUserDetailsReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Error(ctx, nil, err, "[UpdateUserDetailsHandler] json.NewDecoder().Decode() got error")
		handler.SendJSONResponse(w, http.StatusBadRequest, nil, "failed to update user", err)
		return
	}

	caser := cases.Title(language.English)
	err = validate(req)
	if err != nil {
		log.Error(ctx, nil, err, "[UpdateUserDetailsHandler] validate() got error")
		handler.SendJSONResponse(w, http.StatusBadRequest, nil, "failed to update user", err)
		return
	}

	userID := ctx.Value(auth.ContextKeyUserID).(int64)
	err = h.user.UpdateUserDetails(ctx, user.UpdateUserDetailsReq{
		UserID:      userID,
		Email:       strings.ToLower(req.Email),
		DisplayName: caser.String(req.DisplayName),
		Username:    req.Username,
	})
	if err != nil {
		log.Error(ctx, nil, err, "[UpdateUserDetailsHandler] h.user.UpdateUserDetails() got error")
		handler.SendJSONResponse(w, http.StatusInternalServerError, nil, "failed to update user", err)
		return
	}

	handler.SendJSONResponse(w, http.StatusNoContent, nil, handler.SuccessMessage, nil)
}
