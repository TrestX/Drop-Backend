package cart

import (
	"Drop/DropCart/api"
	entity "Drop/DropCart/entities"
	"Drop/DropCart/repository/cart"
	"errors"
	"strings"
	"time"

	"github.com/aekam27/trestCommon"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = cart.NewCartRepository("cart")
)

type cartService struct{}

func NewCartService(repository cart.UserRepository) CartService {
	repo = repository
	return &cartService{}
}

func (r *cartService) AddCart(shopID, userId, cartID string, items Items, token string) (string, string, error) {
	var cartEntity entity.CartDB
	if userId == "" {
		return "", "", errors.New("something went wrong")
	}

	cartEntity.ID = primitive.NewObjectID()
	cartEntity.UserID = userId
	keys := make(map[string]bool)
	nilist := []entity.Item{}
	for i := 0; i < len(items.Items); i++ {
		if _, value := keys[items.Items[i].ItemID]; !value {
			keys[items.Items[i].ItemID] = true
			nilist = append(nilist, items.Items[i])
		}
	}
	cartEntity.Items = nilist
	cartEntity.AddOns = items.AddOns
	cartEntity.ShopID = shopID
	cartEntity.Status = "ADDED"
	cartEntity.AddedTime = time.Now()
	cart, err := repo.FindOne(bson.M{"shopid": shopID, "status": "ADDED", "user_id": userId}, bson.M{})
	if err == nil {
		for i := 0; i < len(cartEntity.Items); i++ {
			for j := 0; j < len(cart.Items); j++ {
				if cart.Items[j].ItemID == cartEntity.Items[i].ItemID {
					cart.Items[j].Quantity = cartEntity.Items[i].Quantity
					break
				}
			}
		}
		cart.Items = append(cart.Items, cartEntity.Items...)
		l := []entity.Item{}
		ids := []string{}
		for i := 0; i < len(cart.Items); i++ {
			if !contains(ids, cart.Items[i].ItemID) && cart.Items[i].Quantity > 0 {
				ids = append(ids, cart.Items[i].ItemID)
				l = append(l, cart.Items[i])
			}
		}

		set := bson.M{"$set": bson.M{"items": l, "updated_time": time.Now()}}
		if len(l) == 0 {
			set = bson.M{"$set": bson.M{"items": l, "status": "Deleted", "updated_time": time.Now()}}
		}
		if len(cart.AddOns) > 0 {
			set = bson.M{"$set": bson.M{"items": l, "updated_time": time.Now(), "addons": cartEntity.AddOns}}
		}
		filter := bson.M{"_id": cart.ID}
		res, err := repo.UpdateOne(filter, set)
		if err != nil {
			return "", "", err
		}
		return res, cart.ID.Hex(), nil
	}
	_ = getPrevioudCart(userId, r)
	res, err := repo.InsertOne(cartEntity)
	if err != nil {
		return "", "", err
	}
	return res, cartEntity.ID.Hex(), nil
}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func getPrevioudCart(userId string, r *cartService) error {
	cart, err := repo.FindOne(bson.M{"user_id": userId, "status": "ADDED"}, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"Get Cart section",
			err,
		)
		return err
	}
	if cart.UserID == "" {
		return nil
	} else {
		_, err = r.UpdateCart(userId, "Deleted")
		if err != nil {
			return err
		}
		return nil
	}
}

func (*cartService) UpdateCart(userId, status string) (string, error) {
	if userId == "" {
		return "", errors.New("something went wrong")
	}
	set := bson.M{"$set": bson.M{"status": status, "updated_time": time.Now()}}
	filter := bson.M{"user_id": userId, "status": "ADDED"}
	return repo.UpdateOne(filter, set)
}

func (*cartService) GetCart(userId string) (CartGetSchema, error) {

	cart, err := repo.FindOne(bson.M{"user_id": userId, "status": "ADDED"}, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"Get Cart section",
			err,
		)
		return CartGetSchema{}, err
	}
	var itemList []entity.ItemDB
	for i := 0; i < len(cart.Items); i++ {
		itemDetails, _ := api.GetItem(cart.Items[i].ItemID)
		itemDetails.Quantity = cart.Items[i].Quantity
		itemDetails.AddOns = cart.AddOns
		itemDetails.ChoiceSelected = cart.Items[i].ChoiceSelected
		itemDetails.Size = cart.Items[i].Size
		newdownloadurl := createPreSignedDownloadUrl(itemDetails.Images[0])
		itemDetails.Images = []string{newdownloadurl}
		itemList = append(itemList, itemDetails)
	}
	body := constructOutput(cart, itemList)

	return body, nil
}

func constructOutput(cart entity.CartDB, itemList []entity.ItemDB) CartGetSchema {
	var cartOutput CartGetSchema
	cartOutput.ID = cart.ID.Hex()
	cartOutput.Items = itemList
	cartOutput.UserID = cart.UserID
	cartOutput.ShopID = cart.ShopID
	cartOutput.Status = cart.Status
	cartOutput.AddOns = cart.AddOns
	cartOutput.AddedTime = cart.AddedTime
	cartOutput.UpdatedTime = cart.UpdatedTime
	return cartOutput
}
func createPreSignedDownloadUrl(url string) string {
	s := strings.Split(url, "?")
	if len(s) > 0 {
		o := strings.Split(s[0], "/")
		if len(o) > 3 {
			fileName := o[4]
			path := o[3]
			downUrl, _ := trestCommon.PreSignedDownloadUrl(fileName, path)
			return downUrl
		}
	}
	return ""
}
func (*cartService) GetCartWithIDs(orderIds []string) ([]CartGetSchema, error) {
	subFilter := bson.A{}
	for _, item := range orderIds {
		id, _ := primitive.ObjectIDFromHex(item)
		subFilter = append(subFilter, bson.M{"_id": id})
	}
	filter := bson.M{"$or": subFilter}
	carts, err := repo.FindWithIDs(filter, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"Get Carts section",
			err,
		)
		return []CartGetSchema{}, err
	}
	var oPList []CartGetSchema
	for i := 0; i < len(carts); i++ {
		var itemList []entity.ItemDB
		for j := 0; j < len(carts[i].Items); j++ {
			itemDetails, _ := api.GetItem(carts[i].Items[j].ItemID)
			itemDetails.Quantity = carts[i].Items[j].Quantity
			itemDetails.AddOns = carts[i].AddOns
			itemDetails.ChoiceSelected = carts[i].Items[j].ChoiceSelected
			itemDetails.Size = carts[i].Items[j].Size
			newdownloadurl := createPreSignedDownloadUrl(itemDetails.Images[0])
			itemDetails.Images = []string{newdownloadurl}
			itemList = append(itemList, itemDetails)
		}
		body := constructOutput(carts[i], itemList)
		oPList = append(oPList, body)
	}

	return oPList, nil
}
