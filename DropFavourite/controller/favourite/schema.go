package favourite

import entity "Drop/DropFavourite/entities"

type FavouriteService interface {
	AddFavourite(itemID, userId string) (string, error)
	DeleteFavourite(itemID, userId string) error
	GetFavourite(userId string, limit, skip int) ([]entity.ItemDB, error)
}
