package item

import (
	"Drop/DropItems/api"
	entity "Drop/DropItems/entities"
	"math/rand"
	"strings"

	"Drop/DropItems/repository/item"
	"errors"
	"time"

	"github.com/aekam27/trestCommon"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = item.NewItemRepository("item")
)

type itemService struct{}

func NewItemService(repository item.UserRepository) ItemService {
	repo = repository
	return &itemService{}
}

func (*itemService) AddItem(item *Item, userId string) (string, error) {
	var itemEntity entity.ItemDB
	if item == nil {
		err := errors.New("item id missing")
		trestCommon.ECLog2(
			"add item section",
			err,
		)
		return "", err
	}
	itemEntity.ID = primitive.NewObjectID()
	itemEntity.Category = item.Category
	itemEntity.Name = item.Name
	itemEntity.SellerID = item.SellerID
	itemEntity.ShopID = item.ShopID
	newtoken, _ := trestCommon.CreateToken(item.SellerID, "", "", "")
	shopDetails, _ := api.GetShopDetails(item.ShopID, newtoken)
	itemEntity.ShopType = shopDetails.Type
	itemEntity.Images = item.Images
	itemEntity.Description = item.Description
	itemEntity.Quantity = item.Quantity
	itemEntity.AddOns = item.AddOns
	itemEntity.Description = item.Description
	itemEntity.CreatedTime = time.Now()
	itemEntity.Type = item.Type
	itemEntity.Price = item.Price
	itemEntity.Matrix = item.Matrix
	itemEntity.Sizes = item.Sizes
	itemEntity.Choices = item.Choices
	return repo.InsertOne(itemEntity)
}

func (*itemService) UpdateItem(item *Item, itemid string) (string, error) {
	if itemid == "" {
		err := errors.New("item id missing")
		trestCommon.ECLog2(
			"update item section",
			err,
		)
		return "", err
	}
	id, _ := primitive.ObjectIDFromHex(itemid)
	setParameters := bson.M{}
	if item.Category != "" {
		setParameters["category"] = item.Category
	}
	if item.Name != "" {
		setParameters["sub_category"] = item.Name
	}
	if item.Deal != "" {
		setParameters["deal"] = item.Deal
	}
	if len(item.AddOns) > 0 {
		setParameters["add_ons"] = item.AddOns
	}
	if item.Description != "" {
		setParameters["description"] = item.Description
	}
	if item.Quantity != 0 {
		setParameters["quantity"] = item.Quantity
	}
	if len(item.Images) > 0 {
		setParameters["images"] = item.Images
	}
	if item.Approved {
		setParameters["approved"] = true
	}
	if item.Rejected {
		setParameters["rejected"] = true
	}
	if item.Featured {
		setParameters["featured"] = true
	}
	if !item.Featured {
		setParameters["featured"] = false
	}
	if item.FeaturedApp {
		setParameters["featured_app"] = true
	}
	if !item.FeaturedApp {
		setParameters["featured_app"] = false
	}
	if item.Type != "" {
		setParameters["type"] = item.Type
	}
	if item.Price != 0 {
		setParameters["price"] = item.Price
	}
	if len(item.Choices) > 0 {
		setParameters["choices"] = item.Choices
	}
	setParameters["updated_time"] = time.Now()
	filter := bson.M{"_id": id}
	set := bson.M{
		"$set": setParameters,
	}
	result, err := repo.UpdateOne(filter, set)
	if err != nil {
		trestCommon.ECLog3(
			"update item section",
			err,
			logrus.Fields{
				"item_id": itemid,
				"item":    item,
			})
		return "", err
	}

	return result, nil
}

func (*itemService) GetItem(shopID, category, name, typee, sellerId, search, featured string, limit, skip int) ([]entity.ItemDB, error) {
	filter := bson.M{}
	if category != "" {
		filter["category"] = category
	}
	if name != "" {
		filter["name"] = name
	}
	if sellerId != "" {
		filter["seller_id"] = sellerId
	}
	if shopID != "" {
		filter["shop_id"] = shopID
	}
	if typee != "" {
		filter["type"] = typee
	}
	if featured == "1" {
		filter["featured"] = true
	}
	if search != "" {
		filter["$or"] = bson.A{
			bson.M{"category": bson.M{"$regex": search, "$options": "i"}},
			bson.M{"name": bson.M{"$regex": search, "$options": "i"}},
		}
	}
	item, err := repo.Find(filter, bson.M{}, limit, skip)
	if err != nil {
		trestCommon.ECLog2(
			"GetItem section",
			err,
		)
		return item, err
	}
	return item, nil
}
func (*itemService) GetSellerItem(sellerId, shopID string, limit, skip int) ([]entity.ItemDB, error) {

	item, err := repo.Find(bson.M{"seller_id": sellerId, "shop_id": shopID}, bson.M{}, limit, skip)
	if err != nil {
		trestCommon.ECLog2(
			"GetItem section",
			err,
		)
		return item, err
	}
	return item, nil
}
func (*itemService) GetItemUsingID(itemId string) (entity.ItemDB, error) {
	if itemId == "" {
		err := errors.New("item id missing")
		trestCommon.ECLog2(
			"update item section",
			err,
		)
		return entity.ItemDB{}, err
	}
	id, _ := primitive.ObjectIDFromHex(itemId)
	return repo.FindOne(bson.M{"_id": id}, bson.M{})

}
func (*itemService) GetItemWithIDs(itemIds []string) ([]entity.ItemDB, error) {
	subFilter := bson.A{}
	for _, item := range itemIds {
		id, _ := primitive.ObjectIDFromHex(item)
		subFilter = append(subFilter, bson.M{"_id": id})
	}
	filter := bson.M{"$or": subFilter}
	return repo.FindWithIDs(filter, bson.M{})
}
func (*itemService) GetFeaturedItem(shopID, category, shopType string, limit, skip int) ([]entity.ItemDB, error) {
	filter := bson.M{"featured_app": true}
	if category != "" {
		if strings.Contains(category, ",") {
			l := strings.Split(category, ",")
			subFilter := bson.A{}
			for i := 0; i < len(l); i++ {
				subFilter = append(subFilter, bson.M{"category": bson.M{"$regex": l[i], "$options": "i"}})
			}
			filter = bson.M{"$or": subFilter}
		} else {
			filter["category"] = category
		}

	}
	if shopType != "" {
		if strings.Contains(shopType, ",") {
			l := strings.Split(shopType, ",")
			subFilter := bson.A{}
			for i := 0; i < len(l); i++ {
				subFilter = append(subFilter, bson.M{"shop_type": bson.M{"$regex": l[i], "$options": "i"}})
			}
			filter = bson.M{"$or": subFilter}
		} else {
			filter["shop_type"] = shopType
		}

	}
	if shopID != "" {
		if strings.Contains(shopID, ",") {
			l := strings.Split(shopID, ",")
			subFilter := bson.A{}
			for i := 0; i < len(l); i++ {
				subFilter = append(subFilter, bson.M{"shop_id": bson.M{"$regex": l[i], "$options": "i"}})
			}
			filter = bson.M{"$or": subFilter}
		} else {
			filter["shop_id"] = shopID
		}
	}
	item, err := repo.Find(filter, bson.M{}, limit, skip)
	if err != nil {
		trestCommon.ECLog2(
			"GetItem section",
			err,
		)
		return item, err
	}
	return item, nil
}
func (*itemService) GetShopFeaturedItem(shopID, category, shopType string, limit, skip int) ([]entity.ItemDB, error) {
	filter := bson.M{"featured": true}
	if category != "" {
		if strings.Contains(category, ",") {
			l := strings.Split(category, ",")
			subFilter := bson.A{}
			for i := 0; i < len(l); i++ {
				subFilter = append(subFilter, bson.M{"category": bson.M{"$regex": l[i], "$options": "i"}})
			}
			filter = bson.M{"$or": subFilter}
		} else {
			filter["category"] = category
		}

	}
	if shopType != "" {
		if strings.Contains(shopType, ",") {
			l := strings.Split(shopType, ",")
			subFilter := bson.A{}
			for i := 0; i < len(l); i++ {
				subFilter = append(subFilter, bson.M{"shop_type": bson.M{"$regex": l[i], "$options": "i"}})
			}
			filter = bson.M{"$or": subFilter}
		} else {
			filter["shop_type"] = shopType
		}

	}
	if shopID != "" {
		if strings.Contains(shopID, ",") {
			l := strings.Split(shopID, ",")
			subFilter := bson.A{}
			for i := 0; i < len(l); i++ {
				subFilter = append(subFilter, bson.M{"shop_id": bson.M{"$regex": l[i], "$options": "i"}})
			}
			filter = bson.M{"$or": subFilter}
		} else {
			filter["shop_id"] = shopID
		}
	}
	item, err := repo.Find(filter, bson.M{}, limit, skip)
	if err != nil {
		trestCommon.ECLog2(
			"GetItem section",
			err,
		)
		return item, err
	}
	return item, nil
}

func (*itemService) GetPopularItem(shopID, category, shopType string, limit, skip int) ([]entity.ItemDB, error) {
	filter := bson.M{}
	if category != "" {
		if strings.Contains(category, ",") {
			l := strings.Split(category, ",")
			subFilter := bson.A{}
			for i := 0; i < len(l); i++ {
				subFilter = append(subFilter, bson.M{"category": bson.M{"$regex": l[i], "$options": "i"}})
			}
			filter = bson.M{"$or": subFilter}
		} else {
			filter["category"] = category
		}

	}
	if shopType != "" {
		if strings.Contains(shopType, ",") {
			l := strings.Split(shopType, ",")
			subFilter := bson.A{}
			for i := 0; i < len(l); i++ {
				subFilter = append(subFilter, bson.M{"shop_type": bson.M{"$regex": l[i], "$options": "i"}})
			}
			filter = bson.M{"$or": subFilter}
		} else {
			filter["shop_type"] = shopType
		}

	}
	if shopID != "" {
		if strings.Contains(shopID, ",") {
			l := strings.Split(shopID, ",")
			subFilter := bson.A{}
			for i := 0; i < len(l); i++ {
				subFilter = append(subFilter, bson.M{"shop_id": bson.M{"$regex": l[i], "$options": "i"}})
			}
			filter = bson.M{"$or": subFilter}
		} else {
			filter["shop_id"] = shopID
		}
	}
	item, err := repo.Find(filter, bson.M{}, limit, skip)
	if err != nil {
		trestCommon.ECLog2(
			"GetItem section",
			err,
		)
		return item, err
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(item), func(i, j int) {
		item[i], item[j] = item[j], item[i]
	})
	return item, nil
}

func (*itemService) GetTopRatedItems(shopID, category, shopType string, limit, skip int) ([]entity.ItemDB, error) {
	filter := bson.M{}
	if category != "" {
		if strings.Contains(category, ",") {
			l := strings.Split(category, ",")
			subFilter := bson.A{}
			for i := 0; i < len(l); i++ {
				subFilter = append(subFilter, bson.M{"category": bson.M{"$regex": l[i], "$options": "i"}})
			}
			filter = bson.M{"$or": subFilter}
		} else {
			filter["category"] = category
		}

	}
	if shopType != "" {
		if strings.Contains(shopType, ",") {
			l := strings.Split(shopType, ",")
			subFilter := bson.A{}
			for i := 0; i < len(l); i++ {
				subFilter = append(subFilter, bson.M{"shop_type": bson.M{"$regex": l[i], "$options": "i"}})
			}
			filter = bson.M{"$or": subFilter}
		} else {
			filter["shop_type"] = shopType
		}

	}
	if shopID != "" {
		if strings.Contains(shopID, ",") {
			l := strings.Split(shopID, ",")
			subFilter := bson.A{}
			for i := 0; i < len(l); i++ {
				subFilter = append(subFilter, bson.M{"shop_id": bson.M{"$regex": l[i], "$options": "i"}})
			}
			filter = bson.M{"$or": subFilter}
		} else {
			filter["shop_id"] = shopID
		}
	}
	item, err := repo.Find(filter, bson.M{}, limit, skip)
	if err != nil {
		trestCommon.ECLog2(
			"GetItem section",
			err,
		)
		return item, err
	}
	return item, nil
}
func (*itemService) GetItemCategoryStructured(shopID, popular, category, deal, name, typee, sellerId, search, featured, shoptype string, limit, skip int) (map[string][]interface{}, error) {
	filter := bson.M{}
	if deal != "" {
		if strings.Contains(deal, ",") {
			l := strings.Split(deal, ",")
			subFilter := bson.A{}
			for i := 0; i < len(l); i++ {
				subFilter = append(subFilter, bson.M{"deal": bson.M{"$regex": l[i], "$options": "i"}})
			}
			filter = bson.M{"$or": subFilter}
		} else {
			filter["deal"] = bson.M{"$regex": deal, "$options": "i"}
		}
	}
	if category != "" {
		if strings.Contains(category, ",") {
			l := strings.Split(category, ",")
			subFilter := bson.A{}
			for i := 0; i < len(l); i++ {
				subFilter = append(subFilter, bson.M{"category": bson.M{"$regex": l[i], "$options": "i"}})
			}
			filter = bson.M{"$or": subFilter}
		} else {
			filter["category"] = category
		}
	}
	if shoptype != "" {
		if strings.Contains(shoptype, ",") {
			l := strings.Split(shoptype, ",")
			subFilter := bson.A{}
			for i := 0; i < len(l); i++ {
				subFilter = append(subFilter, bson.M{"shop_type": bson.M{"$regex": l[i], "$options": "i"}})
			}
			filter = bson.M{"$or": subFilter}
		} else {
			filter["shop_type"] = shoptype
		}

	}
	if shopID != "" {
		if strings.Contains(shopID, ",") {
			l := strings.Split(shopID, ",")
			subFilter := bson.A{}
			for i := 0; i < len(l); i++ {
				subFilter = append(subFilter, bson.M{"shop_id": bson.M{"$regex": l[i], "$options": "i"}})
			}
			filter = bson.M{"$or": subFilter}
		} else {
			filter["shop_id"] = shopID
		}
	}
	if name != "" {
		if strings.Contains(name, ",") {
			l := strings.Split(name, ",")
			subFilter := bson.A{}
			for i := 0; i < len(l); i++ {
				subFilter = append(subFilter, bson.M{"name": bson.M{"$regex": l[i], "$options": "i"}})
			}
			filter = bson.M{"$or": subFilter}
		} else {
			filter["name"] = name
		}
	}
	if sellerId != "" {
		if strings.Contains(sellerId, ",") {
			l := strings.Split(sellerId, ",")
			subFilter := bson.A{}
			for i := 0; i < len(l); i++ {
				subFilter = append(subFilter, bson.M{"seller_id": bson.M{"$regex": l[i], "$options": "i"}})
			}
			filter = bson.M{"$or": subFilter}
		} else {
			filter["seller_id"] = sellerId
		}

	}
	if typee != "" {
		if strings.Contains(typee, ",") {
			l := strings.Split(typee, ",")
			subFilter := bson.A{}
			for i := 0; i < len(l); i++ {
				subFilter = append(subFilter, bson.M{"type": bson.M{"$regex": l[i], "$options": "i"}})
			}
			filter = bson.M{"$or": subFilter}
		} else {
			filter["type"] = typee
		}
	}
	if featured == "1" {
		filter["featured"] = true
	}
	if search != "" {
		filter["$or"] = bson.A{
			bson.M{"category": bson.M{"$regex": search, "$options": "i"}},
			bson.M{"name": bson.M{"$regex": search, "$options": "i"}},
		}
	}

	op := make(map[string][]interface{})
	item, err := repo.Find(filter, bson.M{}, limit, skip)
	var shopIds []string
	var ratingReview []string
	for i := 0; i < len(item); i++ {
		shopIds = append(shopIds, item[i].ShopID)
		ratingReview = append(ratingReview, item[i].ID.Hex())
	}
	shopDetails, _ := api.GetShopDetailsByIDs(shopIds)
	for i := 0; i < len(item); i++ {
		newdownloadurl := createPreSignedDownloadUrl(item[i].Images[0])
		item[i].Images[0] = newdownloadurl
		if _, ok := op[item[i].Category]; ok {
			var l = op[item[i].Category]
			var stru UOPStruct
			stru.Item = item[i]
			for _, shopDetail := range shopDetails {
				if shopDetail.ID.Hex() == item[i].ShopID {
					newdownloadurl := createPreSignedDownloadUrl(shopDetail.ShopLogo)
					shopDetail.ShopLogo = newdownloadurl
					newdownloadurl = createPreSignedDownloadUrl(shopDetail.ShopBanner)
					shopDetail.ShopBanner = newdownloadurl
					stru.Shop = shopDetail
					break
				}
			}
			review, _ := api.GetOrderReview(item[i].ID.Hex(), " ")
			arview := 0
			if len(review) > 0 {
				rview := 0
				for k := 0; k < len(review); k++ {
					rview = rview + int(review[k].Rating)
				}
				arview = rview / len(review)
			}
			stru.AvgRating = arview
			stru.Ratings = review
			l = append(l, stru)
			op[item[i].Category] = l
			if i%3 == 0 {
				op["Popular"] = l
			}
		} else {
			var l []interface{}
			var stru UOPStruct
			stru.Item = item[i]
			for _, shopDetail := range shopDetails {
				if shopDetail.ID.Hex() == item[i].ShopID {
					newdownloadurl := createPreSignedDownloadUrl(shopDetail.ShopLogo)
					shopDetail.ShopLogo = newdownloadurl
					newdownloadurl = createPreSignedDownloadUrl(shopDetail.ShopBanner)
					shopDetail.ShopBanner = newdownloadurl
					stru.Shop = shopDetail
					break
				}
			}
			review, _ := api.GetOrderReview(item[i].ID.Hex(), " ")
			arview := 0
			if len(review) > 0 {
				rview := 0
				for k := 0; k < len(review); k++ {
					rview = rview + int(review[k].Rating)
				}
				arview = rview / len(review)
			}
			stru.AvgRating = arview
			stru.Ratings = review
			l = append(l, stru)
			st := item[i].Category
			op[st] = l
			if i%3 == 0 {
				op["Popular"] = l
			}
		}
	}
	if err != nil {
		trestCommon.ECLog2(
			"GetItem section",
			err,
		)
		return op, err
	}

	return op, nil
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
