package db

import (
	"time"
)

// Event represents information generated by ansible or api action captured to the database during execution
type Event struct {
	ID          int              `db:"id" json:"-"`
	UserID      *int             `db:"user_id" json:"user_id"`
	ProjectID   *int             `db:"project_id" json:"project_id"`
	ObjectID    *int             `db:"object_id" json:"object_id"`
	ObjectType  *EventObjectType `db:"object_type" json:"object_type"`
	Description *string          `db:"description" json:"description"`
	Created     time.Time        `db:"created" json:"created"`

	ObjectName  string  `db:"-" json:"object_name"`
	ProjectName *string `db:"project_name" json:"project_name"`
	Username    *string `db:"-" json:"username"`
}

type EventObjectType string

const (
	EventTask        EventObjectType = "task"
	EventEnvironment EventObjectType = "environment"
	EventInventory   EventObjectType = "inventory"
	EventKey         EventObjectType = "key"
	EventProject     EventObjectType = "project"
	EventRepository  EventObjectType = "repository"
	EventSchedule    EventObjectType = "schedule"
	EventTemplate    EventObjectType = "template"
	EventUser        EventObjectType = "user"
	EventView        EventObjectType = "view"
)

func FillEvents(d Store, events []Event) (err error) {
	usernames := make(map[int]string)

	for i, evt := range events {
		var objName string
		objName, err = getEventObjectName(d, evt)

		if err != nil {
			return
		}

		if objName != "" {
			events[i].ObjectName = objName
		}

		if evt.UserID == nil {
			continue
		}

		var username string

		username, ok := usernames[*evt.UserID]

		if !ok {
			username, err = getEventUsername(d, evt)

			if err != nil {
				return
			}

			if username == "" {
				continue
			}

			usernames[*evt.UserID] = username
		}

		events[i].Username = &username
	}

	return
}

func getEventObjectName(d Store, evt Event) (string, error) {
	if evt.ObjectID == nil || evt.ObjectType == nil {
		return "", nil
	}
	switch *evt.ObjectType {
	case EventTask:
		task, err := d.GetTask(*evt.ProjectID, *evt.ObjectID)
		if err != nil {
			// Task can be deleted, it is ok, just return empty string
			return "", nil
		}
		return task.Playbook, nil
	default:
		return "", nil
	}
}

func getEventUsername(d Store, evt Event) (username string, err error) {
	if evt.UserID == nil {
		return "", nil
	}

	user, err := d.GetUser(*evt.UserID)

	if err != nil {
		return "", err
	}

	return user.Username, nil
}
