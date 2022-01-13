package ratingreview

import entity "Drop/DropRatingReview/entities"

type RatingReviewService interface {
	UpdateRatingReview(Id string, ratingReview RatingReviewSchema) (string, error)
	AddRatingReview(userId, entityId, token string, ratingReview RatingReviewSchema) (string, error)
	GetRatingsReview(userId, entityId string, limit, skip int) ([]entity.RatingReviewDB, error)
	GetReviewRatingsWithIDs(userIds []string) ([]entity.RatingReviewDB, error)
}

type RatingReviewSchema struct {
	For    string  `bson:"for" json:"for"`
	Rating float64 `bson:"rating" json:"rating,omitempty"`
	Review string  `bson:"review" json:"review,omitempty"`
}

type Shop struct {
	Rating       float64 `bson:"rating" json:"rating,omitempty"`
	NumbofRating int64   `bson:"nrating" json:"nrating,omitempty"`
}
