# FritzPowerMonitor
Reads the state of AVM FRITZ!DECT 200 smart plugs and sends messages using Telegram.

## Config
A config.json should contain the password of your Fritz box, the Telegram bot id and a chat id of that bot. Messages will be send there. Here's an example:
```
{
  "FritzPassword": "123",
  "TelegramBotId": "123456789:ABCDEFGHIJKLMNOPQRSTUVWXYZ123456789",
  "TelegramChatId": "123456789"
}
```

## REST Endpoints
```
/listDevices
```
Returns a list of devices. Useful for getting the name of a smart plug.

---

```
/getPowerMeter

Body:
  {
    "name": "FRITZ!DECT 200 #1"
  }
```
Returns power and energy of the given smart plug.

---

```
/monitor

Body:
  {
    "name": "FRITZ!DECT 200 #1",
    "timeToLive": "60",
    "interval": "5",
    "powerThreshold": "20000",
    "energyThreshold": "0"
  }
```

Querries the smart plug every `5` seconds for `1` minute. If the power falls below `20000`, a notification will be send using Telegram.
