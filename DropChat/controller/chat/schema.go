package chat

import (
	entity "Drop/DropChat/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatService interface {
	AddChat(userId, receiverId string, chat Chats) (string, error)
	UpdateChat(userId string, chat Chats) (string, error)
	GetChat(chatId string) (entity.ChatDB, error)
	GetChats(userId, chatId, status string) ([]OP, error)
	GetChatForUser(user_id, delivery_person_id, status string) (entity.ChatDB, error)
}

type Chats struct {
	ChatID     string `bson:"_id" json:"chat_id"`
	SenderID   string `bson:"sender_id" json:"sender_id"`
	ReceiverID string `bson:"receiver_id" json:"receiver_id"`

	Chat      []entity.Chat `bson:"chat" json:"chat"`
	ChatToken string        `bson:"chat_token" json:"chat_token"`
	Status    string        `bson:"status" json:"status"`
}
type Chat struct {
	Sender   string    `bson:"key" json:"key"`
	Message  string    `bson:"message" json:"message"`
	Time     time.Time `bson:"time" json:"time"`
	Initials string    `bson:"initials" json:"initials"`
}

type OP struct {
	ID               primitive.ObjectID `bson:"_id" json:"hat_id"`
	SenderID         string             `bson:"sender_id" json:"sender_id"`
	ReceiverID       string             `bson:"receiver_id" json:"receiver_id"`
	ReceiverJoinTime time.Time          `bson:"receiver_join_time" json:"receiver_join_time"`
	StopTime         time.Time          `bson:"stop_time" json:"stop_time"`
	Chat             []entity.Chat      `bson:"chat" json:"chat"`
	Status           string             `bson:"status" json:"status"`
	UpdatedTime      time.Time          `bson:"updated_time" json:"updated_time"`
	AddedTime        time.Time          `bson:"added_time" json:"added_time"`
	SenderEmail      string             `bson:"email" json:"email"`
	ReceiverEmail    string             `bson:"receiver_email" json:"receiver_email"`
	SenderName       string             `bson:"sender_name" json:"sender_name"`
	ReceiverName     string             `bson:"receiver_name" json:"receiver_name"`
	SenderImage      string             `bson:"sender_image" json:"sender_image"`
	ReceiverImage    string             `bson:"receiver_image" json:"receiver_image"`
}
