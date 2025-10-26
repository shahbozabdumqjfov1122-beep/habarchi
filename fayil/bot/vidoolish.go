package bot

import (
	tele "gopkg.in/telebot.v3"
)

func HandleVideo(c tele.Context) error {
	video := c.Message().Video

	step := userState[c.Sender().ID]
	if step == 0 {
		return nil
	}

	ad := userAds[c.Sender().ID]

	switch step {
	case 5:
		ad.Video = video

		userState[c.Sender().ID] = 0

		menu := &tele.ReplyMarkup{ResizeKeyboard: true}
		btnOK := menu.Text("✅ Tasdiqlayman")
		btnCancle := menu.Text("❌ Bekor qilish")

		menu.Reply(menu.Row(btnOK), menu.Row(btnCancle))

		caption :=
			"📝 Anime nomi: " + ad.Content + "\n\n" +
				"📚 Janri: " + ad.Phone + "\n\n" +
				"🎬 Qisimlar soni: " + ad.Salary + "\n\n" +
				"🔖 anime ko'di: " + ad.Comment + "\n\n" +
				"📩 bizni bot : @anmelaruzb_bot"

		return c.Send(&tele.Video{
			File:    ad.Video.File,
			Caption: caption,
		}, menu)
	}
	return c.Send("bot ishaga tushti")
}
