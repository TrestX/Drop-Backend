package wallet

import (
	"Drop/DropWallet/api"
	entity "Drop/DropWallet/entities"
	"Drop/DropWallet/repository/wallet"
	"errors"
	"strconv"
	"time"

	"github.com/aekam27/trestCommon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = wallet.NewWalletRepository("wallet")
)

type walletService struct{}

func NewWalletService(repository wallet.WalletRepository) WalletService {
	repo = repository
	return &walletService{}
}

func (r *walletService) AddWalletTransaction(wallet Wallet, token string) (string, error) {
	if wallet.UserID == "" {
		return "", errors.New("something went wrong")
	}
	data, err := checkWallet(wallet.UserID)
	newToken, _ := trestCommon.CreateToken(wallet.UserID, "", "", "")
	if err == nil {
		newHistory := data.History
		var addHistory entity.History
		addHistory.Amount = wallet.Amount
		addHistory.Description = wallet.Description
		addHistory.Type = wallet.Type
		newHistory = append(newHistory, addHistory)
		set := bson.M{"$set": bson.M{"history": newHistory, "updated_time": time.Now()}}
		filter := bson.M{"user_id": wallet.UserID}
		_, err := repo.UpdateOne(filter, set)
		if err != nil {
			return "", err
		}
	} else {
		var walletEntity entity.WalletDB
		walletEntity.ID = primitive.NewObjectID()
		walletEntity.UserID = wallet.UserID
		walletEntity.Status = "Active"
		walletEntity.CreatedTime = time.Now()
		var addHistory entity.History
		addHistory.Amount = wallet.Amount
		addHistory.Description = wallet.Description
		addHistory.Type = wallet.Type
		walletEntity.History = []entity.History{addHistory}
		_, err := repo.InsertOne(walletEntity)
		if err != nil {
			return "", err
		}
	}
	getUserDetails, err := api.GetUserDetails(newToken)
	if err != nil {
		return "", err
	}
	var profile api.Profile
	amt, _ := strconv.Atoi(getUserDetails.Wallet)
	wamt, _ := strconv.Atoi(wallet.Amount)
	if wallet.Type == "Add" {
		tA := amt + wamt
		profile.Wallet = strconv.Itoa(tA)
	} else if wallet.Type == "Subtract" {
		tA := amt - wamt
		profile.Wallet = strconv.Itoa(tA)
	}
	_, err = api.UpdateUserWallet(profile, wallet.UserID, newToken)
	if err != nil {
		return "", err
	}
	return "Success", nil
}

func checkWallet(userId string) (entity.WalletDB, error) {
	return repo.FindOne(bson.M{"user_id": userId}, bson.M{})
}

func (*walletService) GetWallet(userId, status string) (entity.WalletDB, error) {
	filter := bson.M{"user_id": userId}
	if status != "" {
		filter["status"] = status
	}
	wallet, err := repo.FindOne(filter, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"Get Wallet section",
			err,
		)
		return entity.WalletDB{}, err
	}
	return wallet, nil
}

func (*walletService) GetWalletByUserId(userId string) (entity.WalletDB, error) {
	filter := bson.M{"user_id": userId}
	wallet, err := repo.FindOne(filter, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"Get Wallet section",
			err,
		)
		return entity.WalletDB{}, err
	}
	return wallet, nil
}

func (*walletService) GetWalletWithIDs(walletIds []string) ([]entity.WalletDB, error) {
	subFilter := bson.A{}
	for _, item := range walletIds {
		id, _ := primitive.ObjectIDFromHex(item)
		subFilter = append(subFilter, bson.M{"_id": id})
	}
	filter := bson.M{"$or": subFilter}
	wallets, err := repo.FindWithIDs(filter, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"Get Wallet section",
			err,
		)
		return []entity.WalletDB{}, err
	}
	return wallets, nil
}
