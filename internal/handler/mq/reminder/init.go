package reminder

import "context"

type reminderUseCaseProvider interface {
	ProcessDueReminder(ctx context.Context) error
}

type Handler struct {
	reminder reminderUseCaseProvider
}

func NewHandler(reminder reminderUseCaseProvider) *Handler {
	return &Handler{
		reminder: reminder,
	}
}