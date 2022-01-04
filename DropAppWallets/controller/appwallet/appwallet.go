package appwallet

import (
	"Drop/DropAppWallets/api"
	entity "Drop/DropAppWallets/entities"
	"Drop/DropAppWallets/repository/appwallet"
	"errors"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = appwallet.NewAppWalletRepository("apptransactions")
)

type appwalletService struct{}

func NewAppWalletService(repository appwallet.AppWalletRepository) AppWalletService {
	repo = repository
	return &appwalletService{}
}
func (r *appwalletService) AddTransaction(transactionDetails AppWalletSchema, token string) (string, error) {
	var appWalletEntity entity.AppWallet
	if transactionDetails.DeliveryID != "not assigned" && transactionDetails.DeliveryID != "" {
		res, _ := repo.FindOne(bson.M{"order_id": transactionDetails.OrderId}, bson.M{})
		iD := res.ID.Hex()
		id, _ := primitive.ObjectIDFromHex(iD)
		set := bson.M{"$set": bson.M{"delivery_id": transactionDetails.DeliveryID, "updated_time": time.Now()}}
		filter := bson.M{"_id": id}
		return repo.UpdateOne(filter, set)
	}
	if transactionDetails.TipAmount > 0 {
		res, _ := repo.FindOne(bson.M{"order_id": transactionDetails.OrderId}, bson.M{})
		iD := res.ID.Hex()
		id, _ := primitive.ObjectIDFromHex(iD)
		set := bson.M{"$set": bson.M{"tip_amount": transactionDetails.TipAmount, "updated_time": time.Now()}}
		filter := bson.M{"_id": id}
		return repo.UpdateOne(filter, set)
	}
	if transactionDetails.OrderId == "" {
		return "", errors.New("something went wrong")
	}
	appWalletEntity.ID = primitive.NewObjectID()
	appWalletEntity.OrderId = transactionDetails.OrderId
	appWalletEntity.ShopId = transactionDetails.ShopId
	appWalletEntity.SellerID = transactionDetails.SellerID
	appWalletEntity.DeliveryCharge = transactionDetails.DeliveryCharge
	appWalletEntity.DeliveryID = transactionDetails.DeliveryID
	appWalletEntity.DeliveryPersonCut = transactionDetails.DeliveryPersonCut
	appWalletEntity.DropAmount = transactionDetails.DropAmount
	appWalletEntity.TipAmount = transactionDetails.TipAmount
	appWalletEntity.SellerAmount = transactionDetails.SellerAmount
	var paymentId = transactionDetails.PaymentId
	if transactionDetails.PaymentId == "" {
		order, err := api.GetOrder(transactionDetails.OrderId, token)
		if err != nil {
			return "", errors.New("unable to get order details for order id =>" + transactionDetails.OrderId)
		}
		paymentId = order.PaymentID
		appWalletEntity.ShopId = order.ShopID
	}
	payment, err := api.GetPaymentByIds([]string{paymentId})
	if err != nil {
		return "", errors.New("unable to get payment id for order id => " + transactionDetails.OrderId)
	}
	if len(payment) == 0 {
		return "", errors.New("no payment details exist for order id => " + transactionDetails.OrderId)
	}
	appWalletEntity.OrderAmount = payment[0].Amount
	appWalletEntity.Status = "Pending"
	appWalletEntity.AddedTime = time.Now()
	return repo.InsertOne(appWalletEntity)
}

func (*appwalletService) UpdateTrans(transactionId, status string) (string, error) {
	if transactionId == "" {
		return "", errors.New("something went wrong")
	}
	id, _ := primitive.ObjectIDFromHex(transactionId)
	if status == "Settled" {
		set := bson.M{"$set": bson.M{"status": status, "updated_time": time.Now(), "settled_time": time.Now()}}
		filter := bson.M{"_id": id}
		return repo.UpdateOne(filter, set)
	} else {
		set := bson.M{"$set": bson.M{"status": status, "updated_time": time.Now()}}
		filter := bson.M{"_id": id}
		return repo.UpdateOne(filter, set)
	}
}

func (*appwalletService) UpdateSellerTransPer(sellerId, per, code string) (int, error) {
	filter := bson.M{}
	if code == "0001" {
		filter["entity_id"] = sellerId
		filter["entity"] = "Seller"
		filter["status"] = "Pending"
	}
	if code == "002729" {
		filter["entity"] = "Seller"
		filter["status"] = "Pending"
	}
	transactions, _ := repo.Find(filter, bson.M{}, 100, 0)
	c := 0
	for i := 0; i < len(transactions); i++ {
		p, _ := strconv.Atoi(per)
		set := bson.M{"$set": bson.M{"settlement_percentage": int64(p), "updated_time": time.Now()}}
		id, _ := primitive.ObjectIDFromHex(transactions[i].ID.Hex())
		filter := bson.M{"_id": id}
		_, err := repo.UpdateOne(filter, set)
		if err == nil {
			c = c + 1
		}
	}
	return c, nil
}

func (*appwalletService) GetTransaction(transactionId, status, entity, entityid, orderid string) (entity.AppWallet, error) {
	filter := bson.M{}
	if status != "" {
		filter["status"] = status
	}
	if orderid != "" {
		filter["order_id"] = orderid
	}
	if transactionId != "" {
		id, _ := primitive.ObjectIDFromHex(transactionId)
		filter["_id"] = id
	}
	return repo.FindOne(filter, bson.M{})
}
func (*appwalletService) GetTransactions(transactionId, status, entity, entityid, orderid string, limit, skip int) ([]entity.AppWallet, error) {
	filter := bson.M{}
	if status != "" {
		filter["status"] = status
	}
	if orderid != "" {
		filter["order_id"] = orderid
	}
	if transactionId != "" {
		id, _ := primitive.ObjectIDFromHex(transactionId)
		filter["_id"] = id
	}
	return repo.Find(filter, bson.M{}, limit, skip)
}

func (*appwalletService) GetDeliveryPersonBalance(transactionId, status, entityid, orderid, token, fromD, endD string, limit, skip int) ([]entity.AppWallet, int64, int64, int64, int64, int64, int64, error) {
	filter := bson.M{}
	if status != "" {
		filter["status"] = status
	}
	if orderid != "" {
		filter["order_id"] = orderid
	}
	if entityid != "" {
		filter["delivery_id"] = entityid
	}
	if transactionId != "" {
		id, _ := primitive.ObjectIDFromHex(transactionId)
		filter["_id"] = id
	}
	dF := bson.M{}
	layout := "2006-01-02T15:04:05.000Z"
	if fromD != "" {
		t, _ := time.Parse(layout, fromD)
		dF["$gt"] = t
		filter["added_time"] = dF
	}
	if endD != "" {
		t, _ := time.Parse(layout, endD)
		dF["$lt"] = t
		filter["added_time"] = dF
	}
	transactions, err := repo.Find(filter, bson.M{}, limit, skip)
	if err != nil {
		return []entity.AppWallet{}, 0, 0, 0, 0, 0, 0, err
	}
	var total int64
	var settled int64
	var unsettled int64
	var totaltip int64
	var settledtip int64
	var unsettledtip int64
	for i := 0; i < len(transactions); i++ {
		apii, _ := api.GetOrder(transactions[i].OrderId, token)
		paymemnt, err := api.GetPaymentByIds([]string{apii.PaymentID})
		if err == nil && len(paymemnt) > 0 {
			transactions[i].Type = paymemnt[0].PaymentMethodTypes
		}
		if transactions[i].Status == "Pending" {
			unsettled = unsettled + transactions[i].DeliveryPersonCut
		}
		if transactions[i].Status == "Settled" {
			settled = settled + transactions[i].DeliveryPersonCut
		}
		if transactions[i].Status == "Pending" {
			unsettledtip = unsettledtip + transactions[i].TipAmount
		}
		if transactions[i].Status == "Settled" {
			settledtip = settledtip + transactions[i].TipAmount
		}
		total = total + transactions[i].DeliveryPersonCut
		totaltip = totaltip + transactions[i].TipAmount
	}
	return transactions, total, settled, unsettled, totaltip, settledtip, unsettledtip, nil
}

func (*appwalletService) GetSellerPersonBalance(transactionId, status, entityid, orderid, fromD, endD string, limit, skip int) ([]entity.AppWallet, int64, int64, int64, error) {
	filter := bson.M{}
	if status != "" {
		filter["status"] = status
	}
	if orderid != "" {
		filter["order_id"] = orderid
	}
	if entityid != "" {
		filter["seller_id"] = entityid
	}
	if transactionId != "" {
		id, _ := primitive.ObjectIDFromHex(transactionId)
		filter["_id"] = id
	}
	dF := bson.M{}
	layout := "2006-01-02T15:04:05.000Z"
	if fromD != "" {
		t, _ := time.Parse(layout, fromD)
		dF["$gt"] = t
		filter["added_time"] = dF
	}
	if endD != "" {
		t, _ := time.Parse(layout, endD)
		dF["$lt"] = t
		filter["added_time"] = dF
	}
	transactions, err := repo.Find(filter, bson.M{}, limit, skip)
	if err != nil {
		return []entity.AppWallet{}, 0, 0, 0, err
	}
	var total int64
	var settled int64
	var unsettled int64
	for i := 0; i < len(transactions); i++ {

		if transactions[i].Status == "Pending" {
			unsettled = unsettled + transactions[i].SellerAmount
		}
		if transactions[i].Status == "Settled" {
			settled = settled + transactions[i].SellerAmount
		}
		total = total + transactions[i].SellerAmount
	}
	return transactions, total, settled, unsettled, nil
}

func (*appwalletService) GetAppEarning() (int64, error) {
	filter := bson.M{}
	filter["entity"] = "App"
	transactions, err := repo.Find(filter, bson.M{}, 10000, 0)
	if err != nil {
		return 0, err
	}
	var total int64
	for i := 0; i < len(transactions); i++ {
		total = total + transactions[i].DropAmount
	}
	return total, nil
}
func (*appwalletService) GetTotalTransactions() (int64, error) {
	filter := bson.M{}
	transactions, err := repo.Find(filter, bson.M{}, 10000, 0)
	if err != nil {
		return 0, err
	}
	var total int64
	for i := 0; i < len(transactions); i++ {
		total = total + transactions[i].OrderAmount
	}
	return total, nil
}

func (*appwalletService) GetSellerPersonS(token, fromD, endD string, limit, skip int) ([]TransactionCustomOutput, error) {
	sellerId, _ := api.GetSellers(token)
	if len(sellerId) < 1 {
		return []TransactionCustomOutput{}, errors.New("unable to retrieve seller ids")
	}
	var trancustO []TransactionCustomOutput
	for i := 0; i < len(sellerId); i++ {
		filter := bson.M{}
		filter["seller_id"] = sellerId[i].ID.Hex()
		dF := bson.M{}
		layout := "2006-01-02T15:04:05.000Z"
		if fromD != "" {
			t, _ := time.Parse(layout, fromD)
			dF["$gt"] = t
			filter["added_time"] = dF
		}
		if endD != "" {
			t, _ := time.Parse(layout, endD)
			dF["$lt"] = t
			filter["added_time"] = dF
		}
		transactions, _ := repo.Find(filter, bson.M{}, 100, 0)
		var trancustIO TransactionCustomOutput
		trancustIO.EntityID = sellerId[i].ID.Hex()
		trancustIO.Email = sellerId[i].Email
		trancustIO.Name = sellerId[i].Name
		trancustIO.PhoneNumber = sellerId[i].PhoneNo
		trancustIO.ProfilePhoto = sellerId[i].ProfilePhoto
		trancustIO.Type = sellerId[i].Type
		dAmt := int64(0)
		tAmt := int64(0)
		usAmt := int64(0)
		sAmt := int64(0)
		for j := 0; j < len(transactions); j++ {
			if transactions[j].Status == "Pending" {
				usAmt = usAmt + transactions[j].SellerAmount
			}
			if transactions[j].Status == "Settled" {
				sAmt = sAmt + transactions[j].SellerAmount
			}
			dAmt = dAmt + transactions[j].DropAmount
			tAmt = tAmt + transactions[j].SellerAmount
		}
		trancustIO.UnSettledAmount = usAmt
		trancustIO.SettledAmount = sAmt
		trancustIO.DropAmount = dAmt
		trancustIO.TotalAmount = tAmt
		if tAmt > 0 {
			trancustO = append(trancustO, trancustIO)
		}

	}
	return trancustO, nil
}

func (*appwalletService) GetDeliveryPersonS(token, fromD, endD string, limit, skip int) ([]TransactionCustomOutputD, error) {
	deliveryId, _ := api.GetDelivery(token)
	if len(deliveryId) < 1 {
		return []TransactionCustomOutputD{}, errors.New("unable to retrieve delivery ids")
	}
	var trancustO []TransactionCustomOutputD
	for i := 0; i < len(deliveryId); i++ {
		filter := bson.M{}
		filter["delivery_id"] = deliveryId[i].ID.Hex()
		dF := bson.M{}
		layout := "2006-01-02T15:04:05.000Z"
		if fromD != "" {
			t, _ := time.Parse(layout, fromD)
			dF["$gt"] = t
			filter["added_time"] = dF
		}
		if endD != "" {
			t, _ := time.Parse(layout, endD)
			dF["$lt"] = t
			filter["added_time"] = dF
		}
		transactions, _ := repo.Find(filter, bson.M{}, 100, 0)
		var trancustIO TransactionCustomOutputD
		trancustIO.EntityID = deliveryId[i].ID.Hex()
		trancustIO.Email = deliveryId[i].Email
		trancustIO.Name = deliveryId[i].Name
		trancustIO.PhoneNumber = deliveryId[i].PhoneNo
		trancustIO.ProfilePhoto = deliveryId[i].ProfilePhoto
		trancustIO.Type = deliveryId[i].Type
		usAmt := int64(0)
		sAmt := int64(0)
		tAmt := int64(0)
		ustAmt := int64(0)
		stAmt := int64(0)
		ttAmt := int64(0)
		cashSettled := int64(0)
		cashUnsettled := int64(0)
		for j := 0; j < len(transactions); j++ {
			apii, _ := api.GetOrder(transactions[j].OrderId, token)
			paymemnt, err := api.GetPaymentByIds([]string{apii.PaymentID})
			if err == nil && len(paymemnt) > 0 {
				if paymemnt[0].PaymentMethodTypes == "cod" {
					if transactions[j].Status == "Pending" {
						cashUnsettled = cashUnsettled + transactions[j].OrderAmount
					}
					if transactions[j].Status == "Settled" {
						cashSettled = cashSettled + transactions[j].DeliveryPersonCut
					}
				}
				if paymemnt[0].PaymentMethodTypes == "card" {
					if transactions[j].Status == "Pending" {
						usAmt = usAmt + transactions[j].DeliveryPersonCut
					}
					if transactions[j].Status == "Settled" {
						sAmt = sAmt + transactions[j].DeliveryPersonCut
					}
				}
				tAmt = tAmt + transactions[j].DeliveryPersonCut
			}
			if transactions[i].Status == "Pending" {
				ustAmt = ustAmt + transactions[i].TipAmount
			}
			if transactions[i].Status == "Settled" {
				stAmt = stAmt + transactions[i].TipAmount
			}
			ttAmt = ttAmt + transactions[i].TipAmount

		}
		trancustIO.CashUnSettled = cashUnsettled
		trancustIO.CashSettled = cashSettled
		trancustIO.UnSettledAmount = usAmt
		trancustIO.SettledAmount = sAmt
		trancustIO.TotalAmount = tAmt
		trancustIO.SettleTipdAmount = stAmt
		trancustIO.UnSettledTipAmount = ustAmt
		trancustIO.TotalTipAmount = ttAmt

		if tAmt > 0 {
			trancustO = append(trancustO, trancustIO)
		}

	}
	return trancustO, nil
}

func (*appwalletService) GetSellerPersonShops(token, sid, fromD, endD string, limit, skip int) ([]TransactionCustomShopOutput, error) {
	var trancustO []TransactionCustomShopOutput
	shops, _ := api.GetShopBySId(sid, token)
	if len(shops) < 1 {
		return []TransactionCustomShopOutput{}, errors.New("unable to retrieve shop ids")
	}
	for j := 0; j < len(shops); j++ {
		filter := bson.M{}
		filter["shop_id"] = shops[j].ID.Hex()
		dF := bson.M{}
		layout := "2006-01-02T15:04:05.000Z"
		if fromD != "" {
			t, _ := time.Parse(layout, fromD)
			dF["$gt"] = t
			filter["added_time"] = dF
		}
		if endD != "" {
			t, _ := time.Parse(layout, endD)
			dF["$lt"] = t
			filter["added_time"] = dF
		}
		transactions, _ := repo.Find(filter, bson.M{}, 100, 0)
		var trancustIO TransactionCustomShopOutput
		trancustIO.ShopName = shops[j].ShopName
		trancustIO.ProfilePhoto = shops[j].ShopLogo
		trancustIO.ShopType = shops[j].Type
		tAmt := int64(0)
		usAmt := int64(0)
		sAmt := int64(0)
		for k := 0; k < len(transactions); k++ {
			if transactions[k].Status == "Pending" {
				usAmt = usAmt + transactions[k].SellerAmount
			}
			if transactions[k].Status == "Settled" {
				sAmt = sAmt + transactions[k].SellerAmount
			}
			tAmt = tAmt + transactions[k].SellerAmount
		}
		trancustIO.UnSettledAmount = usAmt
		trancustIO.SettledAmount = sAmt
		trancustIO.TotalAmount = tAmt
		if tAmt > 0 {
			trancustO = append(trancustO, trancustIO)
		}
	}
	return trancustO, nil
}

func (*appwalletService) UpdateSellerPaymentHistory(sid, name, email, phone, doneby, token string, amount int64) (string, error) {
	var ph entity.SettingPaymentHistoryDB
	ph.ID = primitive.NewObjectID()
	ph.Amount = amount
	ph.DoneBy = doneby
	ph.DoneAt = time.Now()
	ph.Name = name
	ph.PhoneNo = phone
	ph.Type = "Seller"
	ph.Email = email
	_, err := repo.InsertPHistory(ph)
	if err != nil {
		return "", err
	}
	shops, _ := api.GetShopBySId(sid, token)
	if len(shops) < 1 {
		return "", errors.New("unable to retrieve shop ids")
	}
	for j := 0; j < len(shops); j++ {
		filter := bson.M{}
		filter["shop_id"] = shops[j].ID.Hex()
		transactions, _ := repo.Find(filter, bson.M{}, 10000, 0)
		for k := 0; k < len(transactions); k++ {
			if transactions[k].Status == "Pending" {
				if amount >= transactions[k].SellerAmount {
					id, _ := primitive.ObjectIDFromHex(transactions[k].ID.Hex())
					set := bson.M{"$set": bson.M{"status": "Settled", "updated_time": time.Now(), "settled_time": time.Now()}}
					filter := bson.M{"_id": id}
					repo.UpdateOne(filter, set)
					amount = amount - transactions[k].SellerAmount
				}
			}
		}
	}
	return "success", nil
}

func (*appwalletService) UpdateShopPaymentHistory(shid, name, email, phone, doneby, token string, amount int64) (string, error) {
	var ph entity.SettingPaymentHistoryDB
	ph.ID = primitive.NewObjectID()
	ph.Amount = amount
	ph.DoneBy = doneby
	ph.DoneAt = time.Now()
	ph.Name = name
	ph.PhoneNo = phone
	ph.Email = email
	ph.Type = "Shop"
	_, err := repo.InsertPHistory(ph)
	if err != nil {
		return "", err
	}
	filter := bson.M{}
	filter["shop_id"] = shid
	transactions, _ := repo.Find(filter, bson.M{}, 10000, 0)
	for k := 0; k < len(transactions); k++ {
		if transactions[k].Status == "Pending" {
			if amount >= transactions[k].SellerAmount {
				id, _ := primitive.ObjectIDFromHex(transactions[k].ID.Hex())
				set := bson.M{"$set": bson.M{"status": "Settled", "updated_time": time.Now(), "settled_time": time.Now()}}
				filter := bson.M{"_id": id}
				repo.UpdateOne(filter, set)
				amount = amount - transactions[k].SellerAmount
			}
		}
	}
	return "success", nil
}

func (*appwalletService) UpdateDeliveryPaymentHistory(did, name, email, phone, doneby, dtype, token string, amount int64) (string, error) {
	var ph entity.SettingPaymentHistoryDB
	ph.ID = primitive.NewObjectID()
	ph.Amount = amount
	ph.DoneBy = doneby
	ph.DoneAt = time.Now()
	ph.Name = name
	ph.PhoneNo = phone
	ph.Email = email
	ph.Type = "Delivery - " + dtype
	_, err := repo.InsertPHistory(ph)
	if err != nil {
		return "", err
	}
	deliveryId, _ := api.GetDelivery(token)
	if len(deliveryId) < 1 {
		return "", errors.New("unable to retrieve delivery ids")
	}
	for i := 0; i < len(deliveryId); i++ {
		filter := bson.M{}
		filter["delivery_id"] = deliveryId[i].ID.Hex()
		transactions, _ := repo.Find(filter, bson.M{}, 10000, 0)
		for k := 0; k < len(transactions); k++ {
			apii, _ := api.GetOrder(transactions[k].OrderId, token)
			paymemnt, err := api.GetPaymentByIds([]string{apii.PaymentID})
			if err == nil && len(paymemnt) > 0 {
				if paymemnt[0].PaymentMethodTypes == "cod" {
					if transactions[k].Status == "Pending" && dtype == "CashReturned" {
						if amount >= transactions[k].SellerAmount {
							id, _ := primitive.ObjectIDFromHex(transactions[k].ID.Hex())
							set := bson.M{"$set": bson.M{"status": "Settled", "updated_time": time.Now(), "settled_time": time.Now()}}
							filter := bson.M{"_id": id}
							repo.UpdateOne(filter, set)
							amount = amount - transactions[k].OrderAmount
						}
					}
				}
				if paymemnt[0].PaymentMethodTypes == "card" {
					if transactions[k].Status == "Pending" && dtype == "SettlementAmount" {
						if amount >= transactions[k].SellerAmount {
							id, _ := primitive.ObjectIDFromHex(transactions[k].ID.Hex())
							set := bson.M{"$set": bson.M{"status": "Settled", "updated_time": time.Now(), "settled_time": time.Now()}}
							filter := bson.M{"_id": id}
							repo.UpdateOne(filter, set)
							amount = amount - transactions[k].DeliveryPersonCut
						}
					}
				}
			}
		}
	}
	return "success", nil
}

func (*appwalletService) GetPAymentsHistory(limit, skip int) ([]entity.SettingPaymentHistoryDB, error) {
	filter := bson.M{}
	return repo.FindPHistory(filter, bson.M{}, 10000, 0)
}
