package category

import entity "Drop/DropCategories/entities"

type CategoryService interface {
	AddCategory(Categorys *Categorys) (string, error)
	UpdateCategoryStatus(Categorys *Categorys, CategoryId string) (string, error)
	GetActiveCategorys(limit, skip int) ([]entity.CategoryDB, error)
	GetAllCategorys(token string, limit, skip int, sType string) ([]entity.CategoryDB, error)
}

type Categorys struct {
	PresignedUrl string `bson:"presignedurl" json:"presignedurl,omitempty"`
	Status       string `bson:"status" json:"status,omitempty"`
	Type         string `bson:"type" json:"type,omitempty"`
	DealType     string `bson:"deal_type" json:"deal_type,omitempty"`
	Deal         string `bson:"deal" json:"deal,omitempty"`
}
