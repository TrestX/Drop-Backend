package util

import (
	entity "Drop/Droporder/entities"
)

type NotificationService interface {
	SendNotificationWithTopic(title, body, topic, userid string) (string, error)
	GetNotifications(limit, skip int, status, userid, topic, title string) ([]entity.MessageData, error)
	DeleteNotifications(limit, skip int, status, userid, topic, title string) (string, error)
}

type Notification struct {
	Title  string `bson:"title" json:"title"`
	Body   string `bson:"body" json:"body"`
	Topic  string `bson:"topic" json:"topic"`
	UserId string `bson:"userId" json:"userId"`
}
