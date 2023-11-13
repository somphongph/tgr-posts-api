package entities

import "time"

type Entity struct {
	Status    string    `json:"status"`
	CreatedBy string    `json:"createdBy"`
	CreatedOn time.Time `json:"createdOn" bson:"createdOn"`
	UpdatedBy string    `json:"updatedBy"`
	UpdatedOn time.Time `json:"updatedOn" bson:"updatedOn"`
}
