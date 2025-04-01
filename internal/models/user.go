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
	ID        UserID  `json:"id" xorm:"'userid' pk not null"`
	Name      string  `json:"name,omitempty"`
	Age       int32   `json:"age,omitempty"`
	Gender    string  `json:"gender,omitempty"`
	Email     string  `json:"email,omitempty" xorm:"unique"`
	Phone     string  `json:"phone"`
	Address   string  `json:"address,omitempty"`
	CreatedAt *string `json:"created_at,omitempty" xorm:"'created'"`
	UpdatedAt *string `json:"updated_at,omitempty" xorm:"'updated'"`
}
