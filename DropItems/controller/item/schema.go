package item

import entity "Drop/DropItems/entities"

type ItemService interface {
	AddItem(item *Item, userId string) (string, error)
	UpdateItem(item *Item, userId string) (string, error)
	GetItem(shopID, category, name, typee, sellerId, search, featured string, limit, skip int) ([]entity.ItemDB, error)
	GetFeaturedItem(shopID, category, shopType string, limit, skip int) ([]entity.ItemDB, error)
	GetTopRatedItems(shopID, category, shopType string, limit, skip int) ([]entity.ItemDB, error)
	GetItemUsingID(itemId string) (entity.ItemDB, error)
	GetSellerItem(sellerId, shopID string, limit, skip int) ([]entity.ItemDB, error)
	GetItemWithIDs(itemId []string) ([]entity.ItemDB, error)
	GetShopFeaturedItem(shopID, category, shopType string, limit, skip int) ([]entity.ItemDB, error)
	GetPopularItem(shopID, category, shopType string, limit, skip int) ([]entity.ItemDB, error)
	GetItemCategoryStructured(shopID, category, name, typee, sellerId, search, featured, stypee string, limit, skip int) (map[string][]interface{}, error)
}

type Item struct {
	UserID      string            `bson:"user_id" json:"user_id,omitempty"`
	SellerID    string            `bson:"seller_id" json:"seller_id,omitempty"`
	ShopID      string            `bson:"shop_id" json:"shop_id,omitempty"`
	Category    string            `bson:"category" json:"category,omitempty"`
	ShopType    string            `bson:"shop_type" json:"shop_type"`
	Name        string            `bson:"name" json:"name,omitempty"`
	Description string            `bson:"description" json:"description,omitempty"`
	Images      []string          `bson:"images" json:"images,omitempty"`
	Approved    bool              `bson:"approved" json:"approved,omitempty"`
	Rejected    bool              `bson:"rejected" json:"rejected,omitempty"`
	AddOns      []entity.ItemAdOn `bson:"add_ons" json:"add_ons"`
	Quantity    int64             `json:"quantity,omitempty"`
	Featured    bool              `bson:"featured" json:"featured,omitempty"`
	FeaturedApp bool              `bson:"featured_app" json:"featured_app,omitempty"`
	Price       int64             `bson:"price" json:"price,omitempty"`
	Type        string            `bson:"type" json:"type,omitempty"`
	Deal        string            `bson:"deal" json:"deal,omitempty"`
	Sizes       []entity.Optname  `bson:"sizes" json:"sizes"`
	Matrix      string            `bson:"matrix" json:"matrix"`
	Choices     []entity.Choices  `bson:"choices" json:"choices"`
}

type UOPStruct struct {
	Item      entity.ItemDB           `json:"item"`
	Shop      entity.ShopDB           `json:"shop"`
	Ratings   []entity.RatingReviewDB `json:"ratings"`
	AvgRating int                     `json:"avg_rating"`
}
