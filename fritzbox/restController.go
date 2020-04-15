package fritzbox

import (
    "github.com/JulianSauer/FritzPowerMonitor/dto"
    "github.com/labstack/echo"
    "log"
    "net/http"
)

// Monitors a device sending a telegram message when the given threshold is met
// Body: dto.monitorDevice
func MonitorDevice(context echo.Context) error {
    if e := homeAutomation.Login(); e != nil {
        return context.String(http.StatusBadRequest, e.Error())
    }

    deviceToMonitor, e := getDeviceToMonitorFromParameter(context)
    if e != nil {
        return context.String(http.StatusBadRequest, e.Error())
    }
    if e := monitorDevice(deviceToMonitor); e != nil {
        return context.String(http.StatusBadRequest, e.Error())
    }

    return context.String(http.StatusOK, "")
}

// Returns a list of devices connected to the Fritz box
func ListDevices(context echo.Context) error {
    if e := homeAutomation.Login(); e != nil {
        return context.String(http.StatusBadRequest, e.Error())
    }

    deviceList, e := homeAutomation.List()
    if e != nil {
        return context.String(http.StatusBadRequest, e.Error())
    }

    return context.JSON(http.StatusOK, *deviceList)
}

// Shows power and energy of a specific device
// Body: dto.Device
func GetPowerMeter(context echo.Context) error {
    if e := homeAutomation.Login(); e != nil {
        return context.String(http.StatusBadRequest, e.Error())
    }

    deviceParameter, e := getDeviceFromParameter(context)
    if e != nil {
        return context.String(http.StatusBadRequest, e.Error())
    }

    device, e := getDeviceByName(deviceParameter.Name)
    if e != nil {
        return context.String(http.StatusBadRequest, e.Error())
    }
    return context.JSON(http.StatusOK, device.Powermeter)
}

func getDeviceFromParameter(context echo.Context) (*dto.Device, error) {
    device := new(dto.Device)

    if e := context.Bind(device); e != nil {
        log.Println(e.Error())
        return nil, context.String(http.StatusBadRequest, e.Error())
    }
    return device, nil
}

func getDeviceToMonitorFromParameter(context echo.Context) (*dto.MonitorDevice, error) {
    deviceToMonitor := new(dto.MonitorDevice)

    if e := context.Bind(deviceToMonitor); e != nil {
        log.Println(e.Error())
        return nil, context.String(http.StatusBadRequest, e.Error())
    }
    return deviceToMonitor, nil
}
