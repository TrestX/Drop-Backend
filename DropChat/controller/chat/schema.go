package chat

import (
	entity "Drop/DropChat/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatService interface {
	AddChat(userId string, chat Chats) (string, error)
	UpdateChat(userId string, chat Chats) (string, error)
	GetChat(chatId string) (entity.ChatDB, error)
	GetChats(userId, chatId, status string) ([]OP, error)
	GetChatForUser(user_id, delivery_person_id string) (entity.ChatDB, error)
}

type Chats struct {
	ChatID     string `bson:"_id" json:"chat_id,omitempty"`
	SenderID   string `bson:"sender_id" json:"sender_id,omitempty"`
	ReceiverID string `bson:"receiver_id" json:"receiver_id,omitempty"`
	Chat       []Chat `bson:"chat" json:"chat"`
	ChatToken  string `bson:"chat_token" json:"chat_token"`
	Status     string `bson:"status" json:"status"`
}
type Chat struct {
	Sender   string    `bson:"key" json:"key,omitempty"`
	Message  string    `bson:"message" json:"message,omitempty"`
	Time     time.Time `bson:"time" json:"time,omitempty"`
	Initials string    `bson:"initials" json:"initials"`
}

type OP struct {
	ID               primitive.ObjectID `bson:"_id" json:"hat_id,omitempty"`
	SenderID         string             `bson:"sender_id" json:"sender_id,omitempty"`
	ReceiverID       string             `bson:"receiver_id" json:"receiver_id,omitempty"`
	ReceiverJoinTime time.Time          `bson:"receiver_join_time" json:"receiver_join_time"`
	StopTime         time.Time          `bson:"stop_time" json:"stop_time"`
	Chat             []entity.Chat      `bson:"chat" json:"chat"`
	Status           string             `bson:"status" json:"status"`
	UpdatedTime      time.Time          `bson:"updated_time" json:"updated_time,omitempty"`
	AddedTime        time.Time          `bson:"added_time" json:"added_time,omitempty"`
	Email            string             `bson:"email,omitempty" json:"email,omitempty"`
	Name             string             `bson:"name" json:"name,omitempty"`
	ProfilePhoto     string             `bson:"profile_photo" json:"profile_photo,omitempty"`
}
