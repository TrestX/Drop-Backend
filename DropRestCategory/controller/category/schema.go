package category

import entity "Drop/DropItemCategories/entities"

type CategoryService interface {
	AddCategory(Categorys *Categorys) (string, error)
	UpdateCategoryStatus(Categorys *Categorys, CategoryId string) (string, error)
	GetActiveCategorys(shoptype, name string, limit, skip int) ([]entity.CategoryDB, error)
	GetAllCategorys(token string, limit, skip int, sType string) ([]entity.CategoryDB, error)
	DeleteCategory(categoryID string) (string, error)
	GetAllTags(limit, skip int, sType string) ([]entity.TagDB, error)
	DeleteTag(categoryID string) (string, error)
}

type Categorys struct {
	PresignedUrl string `bson:"presignedurl" json:"presignedurl,omitempty"`
	Status       string `bson:"status" json:"status,omitempty"`
	ShopType     string `bson:"shop_type" json:"shop_type,omitempty"`
	CategoryName string `bson:"category_name" json:"category_name,omitempty"`
	TagName      string `bson:"tag_name" json:"tag_name,omitempty"`
}
