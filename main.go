package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatalf("Error during initialization bot: %s", err.Error())
	}
	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		switch update.Message.Text {
		case "/start":
			msg.Text = `Что умеет этот бот?
Этот бот - интерактивное резюме его создателя. 
Создержит нескольк разделов и умеет переключаться между ними.`
			msg.ReplyMarkup = mainKeyboard
		case "Обо мне":
			msg.Text = `IT и Linux энтузиаст. Программирование - это мое хобби, занимаюсь своими проектами во внеучебное время, также люблю изучать мир декстопного Linux и кастомизировать свою систему.
				Могу быстро обучится любой технологии/библиотеки/фреймворку.
				В поисках первого опыта реальной работы, потому что небольшие пет-проекты перестали быть интересными, хочется стать частью более масштабного и сложного проекта.
				Готов к работе в офисе либо удаленно.
				Занимаюсь спортом, играю в шахматы, не женат.
				`
			msg.ReplyMarkup = aboutKeyboard
		case "Проекты":
			msg.Text = `Сервис автоматизации игрового процесса
				Количество человек: 2
				Стек технологий: Go, Fiber, PostreSQL, Docker + Docker Compose
				Сервис позволяет автоматизировать рутинные действия в одной онлайн-игре. Поддерживает работу с несколькими аккаунтами одновременно, имеет сайт для взаимодействия с пользователем.
				В ходе разработки научился применять на практике инструменты конкурентного программирования Go: горутины, каналы, пакет sync стандартной библиотеки, улучшил навыки написания Web API на Go при помощи фреймворка Fiber и использования базы данных в проекте

				Модель распределенного файлового хранилища
				Количество человек: 1
				Стек технологий: Go, Fiber, PostreSQL, Docker + Docker Compose
				Сервис предоставляет CLI интерфейс и Web API для сохранения, получения и удаления файлов в распределенном хранилище. Оно представляет собой множество файловых нод, работу которых организует центральный сервер. Общение между нодами может происходить как по протоколу HTTP, так и на чистом TCP, на выбор пользователя.
				Познакомился с фреймворком Fiber, инструментами стандартной библиотеки для работы с базами данных при помощи сторонних драйверов (pgx).

				Сервис организаций турниров
				Количество человек: 3 команды по 3 человека
				Роль: тимлид команды микросервиса ставок
				Стек технологий: C#, ASP.NET, MS SQL, Entity Framework, Docker
				Сервис позволяет создавать, запускать турниры по различным сеткам(круговая, швейцарская и т. д.), есть система регистрации пользователей, а также возможность ставить ставки на исход матча/турнира.
				Первый опыт разработки бекенда в виде микросервисов, научился использовать базу данных в проектах, писать и отлаживать взаимодействие микросервисов. Первый опыт использования контейнеризации.
				Научился оценивать и фрагментировать бизнес-задачи на технические задания, определять сроки выполнения и проводить код-ревью.

				Графический редактор
				Количество человек: 3
				Роль: тимлид
				Стек технологий: С++, Qt6, libdeflate
				Графический редактор, написанный с нуля, за исключением использования стандартных виджетов Qt6 и кастомного, принимающего набор значений пикселей в картинке для ее отображения и libdeflate для использования реализации алгоритма сжатия deflate в png формате.
				Спроектировал расширяемую архитектуру проекта около 10000 строк кода, производил декомпозицию лабораторных работ на технические задачи для членов команды.

				Кодогенерация http-клиента по исходному коду сервера
				Количество человек: 1
				Стек технологий: C#, Roslyn API, Python, FastApi
				Был написан кодогенератор, который по исходному коду сервера на FastApi, создает http-клиент с консольным интерфейсом для взаимодействия с сервером и добавляет его в компиляцию
				Познакомился с такой технологией как кодогенерация, улучшил понимание внутреннего строения языка программирования, оперировал семантическим и семантическим деревом проекта.
				`
			msg.ReplyMarkup = projectsKeyboard
		case "Навыки":
			msg.Text = `- Программирование, отладка и профилирование на таких языках как Go, C#, Python, C/C++, JavaScript.
			- Разработка приложений с использованием конкурентного программирования
			- Понимание алгоритмов и структур данных
			- Использование модульных и интеграционных тестов
			- Системы контроля версий Git, хостинги GitHub, Gitlab
			-Контейнеризация при помощи Docker в своих проектах
			- Продвинутый пользователь Linux, написание shell скриптов
			- ООП и паттерны проектирования
			- Базовый уровень верстки на HTML и CSS
			`
			msg.ReplyMarkup = skillsKeyboard
		case "Качества":
			msg.Text = `- Ответственный
			- Энтузиаст
			- Усидчивый
			- Высокая обучаемость
			- Открыт к критике
			- Аналитическое мышление
			- Умение аргументировано высказывать свою точку зрения
			`
			msg.ReplyMarkup = qualitiesKeyboard
		case "Главная":
			msg.Text = "Переходим на главную..."
			msg.ReplyMarkup = mainKeyboard
		case "Тоже главная":
			msg.Text = "Тоже переходим на главную..."
			msg.ReplyMarkup = mainKeyboard
		default:
			msg.Text = "Пожалуйста, выберите команду из списка"
		}
		if _, err := bot.Send(msg); err != nil {
			log.Printf("ERR: %s", err.Error())
		}
	}
}
