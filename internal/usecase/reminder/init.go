package reminder

import (
	"context"
	"time"

	"github.com/arifinhermawan/probi/internal/lib/configuration"
	"github.com/arifinhermawan/probi/internal/service/reminder"
)

type libProvider interface {
	GetConfig() *configuration.AppConfig
	GetTimeGMT7() time.Time
}

type reminderServiceProvider interface {
	CreateReminder(ctx context.Context, req reminder.CreateReminderReq) error
	GetDueReminderIDs(ctx context.Context) ([]int64, error)
	GetUserActiveReminder(ctx context.Context, userID int64) ([]reminder.Reminder, error)
	UpdateReminder(ctx context.Context, req reminder.UpdateReminderReq) error
}

type UseCase struct {
	lib      libProvider
	reminder reminderServiceProvider
}

func NewUseCase(lib libProvider, reminder reminderServiceProvider) *UseCase {
	return &UseCase{
		lib:      lib,
		reminder: reminder,
	}
}
