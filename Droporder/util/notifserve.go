package util

import (
	entity "Drop/Droporder/entities"
	notification "Drop/Droporder/repository/order/notificationrepo"
	"context"
	"errors"
	"log"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/api/option"
)

var app *firebase.App

func init() {
	opt := option.WithCredentialsFile("for-poc-325210-a7e014fe2cab.json")
	var err error
	app, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {

	}
	log.Println("App connect Successfull")

}

var (
	repo = notification.NewNotificationRepository("notification")
)

type notificationService struct{}

func NewNotificationService(repository notification.NotificationRepository) NotificationService {
	repo = repository
	return &notificationService{}
}

func (*notificationService) SendNotificationWithTopic(title, body, topic, userid string) (string, error) {
	var msg entity.MessageData
	msg.ID = primitive.NewObjectID()
	msg.Title = title
	msg.Topic = topic
	msg.Body = body
	msg.UserId = userid
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		return "", errors.New("unable to send the notification")
	}
	oneHour := time.Duration(1) * time.Hour

	data := map[string]string{
		"topic":        topic,
		"userid":       userid,
		"Title":        title,
		"Body":         body,
		"click_action": "FLUTTER_NOTIFICATION_CLICK",
	}
	message := &messaging.Message{
		Data: data,
		Android: &messaging.AndroidConfig{
			TTL:      &oneHour,
			Priority: "high",
			Notification: &messaging.AndroidNotification{
				Title:       title,
				Body:        body,
				ClickAction: "FLUTTER_NOTIFICATION_CLICK",
			},
		},
		APNS: &messaging.APNSConfig{
			Headers: map[string]string{
				"apns-priority": "10",
			},
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{
					Alert: &messaging.ApsAlert{
						Title: title,
						Body:  body,
					},
				},
			},
		},
		Topic: topic,
	}
	_, err = client.Send(ctx, message)
	if err != nil {
		msg.Status = "Failed"
		_, err = repo.InsertOne(msg)
		return "", errors.New("unable to send the notification")
	}
	msg.Status = "Success"
	msg.SentTime = time.Now()
	return repo.InsertOne(msg)
}

func (*notificationService) GetNotifications(limit, skip int, status, userid, topic, title string) ([]entity.MessageData, error) {
	filter := bson.M{}
	if status != "" {
		filter["status"] = status
	}
	if userid != "" {
		filter["user_id"] = userid
	}
	if topic != "" {
		filter["topic"] = topic
	}
	if title != "" {
		filter["title"] = title
	}

	return repo.Find(filter, bson.M{}, 100, 0)
}
