package models

type AlertType string

const (
	Normal       AlertType = "Normal"
	LowPressure  AlertType = "LowPressure"
	HighPressure AlertType = "HighPressure"
)

type Position struct {
	UID          UserID  `json:"uid,omitempty"`
	Latitude     float32 `json:"latitude,omitempty"`
	Longitude    float32 `json:"longitude,omitempty"`
	CheckPointID uint64  `json:"checkPointId,omitempty"`
}

type BloodPressure struct {
	Systolic  int32 `json:"systolic,omitempty"`
	Diastolic int32 `json:"diastolic,omitempty"`
}

type BloodPressureWithUserID struct {
	UserID UserID        `json:"user_id,omitempty"`
	BP     BloodPressure `json:"bp,omitempty"`
}

type Alert struct {
	Type    AlertType `json:"type,omitempty"`
	Message string    `json:"message,omitempty"`
}
