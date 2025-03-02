package handler

import (
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/models"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/event"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/tracker"
	protoUser "github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/user"
	"gomodules.xyz/pointer"
	"google.golang.org/genproto/googleapis/type/datetime"
	"time"
)

// user | model <--> proto conversion
func userProtoToModel(user *protoUser.User) *models.User {
	return &models.User{
		ID:   (models.UserID)(user.Id.GetId()),
		Name: user.Name,
		Age:  user.Age,
		ContactInfo: models.ContactInfo{
			UserID:  (*models.UserID)(pointer.Uint64P(user.Id.GetId())),
			Email:   user.Email,
			Phone:   user.PhoneNo,
			Address: user.Address,
		},
		Factors: models.Factors{
			UserID:        (*models.UserID)(pointer.Uint64P(user.Id.GetId())),
			Height:        user.Height,
			Weight:        user.Weight,
			DiabeticLevel: user.DiabeticLevel,
		},
		Gender: user.Gender.String(),
	}
}

func userModelToProto(user *models.User) *protoUser.User {
	return &protoUser.User{
		Id:            &protoUser.UserID{Id: uint64(user.ID)},
		Name:          user.Name,
		Age:           user.Age,
		Email:         user.ContactInfo.Email,
		PhoneNo:       user.ContactInfo.Phone,
		Address:       user.ContactInfo.Address,
		Height:        user.Factors.Height,
		Weight:        user.Factors.Weight,
		DiabeticLevel: user.Factors.DiabeticLevel,
		Gender:        protoUser.Gender(protoUser.Gender_value[user.Gender]),
	}
}

// tracker | model <--> proto conversion
func positionModelToProto(pos *models.Position) *tracker.Position {
	return &tracker.Position{
		Longitude: pos.Longitude,
		Latitude:  pos.Latitude,
		UserId:    uint64(pos.UID),
		Time:      TimeToProtoDateTime(pos.Time),
	}
}
func positionProtoToModel(pos *tracker.Position) *models.Position {
	return &models.Position{
		Longitude:    pos.Longitude,
		Latitude:     pos.Latitude,
		CheckPointID: pos.CkId,
		Time:         ProtoDateTimeToTime(pos.Time),
	}
}

func pulseRateWithUserIDModelToProto(pr *models.PulseRateWithUserID) *tracker.PulseRateWithUserId {
	return &tracker.PulseRateWithUserId{
		UserId:    uint64(pr.UserID),
		PulseRate: pr.PulseRate,
	}
}

func pulseRateWithUserIDProtoToModel(pr *tracker.PulseRateWithUserId) *models.PulseRateWithUserID {
	return &models.PulseRateWithUserID{
		UserID:    models.UserID(pr.UserId),
		PulseRate: pr.PulseRate,
	}
}

func alertModelToProto(al *models.Alert) *tracker.Alert {
	return &tracker.Alert{
		Alert:  alertTypeModelsToProto(al.Type),
		Advice: al.Message,
	}
}

func alertTypeModelsToProto(at models.AlertType) tracker.AlertType {
	switch at {
	case models.Normal:
		return 0
	case models.HighPulseRate:
		return 1
	case models.LowPulseRate:
		return 2
	default:
		return -1
	}
}

func checkpointToAndFromProtoToModel(ctf *tracker.CheckpointToAndFrom) *models.CheckpointToAndFrom {
	return &models.CheckpointToAndFrom{
		To:   ctf.To,
		From: ctf.From,
	}
}

// event | model <--> proto conversion

func eventProtoToModel(event *event.Event) *models.Event {
	return &models.Event{
		EventID:     event.EId,
		GroupID:     event.GId,
		PublisherID: (*models.UserID)(&event.Publisher),
		State:       eventStateProtoToModel(event.State),
		Interested: func() []*models.UserID {
			if event.Interested == nil {
				return nil
			}
			var ids []*models.UserID
			for _, id := range event.Interested {
				ids = append(ids, (*models.UserID)(&id.Id))
			}
			return ids
		}(),
		Going: func() []*models.UserID {
			if event.Going == nil {
				return nil
			}
			var ids []*models.UserID
			for _, id := range event.Going {
				ids = append(ids, (*models.UserID)(&id.Id))
			}
			return ids
		}(),
		NotInterested: func() []*models.UserID {
			if event.NotInterested == nil {
				return nil
			}
			var ids []*models.UserID
			for _, id := range event.NotInterested {
				ids = append(ids, (*models.UserID)(&id.Id))
			}
			return ids
		}(),
		EventDesc: models.EventDescription{
			Description: event.EventDesc.GetDesc(),
			Name:        event.EventDesc.GetName(),
		},
		EventDateTime: func() *time.Time {
			return ProtoDateTimeToTime(event.EventDateTime)
		}(),
	}
}

func eventModelToProto(e *models.Event) *event.Event {
	return &event.Event{
		EId:       e.EventID,
		GId:       e.GroupID,
		Publisher: (uint64)(*e.PublisherID),
		State:     eventStateModelToProto(e.State),
		Interested: func() []*protoUser.UserID {
			if e.Interested == nil {
				return nil
			}
			var ids []*protoUser.UserID
			for _, id := range e.Interested {
				ids = append(ids, &protoUser.UserID{Id: uint64(*id)})
			}
			return ids
		}(),
		Going: func() []*protoUser.UserID {
			if e.Going == nil {
				return nil
			}
			var ids []*protoUser.UserID
			for _, id := range e.Going {
				ids = append(ids, &protoUser.UserID{Id: uint64(*id)})
			}
			return ids
		}(),
		NotInterested: func() []*protoUser.UserID {
			if e.NotInterested == nil {
				return nil
			}
			var ids []*protoUser.UserID
			for _, id := range e.NotInterested {
				ids = append(ids, &protoUser.UserID{Id: uint64(*id)})
			}
			return ids
		}(),
		EventDesc: &event.EventDescription{
			Desc: e.EventDesc.Description,
			Name: e.EventDesc.Name,
		},
		EventDateTime: TimeToProtoDateTime(e.EventDateTime),
	}
}

func ProtoDateTimeToTime(dt *datetime.DateTime) *time.Time {
	if dt == nil {
		return nil
	}
	nanos := time.Duration(dt.Nanos) * time.Nanosecond
	t := time.Date(
		int(dt.Year),
		time.Month(dt.Month),
		int(dt.Day),
		int(dt.Hours),
		int(dt.Minutes),
		int(dt.Seconds),
		int(nanos),
		time.UTC,
	)
	return &t
}

func TimeToProtoDateTime(t *time.Time) *datetime.DateTime {
	if t == nil {
		return nil
	}
	return &datetime.DateTime{
		Year:    int32(t.Year()),
		Month:   int32(t.Month()),
		Day:     int32(t.Day()),
		Hours:   int32(t.Hour()),
		Minutes: int32(t.Minute()),
		Seconds: int32(t.Second()),
		Nanos:   int32(t.Nanosecond()),
	}
}

func eventStateProtoToModel(state event.EventState) models.EventState {
	switch state {
	case 0:
		return models.EventOngoing
	case 1:
		return models.EventClosed
	case 2:
		return models.EventUpcoming
	default:
		return models.EventUnknown
	}
}

func eventStateModelToProto(e models.EventState) event.EventState {
	switch e {
	case models.EventOngoing:
		return event.EventState_ongoing
	case models.EventClosed:
		return event.EventState_closed
	case models.EventUpcoming:
		return event.EventState_upcoming
	}
	return event.EventState_unknown
}
