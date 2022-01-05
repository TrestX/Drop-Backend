package chat

import (
	"Drop/DropChat/api"
	entity "Drop/DropChat/entities"
	"Drop/DropChat/log"
	"Drop/DropChat/repository/chat"
	util "Drop/DropChat/utils"
	"errors"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = chat.NewChatRepository("chat")
)

type chatService struct{}

func NewChatService(repository chat.ChatRepository) ChatService {
	repo = repository
	return &chatService{}
}

func (r *chatService) AddChat(userId string, chat Chats) (string, error) {
	if userId == "" {
		return "", errors.New("something went wrong")
	}
	var entityChat entity.ChatDB
	entityChat.ID = primitive.NewObjectID()
	entityChat.SenderID = userId
	entityChat.AddedTime = time.Now()
	entityChat.ReceiverID = ""
	entityChat.Status = "Added"
	util.SendNotificationWS("chat", "New", "broadcast", userId)
	_, err := repo.InsertOne(entityChat)
	if err != nil {
		return "", err
	}
	return entityChat.ID.Hex(), nil
}

func (r *chatService) UpdateChat(userId string, chat Chats) (string, error) {
	if userId == "" {
		return "", errors.New("something went wrong")
	}
	data, err := checkForChatSessionID(chat.ChatID)
	if err != nil {
		return "", errors.New("somethingwent wrong")
	}
	setFil := bson.M{}
	if data.SenderID != userId {
		if data.ReceiverID == "" {
			setFil["receiver_id"] = userId
			setFil["updated_time"] = time.Now()
			setFil["receiver_join_time"] = time.Now()
			setFil["status"] = "Started"
			util.SendNotificationWS("chat", "Taken", "broadcast", userId)
		}
	}
	if chat.Status != "" {
		setFil["status"] = chat.Status
	}
	if len(chat.Chat) > 0 {
		var newChat entity.Chat
		newChat.Sender = chat.Chat[0].Sender
		newChat.Message = chat.Chat[0].Message
		newChat.Time = time.Now()
		token, _ := util.CreateToken(chat.Chat[0].Sender, "", "", "")
		user, _ := api.GetUserDetails(token)
		if err != nil {
			newChat.Name = "U"
		} else {
			newChat.Name = user.Name
		}
		data.Chat = append(data.Chat, newChat)
		setFil["chat"] = data.Chat
		setFil["updated_time"] = time.Now()
	}
	set := bson.M{"$set": setFil}
	id, _ := primitive.ObjectIDFromHex(chat.ChatID)
	filter := bson.M{"_id": id}
	_, err = repo.UpdateOne(filter, set)
	if err != nil {
		return "", err
	}
	nCdata, err := checkForChatSessionID(chat.ChatID)
	if err != nil {
		return "", errors.New("somethingwent wrong")
	}
	if nCdata.ReceiverID != "" && nCdata.SenderID != "" {
		util.SendNotificationWS("chat", "started", chat.ChatID, userId)
		return "started", nil
	} else {
		util.SendNotificationWS("chat", "New", chat.ChatID, userId)
	}
	return "updated", nil
}

func checkForChatSessionID(chatId string) (entity.ChatDB, error) {
	id, _ := primitive.ObjectIDFromHex(chatId)
	return repo.FindOne(bson.M{"_id": id}, bson.M{})
}

func (*chatService) GetChatForUser(user_id, delivery_person_id string) (entity.ChatDB, error) {
	filter := bson.M{"$or": bson.A{bson.M{"sender_id": user_id, "receiver_id": delivery_person_id}, bson.M{"receiver_id": user_id, "sender_id": delivery_person_id}}}
	chats, err := repo.FindOne(filter, bson.M{})
	if err != nil {
		log.ECLog2(
			"Get Cart section",
			err,
		)
		return entity.ChatDB{}, err
	}
	return chats, nil
}
func (*chatService) GetChat(chatId string) (entity.ChatDB, error) {
	id, _ := primitive.ObjectIDFromHex(chatId)
	chats, err := repo.FindOne(bson.M{"_id": id}, bson.M{})
	if err != nil {
		log.ECLog2(
			"Get Cart section",
			err,
		)
		return entity.ChatDB{}, err
	}
	return chats, nil
}
func (*chatService) GetChats(userId, chatId, status string) ([]OP, error) {
	filter := bson.M{}
	if userId != "" {
		filter["sender_id"] = userId
	}
	if status != "" {
		filter["status"] = status
	}
	if chatId != "" {
		id, _ := primitive.ObjectIDFromHex(chatId)
		filter["_id"] = id
	}
	chats, err := repo.Find(filter, bson.M{}, 100, 0)
	if err != nil {
		log.ECLog2(
			"Get Cart section",
			err,
		)
		return []OP{}, err
	}
	uIList := []string{}
	for i := 0; i < len(chats); i++ {
		uIList = append(uIList, chats[i].SenderID)
	}
	user, _ := api.GetUsersDetailsByIDs(uIList)
	var oPList []OP
	for i := 0; i < len(chats); i++ {
		var eTC OP
		eTC.ID = chats[i].ID
		eTC.AddedTime = chats[i].AddedTime
		eTC.Chat = chats[i].Chat
		eTC.ReceiverID = chats[i].ReceiverID
		eTC.ReceiverJoinTime = chats[i].ReceiverJoinTime
		eTC.SenderID = chats[i].SenderID
		eTC.Status = chats[i].Status
		eTC.StopTime = chats[i].StopTime
		eTC.UpdatedTime = chats[i].UpdatedTime
		for j := 0; j < len(user); j++ {
			if user[j].ID.Hex() == chats[i].SenderID {
				eTC.Email = user[j].Email
				eTC.Name = user[j].Name
				newPdownloadurl := createPreSignedDownloadUrl(user[j].ProfilePhoto)
				eTC.ProfilePhoto = newPdownloadurl
				break
			}
		}
		oPList = append(oPList, eTC)
	}
	return oPList, nil
}

func createPreSignedDownloadUrl(url string) string {
	s := strings.Split(url, "?")
	if len(s) > 0 {
		o := strings.Split(s[0], "/")
		if len(o) > 3 {
			fileName := o[4]
			path := o[3]
			downUrl, _ := util.PreSignedDownloadUrl(fileName, path)
			return downUrl
		}
	}
	return ""
}
