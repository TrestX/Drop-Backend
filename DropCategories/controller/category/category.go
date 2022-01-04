package category

import (
	"Drop/DropCategories/api"
	entity "Drop/DropCategories/entities"

	"Drop/DropCategories/repository/admin"

	"errors"
	"strings"
	"time"

	"github.com/aekam27/trestCommon"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = admin.NewAdminRepository("category")
)

type categoryService struct{}

var dealImageUrl = map[string]string{
	"50% off":         "https://dropappfiles.s3.amazonaws.com/test/1631515513testimg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAXCRKKLUYPCH4VXRG%2F20210913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20210913T064513Z&X-Amz-Expires=1500&X-Amz-SignedHeaders=host&X-Amz-Signature=b030747b9e0876169f06a2103641bbfeb5030f644010ecfc02bb06d205272b35",
	"Free delivery":   "https://dropappfiles.s3.amazonaws.com/test/1631515513testimg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAXCRKKLUYPCH4VXRG%2F20210913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20210913T064513Z&X-Amz-Expires=1500&X-Amz-SignedHeaders=host&X-Amz-Signature=b030747b9e0876169f06a2103641bbfeb5030f644010ecfc02bb06d205272b35",
	"Top rated":       "https://dropappfiles.s3.amazonaws.com/test/1631515513testimg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAXCRKKLUYPCH4VXRG%2F20210913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20210913T064513Z&X-Amz-Expires=1500&X-Amz-SignedHeaders=host&X-Amz-Signature=b030747b9e0876169f06a2103641bbfeb5030f644010ecfc02bb06d205272b35",
	"Fast delivery":   "https://dropappfiles.s3.amazonaws.com/test/1631515513testimg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAXCRKKLUYPCH4VXRG%2F20210913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20210913T064513Z&X-Amz-Expires=1500&X-Amz-SignedHeaders=host&X-Amz-Signature=b030747b9e0876169f06a2103641bbfeb5030f644010ecfc02bb06d205272b35",
	"Pizza":           "https://dropappfiles.s3.amazonaws.com/test/1631515513testimg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAXCRKKLUYPCH4VXRG%2F20210913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20210913T064513Z&X-Amz-Expires=1500&X-Amz-SignedHeaders=host&X-Amz-Signature=b030747b9e0876169f06a2103641bbfeb5030f644010ecfc02bb06d205272b35",
	"Burgers":         "https://dropappfiles.s3.amazonaws.com/test/1631515513testimg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAXCRKKLUYPCH4VXRG%2F20210913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20210913T064513Z&X-Amz-Expires=1500&X-Amz-SignedHeaders=host&X-Amz-Signature=b030747b9e0876169f06a2103641bbfeb5030f644010ecfc02bb06d205272b35",
	"Chinese":         "https://dropappfiles.s3.amazonaws.com/test/1631515513testimg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAXCRKKLUYPCH4VXRG%2F20210913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20210913T064513Z&X-Amz-Expires=1500&X-Amz-SignedHeaders=host&X-Amz-Signature=b030747b9e0876169f06a2103641bbfeb5030f644010ecfc02bb06d205272b35",
	"Sushi":           "https://dropappfiles.s3.amazonaws.com/test/1631515513testimg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAXCRKKLUYPCH4VXRG%2F20210913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20210913T064513Z&X-Amz-Expires=1500&X-Amz-SignedHeaders=host&X-Amz-Signature=b030747b9e0876169f06a2103641bbfeb5030f644010ecfc02bb06d205272b35",
	"Christmas deals": "https://dropappfiles.s3.amazonaws.com/test/1631515513testimg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAXCRKKLUYPCH4VXRG%2F20210913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20210913T064513Z&X-Amz-Expires=1500&X-Amz-SignedHeaders=host&X-Amz-Signature=b030747b9e0876169f06a2103641bbfeb5030f644010ecfc02bb06d205272b35",
	"Valentine deals": "https://dropappfiles.s3.amazonaws.com/test/1631515513testimg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAXCRKKLUYPCH4VXRG%2F20210913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20210913T064513Z&X-Amz-Expires=1500&X-Amz-SignedHeaders=host&X-Amz-Signature=b030747b9e0876169f06a2103641bbfeb5030f644010ecfc02bb06d205272b35",
	"Eid adha deals":  "https://dropappfiles.s3.amazonaws.com/test/1631515513testimg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAXCRKKLUYPCH4VXRG%2F20210913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20210913T064513Z&X-Amz-Expires=1500&X-Amz-SignedHeaders=host&X-Amz-Signature=b030747b9e0876169f06a2103641bbfeb5030f644010ecfc02bb06d205272b35",
	"Eid futr deals":  "https://dropappfiles.s3.amazonaws.com/test/1631515513testimg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAXCRKKLUYPCH4VXRG%2F20210913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20210913T064513Z&X-Amz-Expires=1500&X-Amz-SignedHeaders=host&X-Amz-Signature=b030747b9e0876169f06a2103641bbfeb5030f644010ecfc02bb06d205272b35",
	"Eastern deals":   "https://dropappfiles.s3.amazonaws.com/test/1631515513testimg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAXCRKKLUYPCH4VXRG%2F20210913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20210913T064513Z&X-Amz-Expires=1500&X-Amz-SignedHeaders=host&X-Amz-Signature=b030747b9e0876169f06a2103641bbfeb5030f644010ecfc02bb06d205272b35",
}

func NewCategoryService(repository admin.AdminRepository) CategoryService {
	repo = repository
	return &categoryService{}
}
func (add *categoryService) AddCategory(Categorys *Categorys) (string, error) {
	var CategoryEntity entity.CategoryDB
	CategoryEntity.Status = Categorys.Status
	CategoryEntity.CreatedTime = time.Now()
	CategoryEntity.ID = primitive.NewObjectID()
	CategoryEntity.Deal = Categorys.Deal
	CategoryEntity.Type = Categorys.Type
	CategoryEntity.DealType = Categorys.DealType
	CategoryEntity.PresignedDownloadUrl = dealImageUrl[Categorys.Deal]
	return repo.InsertOne(CategoryEntity)
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
func (*categoryService) UpdateCategoryStatus(Categorys *Categorys, Categoryid string) (string, error) {
	if Categoryid == "" {
		err := errors.New("Category id missing")
		trestCommon.ECLog2(
			"update Category location",
			err,
		)
		return "", err
	}
	id, _ := primitive.ObjectIDFromHex(Categoryid)
	_, err := checkByCategoryID(id)
	if err != nil {
		return "", errors.New("invalid Category Id")
	}
	setParameters := bson.M{}
	if Categorys.Status != "" {
		setParameters["status"] = Categorys.Status
	}
	if Categorys.Deal != "" {
		setParameters["deal"] = Categorys.Deal
	}
	if Categorys.Type != "" {
		setParameters["type"] = Categorys.Type
	}
	if Categorys.PresignedUrl != "" {
		setParameters["presignedurl"] = createPreSignedDownloadUrl(Categorys.PresignedUrl)
	}
	setParameters["updated_time"] = time.Now()
	filter := bson.M{"_id": id}
	set := bson.M{
		"$set": setParameters,
	}
	result, err := repo.UpdateOne(filter, set)
	if err != nil {
		trestCommon.ECLog3(
			"update Category location",
			err,
			logrus.Fields{
				"Category_id": Categoryid,
			})
		return "", err
	}

	return result, nil
}

func checkByCategoryID(id primitive.ObjectID) (entity.CategoryDB, error) {
	Category, err := repo.FindOne(bson.M{"_id": id}, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"Get Category Details section",
			err,
		)
		return Category, err
	}
	return Category, nil
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
func (*categoryService) GetAllCategorys(token string, limit, skip int, stype string) ([]entity.CategoryDB, error) {
	err := getUserDetails(token)
	if err != nil {
		return []entity.CategoryDB{}, err
	}
	filter := bson.M{}
	if stype != "" {
		filter["type"] = stype
	}
	Categorys, err := repo.Find(filter, bson.M{}, limit, skip)
	if err != nil {
		trestCommon.ECLog2(
			"Get Categorys section",
			err,
		)
		return Categorys, err
	}
	for i := 0; i < len(Categorys); i++ {
		newdownloadurl := createPreSignedDownloadUrl(Categorys[i].PresignedDownloadUrl)
		Categorys[i].PresignedDownloadUrl = newdownloadurl
	}
	return Categorys, nil
}

func (*categoryService) GetActiveCategorys(limit, skip int) ([]entity.CategoryDB, error) {
	Categorys, err := repo.Find(bson.M{"status": "Active"}, bson.M{}, limit, skip)
	if err != nil {
		trestCommon.ECLog2(
			"Get Categorys Details section",
			err,
		)
		return Categorys, err
	}
	for i := 0; i < len(Categorys); i++ {
		newdownloadurl := createPreSignedDownloadUrl(Categorys[i].PresignedDownloadUrl)
		Categorys[i].PresignedDownloadUrl = newdownloadurl
	}
	return Categorys, nil
}
