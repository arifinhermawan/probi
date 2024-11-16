package reminder

import (
	"context"

	"github.com/arifinhermawan/probi/internal/usecase/reminder"
)

type reminderUseCaseProvider interface {
	CreateReminder(ctx context.Context, req reminder.CreateReminderReq) error
}

type Handler struct {
	reminder reminderUseCaseProvider
}

func NewHandler(reminder reminderUseCaseProvider) *Handler {
	return &Handler{
		reminder: reminder,
	}
}
