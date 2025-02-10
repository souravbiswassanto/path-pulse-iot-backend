package models

type Position struct {
	UID       UserID  `json:"uid,omitempty"`
	Latitude  float32 `json:"latitude,omitempty"`
	Longitude float32 `json:"longitude,omitempty"`
}

type Checkpoint struct {
	CID uint64   `json:"c_id,omitempty"`
	UID UserID   `json:"u_id,omitempty"`
	Pos Position `json:"pos,omitempty"`
}

type BloodPressure struct {
	Systolic  int32 `json:"systolic,omitempty"`
	Diastolic int32 `json:"diastolic,omitempty"`
}

type BloodPressureWithUserID struct {
	UserID UserID        `json:"user_id,omitempty"`
	BP     BloodPressure `json:"bp,omitempty"`
}
