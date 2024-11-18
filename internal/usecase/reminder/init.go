package reminder

import (
	"context"
	"time"

	"github.com/arifinhermawan/probi/internal/lib/configuration"
	"github.com/arifinhermawan/probi/internal/service/reminder"
	"github.com/arifinhermawan/probi/internal/service/rmq"
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

type rmqServiceProvider interface {
	PublishMessage(ctx context.Context, req rmq.PublishMessageReq) error
}

type UseCase struct {
	lib      libProvider
	reminder reminderServiceProvider
	rmq      rmqServiceProvider
}

func NewUseCase(lib libProvider, reminder reminderServiceProvider, rmq rmqServiceProvider) *UseCase {
	return &UseCase{
		lib:      lib,
		reminder: reminder,
		rmq:      rmq,
	}
}
