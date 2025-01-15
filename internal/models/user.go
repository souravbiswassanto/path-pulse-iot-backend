package models

import "time"

type FitnessGoal string

const (
	FitnessGoalUnknown     = "unknown"
	FitnessGOalWeightLost  = "WeightLoss"
	FitnessGoalMuscleGain  = "MuscleGain"
	FitnessGoalMaintenance = "Maintenance"
)

type User struct {
	ID          int32       `json:"id"`
	Name        string      `json:"name,omitempty"`
	Age         int         `json:"age,omitempty"`
	Gender      string      `json:"gender,omitempty"`
	ContactInfo ContactInfo `json:"contactInfo,omitempty"`
	Factors     Factors     `json:"factors,omitempty"`
	CreatedAt   *time.Time  `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time  `json:"updatedAt,omitempty"`
	Fitness     Fitness
}

type ContactInfo struct {
	UserID  int32  `json:"user_id,omitempty"`
	Email   string `json:"email,omitempty"`
	Phone   string `json:"phone"`
	Address string `json:"address,omitempty"`
}

type Factors struct {
	UserID        int32         `json:"userID,omitempty"`
	Height        float32       `json:"height,omitempty"`
	Weight        float32       `json:"weight,omitempty"`
	DiabeticLevel float32       `json:"diabeticLevel,omitempty"`
	BP            BloodPressure `json:"BP,omitempty"`
}

type BloodPressure struct {
	Systolic  int `json:"systolic,omitempty"`
	Diastolic int `json:"diastolic,omitempty"`
}

type Fitness struct {
	Goal          FitnessGoal `json:"goal,omitempty"`
	CalorieBurned int32       `json:"calorieBurned,omitempty"`
}
