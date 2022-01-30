package controllers

import (
	"github.com/dParikesit/dimsBot/models"
	"github.com/dParikesit/dimsBot/utils"
)

func GetAll() []models.Reminder {
	var reminders []models.Reminder
	utils.Db.Find(&reminders)
	return reminders
}

func GetBool(isDone bool) []models.Reminder {
	var reminders []models.Reminder

	if isDone == true {
		utils.Db.Where("done = ?", "true").Find(&reminders)
	} else {
		utils.Db.Where("done = ?", "false").Find(&reminders)
	}

	return reminders
}

func InsertOne(data *models.Reminder) error {
	result := utils.Db.Create(data)
	return result.Error
}
