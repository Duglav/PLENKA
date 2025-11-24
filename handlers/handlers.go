package handlers

import (
	"fmt"
	"plenkabot/config"
	"plenkabot/keyboard"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleUpdate(update tgbotapi.Update) {
	if update.Message != nil {
		handleMessage(update)
	}
}

func handleMessage(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	text := update.Message.Text

	switch text {
	case "/start":
		msg := tgbotapi.NewMessage(chatID, "🎬Добро пожаловать в ПЛЁНКА!\nВыберите нужный раздел:")
		msg.ReplyMarkup = keyboard.MainMenu()
		config.Bot.Send(msg)

	case "🎬Вернуться в меню":
		msg := tgbotapi.NewMessage(chatID, "🎬Добро пожаловать в ПЛЁНКА!\nВыберите нужный раздел:")
		msg.ReplyMarkup = keyboard.MainMenu()
		config.Bot.Send(msg)

	case "📸Услуги":
		msg := tgbotapi.NewMessage(chatID, "⭐️Выберите нужную услугу:")
		msg.ReplyMarkup = keyboard.ServiceMenu()
		config.Bot.Send(msg)

	case "📂Тарифы":
		msg := tgbotapi.NewMessage(chatID, "⭐️Выберите нужный тариф:")
		msg.ReplyMarkup = keyboard.ModeMenu()
		config.Bot.Send(msg)

	case "⭐️Стандарт":
		msg := tgbotapi.NewMessage(chatID,
			`💼<b>Тариф СТАНДАРТ</b> — <b>$210</b>

			📽3 рилс (Instagram Reels)
			📝4 поста (Instagram Posts)
			📲10 историй (Instagram Stories)
			🧑‍💼Ведение аккаунта

			<b>Указанная цена включает услуги на месяц</b>`)
		msg.ParseMode = "HTML"
		config.Bot.Send(msg)

	case "📸Съемка рекламы":
		msg := tgbotapi.NewMessage(chatID,
			`🎬<b>Услуга: СЪЕМКА РЕКЛАМЫ</b>

			📸Мы создаем креативные и стильные рекламные ролики для твоего бренда или продукта.
			💡Сценарий, съёмка, монтаж — всё под ключ.
			📍Адаптируем под Reels, Stories, TikTok или рекламу в ленте.

			💰<b>Цена рассчитывается индивидуально</b>, в зависимости от объёма и сложности.`)
		msg.ParseMode = "HTML"
		config.Bot.Send(msg)

	case "🤖Создание бота":
		msg := tgbotapi.NewMessage(chatID,
			`🤖<b>Услуга: СОЗДАНИЕ БОТА</b>

			📲Telegram-боты для бизнеса, автоматизации и удобного общения с клиентами.
			⚙️Простые решения: меню, заявки, бронирование, чат-боты.
			📈Разработка и поддержка под твои задачи.

			💰<b>Цена рассчитывается индивидуально</b>, в зависимости от функционала.`)
		msg.ParseMode = "HTML"
		config.Bot.Send(msg)

	case "⭐️Премиум":
		msg := tgbotapi.NewMessage(chatID,
			`💼<b>Тариф ПРЕМИУМ</b> — <b>$450</b>

			📽5 рилс (Instagram Reels)
			📝7 постов (Instagram Posts)
			📲16 историй (Instagram Stories)
			🧑‍💼Ведение аккаунта
			🎯Настройка таргета

			<b>Указанная цена включает услуги на месяц</b>`)
		msg.ParseMode = "HTML"
		config.Bot.Send(msg)

	case "⭐️Плёнка":
		msg := tgbotapi.NewMessage(chatID,
			`💼<b>Тариф ПЛЁНКА</b> — <b>$700</b>

			📽8 рилс (Instagram Reels)
			📝10 постов (Instagram Posts)
			📲22 истории (Instagram Stories)
			🧑‍💼Ведение аккаунта
			🎯Настройка таргета
			🤖Настройка бота и его поддержка

			<b>Указанная цена включает услуги на месяц</b>`)
		msg.ParseMode = "HTML"
		config.Bot.Send(msg)

	case "💬Обратная связь":
		msg := tgbotapi.NewMessage(chatID, "✅Ваша заявка отправлена! Совсем скоро мы с вами свяжемся!")
		config.Bot.Send(msg)

		adminMsg := fmt.Sprintf(
			"📬<b>Новая заявка на обратную связь</b>\n\n👤Имя: %s\n🔗Username: @%s",
			update.Message.From.FirstName,
			update.Message.From.UserName,
		)

		groupID := int64(-1003319244089)
		msgToGroup := tgbotapi.NewMessage(groupID, adminMsg)
		msgToGroup.ParseMode = "HTML"
		config.Bot.Send(msgToGroup)
	}
}
