package fritzbox

import (
    "github.com/JulianSauer/FritzPowerMonitor/dto"
    "github.com/JulianSauer/FritzPowerMonitor/telegram"
    "log"
    "strconv"
    "time"
)

func monitorDevice(deviceToMonitor *dto.MonitorDevice) error {
    _, e := getDeviceByName(deviceToMonitor.Name)
    if e != nil {
        return e
    }

    go func(deviceToMonitor *dto.MonitorDevice) {
        device, _ := getDeviceByName(deviceToMonitor.Name)
        for start := time.Now(); int64(time.Now().Sub(start)/time.Second) < deviceToMonitor.TTL; time.Sleep(time.Duration(deviceToMonitor.Interval) * time.Second) {
            log.Println("Checking on device " + deviceToMonitor.Name)

            energy, e := strconv.ParseInt(device.Powermeter.Energy, 10, 64)
            if e != nil {
                log.Println(e.Error())
                return
            }

            power, e := strconv.ParseInt(device.Powermeter.Power, 10, 64)
            if e != nil {
                log.Println(e.Error())
                return
            }

            if energy < deviceToMonitor.EnergyThreshold || power < deviceToMonitor.PowerThreshold {
                message := "Threshold reached for " + device.Name
                log.Println(message)
                if e := telegram.SendMessage(message); e != nil {
                    log.Println(e.Error())
                }
                return
            }
            device, _ = getDeviceByName(deviceToMonitor.Name)
        }
        message := "Reached timeout for " + device.Name
        log.Println(message)
        if e := telegram.SendMessage(message); e != nil {
            log.Println(e.Error())
        }
    }(deviceToMonitor)

    return nil
}
