package fayil

import (
	tele "gopkg.in/telebot.v3"
	bot2 "kanalga_habar_yuborish_bot/fayil/bot"
	"log"
	"time"
)

func Bot() {
	pref := tele.Settings{
		Token:  "7802537780:AAFjVntjUzJCS7uJ4dd_oxASXXKg-WMzdO4", // tokeningizni shu yerga yozing
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}
	b.Handle(tele.OnText, bot2.HandleText)
	b.Handle(tele.OnPhoto, bot2.HandlePhoto)
	b.Handle(tele.OnVideo, bot2.HandleVideo)
	log.Println("✅ Bot ishga tushdi!")

	b.Start()
}
