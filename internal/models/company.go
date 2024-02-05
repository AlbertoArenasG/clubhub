package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Location struct {
	City    string `json:"city" bson:"city"`
	Country string `json:"country" bson:"country"`
	Address string `json:"address" bson:"address"`
	ZipCode string `json:"zip_code" bson:"zip_code"`
}

type Contact struct {
	Email    string   `json:"email" bson:"email"`
	Phone    string   `json:"phone" bson:"phone"`
	Location Location `json:"location" bson:"location"`
}

type Owner struct {
	FirstName string  `json:"first_name" bson:"first_name"`
	LastName  string  `json:"last_name" bson:"last_name"`
	Contact   Contact `json:"contact" bson:"contact"`
}

type Franchise struct {
	Name     string   `json:"name" bson:"name"`
	URL      string   `json:"url" bson:"url"`
	Location Location `json:"location" bson:"location"`
}

type Information struct {
	Name      string   `json:"name" bson:"name"`
	TaxNumber string   `json:"tax_number" bson:"tax_number"`
	Location  Location `json:"location" bson:"location"`
}

type Company struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Owner       Owner              `json:"owner" bson:"owner"`
	Information Information        `json:"information" bson:"information"`
	Franchises  []Franchise        `json:"franchises" bson:"franchises"`
	CreatedAt   time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
