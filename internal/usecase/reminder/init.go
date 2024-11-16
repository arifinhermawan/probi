package reminder

import (
	"context"

	"github.com/arifinhermawan/probi/internal/service/reminder"
)

type reminderServiceProvider interface {
	CreateReminder(ctx context.Context, req reminder.CreateReminderReq) error
	GetUserActiveReminder(ctx context.Context, userID int64) ([]reminder.Reminder, error)
}

type UseCase struct {
	reminder reminderServiceProvider
}

func NewUseCase(reminder reminderServiceProvider) *UseCase {
	return &UseCase{
		reminder: reminder,
	}
}
