package shop

import (
	"Drop/DropShop/api"
	entity "Drop/DropShop/entities"
	"Drop/DropShop/repository/shop"
	"strconv"

	"github.com/aekam27/trestCommon"

	"errors"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = shop.NewShopRepository("shop")
)

type shopService struct{}

func NewShopService(repository shop.ShopRepository) ShopService {
	repo = repository
	return &shopService{}
}

func (add *shopService) AddShop(shop *Shop, sellerId string) (string, error) {
	var shopEntity entity.ShopDB
	if shop == nil {
		err := errors.New("shop id missing")
		trestCommon.ECLog2(
			"add shop section",
			err,
		)
		return "", err
	}
	data, err := add.GetPrimaryShop(sellerId)
	if err == nil && shop.Primary {
		id := data.ID
		setParameters := bson.M{"primary": false}

		setParameters["updated_time"] = time.Now()
		filter := bson.M{"_id": id}
		set := bson.M{
			"$set": setParameters,
		}
		_, err = repo.UpdateOne(filter, set)
		if err != nil {
			trestCommon.ECLog3(
				"update shop section",
				err,
				logrus.Fields{
					"shop_id": data.ID,
				})
		}

	}
	shopEntity.ID = primitive.NewObjectID()
	shopEntity.SellerID = sellerId
	shopEntity.Address = shop.Address
	shopEntity.City = shop.City
	shopEntity.Country = shop.Country
	shopEntity.State = shop.State
	shopEntity.Pin = shop.Pin
	geoLocation := []float64{shop.Longitude, shop.Latitude}
	shopEntity.GeoLocation = bson.M{"type": "Point", "coordinates": geoLocation}
	shopEntity.Primary = shop.Primary
	shopEntity.Rating = 0
	shopEntity.Tags = shop.Tags
	shopEntity.MinOrderAmount = shop.MinOrderAmount
	shopEntity.CreatedTime = time.Now()
	if shop.ShopDescription != "" {
		shopEntity.ShopDescription = shop.ShopDescription
	}
	if shop.Timing != "" {
		shopEntity.Timing = shop.Timing
	}
	if shop.ShopName != "" {
		shopEntity.ShopName = shop.ShopName
	}
	if shop.ShopStatus != "" {
		shopEntity.ShopStatus = shop.ShopStatus
	}
	shopPics := []string{}
	if len(shop.ShopPhotos) > 0 {
		l := []string{}
		for i := 0; i < len(shop.ShopPhotos); i++ {
			url := createPreSignedDownloadUrl(shop.ShopPhotos[i])
			l = append(l, url)
		}
		shopPics = l
	}
	shopEntity.ShopPhotos = shopPics
	if shop.ShopLogo != "" {
		url := createPreSignedDownloadUrl(shop.ShopLogo)
		shopEntity.ShopLogo = url
	}
	if shop.ShopBanner != "" {
		url := createPreSignedDownloadUrl(shop.ShopBanner)
		shopEntity.ShopBanner = url
	}
	shopEntity.Type = shop.Type
	shopEntity.Featured = shop.Featured
	result, err := repo.InsertOne(shopEntity)
	if err != nil {
		return "", err
	}
	return result, nil
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
func (*shopService) UpdateShop(shop *Shop, shopid string) (string, error) {
	if shopid == "" {
		err := errors.New("shop id missing")
		trestCommon.ECLog2(
			"update shop section",
			err,
		)
		return "", err
	}
	id, _ := primitive.ObjectIDFromHex(shopid)

	setParameters := bson.M{}

	if shop.Address != "" {
		setParameters["address"] = shop.Address
	}
	if shop.Featured {
		setParameters["featured"] = true
	}
	if !shop.Featured {
		setParameters["featured"] = false
	}
	if shop.Type != "" {
		setParameters["type"] = shop.Type
	}
	if shop.Deal != "" {
		setParameters["deal"] = shop.Deal
	}
	if shop.DeliveryType != "" {
		setParameters["delivery"] = shop.DeliveryType
	}
	if shop.Cuisine != "" {
		setParameters["tags"] = shop.Tags
	}
	if shop.State != "" {
		setParameters["state"] = shop.State
	}
	if shop.City != "" {
		setParameters["city"] = shop.City
	}
	if shop.Country != "" {
		setParameters["country"] = shop.Country
	}
	if shop.Pin != "" {
		setParameters["pin"] = shop.Pin
	}
	if shop.Timing != "" {
		setParameters["timing"] = shop.Timing
	}
	if shop.ShopDescription != "" {
		setParameters["description"] = shop.ShopDescription
	}
	if shop.ShopStatus != "" {
		setParameters["status"] = shop.ShopStatus
	}
	if shop.Latitude != 0 && shop.Longitude != 0 {
		geoLocation := []float64{shop.Longitude, shop.Latitude}
		setParameters["geo_location"] = bson.M{"type": "Point", "coordinates": geoLocation}
	}
	if len(shop.ShopPhotos) > 0 {
		l := []string{}
		for i := 0; i < len(shop.ShopPhotos); i++ {
			url := createPreSignedDownloadUrl(shop.ShopPhotos[i])
			l = append(l, url)
		}
		setParameters["shop_photos"] = l
	}
	if shop.ShopLogo != "" {
		url := createPreSignedDownloadUrl(shop.ShopLogo)
		setParameters["shop_logo"] = url
	}
	if shop.ShopBanner != "" {
		url := createPreSignedDownloadUrl(shop.ShopBanner)
		setParameters["shop_banner"] = url
	}
	if shop.Type != "" {
		setParameters["type"] = shop.Type
	}
	if shop.Rating > 0 {
		pshopR, _ := repo.FindOne(bson.M{"_id": id}, bson.M{})
		setParameters["rating"] = (pshopR.Rating*float64(shop.NumbofRating) + shop.Rating) / float64(shop.NumbofRating+1)
	}
	setParameters["updated_time"] = time.Now()
	filter := bson.M{"_id": id}
	set := bson.M{
		"$set": setParameters,
	}
	result, err := repo.UpdateOne(filter, set)
	if err != nil {
		trestCommon.ECLog3(
			"update shop section",
			err,
			logrus.Fields{
				"shop_id": shopid,
				"shop":    shop,
			})
		return "", err
	}

	return result, nil
}

func (add *shopService) PrimaryShop(shopid string, sellerId string) (string, error) {
	if shopid == "" {
		err := errors.New("shop id missing")
		trestCommon.ECLog2(
			"update shop section",
			err,
		)
		return "", err
	}
	data, err := add.GetPrimaryShop(sellerId)
	if err == nil {
		id := data.ID
		setParameters := bson.M{"primary": false}

		setParameters["updated_time"] = time.Now()
		filter := bson.M{"_id": id}
		set := bson.M{
			"$set": setParameters,
		}
		_, err = repo.UpdateOne(filter, set)
		if err != nil {
			trestCommon.ECLog3(
				"update shop section",
				err,
				logrus.Fields{
					"shop_id": shopid,
				})
		}

	}
	id, _ := primitive.ObjectIDFromHex(shopid)

	setParameters := bson.M{"primary": true}

	setParameters["updated_time"] = time.Now()
	filter := bson.M{"_id": id}
	set := bson.M{
		"$set": setParameters,
	}
	return repo.UpdateOne(filter, set)

}

func (*shopService) GetShop(sellerId string, limit, skip int) ([]entity.ShopDB, error) {

	shop, err := repo.Find(bson.M{"seller_id": sellerId}, bson.M{}, limit, skip)
	if err != nil {
		trestCommon.ECLog2(
			"GetShop section",
			err,
		)
		return shop, err
	}
	for i := 0; i < len(shop); i++ {
		nerBannerUrl := createPreSignedDownloadUrl(shop[i].ShopBanner)
		nerLogoUrl := createPreSignedDownloadUrl(shop[i].ShopLogo)
		shop[i].ShopBanner = nerBannerUrl
		shop[i].ShopLogo = nerLogoUrl
		newShop := []string{}
		for o := 0; o < len(shop[i].ShopPhotos); o++ {
			nerShopUrl := createPreSignedDownloadUrl(shop[i].ShopPhotos[o])
			newShop = append(newShop, nerShopUrl)
		}
		shop[i].ShopPhotos = newShop
	}
	return shop, nil
}
func (*shopService) GetShopUsingID(shopId string, userID string) (entity.ShopDB, error) {
	if shopId == "" {
		err := errors.New("shop id missing")
		trestCommon.ECLog2(
			"update shop section",
			err,
		)
		return entity.ShopDB{}, err
	}
	id, _ := primitive.ObjectIDFromHex(shopId)
	shop, err := repo.FindOne(bson.M{"_id": id}, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"GetShop section",
			err,
		)
		return shop, err
	}
	nerBannerUrl := createPreSignedDownloadUrl(shop.ShopBanner)
	nerLogoUrl := createPreSignedDownloadUrl(shop.ShopLogo)
	shop.ShopBanner = nerBannerUrl
	shop.ShopLogo = nerLogoUrl
	newShop := []string{}
	for o := 0; o < len(shop.ShopPhotos); o++ {
		nerShopUrl := createPreSignedDownloadUrl(shop.ShopPhotos[o])
		newShop = append(newShop, nerShopUrl)
	}
	shop.ShopPhotos = newShop
	return shop, err
}
func (*shopService) GetPrimaryShop(sellerID string) (entity.ShopDB, error) {
	shop, err := repo.FindOne(bson.M{"seller_id": sellerID, "primary": true}, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"GetShop section",
			err,
		)
		return shop, err
	}
	nerBannerUrl := createPreSignedDownloadUrl(shop.ShopBanner)
	nerLogoUrl := createPreSignedDownloadUrl(shop.ShopLogo)
	shop.ShopBanner = nerBannerUrl
	shop.ShopLogo = nerLogoUrl
	newShop := []string{}
	for o := 0; o < len(shop.ShopPhotos); o++ {
		nerShopUrl := createPreSignedDownloadUrl(shop.ShopPhotos[o])
		newShop = append(newShop, nerShopUrl)
	}
	shop.ShopPhotos = newShop
	return shop, err
}
func (*shopService) GetFeaturedShop(limit, skip int) ([]entity.ShopDB, error) {
	shop, err := repo.Find(bson.M{"featured": true}, bson.M{}, limit, skip)
	if err != nil {
		trestCommon.ECLog2(
			"GetShop section",
			err,
		)
		return shop, err
	}
	for i := 0; i < len(shop); i++ {
		nerBannerUrl := createPreSignedDownloadUrl(shop[i].ShopBanner)
		nerLogoUrl := createPreSignedDownloadUrl(shop[i].ShopLogo)
		shop[i].ShopBanner = nerBannerUrl
		shop[i].ShopLogo = nerLogoUrl
		newShop := []string{}
		for o := 0; o < len(shop[i].ShopPhotos); o++ {
			nerShopUrl := createPreSignedDownloadUrl(shop[i].ShopPhotos[o])
			newShop = append(newShop, nerShopUrl)
		}
		shop[i].ShopPhotos = newShop
	}
	return shop, nil
}
func (*shopService) SearchShopByType(typ string, limit, skip int) ([]entity.ShopDB, error) {
	shop, err := repo.Find(bson.M{"type": typ}, bson.M{}, limit, skip)
	if err != nil {
		trestCommon.ECLog2(
			"GetShop section",
			err,
		)
		return shop, err
	}
	for i := 0; i < len(shop); i++ {
		nerBannerUrl := createPreSignedDownloadUrl(shop[i].ShopBanner)
		nerLogoUrl := createPreSignedDownloadUrl(shop[i].ShopLogo)
		shop[i].ShopBanner = nerBannerUrl
		shop[i].ShopLogo = nerLogoUrl
		newShop := []string{}
		for o := 0; o < len(shop[i].ShopPhotos); o++ {
			nerShopUrl := createPreSignedDownloadUrl(shop[i].ShopPhotos[o])
			newShop = append(newShop, nerShopUrl)
		}
		shop[i].ShopPhotos = newShop
	}
	return shop, nil
}

func (add *shopService) AddShopAdmin(shop *Shop) (string, error) {
	var shopEntity entity.ShopDB
	if shop == nil {
		err := errors.New("shop id missing")
		trestCommon.ECLog2(
			"add shop section",
			err,
		)
		return "", err
	}
	data, err := add.GetPrimaryShop(shop.SellerID)
	if err == nil && shop.Primary {
		id := data.ID
		setParameters := bson.M{"primary": false}

		setParameters["updated_time"] = time.Now()
		filter := bson.M{"_id": id}
		set := bson.M{
			"$set": setParameters,
		}
		_, err = repo.UpdateOne(filter, set)
		if err != nil {
			trestCommon.ECLog3(
				"update shop section",
				err,
				logrus.Fields{
					"shop_id": data.ID,
				})
		}

	}
	shopEntity.ID = primitive.NewObjectID()
	shopEntity.SellerID = shop.SellerID
	shopEntity.Address = shop.Address
	shopEntity.City = shop.City
	shopEntity.Country = shop.Country
	shopEntity.State = shop.State
	shopEntity.Pin = shop.Pin
	geoLocation := []float64{shop.Longitude, shop.Latitude}
	shopEntity.GeoLocation = bson.M{"type": "Point", "coordinates": geoLocation}
	shopEntity.Primary = shop.Primary
	shopEntity.CreatedTime = time.Now()
	if shop.ShopDescription != "" {
		shopEntity.ShopDescription = shop.ShopDescription
	}
	if shop.Timing != "" {
		shopEntity.Timing = shop.Timing
	}
	if shop.ShopName != "" {
		shopEntity.ShopName = shop.ShopName
	}
	if shop.ShopStatus != "" {
		shopEntity.ShopStatus = shop.ShopStatus
	}
	if shop.Deal != "" {
		shopEntity.Deal = shop.Deal
	}
	if shop.DeliveryType != "" {
		shopEntity.DeliveryType = shop.DeliveryType
	}
	if shop.Cuisine != "" {
		shopEntity.Cuisine = shop.Cuisine
	}
	shopPics := []string{}
	if len(shop.ShopPhotos) > 0 {
		l := []string{}
		for i := 0; i < len(shop.ShopPhotos); i++ {
			url := createPreSignedDownloadUrl(shop.ShopPhotos[i])
			l = append(l, url)
		}
		shopPics = l
	}
	shopEntity.ShopPhotos = shopPics
	if shop.ShopLogo != "" {
		url := createPreSignedDownloadUrl(shop.ShopLogo)
		shopEntity.ShopLogo = url
	}
	if shop.ShopBanner != "" {
		url := createPreSignedDownloadUrl(shop.ShopBanner)
		shopEntity.ShopBanner = url
	}
	shopEntity.Type = shop.Type
	shopEntity.Featured = shop.Featured
	result, err := repo.InsertOne(shopEntity)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (*shopService) GetShopAdmin(limit, skip int, sellerId, sType, status, featured, deal, rating, priceu, pricel, lowest string, lat, long float64) ([]OpSchema, error) {
	filter := bson.M{}
	if sellerId != "" {
		if strings.Contains(sellerId, ",") {
			l := strings.Split(sellerId, ",")
			subFilter := bson.A{}
			for i := 0; i < len(l); i++ {
				subFilter = append(subFilter, bson.M{"seller_id": l[i]})
			}
			filter = bson.M{"$or": subFilter}
		} else {
			filter["seller_id"] = sellerId
		}

	}
	if sType != "" {
		if strings.Contains(sType, ",") {
			l := strings.Split(sType, ",")
			subFilter := bson.A{}
			for i := 0; i < len(l); i++ {
				subFilter = append(subFilter, bson.M{"type": l[i]})
			}
			filter = bson.M{"$or": subFilter}
		} else {
			filter["type"] = sType
		}
	}
	if featured != "" {
		if featured == "1" {
			filter["featured"] = true
		}
	}
	if status != "" {
		filter["shop_status"] = status
	}
	if deal != "" {
		if strings.Contains(deal, ",") {
			l := strings.Split(deal, ",")
			subFilter := bson.A{}
			for i := 0; i < len(l); i++ {
				subFilter = append(subFilter, bson.M{"deal": bson.M{"$regex": "/" + l[i] + "/", "$options": "i"}})
			}
			filter = bson.M{"$or": subFilter}
		} else {
			filter["deal"] = bson.M{"$regex": "/" + deal + "/", "$options": "i"}
		}
	}
	if rating != "" {
		rat, _ := strconv.Atoi(rating)
		filter["deal"] = bson.M{"$gt": rat}
	}
	if pricel != "" && priceu != "" {
		pl, _ := strconv.Atoi(pricel)
		pu, _ := strconv.Atoi(priceu)
		filter["minorderamount"] = bson.M{"$lt": pl, "$gt": pu}
	}

	if lowest != "" {
		low, err := strconv.Atoi(lowest)
		if err == nil && low > 0 {
			settings, _ := repo.FindOneSetting(bson.M{"current": true}, bson.M{})
			rang := 0
			for _, value := range settings.DeliveryCharge {
				md, _ := value.(map[string]interface{})
				if md["charge"].(int) <= low {
					rang = md["range"].(int)
				}
			}
			if lat > float64(0) && long > float64(0) {
				filter["geo_location"] = bson.M{
					"$near": bson.M{
						"$geometry": bson.M{
							"type":        "Point",
							"coordinates": []float64{lat, long},
						},
						"$maxDistance": rang,
					},
				}
			}
		}
	}

	shops, err := repo.Find(filter, bson.M{}, limit, skip)
	if err != nil {
		return []OpSchema{}, err
	}
	l := []string{}
	for i := 0; i < len(shops); i++ {
		l = append(l, shops[i].SellerID)
	}
	sellerList, err := api.GetUsersDetailsByIDs(l)
	if err != nil {
		return []OpSchema{}, err
	}
	opList := []OpSchema{}
	for j := 0; j < len(shops); j++ {
		for k := 0; k < len(sellerList); k++ {
			if shops[j].SellerID == sellerList[k].ID.Hex() {
				var body OpSchema
				body.Address = shops[j].Address
				body.City = shops[j].City
				body.State = shops[j].State
				body.Country = shops[j].Country
				body.CreatedTime = shops[j].CreatedTime
				body.UpdatedTime = shops[j].UpdatedTime
				body.Featured = shops[j].Featured
				body.GeoLocation = shops[j].GeoLocation
				body.ID = shops[j].ID
				body.Pin = shops[j].Pin
				body.Deal = shops[j].Deal
				body.DeliveryType = shops[j].DeliveryType
				body.Cuisine = shops[j].Cuisine
				body.Primary = shops[j].Primary
				body.SellerID = shops[j].SellerID
				nerBannerUrl := createPreSignedDownloadUrl(shops[j].ShopBanner)
				body.ShopBanner = nerBannerUrl
				body.Tags = shops[j].Tags
				body.MinOrderAmount = shops[j].MinOrderAmount
				body.ShopDescription = shops[j].ShopDescription
				nerLogoUrl := createPreSignedDownloadUrl(shops[j].ShopLogo)
				body.ShopLogo = nerLogoUrl
				body.ShopName = shops[j].ShopName
				newShop := []string{}
				for o := 0; o < len(shops[j].ShopPhotos); o++ {
					nerShopUrl := createPreSignedDownloadUrl(shops[j].ShopPhotos[o])
					newShop = append(newShop, nerShopUrl)
				}
				body.ShopPhotos = newShop
				body.Rating = shops[j].Rating
				body.ShopStatus = shops[j].ShopStatus
				body.Timing = shops[j].Timing
				body.Type = shops[j].Type
				body.SellerEmail = sellerList[k].Email
				opList = append(opList, body)
				break
			}
		}
	}
	return opList, nil
}

func (*shopService) GetNearestShopAdmin(limit, skip int, sellerId, sType, status string, lat, long float64) ([]OpSchema, error) {
	filter := bson.M{}
	if sellerId != "" {
		filter["seller_id"] = sellerId
	}
	if sType != "" {
		filter["type"] = sType
	}
	if status != "" {
		filter["shop_status"] = status
	}
	if lat > float64(0) && long > float64(0) {
		filter["geo_location"] = bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{lat, long},
				},
				"$maxDistance": 7000,
			},
		}
	}
	shops, err := repo.Find(filter, bson.M{}, limit, skip)
	if err != nil {
		return []OpSchema{}, err
	}
	l := []string{}
	for i := 0; i < len(shops); i++ {
		l = append(l, shops[i].SellerID)
	}
	sellerList, err := api.GetUsersDetailsByIDs(l)
	if err != nil {
		return []OpSchema{}, err
	}
	opList := []OpSchema{}
	for j := 0; j < len(shops); j++ {
		for k := 0; k < len(sellerList); k++ {
			if shops[j].SellerID == sellerList[k].ID.Hex() {
				var body OpSchema
				body.Address = shops[j].Address
				body.City = shops[j].City
				body.State = shops[j].State
				body.Country = shops[j].Country
				body.CreatedTime = shops[j].CreatedTime
				body.UpdatedTime = shops[j].UpdatedTime
				body.Featured = shops[j].Featured
				body.GeoLocation = shops[j].GeoLocation
				body.ID = shops[j].ID
				body.Pin = shops[j].Pin
				body.Deal = shops[j].Deal
				body.Tags = shops[j].Tags
				body.MinOrderAmount = shops[j].MinOrderAmount
				body.DeliveryType = shops[j].DeliveryType
				body.Cuisine = shops[j].Cuisine
				body.Primary = shops[j].Primary
				body.SellerID = shops[j].SellerID
				nerBannerUrl := createPreSignedDownloadUrl(shops[j].ShopBanner)
				body.ShopBanner = nerBannerUrl
				body.ShopDescription = shops[j].ShopDescription
				nerLogoUrl := createPreSignedDownloadUrl(shops[j].ShopLogo)
				body.ShopLogo = nerLogoUrl
				body.ShopName = shops[j].ShopName
				body.Rating = shops[j].Rating
				newShop := []string{}
				for o := 0; o < len(shops[j].ShopPhotos); o++ {
					nerShopUrl := createPreSignedDownloadUrl(shops[j].ShopPhotos[o])
					newShop = append(newShop, nerShopUrl)
				}
				body.ShopPhotos = newShop
				body.ShopStatus = shops[j].ShopStatus
				body.Timing = shops[j].Timing
				body.Type = shops[j].Type
				body.SellerEmail = sellerList[k].Email
				opList = append(opList, body)
				break
			}
		}
	}
	return opList, nil
}

func (*shopService) GetTopRatedShopAdmin(limit, skip int, sellerId, sType, status string) ([]OpSchema, error) {
	filter := bson.M{}
	if sellerId != "" {
		filter["seller_id"] = sellerId
	}
	if sType != "" {
		filter["type"] = sType
	}
	if status != "" {
		filter["shop_status"] = status
	}
	shops, err := repo.FindSort(filter, bson.M{}, bson.M{"rating": -1}, 100, 0)
	if err != nil {
		return []OpSchema{}, err
	}
	l := []string{}
	for i := 0; i < len(shops); i++ {
		l = append(l, shops[i].SellerID)
	}
	sellerList, err := api.GetUsersDetailsByIDs(l)
	if err != nil {
		return []OpSchema{}, err
	}
	opList := []OpSchema{}
	for j := 0; j < len(shops); j++ {
		for k := 0; k < len(sellerList); k++ {
			if shops[j].SellerID == sellerList[k].ID.Hex() {
				var body OpSchema
				body.Address = shops[j].Address
				body.City = shops[j].City
				body.State = shops[j].State
				body.Country = shops[j].Country
				body.CreatedTime = shops[j].CreatedTime
				body.UpdatedTime = shops[j].UpdatedTime
				body.Featured = shops[j].Featured
				body.GeoLocation = shops[j].GeoLocation
				body.ID = shops[j].ID
				body.Pin = shops[j].Pin
				body.Deal = shops[j].Deal
				body.DeliveryType = shops[j].DeliveryType
				body.Cuisine = shops[j].Cuisine
				body.Tags = shops[j].Tags
				body.Tags = shops[j].Tags
				body.MinOrderAmount = shops[j].MinOrderAmount
				body.Primary = shops[j].Primary
				body.SellerID = shops[j].SellerID
				nerBannerUrl := createPreSignedDownloadUrl(shops[j].ShopBanner)
				body.ShopBanner = nerBannerUrl
				body.ShopDescription = shops[j].ShopDescription
				nerLogoUrl := createPreSignedDownloadUrl(shops[j].ShopLogo)
				body.ShopLogo = nerLogoUrl
				body.ShopName = shops[j].ShopName
				body.Rating = shops[j].Rating
				newShop := []string{}
				for o := 0; o < len(shops[j].ShopPhotos); o++ {
					nerShopUrl := createPreSignedDownloadUrl(shops[j].ShopPhotos[o])
					newShop = append(newShop, nerShopUrl)
				}
				body.ShopPhotos = newShop
				body.ShopStatus = shops[j].ShopStatus
				body.Timing = shops[j].Timing
				body.Type = shops[j].Type
				body.SellerEmail = sellerList[k].Email
				opList = append(opList, body)
				break
			}
		}
	}
	return opList, nil
}
func (*shopService) GetAdminUsersWithIDs(shopIds []string) ([]entity.ShopDB, error) {
	subFilter := bson.A{}
	for _, item := range shopIds {
		id, _ := primitive.ObjectIDFromHex(item)
		subFilter = append(subFilter, bson.M{"_id": id})
	}
	filter := bson.M{"$or": subFilter}
	users, err := repo.FindWithIDs(filter, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"Get Carts section",
			err,
		)
		return []entity.ShopDB{}, err
	}
	for i := 0; i < len(users); i++ {
		newPdownloadurl := createPreSignedDownloadUrl(users[i].ShopLogo)
		users[i].ShopLogo = newPdownloadurl
		newNdownloadurl := createPreSignedDownloadUrl(users[i].ShopBanner)
		users[i].ShopBanner = newNdownloadurl
	}
	return users, nil
}
