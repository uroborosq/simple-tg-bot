package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var mainKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Обо мне"),
		tgbotapi.NewKeyboardButton("Проекты"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Навыки"),
		tgbotapi.NewKeyboardButton("Качества"),
	),
)

var aboutKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Проекты"),
		tgbotapi.NewKeyboardButton("Главная"),
	),
)

var projectsKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Навыки"),
		tgbotapi.NewKeyboardButton("Главная"),
	),
)

var skillsKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Качества"),
		tgbotapi.NewKeyboardButton("Главная"),
	),
)
var qualitiesKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Главная"),
		tgbotapi.NewKeyboardButton("Тоже главная"),
	),
)
