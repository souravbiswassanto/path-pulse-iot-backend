package models

import (
	"time"
)

type EventType int
type EventState string

const (
	EventOngoing  EventState = "Ongoing"
	EventClosed   EventState = "Closed"
	EventUpcoming EventState = "Upcoming"
	EventUnknown  EventState = "Unknown"
	Running       EventType  = iota
	Walking
	Exercise
	Discussion
	Hiking
	Travelling
	Game
)

type EventDescription struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type Event struct {
	EventID       uint64           `json:"event_id,omitempty"`
	GroupID       uint64           `json:"group_id,omitempty"`
	PublisherID   *UserID          `json:"publisher_id,omitempty"`
	State         EventState       `json:"state,omitempty"`
	Interested    []*UserID        `json:"interested,omitempty"`
	Going         []*UserID        `json:"going,omitempty"`
	NotInterested []*UserID        `json:"not_interested,omitempty"`
	EventDesc     EventDescription `json:"event_desc,omitempty"`
	EventDateTime *time.Time       `json:"event_date_time,omitempty"`
}
