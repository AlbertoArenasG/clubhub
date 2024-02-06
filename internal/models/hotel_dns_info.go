package models

import (
	whoisparser "github.com/likexian/whois-parser"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelDnsInfo struct {
	ID              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CompanyID       primitive.ObjectID `json:"company_id,omitempty" bson:"company_id,omitempty"`
	Url             string             `json:"url" bson:"url"`
	Registrar       string             `json:"registrar" bson:"registrar"`
	DomainStatus    string             `json:"domainStatus" bson:"domainStatus"`
	CreatedDate     string             `json:"createdDate" bson:"createdDate"`
	ExpirationDate  string             `json:"expirationDate" bson:"expirationDate"`
	RegistrarName   string             `json:"registrarName" bson:"registrarName"`
	RegistrantName  string             `json:"registrantName" bson:"registrantName"`
	RegistrantEmail string             `json:"registrantEmail" bson:"registrantEmail"`
}

func MapWhoisDataToDnsInfo(companyID primitive.ObjectID, url string, whoisData *whoisparser.WhoisInfo) *HotelDnsInfo {
	return &HotelDnsInfo{
		CompanyID:       companyID,
		Url:             url,
		Registrar:       whoisData.Registrar.Name,
		DomainStatus:    whoisData.Domain.Status[0],
		CreatedDate:     whoisData.Domain.CreatedDate,
		ExpirationDate:  whoisData.Domain.ExpirationDate,
		RegistrarName:   whoisData.Registrar.Name,
		RegistrantName:  whoisData.Registrant.Name,
		RegistrantEmail: whoisData.Registrant.Email,
	}
}
