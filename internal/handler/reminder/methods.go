package reminder

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/arifinhermawan/blib/log"
	"github.com/arifinhermawan/blib/tracer"
	"github.com/arifinhermawan/probi/internal/handler"
	"github.com/arifinhermawan/probi/internal/lib/auth"
	internalTime "github.com/arifinhermawan/probi/internal/lib/time"
	"github.com/arifinhermawan/probi/internal/usecase/reminder"
)

func (h *Handler) CreateReminderHandler(w http.ResponseWriter, r *http.Request) {
	ctx, txn := tracer.StartHTTPTransaction(r.Context(), tracer.Handler+"CreateReminderHandler", r)
	w = txn.SetWebResponse(w)
	defer txn.End()

	var req createReminderReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Error(ctx, nil, err, "[CreateReminderHandler] json.NewDecoder().Decode() got error")
		handler.SendJSONResponse(w, http.StatusBadRequest, nil, "failed to create reminder", err)
		return
	}

	err = validate(req)
	if err != nil {
		log.Error(ctx, nil, err, "[CreateReminderHandler] validate() got error")
		handler.SendJSONResponse(w, http.StatusBadRequest, nil, "failed to create reminder", err)
		return
	}

	startDate, err := time.Parse(internalTime.DateFormat, req.StartDate)
	if err != nil {
		log.Error(ctx, nil, err, "[CreateReminderHandler] time.Parse() got error")
		handler.SendJSONResponse(w, http.StatusBadRequest, nil, "failed to create reminder", err)
		return
	}

	endDate := time.Time{}
	if req.EndDate != "" {
		endDate, err = time.Parse(internalTime.DateFormat, req.EndDate)
		if err != nil {
			log.Error(ctx, nil, err, "[CreateReminderHandler] time.Parse() got error")
			handler.SendJSONResponse(w, http.StatusBadRequest, nil, "failed to create reminder", err)
			return
		}
	}

	if req.Frequency == "" {
		endDate = startDate
	}

	if req.Interval < 0 {
		req.Interval = 0
	}

	userID := ctx.Value(auth.ContextKeyUserID).(int64)
	err = h.reminder.CreateReminder(ctx, reminder.CreateReminderReq{
		UserID:    userID,
		Title:     req.Title,
		Frequency: strings.ToUpper(req.Frequency),
		Interval:  req.Interval,
		StartDate: startDate,
		EndDate:   endDate,
	})
	if err != nil {
		log.Error(ctx, nil, err, "[CreateReminderHandler] h.reminder.CreateReminder() got error")
		handler.SendJSONResponse(w, http.StatusInternalServerError, nil, "failed to create reminder", err)
		return
	}

	handler.SendJSONResponse(w, http.StatusCreated, nil, handler.SuccessMessage, nil)
}

func (h *Handler) GetUserActiveReminderHandler(w http.ResponseWriter, r *http.Request) {
	ctx, txn := tracer.StartHTTPTransaction(r.Context(), tracer.Handler+"GetUserActiveReminderHandler", r)
	w = txn.SetWebResponse(w)
	defer txn.End()

	userID := ctx.Value(auth.ContextKeyUserID).(int64)
	res, err := h.reminder.GetUserActiveReminder(ctx, userID)
	if err != nil {
		log.Error(ctx, nil, err, "[GetUserActiveReminderHandler] h.reminder.GetUserActiveReminder() got error")
		handler.SendJSONResponse(w, http.StatusInternalServerError, nil, "failed to fetch reminder", err)
		return
	}

	reminders := make([]Reminder, len(res))
	for idx, reminder := range res {
		reminders[idx] = Reminder(reminder)
	}

	handler.SendJSONResponse(w, http.StatusOK, reminders, handler.SuccessMessage, nil)
}
