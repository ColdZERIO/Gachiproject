package main

import (
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/telebot.v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env не найден")
	}

	token := os.Getenv("TELEGRAM_API_KEY")

	if token == "" {
		log.Fatal("TELEGRAM_API_KEY is empty")
	}

	// Настройки предзагрузки для бота (опционально)
	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	// Создаем экземпляр бота
	myBot, err := telebot.NewBot(pref)

	if err != nil {
		log.Fatal(err)
		return
	}

	// Обработчик команды /start
	myBot.Handle("/start", func(c telebot.Context) error {
		return c.Send("Ну поздравляю, ты запустил ЗАЛУПНЫЙ бот да еще и на v4.")
	})

	giflist := []string{"https://tenor.com/ru/view/gachimuchi-gif-20116037",
		"https://tenor.com/ru/view/gachi-gachi-muchi-gym-gif-26208724",
		"https://tenor.com/ru/view/ricardo-milos-dance-flex-meme-gif-13919144",
		"https://tenor.com/ru/view/memeblog-gachi-gif-18207743",
		"https://tenor.com/ru/view/yummy-mmm-an-oticed-gif-26590512"}

	myBot.Handle(telebot.OnText, func(a telebot.Context) error {
		text := strings.ToLower(a.Text())
		if strings.Contains(text, "время") {
			return a.Send("Че, время подсказать? да вот: " + time.Now().UTC().String())
		}

		if strings.Contains(text, "гачи") {
			randnum := rand.Intn(len(giflist) - 1)
			log.Println(giflist[randnum])
			gif := &telebot.Animation{
				File: telebot.FromURL(giflist[randnum]),
			}
			return a.Send(gif)
		}

		if strings.Contains(text, "япошка") {
			file := &telebot.Audio{
				File: telebot.FromDisk("./mp3/Yamatekudasai.mp3"),
				Caption: "🎵 Наслаждайся",
				Title: "ЯМАТЕКУДАСАЙ!!!",
			}

			return a.Send(file)
		}

		log.Println(a.Text())
		return a.Send("Ты ляпнул: " + a.Text())
	})

	log.Println("Залупа запущена (v4)")
	myBot.Start()
}
