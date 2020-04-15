package fritzbox

import (
    "github.com/bpicode/fritzctl/fritz"
    "log"
)

var homeAutomation fritz.HomeAuto

func SetPassword(password string) {
    homeAutomation = fritz.NewHomeAuto(fritz.SkipTLSVerify(), fritz.Credentials("", password))
}

func getDeviceByName(name string) (*fritz.Device, error) {
    devices, e := homeAutomation.List()
    if e != nil {
        log.Println(e.Error())
        return nil, nil
    }

    var device *fritz.Device = nil
    for _, d := range devices.Devices {
        if d.Name == name {
            device = &d
            break
        }
    }
    if device == nil || !device.CanMeasurePower() {
        errorMessage := "Power device not found: " + name
        log.Println(errorMessage)
        return nil, nil
    }
    return device, nil
}
