package bot

import tele "gopkg.in/telebot.v3"

func ElonJoylash(c tele.Context) error {
	userState[c.Sender().ID] = 1
	userAds[c.Sender().ID] = &Ad{}

	msg := " anme nomi ::"
	return c.Send(msg)
}
