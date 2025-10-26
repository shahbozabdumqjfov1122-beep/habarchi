package bot

import (
	tele "gopkg.in/telebot.v3"
)

func HandlePhoto(c tele.Context) error {
	photo := c.Message().Photo

	step := userState[c.Sender().ID] // ❌ qayta e’lon qilinmagan, globaldan ishlatiladi
	if step == 0 {
		return nil
	}

	ad := userAds[c.Sender().ID]

	switch step {
	case 5:
		ad.Photo = photo

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

		return c.Send(&tele.Photo{
			File:    ad.Photo.File,
			Caption: caption,
		}, menu)
	}
	return nil
}
