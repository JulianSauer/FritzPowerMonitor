package telegram

import (
    "bytes"
    "encoding/json"
    "errors"
    "github.com/JulianSauer/FritzPowerMonitor/config"
    "github.com/JulianSauer/FritzPowerMonitor/dto"
    "io/ioutil"
    "net/http"
)

const API = "https://api.telegram.org/bot"
const SEND_MESSAGE = "/sendMessage"

var configuration *config.Config = nil

func SendMessage(message string) error {
    if configuration == nil {
        configuration = config.Load()
    }

    return sendMessage(configuration.TelegramBotId, dto.TelegramMessage{ChatId: configuration.TelegramChatId, Text: message})
}

func sendMessage(botId string, message dto.TelegramMessage) error {
    url := API + botId + SEND_MESSAGE

    body, e := json.Marshal(message)
    if e != nil {
        return e
    }

    request, e := http.NewRequest("GET", url, bytes.NewBuffer(body))
    if e != nil {
        return e
    }

    request.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    response, e := client.Do(request)
    if e != nil {
        return e
    }

    if response.StatusCode < 200 || response.StatusCode >= 300 {
        responseBody, e := ioutil.ReadAll(response.Body)
        if e != nil {
            return e
        }
        return errors.New(string(responseBody))
    }

    return nil
}
