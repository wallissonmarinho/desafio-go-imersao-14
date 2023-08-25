package domain

import "gopkg.in/guregu/null.v4"

type RouteRequest struct {
	Name        null.String `json:"name"`
	Source      Coordinates `json:"source"`
	Destination Coordinates `json:"destination"`
}

type Coordinates struct {
	Lat null.Float `json:"lat"`
	Lng null.Float `json:"lng"`
}

type RouteResponse struct {
	ID          null.String `json:"id"`
	Name        null.String `json:"name"`
	Source      Coordinates `json:"source"`
	Destination Coordinates `json:"destination"`
}
