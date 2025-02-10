package models

type FitnessGoal string
type UserID uint64

const (
	FitnessGoalUnknown     FitnessGoal = "unknown"
	FitnessGOalWeightLost  FitnessGoal = "WeightLoss"
	FitnessGoalMuscleGain  FitnessGoal = "MuscleGain"
	FitnessGoalMaintenance FitnessGoal = "Maintenance"
)

type User struct {
	ID          *UserID     `json:"id"`
	Name        string      `json:"name,omitempty"`
	Age         int32       `json:"age,omitempty"`
	Gender      string      `json:"gender,omitempty"`
	ContactInfo ContactInfo `json:"contact_info,omitempty"`
	Factors     Factors     `json:"factors,omitempty"`
	CreatedAt   *string     `json:"created_at,omitempty"`
	UpdatedAt   *string     `json:"updated_at,omitempty"`
	Fitness     Fitness
}

type ContactInfo struct {
	UserID  *UserID `json:"user_id,omitempty"`
	Email   string  `json:"email,omitempty"`
	Phone   string  `json:"phone"`
	Address string  `json:"address,omitempty"`
}

type Factors struct {
	UserID        *UserID       `json:"user_id,omitempty"`
	Height        float32       `json:"height,omitempty"`
	Weight        float32       `json:"weight,omitempty"`
	DiabeticLevel float32       `json:"diabetic_level,omitempty"`
	BP            BloodPressure `json:"bp,omitempty"`
}

type Fitness struct {
	Goal          FitnessGoal `json:"goal,omitempty"`
	CalorieBurned int32       `json:"calorie_burned,omitempty"`
}
