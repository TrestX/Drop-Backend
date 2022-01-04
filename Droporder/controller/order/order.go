package order

import (
	"Drop/Droporder/api"
	entity "Drop/Droporder/entities"
	"math"
	"math/rand"

	notification "Drop/Droporder/repository/order/notificationrepo"
	util "Drop/Droporder/util"
	"fmt"
	"strings"

	"github.com/aekam27/trestCommon"

	"Drop/Droporder/repository/order"
	cart "Drop/Droporder/repository/order/cartrepo"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = order.NewOrderRepository("order")
)

var (
	repocart = cart.NewCartRepository("cart")
)
var (
	notificationService = util.NewNotificationService(notification.NewNotificationRepository("notification"))
)

type orderService struct{}

func NewOrderService(repository order.OrderRepository) OrderService {
	repo = repository
	return &orderService{}
}

func (r *orderService) PlaceOrder(userId, token string, order Order) (AdminOrderOutput, error) {
	var orderEntity entity.OrderDB
	var deliveryEntity entity.DeliveryDB
	if userId == "" {
		return AdminOrderOutput{}, errors.New("user id missing")
	}
	if order.ShopAddressID == "" || order.CartID == "" || order.ShopID == "" || order.PaymentID == "" {
		return AdminOrderOutput{}, errors.New("mandatory ids missing")
	}
	orderEntity.ID = primitive.NewObjectID()
	orderEntity.UserID = userId
	orderEntity.CartID = order.CartID
	orderEntity.ShopID = order.ShopID
	orderEntity.PaymentID = order.PaymentID
	orderEntity.Status = "Ordered"
	orderEntity.AddedTime = time.Now()
	orderEntity.OrderPlacedTime = time.Now()
	shopDetails, _ := api.GetShopDetails(order.ShopID, token)
	orderEntity.SellerID = shopDetails.SellerID
	orderEntity.DeliveryCode = int64(1000 + rand.Intn(9999-1000))
	var shopAddress entity.AddressDB
	shopAddress.Address = shopDetails.Address
	shopAddress.City = shopDetails.City
	shopAddress.Country = shopDetails.Country
	shopAddress.Pin = shopDetails.Pin
	shopAddress.State = shopDetails.State
	shopAddress.GeoLocation = shopDetails.GeoLocation
	deliveryEntity.ShopAddress = shopAddress
	userAddress, err := getAddressDetails(order.UserAddressID, token)
	if err != nil {
		return AdminOrderOutput{}, err
	}
	deliveryEntity.UserAddress = userAddress
	deliveryEntity.AddedTime = time.Now()
	orderEntity.DeliveryDetails = deliveryEntity
	var appWallet AppWallet
	appWallet.OrderId = orderEntity.ID.Hex()
	appWallet.DeliveryCharge = order.DeliveryCharge
	if order.DeliveryCharge == 0 {
		appWallet.DeliveryPersonCut = order.SellerAmount - int64((order.SellerAmount*order.DeliveryPersonCutPer)/100)
	} else {
		appWallet.DeliveryPersonCut = order.DeliveryPersonCut
	}
	appWallet.DropAmount = order.DropAmount
	appWallet.SellerAmount = order.SellerAmount
	appWallet.PaymentId = order.PaymentID
	appWallet.DeliveryID = "not assigned"
	appWallet.TipAmount = 0
	appWallet.ShopId = order.ShopID
	appWallet.SellerID = shopDetails.SellerID
	_, err = api.PostAppWallet(appWallet, token)

	_, err = repo.InsertOne(orderEntity)
	if err != nil {
		return AdminOrderOutput{}, errors.New("unable to place order")
	}
	cartId, _ := primitive.ObjectIDFromHex(order.CartID)
	repocart.UpdateOne(bson.M{"_id": cartId}, bson.M{"$set": bson.M{"status": "Ordered"}})
	util.NewNotificationService(notification.NewNotificationRepository("notification")).SendNotificationWithTopic("Order Placed", "Successfully placed the order", userId, userId)
	util.NewNotificationService(notification.NewNotificationRepository("notification")).SendNotificationWithTopic("New Order", "New Order", "deliverygeneral", "")
	return r.GetLatestOrders(userId, token, 1, 0)

}

func getAddressDetails(addressId, token string) (entity.AddressDB, error) {
	address, err := api.GetUserAddress(addressId, token)
	if err != nil {
		return entity.AddressDB{}, err
	}
	return address, nil
}

func getPaymentsDetails(paymentId, token string) (entity.PaymentDB, error) {
	payment, err := api.GetUserPaymentDetails(paymentId, token)
	if err != nil {
		return entity.PaymentDB{}, err
	}
	return payment.Data, nil
}

func getCartDetails(cartId, token string) (entity.CartDB, error) {
	cartDetails, err := api.GetUserCartDetails(cartId, token)
	if err != nil {
		return entity.CartDB{}, err
	}
	return cartDetails.Data, nil
}

func getOrder(id primitive.ObjectID) (entity.OrderDB, error) {
	return repo.FindOne(bson.M{"_id": id}, bson.M{})
}
func (*orderService) UpdateOrder(orderId string, order Order) (string, error) {
	id, _ := primitive.ObjectIDFromHex(orderId)
	porder, _ := getOrder(id)
	set := bson.M{}
	set["updated_time"] = time.Now()
	if order.Status != "" {
		set["status"] = order.Status
		if order.Status == "Accepted" {
			util.NewNotificationService(notification.NewNotificationRepository("notification")).SendNotificationWithTopic("Order Accepted", "Order has been Accepted by the restaurant", porder.UserID, porder.UserID)
			set["order_accepted_time"] = time.Now()
		}
		if order.Status == "Pickup" {
			util.NewNotificationService(notification.NewNotificationRepository("notification")).SendNotificationWithTopic("Order is on the way", "Order is on the way", porder.UserID, porder.UserID)
			set["order_pickup_time"] = time.Now()
		}
		if order.Status == "Delivered" {
			util.NewNotificationService(notification.NewNotificationRepository("notification")).SendNotificationWithTopic("Order Delivered Successfully", "Order Delivered Successfully", porder.UserID, porder.UserID)
			util.NewNotificationService(notification.NewNotificationRepository("notification")).SendNotificationWithTopic("How was your Experience?", "How was your Delivery and Food Experience?", porder.UserID, porder.UserID)
			set["delivered_time"] = time.Now()
		}
	}
	if order.TrackingID != "" {
		set["delivery_details.tracking_id"] = order.TrackingID
	}
	if order.DeliveryID != "" {
		util.NewNotificationService(notification.NewNotificationRepository("notification")).SendNotificationWithTopic("Delivery Person Assigned", "A Delivery Person Has Been Assigned for your order", porder.UserID, porder.UserID)
		set["delivery_details.delivery_id"] = order.DeliveryID
		var appWallet AppWallet
		appWallet.DeliveryID = order.DeliveryID
		appWallet.OrderId = orderId
		token, _ := trestCommon.CreateToken(order.DeliveryID, "", "", "")
		_, err := api.PostAppWallet(appWallet, token)
		if err != nil {
			return "", errors.New("unable to assign delivery person")
		}
	}
	if order.TipAmount != 0 {
		var appWallet AppWallet
		set["tip_amount"] = order.TipAmount
		appWallet.TipAmount = order.TipAmount
		appWallet.OrderId = orderId
		token, _ := trestCommon.CreateToken(orderId, "", "", "")
		_, err := api.PostAppWallet(appWallet, token)
		if err != nil {
			return "", errors.New("unable to assign delivery person")
		}
	}
	if order.DeliveryCode > 0 && order.DeliveryCode == porder.DeliveryCode {
		set["status"] = "Delivered"
	}
	if order.DeliveryCode > 0 && order.DeliveryCode != porder.DeliveryCode {
		return "", errors.New("deliverycode do not match")
	}
	set = bson.M{"$set": set}
	filter := bson.M{"_id": id}
	return repo.UpdateOne(filter, set)
}

func getUserDetails(token string) error {
	user, err := api.GetUserDetails(token)
	if err != nil {
		return err
	}
	if strings.ToLower(user.AccountType) != "admin" {
		return errors.New("user doesnot have admin privilages")
	}
	return nil
}

func (*orderService) GetAllOrdersAdmin(token, status string, limit, skip int) ([]OrderOutput, error) {
	err := getUserDetails(token)
	if err != nil {
		return []OrderOutput{}, err
	}
	filter := bson.M{}
	if status != "" {
		filter["status"] = status
	}
	orders, err := repo.Find(filter, bson.M{}, limit, skip)
	if err != nil {
		trestCommon.ECLog2(
			"Get Orders section",
			err,
		)
		return []OrderOutput{}, err
	}
	var orderList []OrderOutput
	for i := 0; i < len(orders); i++ {
		token, _ := trestCommon.CreateToken(orders[i].UserID, "", "", "")
		payment, _ := getPaymentsDetails(orders[i].PaymentID, token)
		cartDetails, _ := getCartDetails(orders[i].CartID, token)
		shopDetails, _ := api.GetShopDetails(orders[i].ShopID, token)
		newtoken, _ := trestCommon.CreateToken(shopDetails.SellerID, "", "", "")
		sellerDetails, _ := api.GetUserDetails(newtoken)
		body := constructAdminOutput(orders[i], payment, cartDetails, shopDetails, sellerDetails)
		orderList = append(orderList, body)
	}
	return orderList, nil
}

func constructAdminOutput(order entity.OrderDB, payment entity.PaymentDB, cart entity.CartDB, shopDetails entity.ShopDB, sellerDetails entity.UserDB) OrderOutput {
	var orderOutput OrderOutput
	orderOutput.ID = order.ID.Hex()
	orderOutput.AcceptedAt = order.OrderAcceptedTime
	orderOutput.OrderStatus = order.Status
	orderOutput.OrderedAt = order.AddedTime
	orderOutput.PaymentAmount = payment.Amount
	orderOutput.PaymentCurrency = payment.Currency
	orderOutput.OrderDescription = payment.Description
	orderOutput.DeliveryDetails = payment.Shipping
	orderOutput.ItemsOrdered = cart.Items
	orderOutput.ShopDetails = shopDetails
	orderOutput.SellerDetails = sellerDetails
	return orderOutput
}
func (*orderService) GetOrders(userId, token string, limit, skip int) ([]OrderOutput, error) {
	orders, err := repo.Find(bson.M{"user_id": userId}, bson.M{}, limit, skip)
	if err != nil {
		trestCommon.ECLog2(
			"Get Orders section",
			err,
		)
		return []OrderOutput{}, err
	}
	var orderList []OrderOutput

	for i := 0; i < len(orders); i++ {

		payment, _ := getPaymentsDetails(orders[i].PaymentID, token)
		cartDetails, _ := getCartDetails(orders[i].CartID, token)
		body := constructOutput(orders[i], payment, cartDetails)
		orderList = append(orderList, body)
	}
	return orderList, nil
}

func (*orderService) GetOrderWithIDs(userIds []string) (interface{}, error) {
	subFilter := bson.A{}
	for _, id := range userIds {
		subFilter = append(subFilter, bson.M{"user_id": id})
	}
	filter := bson.M{"$or": subFilter}
	fmt.Println(filter)
	orders, err := repo.FindWithIDs(filter, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"Get Orders section",
			err,
		)
		return nil, err
	}
	var paymentIds []string
	var cartIds []string
	for i := 0; i < len(orders); i++ {
		paymentIds = append(paymentIds, orders[i].PaymentID)
		cartIds = append(cartIds, orders[i].CartID)
	}
	payments, _ := api.GetUserPaymentDetailsPaymentIds(paymentIds)
	carts, _ := api.GetCartsDetails(cartIds)
	var oP OrderInteface
	oP.PaymentList = payments
	oP.CartList = carts
	oP.OrderList = orders
	return oP, nil
}

func (*orderService) GetOrder(orderId string) (entity.OrderDB, error) {
	id, _ := primitive.ObjectIDFromHex(orderId)
	return repo.FindOne(bson.M{"_id": id}, bson.M{})
}

func constructOutput(order entity.OrderDB, payment entity.PaymentDB, cart entity.CartDB) OrderOutput {
	var orderOutput OrderOutput
	orderOutput.ID = order.ID.Hex()
	orderOutput.OrderStatus = order.Status
	orderOutput.OrderedAt = order.AddedTime
	orderOutput.PaymentAmount = payment.Amount
	orderOutput.PaymentCurrency = payment.Currency
	orderOutput.OrderDescription = payment.Description
	newdownloadurl := createPreSignedDownloadUrl(payment.Shipping.ProfilePhoto)
	payment.Shipping.ProfilePhoto = newdownloadurl
	orderOutput.DeliveryDetails = payment.Shipping
	orderOutput.ItemsOrdered = cart.Items
	return orderOutput
}

func (*orderService) GetNewDeliveryOrders(token string, limit, skip int, latitude, longitude float64) ([]OrderOutput, error) {

	orders, err := repo.Find(bson.M{"delivery_details.shop_address.geo_location": bson.M{
		"$near": bson.M{
			"$geometry": bson.M{
				"type":        "Point",
				"coordinates": []float64{latitude, longitude},
			},
			"$maxDistance": 10000,
		},
	}, "status": "Accepted"}, bson.M{}, limit, skip)
	if err != nil {
		trestCommon.ECLog2(
			"Get Order section",
			err,
		)
		return []OrderOutput{}, err
	}
	var orderList []OrderOutput
	for i := 0; i < len(orders); i++ {
		payment, _ := getPaymentsDetails(orders[i].PaymentID, token)
		cartDetails, _ := getCartDetails(orders[i].CartID, token)
		body := constructOutput(orders[i], payment, cartDetails)
		orderList = append(orderList, body)
	}
	return orderList, nil
}

func (*orderService) GetActiveDeliveryOrders(token string, limit, skip int, deliveryID, status string) ([]OrderOutput, error) {
	orders, err := repo.Find(bson.M{"delivery_details.delivery_id": deliveryID, "status": status}, bson.M{}, limit, skip)
	if err != nil {
		trestCommon.ECLog2(
			"Get Order section",
			err,
		)
		return []OrderOutput{}, err
	}
	var orderList []OrderOutput
	for i := 0; i < len(orders); i++ {
		payment, _ := getPaymentsDetails(orders[i].PaymentID, token)
		cartDetails, _ := getCartDetails(orders[i].CartID, token)
		body := constructOutput(orders[i], payment, cartDetails)
		orderList = append(orderList, body)
	}
	return orderList, nil
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
func (*orderService) GetAdminOrders(token string, limit, skip int, shopID, sellerID, deliveryId, userId, status, fromD, endD string) ([]AdminOrderOutput, error) {
	filter := bson.M{}
	if shopID != "" {
		filter["shop_id"] = shopID
	}
	if sellerID != "" {
		filter["seller_id"] = sellerID
	}
	if deliveryId != "" {
		filter["delivery_details.delivery_id"] = deliveryId
	}
	if userId != "" {
		filter["user_id"] = userId
	}
	if status != "" {
		filter["status"] = status
	}
	dF := bson.M{}
	layout := "2006-01-02T15:04:05.000Z"
	if fromD != "" {
		t, _ := time.Parse(layout, fromD)
		dF["$gt"] = t
		filter["order_placed_time"] = dF
	}
	if endD != "" {
		t, _ := time.Parse(layout, endD)
		dF["$lt"] = t
		filter["order_placed_time"] = dF
	}
	orders, err := repo.Find(filter, bson.M{}, limit, skip)
	if err != nil {
		trestCommon.ECLog2(
			"Get Order section",
			err,
		)
		return []AdminOrderOutput{}, err
	}
	var paymentIds []string
	var cartIds []string
	for i := 0; i < len(orders); i++ {
		paymentIds = append(paymentIds, orders[i].PaymentID)
		cartIds = append(cartIds, orders[i].CartID)
	}
	payments, _ := api.GetUserPaymentDetailsPaymentIds(paymentIds)
	carts, _ := api.GetCartsDetails(cartIds)
	var orderOplist []AdminOrderOutput
	for j := 0; j < len(orders); j++ {
		var orderOp AdminOrderOutput
		orderOp.ID = orders[j].ID.Hex()
		orderOp.OrderedAt = orders[j].OrderPlacedTime
		orderOp.TipAmount = orders[j].TipAmount
		orderOp.AcceptedAt = orders[j].OrderAcceptedTime
		orderOp.DeliveredTime = orders[j].DeliveredTime
		orderOp.OrderPickeUpTime = orders[j].OrderPickUpTime
		orderOp.OrderStatus = orders[j].Status
		orderOp.PaymentAmount = payments[j].Amount
		orderOp.PaymentCurrency = payments[j].Currency
		orderOp.OrderDescription = payments[j].Description
		orderOp.CouponCode = payments[j].CouponCode
		orderOp.PaymentMethodTypes = payments[j].PaymentMethodTypes
		orderOp.PaymentStatus = payments[j].Status
		newdownloadurl := createPreSignedDownloadUrl(payments[j].Shipping.ProfilePhoto)
		payments[j].Shipping.ProfilePhoto = newdownloadurl
		orderOp.DeliveryDetails = payments[j].Shipping
		orderOp.ItemsOrdered = carts[j].Items
		shopDetails, _ := api.GetShopDetails(orders[j].ShopID, token)
		orderOp.ShopDetails = shopDetails
		newtoken, _ := trestCommon.CreateToken(shopDetails.SellerID, "", "", "")
		sellerDetails, _ := api.GetUserDetails(newtoken)
		geoLocationU := orders[j].DeliveryDetails.UserAddress.GeoLocation["coordinates"].(primitive.A)
		geoLocationS := orders[j].DeliveryDetails.ShopAddress.GeoLocation["coordinates"].(primitive.A)
		distance := calculateDistance(geoLocationU[0].(float64), geoLocationU[1].(float64), geoLocationS[0].(float64), geoLocationS[1].(float64))
		orderOp.Distance = distance
		if orders[j].Status == "Delivered" {
			review, _ := api.GetOrderReview(orders[j].ID.Hex(), " ")
			orderOp.OrderReview = review[0]
		}
		orderOp.SellerDetails = sellerDetails
		orderOplist = append(orderOplist, orderOp)
	}
	return orderOplist, nil
}

func calculateDistance(lat1, lat2, lng1, lng2 float64) float64 {
	const PI float64 = 3.141592653589793

	radlat1 := float64(PI * lat1 / 180)
	radlat2 := float64(PI * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	if dist > 1 {
		dist = 1
	}
	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515

	return dist
}

func (*orderService) GetLatestOrders(userId, token string, limit, skip int) (AdminOrderOutput, error) {
	filter := bson.M{}
	if userId != "" {
		filter["user_id"] = userId
	}

	orders, err := repo.Find(filter, bson.M{}, limit, skip)
	if err != nil {
		trestCommon.ECLog2(
			"Get Order section",
			err,
		)
		return AdminOrderOutput{}, err
	}
	var paymentIds []string
	var cartIds []string
	for i := 0; i < len(orders); i++ {
		paymentIds = append(paymentIds, orders[i].PaymentID)
		cartIds = append(cartIds, orders[i].CartID)
	}
	payments, _ := api.GetUserPaymentDetailsPaymentIds(paymentIds)
	carts, _ := api.GetCartsDetails(cartIds)
	var orderOplist []AdminOrderOutput
	for j := 0; j < len(orders); j++ {
		var orderOp AdminOrderOutput
		orderOp.ID = orders[j].ID.Hex()
		orderOp.OrderedAt = orders[j].OrderPlacedTime
		orderOp.TipAmount = orders[j].TipAmount
		orderOp.AcceptedAt = orders[j].OrderAcceptedTime
		orderOp.DeliveredTime = orders[j].DeliveredTime
		orderOp.OrderPickeUpTime = orders[j].OrderPickUpTime
		orderOp.OrderStatus = orders[j].Status
		orderOp.PaymentAmount = payments[j].Amount
		orderOp.PaymentCurrency = payments[j].Currency
		orderOp.OrderDescription = payments[j].Description
		orderOp.CouponCode = payments[j].CouponCode
		orderOp.PaymentMethodTypes = payments[j].PaymentMethodTypes
		orderOp.PaymentStatus = payments[j].Status
		newdownloadurl := createPreSignedDownloadUrl(payments[j].Shipping.ProfilePhoto)
		payments[j].Shipping.ProfilePhoto = newdownloadurl
		orderOp.DeliveryDetails = payments[j].Shipping
		orderOp.ItemsOrdered = carts[j].Items
		shopDetails, _ := api.GetShopDetails(orders[j].ShopID, token)
		orderOp.ShopDetails = shopDetails
		newtoken, _ := trestCommon.CreateToken(shopDetails.SellerID, "", "", "")
		sellerDetails, _ := api.GetUserDetails(newtoken)
		geoLocationU := orders[j].DeliveryDetails.UserAddress.GeoLocation["coordinates"].(primitive.A)
		geoLocationS := orders[j].DeliveryDetails.ShopAddress.GeoLocation["coordinates"].(primitive.A)
		distance := calculateDistance(geoLocationU[0].(float64), geoLocationU[1].(float64), geoLocationS[0].(float64), geoLocationS[1].(float64))
		orderOp.Distance = distance
		if orders[j].Status == "Delivered" {
			review, _ := api.GetOrderReview(orders[j].ID.Hex(), " ")
			orderOp.OrderReview = review[0]
		}
		orderOp.SellerDetails = sellerDetails
		orderOplist = append(orderOplist, orderOp)
	}
	return orderOplist[0], nil
}
