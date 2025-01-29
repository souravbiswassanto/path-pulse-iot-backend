package models

type Group struct {
	GID     uint64   `json:"g_id,omitempty"`
	Name    string   `json:"name,omitempty"`
	Members []UserID `json:"members,omitempty"`
}
