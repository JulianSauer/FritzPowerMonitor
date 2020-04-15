package dto

type TelegramMessage struct {
    ChatId string `json:"chat_id"`
    Text   string `json:"text"`
}
