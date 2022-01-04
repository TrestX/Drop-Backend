package payment

import (
	entity "Nailzee/NailzeePayments/entities"
	"errors"

	payment "Nailzee/NailzeePayments/repository/payment"

	"github.com/aekam27/trestCommon"
	"github.com/spf13/viper"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/paymentintent"
	"github.com/stripe/stripe-go/paymentmethod"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = payment.NewPaymentRepository("nailzeepayments")
)

type paymentService struct{}

func NewPaymentService(repository payment.PaymentRepository) PaymentService {
	repo = repository
	return &paymentService{}
}

func (*paymentService) CreatePaymentIntent(userId, token string, payments PaymentIntentSchema) (string, string, error) {
	paymentDB, err := getCustomerDetails(payments.UserId)
	stripe.Key = viper.GetString("stripe.key")
	if err != nil {
		return "customer doesnot exist", "", err
	} else {
		param := createConfirmPaymentIntentParams(payments, paymentDB.CustomerId)
		paymentIntent, err := paymentintent.New(&param)
		if err != nil {
			return "", "", err
		}
		switch status := paymentIntent.Status; status {
		case "requires_action", "requires_source_action":
			return paymentIntent.ClientSecret, paymentIntent.ID, nil
		case "requires_payment_method", "requires_source":
			return paymentIntent.ClientSecret, "", errors.New("your card was denied, please provide a new payment method")
		case "succeeded":
			return paymentIntent.ClientSecret, "Success", nil
		}
		return paymentIntent.ClientSecret, "", errors.New("Error")
	}
}
func (*paymentService) CreateNewPaymentIntent(userId, token string, payments PaymentIntentSchema) (string, string, error) {
	stripe.Key = viper.GetString("stripe.key")
	params := &stripe.CustomerParams{
		Name: stripe.String(payments.Name),
		Address: &stripe.AddressParams{
			Line1:      stripe.String(payments.AddressLine),
			PostalCode: stripe.String(payments.Pin),
			City:       stripe.String(payments.City),
			State:      stripe.String(payments.State),
			Country:    stripe.String(payments.Country),
		},
	}
	c, _ := customer.New(params)
	param := createPaymentIntentParams(payments, c)
	paymentIntent, err := paymentintent.New(&param)
	if err != nil {
		return "", "", err
	}
	if payments.IsSaved {
		var paymentEntity entity.PaymentDB
		paymentEntity.ID = primitive.NewObjectID()
		paymentEntity.CustomerId = c.ID
		paymentEntity.UserId = payments.UserId
		_, err = repo.InsertOne(paymentEntity)
		if err != nil {
			return "", "", err
		}
	}
	switch status := paymentIntent.Status; status {
	case "requires_action", "requires_source_action":
		return paymentIntent.ClientSecret, paymentIntent.ID, nil
	case "requires_payment_method", "requires_source":
		return paymentIntent.ClientSecret, "", errors.New("your card was denied, please provide a new payment method")
	case "succeeded":
		return paymentIntent.ClientSecret, "Success", nil
	}
	return paymentIntent.ClientSecret, "", errors.New("unknown payment status")
}

func (*paymentService) ConfirmPaymentIntent(userId, token string, payments PaymentIntentSchema) (string, error) {
	stripe.Key = viper.GetString("stripe.key")
	if payments.PaymentIntentID != "" {
		_, err := paymentintent.Confirm(payments.PaymentIntentID, nil)
		if err != nil {
			return "success", nil
		}
		return "unable to confirm payment intent", err
	}
	return "", errors.New("provide a valid intent id")
}

func createPaymentIntentParams(payments PaymentIntentSchema, c *stripe.Customer) stripe.PaymentIntentParams {
	params := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(payments.Amount),
		Currency:           stripe.String(payments.Currency),
		PaymentMethod:      stripe.String(payments.PaymentMethodID),
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		Description:        stripe.String(payments.Description),
		ReceiptEmail:       stripe.String(payments.ReceiptEmail),
		ConfirmationMethod: stripe.String(string(stripe.PaymentIntentConfirmationMethodManual)),
		Confirm:            stripe.Bool(true),
		Customer:           stripe.String(c.ID),
		SetupFutureUsage:   stripe.String(string(stripe.PaymentIntentSetupFutureUsageOffSession)),
	}
	return *params
}
func createConfirmPaymentIntentParams(payments PaymentIntentSchema, c string) stripe.PaymentIntentParams {
	p := &stripe.PaymentMethodListParams{
		Customer: stripe.String(c),
		Type:     stripe.String(string(stripe.PaymentMethodTypeCard)),
	}
	i := paymentmethod.List(p)
	var pm *stripe.PaymentMethod
	for i.Next() {
		pm = i.PaymentMethod()
	}
	params := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(payments.Amount),
		Currency:           stripe.String(payments.Currency),
		Description:        stripe.String(payments.Description),
		PaymentMethod:      stripe.String(pm.ID),
		ReceiptEmail:       stripe.String(payments.ReceiptEmail),
		Customer:           stripe.String(c),
		Confirm:            stripe.Bool(true),
		ConfirmationMethod: stripe.String(string(stripe.PaymentIntentConfirmationMethodManual)),
	}
	return *params
}

func getCustomerDetails(userId string) (entity.PaymentDB, error) {
	payment, err := repo.FindOne(bson.M{"user_id": userId}, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"Get Payments Details section",
			err,
		)
		return payment, err
	}
	return payment, nil
}
