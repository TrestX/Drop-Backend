package category

import (
	"Drop/DropItemCategories/api"
	entity "Drop/DropItemCategories/entities"

	tag "Drop/DropItemCategories/repository"
	"Drop/DropItemCategories/repository/admin"
	"errors"
	"strings"
	"time"

	"github.com/aekam27/trestCommon"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = admin.NewAdminRepository("itemcategory")
)
var (
	repo1 = tag.NewTagRepository("itemtags")
)

type categoryService struct{}

func NewCategoryService(repository admin.AdminRepository) CategoryService {
	repo = repository
	return &categoryService{}
}
func (add *categoryService) AddCategory(Categorys *Categorys) (string, error) {
	if Categorys.CategoryName != "" {
		var CategoryEntity entity.CategoryDB
		CategoryEntity.Status = Categorys.Status
		CategoryEntity.CreatedTime = time.Now()
		CategoryEntity.ID = primitive.NewObjectID()
		CategoryEntity.CategoryName = Categorys.CategoryName
		CategoryEntity.ShopType = Categorys.ShopType
		CategoryEntity.PresignedDownloadUrl = Categorys.PresignedUrl
		return repo.InsertOne(CategoryEntity)
	} else {
		var TagEntity entity.TagDB
		TagEntity.Status = Categorys.Status
		TagEntity.CreatedTime = time.Now()
		TagEntity.ID = primitive.NewObjectID()
		TagEntity.TagName = Categorys.TagName
		TagEntity.ShopType = Categorys.ShopType
		return repo1.InsertOne(TagEntity)
	}

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
	if Categorys.CategoryName != "" {
		setParameters["category_name"] = Categorys.CategoryName
	}
	if Categorys.ShopType != "" {
		setParameters["shop_type"] = Categorys.ShopType
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
		filter["shop_type"] = stype
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

func (*categoryService) GetActiveCategorys(shoptype, name string, limit, skip int) ([]entity.CategoryDB, error) {
	filter := bson.M{}
	filter["status"] = "Active"
	if shoptype != "" {
		filter["shop_type"] = shoptype
	}
	if name != "" {
		filter["category_name"] = name
	}
	Categorys, err := repo.Find(filter, bson.M{}, limit, skip)
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

func (*categoryService) DeleteCategory(categoryID string) (string, error) {
	if categoryID == "" {
		err := errors.New("Category id missing")
		trestCommon.ECLog2(
			"Delete Category",
			err,
		)
		return "", err
	}
	id, _ := primitive.ObjectIDFromHex(categoryID)
	filter := bson.M{"_id": id}
	err := repo.DeleteOne(filter)
	if err != nil {
		trestCommon.ECLog2(
			"Delete Category",
			err,
		)
		return "", err
	}
	return "success", nil
}

func (*categoryService) GetAllTags(limit, skip int, stype string) ([]entity.TagDB, error) {
	filter := bson.M{}
	if stype != "" {
		filter["shop_type"] = stype
	}
	Categorys, err := repo1.Find(filter, bson.M{}, limit, skip)
	if err != nil {
		trestCommon.ECLog2(
			"Get Categorys section",
			err,
		)
		return Categorys, err
	}
	return Categorys, nil
}

func (*categoryService) DeleteTag(categoryID string) (string, error) {
	if categoryID == "" {
		err := errors.New("Category id missing")
		trestCommon.ECLog2(
			"Delete Category",
			err,
		)
		return "", err
	}
	id, _ := primitive.ObjectIDFromHex(categoryID)
	filter := bson.M{"_id": id}
	err := repo1.DeleteOne(filter)
	if err != nil {
		trestCommon.ECLog2(
			"Delete Category",
			err,
		)
		return "", err
	}
	return "success", nil
}
