package profile

import (
	"Drop/DropUserAccount/api"
	entity "Drop/DropUserAccount/entities"
	"Drop/DropUserAccount/repository/user"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/aekam27/trestCommon"
)

var (
	repo = user.NewProfileRepository("users")
)

type profileService struct{}

func NewProfileService(repository user.UserRepository) ProfileService {
	repo = repository
	return &profileService{}
}

func (*profileService) UpdateProfile(profile *Profile, userid string) (string, error) {
	if userid == "" {
		err := errors.New("user id missing")
		trestCommon.ECLog2(
			"update profile section",
			err,
		)
		return "", err
	}
	id, _ := primitive.ObjectIDFromHex(userid)
	setParameters := bson.M{}
	_, err := checkUser(userid)
	if err != nil {
		trestCommon.ECLog2(
			"update profile section",
			err,
		)
		return "", err
	}
	if profile.Name != "" {
		name := profile.Name
		setParameters["name"] = name
	}
	if profile.PhoneNo != "" {
		phoneNo := profile.PhoneNo
		setParameters["phone_number"] = phoneNo
	}
	if profile.Wallet != "" {
		wallet := profile.Wallet
		setParameters["wallet"] = wallet
	}
	if profile.DOB != "" {
		setParameters["dob"] = profile.DOB
	}
	if profile.Gender != "" {
		setParameters["gender"] = profile.Gender
	}
	if profile.Availability != "" {
		setParameters["availability"] = profile.Availability
	}
	if profile.ApprovalStatus != "" {
		setParameters["approval_status"] = profile.ApprovalStatus
	}
	if profile.Deleted {
		setParameters["deleted"] = true
	}
	if profile.VehicleNumber != "" {
		setParameters["vehicle_number"] = profile.VehicleNumber
	}
	if profile.VehicleType != "" {
		setParameters["vehicle_type"] = profile.VehicleType
	}
	if profile.ProfilePhoto != "" {
		setParameters["profile_photo"] = profile.ProfilePhoto
	}
	setParameters["update_time"] = time.Now()
	filter := bson.M{"_id": id}
	set := bson.M{
		"$set": setParameters,
	}
	result, err := repo.UpdateOne(filter, set)
	if err != nil {
		trestCommon.ECLog3(
			"update profile section",
			err,
			logrus.Fields{
				"user_id": userid,
				"profile": profile,
			})
		return "", err
	}

	return result, nil
}

func (*profileService) GetProfile(userID, token string) (entity.UserDB, int, int, int, int, error) {
	if userID == "" {
		err := errors.New("user id missing")
		trestCommon.ECLog2(
			"GetProfile section",
			err,
		)
		return entity.UserDB{}, 0, 0, 0, 0, err
	}
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		trestCommon.ECLog3(
			"GetProfile section",
			err,
			logrus.Fields{
				"user_id": userID,
			},
		)
		return entity.UserDB{}, 0, 0, 0, 0, err
	}
	profile, err := repo.FindOne(bson.M{"_id": id}, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"GetProfile section",
			err,
		)
		return profile, 0, 0, 0, 0, err
	}
	profile.Password = ""
	ordered := 0
	delivered := 0
	accepted := 0
	ready := 0
	if strings.ToLower(profile.AccountType) == "user" {
		ordered, _ = api.GetOrders(profile.ID.Hex(), "", token, "Ordered")
		delivered, _ = api.GetOrders(profile.ID.Hex(), "", token, "Delivered")
		accepted, _ = api.GetOrders(profile.ID.Hex(), "", token, "Accepted")
		ready, _ = api.GetOrders(profile.ID.Hex(), "", token, "Ready")
	}
	if strings.ToLower(profile.AccountType) == "delivery" {
		ordered, _ = api.GetOrders("", profile.ID.Hex(), token, "Ordered")
		delivered, _ = api.GetOrders("", profile.ID.Hex(), token, "Delivered")
		accepted, _ = api.GetOrders("", profile.ID.Hex(), token, "Accepted")
		ready, _ = api.GetOrders("", profile.ID.Hex(), token, "Ready")
	}
	return profile, ordered, delivered, accepted, ready, nil
}
func checkUser(userID string) (entity.UserDB, error) {
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		trestCommon.ECLog3(
			"CheckUser section",
			err,
			logrus.Fields{
				"user_id": userID,
			},
		)
		return entity.UserDB{}, err
	}
	return repo.FindOne(bson.M{"_id": id}, bson.M{})
}
func (*profileService) CheckPhoneNumber(phoneNumber string) (entity.UserDB, error) {
	return repo.FindOne(bson.M{"phone_no": phoneNumber}, bson.M{})

}

func (*profileService) GetAllUsers(userID, accountType, token string) ([]AdminOutput, error) {
	if userID == "" {
		err := errors.New("user id missing")
		trestCommon.ECLog2(
			"GetAllProfiles section",
			err,
		)
		return []AdminOutput{}, err
	}
	userDetails, err := checkUser(userID)
	if err != nil {
		trestCommon.ECLog2(
			"GetAllProfiles section",
			err,
		)
		return []AdminOutput{}, err
	}
	if userDetails.AccountType != "Admin" {
		trestCommon.ECLog2(
			"GetAllProfiles section",
			errors.New("User doesnot have admin privilages"),
		)
		return []AdminOutput{}, errors.New("User doesnot have admin privilages")
	}
	filter := bson.M{}
	if accountType != "" {
		filter["account_type"] = accountType
	}
	profiles, err := repo.Find(filter, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"GetAllProfiles section",
			err,
		)
		return []AdminOutput{}, err
	}
	var userIds []string
	for i := 0; i < len(profiles); i++ {
		userIds = append(userIds, profiles[i].ID.Hex())
	}
	res, err := api.GetUserOrders(userIds)
	body := formatOutput(profiles, res)
	return body, nil
}

func formatOutput(profiles []entity.UserDB, details api.OrderInteface) []AdminOutput {
	var OpList []AdminOutput
	for i := 0; i < len(profiles); i++ {
		var adminOp AdminOutput

		adminOp.ID = profiles[i].ID.Hex()
		adminOp.AccountType = profiles[i].AccountType
		adminOp.DOB = profiles[i].DOB
		adminOp.Gender = profiles[i].Gender
		adminOp.Email = profiles[i].Email
		adminOp.Name = profiles[i].Name
		adminOp.Wallet = profiles[i].Wallet
		var orderList []OrdersOP
		for j := 0; j < len(details.OrderList); j++ {
			var orders OrdersOP
			if j > 5 {
				break
			}
			if profiles[i].ID.Hex() == details.OrderList[j].UserID {
				orders.ID = details.OrderList[j].ID.Hex()
				amt := strconv.Itoa(int(details.PaymentList[j].Amount))
				orders.Status = details.OrderList[j].Status
				orders.PaymentAmount = details.PaymentList[j].Currency + " " + amt
				orders.PaymentMethod = details.PaymentList[j].PaymentMethodTypes[0]
				orders.ItemsOrderd = details.CartList[j].Items
				temp := details.PaymentList[j].Shipping.Address
				orders.DeliveryAddress = temp.Address + ", " + temp.City + ", " + temp.State + ", " + temp.Country + ", " + temp.Pin
				orders.DeliveryGeoLocation = temp.GeoLocation
				orderList = append(orderList, orders)
			}
		}
		adminOp.Orders = orderList
		OpList = append(OpList, adminOp)
	}
	return OpList
}

func (*profileService) ChangeProfileStatus(userID, status string) (string, error) {
	if userID == "" {
		err := errors.New("user id missing")
		trestCommon.ECLog2(
			"GetAllProfiles section",
			err,
		)
		return "", err
	}
	id, _ := primitive.ObjectIDFromHex(userID)
	set := bson.M{"status": status, "updated_time": time.Now()}
	filter := bson.M{"_id": id}
	result, err := repo.UpdateOne(filter, bson.M{"$set": set})
	if err != nil {
		trestCommon.ECLog3(
			"update profile section",
			err,
			logrus.Fields{})
		return "", err
	}
	return result, err
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
func (*profileService) GetAdminProfile(typ, status, stype string) ([]entity.UserDB, error) {
	filter := bson.M{}
	if typ != "" {
		filter["account_type"] = typ
	}
	if status != "" {
		filter["status"] = status
	}
	if stype != "" {
		n := bson.M{}
		e := bson.M{}
		n["name"] = stype
		e["$elemMatch"] = n
		filter["type"] = e
	}
	user, err := repo.Find(filter, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"Get banners section",
			err,
		)
		return user, err
	}
	for i := 0; i < len(user); i++ {
		newPdownloadurl := createPreSignedDownloadUrl(user[i].ProfilePhoto)
		user[i].ProfilePhoto = newPdownloadurl
		newNdownloadurl := createPreSignedDownloadUrl(user[i].NationalID)
		user[i].NationalID = newNdownloadurl
		newPidownloadurl := createPreSignedDownloadUrl(user[i].PictureID)
		user[i].PictureID = newPidownloadurl
	}
	return user, nil
}

func (*profileService) GetAdminUsersWithIDs(userIds []string) ([]entity.UserDB, error) {
	subFilter := bson.A{}
	for _, item := range userIds {
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
		return []entity.UserDB{}, err
	}
	for i := 0; i < len(users); i++ {
		newPdownloadurl := createPreSignedDownloadUrl(users[i].ProfilePhoto)
		users[i].ProfilePhoto = newPdownloadurl
		newNdownloadurl := createPreSignedDownloadUrl(users[i].NationalID)
		users[i].NationalID = newNdownloadurl
		newPidownloadurl := createPreSignedDownloadUrl(users[i].PictureID)
		users[i].PictureID = newPidownloadurl
	}
	return users, nil
}
