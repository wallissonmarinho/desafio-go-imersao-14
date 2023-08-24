package domain

import "gopkg.in/guregu/null.v4"

type Response struct {
	Code     null.Int    `json:"code"`
	Response interface{} `json:"response"`
}
