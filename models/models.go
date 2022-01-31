package models

import "time"

type Reminder struct {
	UserId      string    `json:"user_id"`
	Schedule    time.Time `json:"schedule"`
	Description string    `json:"description"`
	Done        bool      `json:"done,omitempty"`
}

type ReminderTemp struct {
	UserId      string `json:"user_id"`
	Schedule    string `json:"schedule"`
	Description string `json:"description"`
	Done        bool   `json:"done,omitempty"`
}

func (r *ReminderTemp) ConvTime() (reminder Reminder, err error) {
	reminder.UserId = r.UserId
	reminder.Description = r.Schedule
	reminder.Done = r.Done
	reminder.Schedule, err = time.Parse("02/01/2006", r.Schedule)
	return reminder, err
}
