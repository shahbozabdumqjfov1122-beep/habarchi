package bot

import (
	"fmt"
	"log"

	tele "gopkg.in/telebot.v3"
)

var bChannelID = int64(-1003050934981)
var correctPassword = "alfa123" // 🔑 Parol
type Ad struct {
	Content string
	Phone   string
	Salary  string
	Comment string
	Photo   *tele.Photo
	Video   *tele.Video
}

var userState = make(map[int64]int)
var userAds = make(map[int64]*Ad)
var authorized = make(map[int64]bool) // ✅ Kim parol kiritganini saqlash
func ElonMenu() *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{}
	btnBot := menu.URL("📩 Bizni bot", "https://t.me/anmelaruzb_bot")
	menu.Inline(
		menu.Row(btnBot),
	)
	return menu
}
func HandleText(c tele.Context) error {
	text := c.Text()
	uid := c.Sender().ID
	// --- 1. Start bosilganda faqat parol so‘rashi ---
	if text == "/start" {
		authorized[uid] = false
		return c.Send("🔑 Botdan foydalanish uchun parolni kiriting:")
	}
	// --- 2. Agar hali parol kiritilmagan bo‘lsa ---
	if !authorized[uid] {
		if text == correctPassword {
			authorized[uid] = true
			menu := &tele.ReplyMarkup{ResizeKeyboard: true}
			btnElon := menu.Text("➕ E’lon joylash")
			btnCancel := menu.Text("❌ Bekor qilish")
			menu.Reply(menu.Row(btnElon), menu.Row(btnCancel))
			return c.Send("✅ Parol to‘g‘ri! Endi botdan foydalanishingiz mumkin.", menu)
		} else {
			return c.Send("❌ Parol noto‘g‘ri. Qayta urinib ko‘ring.")
		}
	}
	// --- 3. Faqat paroldan keyin ishlaydigan qism ---
	switch text {
	case "➕ E’lon joylash":
		return ElonJoylash(c)
	case "✅ Tasdiqlayman":
		ad, ok := userAds[uid]
		if !ok || ad == nil || ad.Content == "" || ad.Phone == "" {
			return c.Send("⚠️ Siz hali e’lon kiritmagansiz. Avval '➕ E’lon joylash' tugmasini bosing.")
		}
		caption :=
			"📝 Anime nomi: " + ad.Content + "\n\n" +
				"📚 Janri: " + ad.Phone + "\n\n" +
				"🎬 Qisimlar soni: " + ad.Salary + "\n\n" +
				"🔖 anime ko'di: " + ad.Comment

		menu := ElonMenu()
		var M *tele.Message
		var err error
		if ad.Photo != nil {
			M, err = c.Bot().Send(tele.ChatID(bChannelID), &tele.Photo{
				File:    ad.Photo.File,
				Caption: caption,
			}, menu)
		} else if ad.Video != nil {
			M, err = c.Bot().Send(tele.ChatID(bChannelID), &tele.Video{
				File:    ad.Video.File,
				Caption: caption,
			}, menu)
		} else {
			M, err = c.Bot().Send(tele.ChatID(bChannelID), caption, menu)
		}
		if err != nil {
			log.Println("Kanalga yuborishda xatolik:", err)
			return c.Send("❌ Kanalga yuborishda xatolik yuz berdi.")
		}
		link := fmt.Sprintf("https://t.me/c/3050934981/%d", M.ID)
		return c.Send(
			fmt.Sprintf("✅ E’lon muvaffaqiyatli kanalga joylandi!\n👉 <a href='%s'>Kanalga o‘tish</a>", link),
			&tele.SendOptions{ParseMode: tele.ModeHTML},
		)
	case "❌ Bekor qilish":
		menu := &tele.ReplyMarkup{ResizeKeyboard: true}
		btnElonJoylash := menu.Text("➕ E’lon joylash")
		menu.Reply(menu.Row(btnElonJoylash))
		userState[uid] = 0
		userAds[uid] = &Ad{}
		return c.Send("❌ Jarayon bekor qilindi", menu)
	}
	step := userState[uid]
	if step == 0 {
		return nil
	}
	ad := userAds[uid]
	switch step {
	case 1:
		ad.Content = text
		userState[uid] = 2
		return c.Send("📚 Janrini kiriting:")
	case 2:
		ad.Phone = text
		userState[uid] = 3
		return c.Send("🎬 Qismlar sonini kiriting:")
	case 3:
		ad.Salary = text
		userState[uid] = 4
		return c.Send("🔖 anime ko'di:")
	case 4:
		ad.Comment = text
		userState[uid] = 5
		return c.Send("📷 Endi rasm yoki 🎥 video yuboring:")
	}
	return nil
}
