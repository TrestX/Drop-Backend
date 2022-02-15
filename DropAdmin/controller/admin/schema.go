package admin

import entity "Drop/DropAdmin/entities"

type AdminService interface {
	AddBanner(banners *Banners) (string, error)
	UpdateBannerStatus(banners *Banners, bannerId string) (string, error)
	GetActivebanners(limit, skip int, bannerType string) ([]entity.BannerDB, error)
	GetAllBanners(token string, limit, skip int) ([]entity.BannerDB, error)
}

type Banners struct {
	Name         string `bson:"name" json:"name,omitempty"`
	PresignedUrl string `bson:"presignedurl" json:"presignedurl,omitempty"`
	Status       string `bson:"status" json:"status,omitempty"`
	Description  string `bson:"description" json:"description,omitempty"`
	Text         string `bson:"text" json:"text,omitempty"`
	ShopName     string `bson:"shop" json:"shop,omitempty"`
	Category     string `bson:"category" json:"category,omitempty"`
	ButtonText   string `bson:"button_text" json:"button_text,omitempty"`
	Dimensions   string `bson:"dimensions" json:"dimensions,omitempty"`
}
