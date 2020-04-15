package main

import (
    "github.com/JulianSauer/MeasureMessage/config"
    "github.com/JulianSauer/MeasureMessage/fritzbox"
    "github.com/labstack/echo"
)

func main() {
    configuration := config.Load()
    fritzbox.SetPassword(configuration.FritzPassword)

    router := echo.New()
    router.GET("listDevices", fritzbox.ListDevices)
    router.POST("monitor", fritzbox.MonitorDevice)
    router.POST("getPowerMeter", fritzbox.GetPowerMeter)
    router.Logger.Fatal(router.Start(":8080"))
}
