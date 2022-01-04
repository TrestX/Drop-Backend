package payment

import (
	entity "Drop/DropPayments/entities"

	payment "Drop/DropPayments/repository/payment"
	"errors"
	"strconv"
	"time"

	"github.com/aekam27/trestCommon"

	"Drop/DropPayments/api"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = payment.NewPaymentRepository("payments")
)

type paymentService struct{}

func NewPaymentService(repository payment.PaymentRepository) PaymentService {
	repo = repository
	return &paymentService{}
}

func (*paymentService) CreatePaymentIntent(userId, token string, payments PaymentIntentSchema) (string, string, error) {

	if userId == "" {
		return "", "", errors.New("UserId missing")
	}

	_, err := verifyPaymentIntentSchema(payments)
	if err != nil {
		return "", "", err
	}
	address, err := getAddressOfUser(payments.UserAddressId, token)
	if err != nil {
		return "", "", err
	}
	name, phoneno, email, profilePhoto, err := getUserDetails(userId, token)
	if err != nil {
		return "", "", err
	}
	paymentEntity := createPaymentEntity(payments, address, name, phoneno, email, profilePhoto)
	if err != nil {
		return "", "", err
	}
	paymentEntity.UserID = userId
	paymentEntity.SellerID = payments.SellerID
	paymentEntity.ShopID = payments.ShopID
	paymentEntity.CouponCode = payments.CouponCode
	if payments.Type == "cod" {
		paymentEntity.PaymentMethodTypes = "cod"
		if err != nil {
			return "", "", err
		}
		_, err = repo.InsertOne(paymentEntity)
		if err != nil {
			return "", "", err
		}
		return "", paymentEntity.ID.Hex(), nil
	} else {
		paymentEntity.PaymentMethodTypes = "card"
		params, err := createPaymentIntentParams(payments, address, name, phoneno, email, paymentEntity)
		if err != nil {
			return "", "", err
		}
		_, err = repo.InsertOne(paymentEntity)
		if err != nil {
			return "", "", err
		}
		return params.Data.Link, paymentEntity.ID.Hex(), nil
	}

}

func createPaymentEntity(payments PaymentIntentSchema, address entity.AddressDB, name, phoneno, email, profilePhoto string) entity.PaymentEntityDB {
	amount, _ := strconv.ParseInt(payments.Amount, 10, 64)
	var shipping entity.ShippingDetails
	shipping.Address = address
	shipping.Name = name
	shipping.Phone = phoneno
	shipping.Email = email
	shipping.ProfilePhoto = profilePhoto
	var paymentEntity entity.PaymentEntityDB
	paymentEntity.ID = primitive.NewObjectID()
	paymentEntity.Amount = amount
	paymentEntity.Currency = payments.Currency
	paymentEntity.Description = payments.Description
	paymentEntity.CustomerEmail = payments.CustomerEmail
	paymentEntity.Shipping = shipping
	paymentEntity.Status = "In Process"
	paymentEntity.AddedTime = time.Now()
	return paymentEntity
}

func createPaymentIntentParams(payments PaymentIntentSchema, address entity.AddressDB, name, phoneno, email string, paymentEntity entity.PaymentEntityDB) (api.PaymentResponse, error) {
	// amount, _ := strconv.ParseInt(payments.Amount, 10, 64)
	// params := &stripe.PaymentIntentParams{
	// 	Amount:             stripe.Int64(amount),
	// 	Currency:           stripe.String(payments.Currency),
	// 	PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
	// 	Description:        stripe.String(payments.Description),
	// 	ReceiptEmail:       stripe.String(payments.ReceiptEmail),
	// 	Shipping: &stripe.ShippingDetailsParams{
	// 		Address: &stripe.AddressParams{
	// 			City:       stripe.String(address.City),
	// 			Country:    stripe.String(address.Country),
	// 			State:      stripe.String(address.State),
	// 			Line1:      stripe.String(address.Address),
	// 			PostalCode: stripe.String(address.Pin),
	// 		},
	// 		Name:  stripe.String(name),
	// 		Phone: stripe.String(phoneno),
	// 	},
	// }
	var paymentIntent entity.PaymentIntent
	paymentIntent.Amount = payments.Amount
	paymentIntent.Currency = payments.Currency
	paymentIntent.PaymentMethod = "card"
	paymentIntent.Redirect_Url = "https://api.drop-deliveryapp.com/docker4/payments/success/" + paymentEntity.ID.Hex()
	paymentIntent.Tx_Ref = "drop-" + paymentEntity.ID.Hex()
	if email != "" {
		paymentIntent.Customer.Email = email
	} else {
		paymentIntent.Customer.Email = "noemailfound@abc.com"
	}
	if name != "" {
		paymentIntent.Customer.Name = name
	} else {
		paymentIntent.Customer.Name = "nonamefound"
	}
	if phoneno != "" {
		paymentIntent.Customer.PhoneNumber = phoneno
	} else {
		paymentIntent.Customer.PhoneNumber = "9876543210"
	}
	paymentIntent.Customizations.Title = "Drop Payment"
	paymentIntent.Customizations.Description = payments.Description
	paymentIntent.Customizations.Logo = "https://thumbs.dreamstime.com/b/digital-wallet-e-payment-logo-design-vector-illustration-online-electronic-bank-network-buy-coin-bitcoin-exchange-background-phone-168953702.jpg"
	data, _ := api.SendPayemnt(paymentIntent)
	if data.Status == "error" {
		return api.PaymentResponse{}, errors.New(data.Message)
	}
	return data, nil
}

func getUserDetails(userId, token string) (string, string, string, string, error) {
	user, err := api.GetUserDetails(token)
	if err != nil {
		return "", "", "", "", err
	}
	return user.Name, user.PhoneNo, user.Email, user.ProfilePhoto, nil
}

func getAddressOfUser(addressId, token string) (entity.AddressDB, error) {
	address, err := api.GetUserAddress(addressId, token)
	if err != nil {
		return entity.AddressDB{}, err
	}
	return address, nil
}

func verifyPaymentIntentSchema(payments PaymentIntentSchema) (string, error) {
	if payments.Amount == "" {
		return "", errors.New("Amount Mandatory")
	}
	if payments.Currency == "" {
		return "", errors.New("Billing Currency Mandatory")
	}
	if payments.CustomerEmail == "" {
		return "", errors.New("Customer Email Mandatory")
	}
	if payments.RedirectUrl == "" {
		return "", errors.New("Redirect URL Mandatory")
	}
	if payments.Description == "" {
		return "", errors.New("Payment Description Mandatory")
	}
	if payments.UserAddressId == "" {
		return "", errors.New("User Address Id Mandatory")
	}
	return "ok", nil
}

func (*paymentService) UpdatePaymentStatus(userId, paymentId, status string) (string, error) {
	if userId == "" {
		return "", errors.New("Mandatory user id missing")
	}
	if paymentId == "" {
		return "", errors.New("Mandatory payment id missing")
	}
	if status == "" {
		return "", errors.New("Mandatory status missing")
	}
	id, _ := primitive.ObjectIDFromHex(paymentId)
	set := bson.M{"$set": bson.M{"status": status, "updated_time": time.Now()}}
	filter := bson.M{"_id": id, "user_id": userId}
	return repo.UpdateOne(filter, set)
}

func (*paymentService) UpdatePaymentStatusSuccess(userId, paymentId string) (string, error) {
	if userId == "" {
		return "", errors.New("Mandatory user id missing")
	}
	if paymentId == "" {
		return "", errors.New("Mandatory payment id missing")
	}
	id, _ := primitive.ObjectIDFromHex(paymentId)
	set := bson.M{"$set": bson.M{"status": "Success", "updated_time": time.Now()}}
	filter := bson.M{"_id": id, "user_id": userId}
	return repo.UpdateOne(filter, set)
}

func (*paymentService) GetPaymentsDetails(userId, status string, limit, skip int) ([]entity.PaymentEntityDB, error) {
	filter := bson.M{"user_id": userId}
	if status != "" {
		filter["status"] = status
	}
	payments, err := repo.Find(filter, bson.M{}, limit, skip)
	if err != nil {
		trestCommon.ECLog2(
			"Get Payments Details section",
			err,
		)
		return payments, err
	}
	return payments, nil
}

func (*paymentService) GetPaymentDetails(userId, paymentId string) (entity.PaymentEntityDB, error) {
	id, _ := primitive.ObjectIDFromHex(paymentId)
	payment, err := repo.FindOne(bson.M{"_id": id}, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"Get Payments Details section",
			err,
		)
		return payment, err
	}
	return payment, nil
}

func (*paymentService) GetPaymentWithIDs(paymentIds []string) ([]entity.PaymentEntityDB, error) {
	subFilter := bson.A{}
	for _, item := range paymentIds {
		id, _ := primitive.ObjectIDFromHex(item)
		subFilter = append(subFilter, bson.M{"_id": id})
	}
	filter := bson.M{"$or": subFilter}
	return repo.FindWithIDs(filter, bson.M{})
}

func (*paymentService) GetAdminPaymentDetails(status, user, token, seller, shop string, limit, skip int) ([]PaymentOP, error) {
	filter := bson.M{}
	if status != "" {
		filter["status"] = status
	}
	if user != "" {
		filter["user_id"] = user
	}
	if seller != "" {
		filter["seller_id"] = seller
	}
	if shop != "" {
		filter["shop_id"] = shop
	}
	payment, err := repo.Find(filter, bson.M{}, limit, skip)
	if err != nil {
		trestCommon.ECLog2(
			"Get Payments Details section",
			err,
		)
		return []PaymentOP{}, err
	}
	var opL []PaymentOP
	for i := 0; i < len(payment); i++ {
		shopDetails, _ := api.GetShopDetails(payment[i].ShopID, token)
		stoken, _ := trestCommon.CreateToken(payment[i].SellerID, "", "", "")
		sellerDetails, _ := api.GetUserDetails(stoken)
		var op PaymentOP
		op.AddedTime = payment[i].AddedTime
		op.ID = payment[i].ID
		op.UserID = payment[i].UserID
		op.SellerID = payment[i].SellerID
		op.ShopID = payment[i].ShopID
		op.Amount = payment[i].Amount
		op.Currency = payment[i].Currency
		op.PaymentMethodTypes = payment[i].PaymentMethodTypes
		op.Description = payment[i].Description
		op.ReceiptEmail = payment[i].CustomerEmail
		op.Shipping = payment[i].Shipping
		op.Status = payment[i].Status
		op.CouponCode = payment[i].CouponCode
		op.UpdatedTime = payment[i].UpdatedTime
		var sD ShopDetails
		sD.AccountNumber = sellerDetails.AccountNumber
		sD.Deal = shopDetails.Deal
		sD.IFSC = sellerDetails.IFSC
		sD.SellerName = sellerDetails.Name
		sD.ShopName = shopDetails.ShopName
		sD.Address = shopDetails.Address + ", " + shopDetails.City + ", " + shopDetails.State + "," + shopDetails.Country + " " + shopDetails.Pin
		op.SellerDetails = sD
		opL = append(opL, op)
	}
	return opL, nil
}
