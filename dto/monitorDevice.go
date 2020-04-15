package dto

type MonitorDevice struct {
    Name            string `json:"name"`
    TTL             int64  `json:"timeToLive"`
    Interval        int64  `json:"interval"`
    PowerThreshold  string `json:"powerThreshold"`
    EnergyThreshold string `json:"energyThreshold"`
}
