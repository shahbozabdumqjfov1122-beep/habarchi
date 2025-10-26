package bot

import tele "gopkg.in/telebot.v3"

func Start(c tele.Context) error {

	menu := &tele.ReplyMarkup{ResizeKeyboard: true}

	btnELON := menu.Text("➕ E’lon joylash")

	menu.Reply(
		menu.Row(btnELON),
	)

	msg := `
👋 Assalomu alaykum!
`

	return c.Send(msg, menu)
}
