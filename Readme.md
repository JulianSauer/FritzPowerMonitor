# FritzPowerMonitor
Reads the state of AVM FRITZ!DECT 200 smart plugs and sends messages using Telegram.

You can run it by using the [prebuild image](https://github.com/JulianSauer/FritzPowerMonitor/packages/674815):
```
docker run --rm -d -p 8080:8080 -v path/to/your/config.json:/bin/config.json docker.pkg.github.com/juliansauer/fritzpowermonitor/fritz-power-monitor:latest
```

or by cloning this repo and running `docker-compose up -d`.

## Config
A config.json must contain the password of your Fritz box, the Telegram bot id and a chat id of that bot. Messages will be send there. Here's an example:
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
    "energyThreshold": "0",
    "message": "This will be displayed in Telegram"
  }
```

Querries the smart plug every `5` seconds for `1` minute. If the power falls below `20000`, a notification will be send using Telegram. The notification can be specified by using the optional `message` field.
