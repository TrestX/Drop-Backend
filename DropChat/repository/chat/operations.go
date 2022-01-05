package chat

import (
	entity "Drop/DropChat/entities"
	"Drop/DropChat/log"
	"Drop/DropChat/model"
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type repo struct {
	CollectionName string
}

//NewFirestoreRepository creates a new repo
func NewChatRepository(collectionName string) ChatRepository {
	return &repo{
		CollectionName: collectionName,
	}
}

//used by signup
func (r *repo) InsertOne(document interface{}) (string, error) {
	_, err := model.InsertOne(document, r.CollectionName)
	if err != nil {
		log.ECLog3(
			"insert item",
			err,
			logrus.Fields{
				"document":        document,
				"collection name": r.CollectionName,
			})
		return "", err
	}
	return "Added Successfully", nil
}

//used by update chat ,login and email verifcation
func (r *repo) UpdateOne(filter, update bson.M) (string, error) {
	result, err := model.UpdateOne(filter, update, r.CollectionName)
	if err != nil {
		log.ECLog3(
			"update chat",
			err,
			logrus.Fields{
				"filter":          filter,
				"update":          update,
				"collection name": r.CollectionName,
			})

		return "", err
	}
	if result.MatchedCount == 0 || result.ModifiedCount == 0 {
		err = errors.New("chat not found(404)")
		log.ECLog3(
			"update chat",
			err,
			logrus.Fields{
				"filter":          filter,
				"update":          update,
				"collection name": r.CollectionName,
			})
		return "", err
	}
	return "updated successfully", nil
}

//used by get chat ,login and email verification
func (r *repo) FindOne(filter, projection bson.M) (entity.ChatDB, error) {
	var chat entity.ChatDB
	err := model.FindOne(filter, projection, r.CollectionName).Decode(&chat)
	if err != nil {
		log.ECLog3(
			"Find chat",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return chat, err
	}
	return chat, err
}

//not used may use in future for gettin list of chat
func (r *repo) Find(filter, projection bson.M, limit, skip int) ([]entity.ChatDB, error) {
	var carts []entity.ChatDB
	cursor, err := model.FindWithLimitAndOffSet(filter, projection, limit, skip, r.CollectionName)
	if err != nil {
		log.ECLog3(
			"Find chat",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var chat entity.ChatDB
		if err = cursor.Decode(&chat); err != nil {
			log.ECLog3(
				"Find chat",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return nil, err
		}
		carts = append(carts, chat)
	}
	return carts, nil
}

//not using
func (r *repo) DeleteOne(filter bson.M) error {
	deleteResult, err := model.DeleteOne(filter, r.CollectionName)
	if err != nil {
		log.ECLog3(
			"delete chat",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	if deleteResult.DeletedCount == 0 {
		err = errors.New("chat not found(404)")
		log.ECLog3(
			"delete chat",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	return nil
}
func (r *repo) DeleteMany(filter bson.M) error {
	deleteResult, err := model.DeleteOne(filter, r.CollectionName)
	if err != nil {
		log.ECLog3(
			"delete chat",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	if deleteResult.DeletedCount == 0 {
		err = errors.New("chat not found(404)")
		log.ECLog3(
			"delete chat",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	return nil
}

func (r *repo) FindWithIDs(filter, projection bson.M) ([]entity.ChatDB, error) {
	var carts []entity.ChatDB
	cursor, err := model.Find(filter, projection, r.CollectionName)
	if err != nil {
		log.ECLog3(
			"Find item",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var chat entity.ChatDB
		if err = cursor.Decode(&chat); err != nil {
			log.ECLog3(
				"Find item",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return carts, nil
		}
		carts = append(carts, chat)
	}
	return carts, nil
}
