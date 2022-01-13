package router

import (
	ratingReviewHandler "Drop/DropRatingReview/handlers/rating-review-handler"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes contains all routes
type Routes []Route

var routes = Routes{
	Route{
		"rating-review",
		"PUT",
		"/ratingsReviews/{ratingReviewID}",
		ratingReviewHandler.UpdateReviewRating,
	},
	Route{
		"rating-review",
		"GET",
		"/ratingsReviews",
		ratingReviewHandler.GetReviewsAndRatings,
	},
	Route{
		"rating-review",
		"POST",
		"/ratingsReviews/{entityID}",
		ratingReviewHandler.AddReviewRating,
	},
	Route{
		"rating-review",
		"GET",
		"/ratingsReviews/list/{reviewIds}",
		ratingReviewHandler.GetReviewRatingsWithIDs,
	},
}
