package i18n

import (
	"fmt"
	"strconv"
	"strings"
	"surdo-bot/internal/i18n/lng"
	"surdo-bot/internal/types"
)

func monospace(s string) string {
	return "`" + s + "`"
}

var Buttons = map[string]map[lng.Language]string{
	"PLACEHOLDER": {
		lng.RUSSIAN:   "TODO",
		lng.ENGLISH:   "TODO",
		lng.UZBEKLAT:  "TODO",
		lng.UZBEKCYR:  "TODO",
		lng.UNDEFINED: "TODO",
	},
	"BTN_OrderHistoryBtn": {
		lng.RUSSIAN:   "🗄 История заказов",
		lng.ENGLISH:   "TODO",
		lng.UZBEKLAT:  "TODO",
		lng.UZBEKCYR:  "TODO",
		lng.UNDEFINED: "TODO",
	},
	"BTN_OrderHistoryBtn_next": {
		lng.RUSSIAN:   "> Следующая",
		lng.ENGLISH:   "> ",
		lng.UZBEKLAT:  "> ",
		lng.UZBEKCYR:  "> ",
		lng.UNDEFINED: "> ",
	},
	"BTN_OrderHistoryBtn_prev": {
		lng.RUSSIAN:   "Предыдущая <",
		lng.ENGLISH:   "<",
		lng.UZBEKLAT:  "<",
		lng.UZBEKCYR:  "<",
		lng.UNDEFINED: "<",
	},
	"BTN_LicenseAgreeBtn": {
		lng.RUSSIAN:   "➡️ Продолжить",
		lng.ENGLISH:   "➡️ Continue",
		lng.UZBEKLAT:  "➡️ Davom ettirish",
		lng.UZBEKCYR:  "➡️ Давом эттириш",
		lng.UNDEFINED: "➡️",
	},
	"BTN_SOSContinueBtn": {
		lng.RUSSIAN:   "➡️ Продолжить",
		lng.ENGLISH:   "➡️ Continue",
		lng.UZBEKLAT:  "➡️ Davom ettirish",
		lng.UZBEKCYR:  "➡️ Давом эттириш",
		lng.UNDEFINED: "➡️",
	},
	"BTN_OrderBtn": {
		lng.RUSSIAN:   "📦 Заказать перевод",
		lng.ENGLISH:   "📦 Order Translation",
		lng.UZBEKLAT:  "📦 Tarjima buyurtma qilish",
		lng.UZBEKCYR:  "📦 Таржима буюртма қилиш",
		lng.UNDEFINED: "📦",
	},
	"BTN_OrderOnlineBtn": {
		lng.RUSSIAN:   "🌐 Заказать онлайн",
		lng.ENGLISH:   "🌐 Order Online",
		lng.UZBEKLAT:  "🌐 Onlayn buyurtma berish",
		lng.UZBEKCYR:  "🌐 Онлайн буюртма бериш",
		lng.UNDEFINED: "🌐",
	},
	"BTN_OrderOCTBtn": {
		lng.RUSSIAN:   "🕒 Выбрать время",
		lng.ENGLISH:   "🕒 Select Time",
		lng.UZBEKLAT:  "🕒 Vaxt tanlash",
		lng.UZBEKCYR:  "🕒 Вақт танлаш",
		lng.UNDEFINED: "🕒",
	},
	"BTN_OrderOnlineASAPBtn": {
		lng.RUSSIAN:   "⏩ Как можно быстрее",
		lng.ENGLISH:   "⏩ As Soon As Possible",
		lng.UZBEKLAT:  "⏩ Imkon qadar tezroq",
		lng.UZBEKCYR:  "⏩ Имкон қадар тезроқ",
		lng.UNDEFINED: "⏩",
	},
	"BTN_OrderOfflineBtn": {
		lng.RUSSIAN:   "🏠 Заказать оффлайн (Выезд)",
		lng.ENGLISH:   "🏠 Order Offline (On-Site)",
		lng.UZBEKLAT:  "🏠 Offlayn buyurtma berish (Chiqish)",
		lng.UZBEKCYR:  "🏠 Оффлайн буюртма бериш (Чиқиш)",
		lng.UNDEFINED: "🏠",
	},
	"BTN_SOSBtn": {
		lng.RUSSIAN:   "🆘 SOS",
		lng.ENGLISH:   "🆘 SOS",
		lng.UZBEKLAT:  "🆘 SOS",
		lng.UZBEKCYR:  "🆘 SOS",
		lng.UNDEFINED: "🆘",
	},
	"BTN_HelpBtn": {
		lng.RUSSIAN:   "❓ Справка и поддержка",
		lng.ENGLISH:   "❓ Help & Support",
		lng.UZBEKLAT:  "❓ Yordam va qo‘llab-quvvatlash",
		lng.UZBEKCYR:  "❓ Ёрдам ва қўллаб-қувватлаш",
		lng.UNDEFINED: "❓",
	},
	"BTN_SupportBtn": {
		lng.RUSSIAN:   "🤝 Поддержка",
		lng.ENGLISH:   "🤝 Support",
		lng.UZBEKLAT:  "🤝 Qo‘llab-quvvatlash",
		lng.UZBEKCYR:  "🤝 Қўллаб-қувватлаш",
		lng.UNDEFINED: "🤝",
	},
	"BTN_BalanceBtn": {
		lng.RUSSIAN:   "💳 Баланс и оплата",
		lng.ENGLISH:   "💳 Balance & Payment",
		lng.UZBEKLAT:  "💳 Balans va to‘lov",
		lng.UZBEKCYR:  "💳 Баланс ва тўлов",
		lng.UNDEFINED: "💳",
	},
	"BTN_SettingsBtn": {
		lng.RUSSIAN:   "⚙️ Параметры",
		lng.ENGLISH:   "⚙️ Settings",
		lng.UZBEKLAT:  "⚙️ Tashkilotlash",
		lng.UZBEKCYR:  "⚙️ Ташкилатлаш",
		lng.UNDEFINED: "⚙️",
	},
	"BTN_MainMenuBtn": {
		lng.RUSSIAN:   "⬅️ Назад",
		lng.ENGLISH:   "⬅️ Back",
		lng.UZBEKLAT:  "⬅️ Orqaga",
		lng.UZBEKCYR:  "⬅️ Орқага",
		lng.UNDEFINED: "⬅️",
	},
	"BTN_MainMenuNoDeleteBtn": {
		lng.RUSSIAN:   "⬅️ Назад",
		lng.ENGLISH:   "⬅️ Back",
		lng.UZBEKLAT:  "⬅️ Orqaga",
		lng.UZBEKCYR:  "⬅️ Орқага",
		lng.UNDEFINED: "⬅️",
	},
	"BTN_PaymentChooseSumBtn": {
		lng.RUSSIAN:   "💸 Пополнить",
		lng.ENGLISH:   "💸 Top Up",
		lng.UZBEKLAT:  "💸 Toldirish",
		lng.UZBEKCYR:  "💸 Толдириш",
		lng.UNDEFINED: "💸",
	},
	"BTN_LongHelpBtn": {
		lng.RUSSIAN:   "ℹ️ Подробнее",
		lng.ENGLISH:   "ℹ️ More Info",
		lng.UZBEKLAT:  "ℹ️ Ko‘proq ma'lumot",
		lng.UZBEKCYR:  "ℹ️ Кўпроқ маълумот",
		lng.UNDEFINED: "ℹ️",
	},
	"BTN_OrderOnlineASAPSlot5Btn": {
		lng.RUSSIAN:   "5 минут",
		lng.ENGLISH:   "",
		lng.UZBEKLAT:  "",
		lng.UZBEKCYR:  "",
		lng.UNDEFINED: "",
	},
	"BTN_OrderOnlineASAPSlot10Btn": {
		lng.RUSSIAN:   "10 минут",
		lng.ENGLISH:   "",
		lng.UZBEKLAT:  "",
		lng.UZBEKCYR:  "",
		lng.UNDEFINED: "",
	},
	"BTN_OrderOnlineASAPSlot15Btn": {
		lng.RUSSIAN:   "15 минут",
		lng.ENGLISH:   "",
		lng.UZBEKLAT:  "",
		lng.UZBEKCYR:  "",
		lng.UNDEFINED: "",
	},
	"BTN_OrderOnlineASAPSlot20Btn": {
		lng.RUSSIAN:   "20 минут",
		lng.ENGLISH:   "",
		lng.UZBEKLAT:  "",
		lng.UZBEKCYR:  "",
		lng.UNDEFINED: "",
	},
	"BTN_OrderOnlineASAPSlot25Btn": {
		lng.RUSSIAN:   "25 минут",
		lng.ENGLISH:   "",
		lng.UZBEKLAT:  "",
		lng.UZBEKCYR:  "",
		lng.UNDEFINED: "",
	},
	"BTN_OrderOnlineASAPSlot30Btn": {
		lng.RUSSIAN:   "30 минут",
		lng.ENGLISH:   "",
		lng.UZBEKLAT:  "",
		lng.UZBEKCYR:  "",
		lng.UNDEFINED: "",
	},
	"BTN_SetWorkTimeBtn": {
		lng.RUSSIAN:   "Настроить время работы",
		lng.ENGLISH:   "",
		lng.UZBEKLAT:  "",
		lng.UZBEKCYR:  "",
		lng.UNDEFINED: "",
	},
	"BTN_TrOrdersHistoryBtn": {
		lng.RUSSIAN:   "🗄 История заказов",
		lng.ENGLISH:   "",
		lng.UZBEKLAT:  "",
		lng.UZBEKCYR:  "",
		lng.UNDEFINED: "",
	},
	"BTN_TrBalanceBtn": {
		lng.RUSSIAN:   "Баланс",
		lng.ENGLISH:   "",
		lng.UZBEKLAT:  "",
		lng.UZBEKCYR:  "",
		lng.UNDEFINED: "",
	},
}

var Messages = map[string]map[lng.Language]string{
	"hello": {
		lng.UNDEFINED: `*Assalomu alaykum!* 👋
Bizning xizmatimizga xush kelibsiz! Sizni professional surdotarjimonlar bilan osongina va qulay bog‘lash uchun shu yerdamiz.

*Здравствуйте!* 👋
Добро пожаловать в наш сервис! Мы здесь, чтобы помочь вам легко и удобно связаться с профессиональными сурдопереводчиками.

*Hello!* 👋
Welcome to our service! We are here to help you connect easily and conveniently with professional sign language interpreters.
`,
	},
	"hello_translator": {
		lng.UNDEFINED: `*Assalomu alaykum!* 👋

*Здравствуйте!* 👋

*Hello!* 👋
`,
	},
	"chooseLanguage": {
		lng.UNDEFINED: "Tilni tanlang/Тилни танланг\nВыберите язык\nChoose language",
	},
	"language": {
		lng.UZBEKLAT: "🇺🇿 Oʻzbekcha",
		lng.UZBEKCYR: "🇺🇿 Ўзбекча",
		lng.ENGLISH:  "🇺🇸 English",
		lng.RUSSIAN:  "🇷🇺 Русский",
	},
	"license": {
		lng.RUSSIAN: `Продолжая использование данного бота вы соглашаетесь с условиями *лицензии и оферты* 
Прочитать условия вы можете по ссылкам ниже:
- [На русском языке 🇷🇺](surdo.x-soft.uz/ru/license)
- [На узбекском языке 🇺🇿](surdo.x-soft.uz/uz/license)
- [На английском языке 🇺🇸](surdo.x-soft.uz/en/license)`,
		lng.UZBEKLAT: `Mazkur botdan foydalanishni davom ettirgan holda siz *litsenziya va taklif shartlariga* rozilik bildirasiz.  
Shartlarni quyidagi havolalar orqali o‘qishingiz mumkin:  
- [Rus tilida 🇷🇺](surdo.x-soft.uz/ru/license)  
- [O‘zbek tilida 🇺🇿](surdo.x-soft.uz/uz/license)  
- [Ingliz tilida 🇺🇸](surdo.x-soft.uz/en/license)`,
		lng.UZBEKCYR: `Мазкур ботдан фойдаланишни давом эттирган ҳолда сиз *лицензия ва таклиф шартларига* розилик билдирасиз.  
Шартларни қуйидаги ҳаволалар орқали ўқишингиз мумкин:  
- [Рус тилида 🇷🇺](surdo.x-soft.uz/ru/license)  
- [Ўзбек тилида 🇺🇿](surdo.x-soft.uz/uz/license)  
- [Инглиз тилида 🇺🇸](surdo.x-soft.uz/en/license)`,
		lng.ENGLISH: `By continuing to use this bot, you agree to the terms of the *license and offer*.  
You can read the terms via the links below:  
- [In Russian 🇷🇺](surdo.x-soft.uz/ru/license)  
- [In Uzbek 🇺🇿](surdo.x-soft.uz/uz/license)  
- [In English 🇺🇸](surdo.x-soft.uz/en/license)`,
	},
	"mainMenu": {
		lng.RUSSIAN:  "📝 *Главное меню*",
		lng.ENGLISH:  "📝 *Main Menu*",
		lng.UZBEKLAT: "📝 *Asosiy menyu*",
		lng.UZBEKCYR: "📝 *Асосий меню*",
	},
	"smallHelp": {
		lng.RUSSIAN: `❓ *Справка* ❓

1. *Пополнение баланса:* в главном меню нажмите на кнопку *«` + Buttons["BTN_BalanceBtn"][lng.RUSSIAN] + `»*, в появившемся сообщении выберите пункт *«` + Buttons["BTN_PaymentBtn"][lng.RUSSIAN] + `»*, используйте удобный способ оплаты

2. *Заказ перевода:*  из главного меню выберите пункт *«` + Buttons["BTN_OrderBtn"][lng.RUSSIAN] + `»* и следуйте дальнейшим инструкциям из сообщений

3. *Кнопка "SOS":* для вызова экстренных служб используйте пункт *«` + Buttons["BTN_SOSBtn"][lng.RUSSIAN] + `»* из главного меню, после нажатия Вас соединят с оператором "112". 
*ВНИМАНИЕ: ложный вызов экстренных служб влечет за собой ответственность*

4. *Техническая поддержка:* Если у Вас возникнут вопросы по работе с данным ботом, вы можете обратиться к тех. поддержке. Для этого из меню нажмите на клавишу *«` + Buttons["BTN_HelpBtn"][lng.RUSSIAN] + `»*, кнопка *«` + Buttons["BTN_SupportBtn"][lng.RUSSIAN] + `»* переведет Вас в чат с поддержкой
`,
		lng.ENGLISH: `**Help**

1. *Top Up:* To top up your balance, press the *«` + Buttons["BTN_BalanceBtn"][lng.ENGLISH] + `»* button in the main menu, then select *«` + Buttons["BTN_PaymentBtn"][lng.ENGLISH] + `»* from the appeared message and use any of the available payment methods  

2. *Order Translation:* To order a translation from the main menu, select *«` + Buttons["BTN_OrderBtn"][lng.ENGLISH] + `»* and follow the further instructions from the messages  

3. *"SOS" Button:* To call emergency services, use *«` + Buttons["BTN_SOSBtn"][lng.ENGLISH] + `»* from the main menu. After pressing this button, you will be connected with the "112" operator. 
*ATTENTION: A false emergency call will lead to responsibility*  

4. *Technical Support:* If you have any questions about using this bot, you can contact technical support. For this, press *«` + Buttons["BTN_HelpBtn"][lng.ENGLISH] + `»* in the main menu, and the *«` + Buttons["BTN_SupportBtn"][lng.ENGLISH] + `»* button will lead you to the support chat
`,
		lng.UZBEKLAT: `**Yordam**

1. **To'ldirish:** Balansni to'ldirish uchun asosiy menyuda *«` + Buttons["BTN_BalanceBtn"][lng.UZBEKLAT] + `»* tugmasini bosing, paydo bo'ladigan xabarda *«` + Buttons["BTN_PaymentBtn"][lng.UZBEKLAT] + `»* bo'limini tanlang, taklif qilingan to'lov usulidan foydalaning  

2. **Tarjima buyurtmasi:** Asosiy menyudan tarjima buyurtmasi berish uchun *«` + Buttons["BTN_OrderBtn"][lng.UZBEKLAT] + `»* bo'limini tanlang va xabarlardagi keyingi ko'rsatmalarga amal qiling  

3. **"SOS" tugmasi:** Favqulodda xizmatlarni chaqirish uchun asosiy menyudan *«` + Buttons["BTN_SOSBtn"][lng.UZBEKLAT] + `»* bo'limini tanlang, ushbu tugmani bosganingizdan so'ng sizni "112" operatori bilan bog'lashadi. 
*DIQQAT: yolg'oncha favqulodda xizmatlar chaqirig'i javobgarlikka olib keladi*  

4. **Texnik yordam:** Agar sizda ushbu bot bilan ishlash bo'yicha savollar bo'lsa, texnik yordamga murojaat qilishingiz mumkin. Buning uchun asosiy menyudan *«` + Buttons["BTN_HelpBtn"][lng.UZBEKLAT] + `»* tugmasini bosing, *«` + Buttons["BTN_SupportBtn"][lng.UZBEKLAT] + `»* tugmasi sizni yordam chatiga olib boradi  
`,
		lng.UZBEKCYR: `**Yordam**

1. **To'ldirish:** Balansni to'ldirish uchun asosiy menyuda *«` + Buttons["BTN_BalanceBtn"][lng.UZBEKLAT] + `»* tugmasini bosing, paydo bo'ladigan xabarda *«` + Buttons["BTN_PaymentBtn"][lng.UZBEKLAT] + `»* bo'limini tanlang, taklif qilingan to'lov usulidan foydalaning  

2. **Tarjima buyurtmasi:** Asosiy menyudan tarjima buyurtmasi berish uchun *«` + Buttons["BTN_OrderBtn"][lng.UZBEKLAT] + `»* bo'limini tanlang va xabarlardagi keyingi ko'rsatmalarga amal qiling  

3. **"SOS" tugmasi:** Favqulodda xizmatlarni chaqirish uchun asosiy menyudan *«` + Buttons["BTN_SOSBtn"][lng.UZBEKLAT] + `»* bo'limini tanlang, ushbu tugmani bosganingizdan so'ng sizni "112" operatori bilan bog'lashadi. 
*DIQQAT: yolg'oncha favqulodda xizmatlar chaqirig'i javobgarlikka olib keladi*  

4.**Texnik yordam:** Agar sizda ushbu bot bilan ishlash bo'yicha savollar bo'lsa, texnik yordamga murojaat qilishingiz mumkin. Buning uchun asosiy menyudan *«` + Buttons["BTN_HelpBtn"][lng.UZBEKLAT] + `»* tugmasini bosing, *«` + Buttons["BTN_SupportBtn"][lng.UZBEKLAT] + `»* tugmasi sizni yordam chatiga olib boradi  
`,
	},
	"longHelp": {
		lng.RUSSIAN:  "Здесь должна быть большая справка, но пока нет",
		lng.ENGLISH:  "Здесь должна быть большая справка, но пока нет",
		lng.UZBEKLAT: "Здесь должна быть большая справка, но пока нет",
		lng.UZBEKCYR: "Здесь должна быть большая справка, но пока нет",
	},
	"order": {
		lng.RUSSIAN: `*📦 Заказать перевод 📦*

`,
		lng.ENGLISH:  ``,
		lng.UZBEKLAT: ``,
		lng.UZBEKCYR: ``,
	},
	"sos": {
		lng.RUSSIAN: `*🆘 SOS 🆘*

*ВНИМАНИЕ:* ложный вызов экстренных служб влечет за собой ответственность. Продолжая, вы соглашаетесь передать свои данные в службу "112"

Если Вы согласны, нажмите *«` + Buttons["BTN_SOSContinueBtn"][lng.RUSSIAN] + `»* и следуйте дальнейшим инструкциям
`,
		lng.ENGLISH:  ``,
		lng.UZBEKLAT: ``,
		lng.UZBEKCYR: ``,
	},
	"sosOrdered": {
		lng.RUSSIAN: `*✅ Вам сейчас помогут ✅*

Сейчас Вам перезвонит оператор, ему переданы ваши данные.
Объясните ему вашу ситуацию и следуйте его инструкциям.
`,
		lng.ENGLISH:  ``,
		lng.UZBEKLAT: ``,
		lng.UZBEKCYR: ``,
	},
	"balance": {
		lng.RUSSIAN: `*💳 Баланс и оплата 💳*
*Ваш баланс:* %d сум
*Заморожено на счету:* %d сум

`,
		lng.ENGLISH:  ``,
		lng.UZBEKLAT: ``,
		lng.UZBEKCYR: ``,
	},
	"settings": {
		lng.RUSSIAN: `*⚙️ Параметры ⚙️*

`,
		lng.ENGLISH:  ``,
		lng.UZBEKLAT: ``,
		lng.UZBEKCYR: ``,
	},
	"orderOnline": {
		lng.RUSSIAN: `*🌐 Заказать онлайн 🌐*

`,
		lng.ENGLISH:  ``,
		lng.UZBEKLAT: ``,
		lng.UZBEKCYR: ``,
	},
	"orderOnlineASAP": {
		lng.RUSSIAN: `*⏩ Как можно быстрее ⏩*

Выберите длительность звонка
*Внимание:* Деньги будут заморожены на время поиска. При отмене деньги будут разморожены. Если переводчик будет найден, при отмене средства будут возвращены за исключением цены минимального тарифа (5 минут)
`,
		lng.ENGLISH:  ``,
		lng.UZBEKLAT: ``,
		lng.UZBEKCYR: ``,
	},
	"ASAP": {
		lng.RUSSIAN:  "Как можно быстрее",
		lng.ENGLISH:  ``,
		lng.UZBEKLAT: ``,
		lng.UZBEKCYR: ``,
	},

	"orderOffline": {
		lng.RUSSIAN: `*🏠 Заказать оффлайн (Выезд) 🏠*

`,
		lng.ENGLISH:  ``,
		lng.UZBEKLAT: ``,
		lng.UZBEKCYR: ``,
	},

	"orderCreatedTmpl": {
		lng.RUSSIAN: `*Спасибо за ваш Заказ*

*Время:* $ORDER_TIME
*Длительность:* $ORDER_DURATION
*Стоимость:* $ORDER_COST сум (заморожено на вашем счету)
*Номер:* $ORDER_ID __(при обращении в поддержку отправьте этот номер оператору)__

За 15 минут до начала Вам придет уведомление о предстоящем заказе. За 3 минуты мы пришлем Вам ссылку на звонок, убедитесь что ваша связь стабильна, заранее установите необходимое для звонков ПО на Ваше устройство

Если переводчик не будет найден, то заказ будет отменен за 10 минут до начала.
Отменить данный заказ можно в истории заказов.

Увидеть этот заказ Вы можете в меню истории заказов.
`,
		lng.ENGLISH:  ``,
		lng.UZBEKLAT: ``,
		lng.UZBEKCYR: ``,
	},
	"SetWorkTime": {
		lng.RUSSIAN: `*Рабочее время*

`,
		lng.ENGLISH:  ``,
		lng.UZBEKLAT: ``,
		lng.UZBEKCYR: ``,
	},
	"TrOrdersHistory": {
		lng.RUSSIAN: `*🗄 История заказов 🗄*

`,
		lng.ENGLISH:  ``,
		lng.UZBEKLAT: ``,
		lng.UZBEKCYR: ``,
	},
	"TrBalance": {
		lng.RUSSIAN: `*Баланс*

`,
		lng.ENGLISH:  ``,
		lng.UZBEKLAT: ``,
		lng.UZBEKCYR: ``,
	},
	"WorktimeHistoryListView": {
		lng.RUSSIAN: `*Мое рабочее время*
Чтобы удалить рабочее время нажмите на кнопку с "❌" и номером из списка ниже
Страница %d из %d

%s`,
	},
	"OrderHistoryListView": {
		lng.RUSSIAN: `*🗄 История заказов 🗄*
Чтобы отменить заказ нажмите на кнопку с "❌" и номером из списка ниже 
Страница %d из %d

%s`,
		lng.ENGLISH: `*История заказов*
Страница %d из %d

%s`,
		lng.UZBEKLAT: `*История заказов*
Страница %d из %d

%s`,
		lng.UZBEKCYR: `*История заказов*
Страница %d из %d

%s`,
		lng.UNDEFINED: `*История заказов*
Страница %d из %d

%s`,
	},
}

func BuildMessageOrderCreated(order types.TOrder, lang lng.Language) string {
	msg := Messages["orderCreatedTmpl"][lang]
	orderTime := ""
	if order.Slot != nil {
		orderTime = order.Slot.Start.String()
		msg = strings.Replace(msg, "$ORDER_DURATION",
			fmt.Sprintf("%.f минут", order.Slot.End.Sub(order.Slot.Start).Minutes()), -1)
	}
	msg = strings.Replace(msg, "$ORDER_TIME", orderTime, -1)

	msg = strings.Replace(msg, "$ORDER_COST", strconv.Itoa(int(order.Cost)), -1)
	msg = strings.Replace(msg, "$ORDER_ID", order.ID.String(), -1)
	return msg
}
