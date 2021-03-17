package dto

type MonitorDevice struct {
    Name            string `json:"name"`
    TTL             int64  `json:"timeToLive,string"`
    Interval        int64  `json:"interval,string"`
    PowerThreshold  int64  `json:"powerThreshold,string"`
    EnergyThreshold int64  `json:"energyThreshold,string"`
    Message         string `json:"message,omitempty"`
}
