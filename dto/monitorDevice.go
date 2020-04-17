package dto

type MonitorDevice struct {
    Name            string `json:"name"`
    TTL             int64  `json:"timeToLive"`
    Interval        int64  `json:"interval"`
    PowerThreshold  int64  `json:"powerThreshold"`
    EnergyThreshold int64  `json:"energyThreshold"`
}
