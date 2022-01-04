package coupon

import (
	entity "Drop/DropCoupons/entities"

	"Drop/DropCoupons/repository/coupon"
	"errors"
	"time"

	"github.com/aekam27/trestCommon"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = coupon.NewCouponRepository("coupon")
)

type couponService struct{}

func NewCouponService(repository coupon.CouponRepository) CouponService {
	repo = repository
	return &couponService{}
}

func (r *couponService) AddCoupon(coupon Coupon, token string) (string, error) {
	var couponEntity entity.CouponDB
	couponEntity.ID = primitive.NewObjectID()
	couponEntity.Status = "Active"
	couponEntity.CreatedTime = time.Now()
	couponEntity.CouponCode = coupon.CouponCode
	couponEntity.Description = coupon.Description
	couponEntity.DiscountPercentage = coupon.DiscountPercentage
	couponEntity.MaxDiscount = coupon.MaxDiscount
	couponEntity.MaximumUsage = coupon.MaximumUsage
	couponEntity.UsagePerDay = coupon.UsagePerDay
	couponEntity.ValidAmount = coupon.ValidAmount
	return repo.InsertOne(couponEntity)
}

func (*couponService) UpdateCoupon(coupon Coupon, couponId string) (string, error) {
	if couponId == "" {
		err := errors.New("Coupon id missing")
		trestCommon.ECLog2(
			"update Coupon location",
			err,
		)
		return "", err
	}
	id, _ := primitive.ObjectIDFromHex(couponId)
	_, err := checkByCouponID(id)
	if err != nil {
		return "", errors.New("invalid Coupon Id")
	}
	setParameters := bson.M{}
	if coupon.Status != "" {
		setParameters["status"] = coupon.Status
	}
	if coupon.Description != "" {
		setParameters["description"] = coupon.Description
	}
	if coupon.DiscountPercentage != "" {
		setParameters["discount_percentage"] = coupon.DiscountPercentage
	}
	if coupon.MaxDiscount != "" {
		setParameters["max_discount"] = coupon.MaxDiscount
	}
	if coupon.MaximumUsage != "" {
		setParameters["max_usage"] = coupon.MaxDiscount
	}
	if coupon.CouponCode != "" {
		setParameters["coupon_code"] = coupon.CouponCode
	}
	if coupon.ValidAmount != "" {
		setParameters["valid_amount"] = coupon.ValidAmount
	}
	if coupon.ValidAmount != "" {
		setParameters["usage_per_day"] = coupon.UsagePerDay
	}
	setParameters["updated_time"] = time.Now()
	filter := bson.M{"_id": id}
	set := bson.M{
		"$set": setParameters,
	}
	result, err := repo.UpdateOne(filter, set)
	if err != nil {
		trestCommon.ECLog3(
			"update Coupon location",
			err,
			logrus.Fields{
				"Coupon_id": couponId,
			})
		return "", err
	}
	return result, nil
}

func checkByCouponID(id primitive.ObjectID) (entity.CouponDB, error) {
	coupon, err := repo.FindOne(bson.M{"_id": id}, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"Get Coupon Details section",
			err,
		)
		return coupon, err
	}
	return coupon, nil
}

func (*couponService) GetCoupons(code, validAmt, maxDis, usagePD, maxUsage, disPer, status, cId string, limit, skip int) ([]entity.CouponDB, error) {
	filter := bson.M{}
	if code != "" {
		id, _ := primitive.ObjectIDFromHex(cId)
		filter["_id"] = id
	}
	if code != "" {
		filter["coupon_code"] = code
	}
	if validAmt != "" {
		filter["valid_amount"] = validAmt
	}
	if maxDis != "" {
		filter["max_discount"] = maxDis
	}
	if usagePD != "" {
		filter["usage_per_day"] = usagePD
	}
	if maxUsage != "" {
		filter["max_usage"] = maxUsage
	}
	if disPer != "" {
		filter["discount_percentage"] = disPer
	}
	if status != "" {
		filter["status"] = status
	}
	coupon, err := repo.Find(filter, bson.M{}, limit, skip)
	if err != nil {
		trestCommon.ECLog2(
			"Get Coupon section",
			err,
		)
		return []entity.CouponDB{}, err
	}
	return coupon, nil
}

func (*couponService) GetCoupon(code, validAmt, maxDis, usagePD, maxUsage, disPer, status, cId string) (entity.CouponDB, error) {
	filter := bson.M{}
	if code != "" {
		id, _ := primitive.ObjectIDFromHex(cId)
		filter["_id"] = id
	}
	if code != "" {
		filter["coupon_code"] = code
	}
	if validAmt != "" {
		filter["valid_amount"] = validAmt
	}
	if maxDis != "" {
		filter["max_discount"] = maxDis
	}
	if usagePD != "" {
		filter["usage_per_day"] = usagePD
	}
	if maxUsage != "" {
		filter["max_usage"] = maxUsage
	}
	if disPer != "" {
		filter["discount_percentage"] = disPer
	}
	if status != "" {
		filter["status"] = status
	}
	coupon, err := repo.FindOne(filter, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"Get Coupon section",
			err,
		)
		return entity.CouponDB{}, err
	}
	return coupon, nil
}

func (*couponService) GetCouponWithIDs(couponIds []string) ([]entity.CouponDB, error) {
	subFilter := bson.A{}
	for _, item := range couponIds {
		id, _ := primitive.ObjectIDFromHex(item)
		subFilter = append(subFilter, bson.M{"_id": id})
	}
	filter := bson.M{"$or": subFilter}
	coupons, err := repo.FindWithIDs(filter, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"Get coupon section",
			err,
		)
		return []entity.CouponDB{}, err
	}
	return coupons, nil
}
