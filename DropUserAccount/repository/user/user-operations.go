package user

import (
	entity "Drop/DropUserAccount/entities"
	"context"
	"errors"

	"github.com/aekam27/trestCommon"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type repo struct {
	CollectionName string
}

//NewFirestoreRepository creates a new repo
func NewProfileRepository(collectionName string) UserRepository {
	return &repo{
		CollectionName: collectionName,
	}
}

//used by signup
func (r *repo) InsertOne(document interface{}) (string, error) {
	user, err := trestCommon.InsertOne(document, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"insert profile",
			err,
			logrus.Fields{
				"document":        document,
				"collection name": r.CollectionName,
			})
		return "", err
	}
	userid := user.InsertedID.(primitive.ObjectID).Hex()
	return userid, nil
}

//used by update profile ,login and email verifcation
func (r *repo) UpdateOne(filter, update bson.M) (string, error) {
	result, err := trestCommon.UpdateOne(filter, update, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"update profile",
			err,
			logrus.Fields{
				"filter":          filter,
				"update":          update,
				"collection name": r.CollectionName,
			})

		return "", err
	}
	if result.MatchedCount == 0 || result.ModifiedCount == 0 {
		err = errors.New("profile not found(404)")
		trestCommon.ECLog3(
			"update profile",
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

//used by get profile ,login and email verification
func (r *repo) FindOne(filter, projection bson.M) (entity.UserDB, error) {
	var profile entity.UserDB
	err := trestCommon.FindOne(filter, projection, r.CollectionName).Decode(&profile)
	if err != nil {
		trestCommon.ECLog3(
			"Find profile",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return profile, err
	}
	return profile, err
}

//not used may use in future for gettin list of profiles
func (r *repo) Find(filter, projection bson.M) ([]entity.UserDB, error) {
	var profiles []entity.UserDB
	cursor, err := trestCommon.Find(filter, projection, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find profiles",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var profile entity.UserDB
		if err = cursor.Decode(&profile); err != nil {
			trestCommon.ECLog3(
				"Find profiles",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return profiles, nil
		}
		profile.Password = ""
		profiles = append(profiles, profile)
	}
	return profiles, nil
}

//not using
func (r *repo) DeleteOne(filter bson.M) error {
	deleteResult, err := trestCommon.DeleteOne(filter, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"delete profile",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	if deleteResult.DeletedCount == 0 {
		err = errors.New("profile not found(404)")
		trestCommon.ECLog3(
			"delete profile",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	return nil
}

func (r *repo) FindWithIDs(filter, projection bson.M) ([]entity.UserDB, error) {
	var users []entity.UserDB
	cursor, err := trestCommon.Find(filter, projection, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find users",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var user entity.UserDB
		if err = cursor.Decode(&user); err != nil {
			trestCommon.ECLog3(
				"Find users",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return users, nil
		}
		users = append(users, user)
	}
	return users, nil
}
