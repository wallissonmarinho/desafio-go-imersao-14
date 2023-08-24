package domain

import "gopkg.in/guregu/null.v4"

type RouteRequest struct {
	Name        null.String `json:"name"`
	Source      coordinates `json:"source"`
	Destination coordinates `json:"destination"`
}

type coordinates struct {
	Lat null.Float `json:"lat"`
	Lng null.Float `json:"lng"`
}

type RouteResponse struct {
	ID          null.String `json:"id"`
	Name        null.String `json:"name"`
	Source      coordinates `json:"source"`
	Destination coordinates `json:"destination"`
}
