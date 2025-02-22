package models

type AlertType string

const (
	Normal        AlertType = "Normal"
	LowPulseRate  AlertType = "LowPulseRate"
	HighPulseRate AlertType = "HighPulseRate"
)

type Position struct {
	UID          UserID  `json:"uid,omitempty"`
	Latitude     float64 `json:"latitude,omitempty"`
	Longitude    float64 `json:"longitude,omitempty"`
	CheckPointID uint64  `json:"checkPointId,omitempty"`
}

type PulseRateWithUserID struct {
	UserID    UserID  `json:"user_id,omitempty"`
	PulseRate float32 `json:"pulse_rate,omitempty"`
}

type Alert struct {
	Type    AlertType `json:"type,omitempty"`
	Message string    `json:"message,omitempty"`
}

type CheckpointToAndFrom struct {
	To   uint64 `json:"to,omitempty"`
	From uint64 `json:"from,omitempty"`
}
