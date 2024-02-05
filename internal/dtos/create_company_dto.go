package dtos

import (
	"github.com/AlbertoArenasG/clubhub/internal/models"
	"github.com/go-playground/validator/v10"
)

type Location struct {
	City    string `json:"city" validate:"required"`
	Country string `json:"country" validate:"required"`
	Address string `json:"address" validate:"required"`
	ZipCode string `json:"zip_code" validate:"required"`
}

type Contact struct {
	Email    string   `json:"email" validate:"required,email"`
	Phone    string   `json:"phone" validate:"required"`
	Location Location `json:"location" validate:"required"`
}

type Owner struct {
	FirstName string  `json:"first_name" validate:"required"`
	LastName  string  `json:"last_name" validate:"required"`
	Contact   Contact `json:"contact" validate:"required"`
}

type Franchise struct {
	Name     string   `json:"name" validate:"required"`
	URL      string   `json:"url" validate:"required,url"`
	Location Location `json:"location" validate:"required"`
}

type Information struct {
	Name      string   `json:"name" bson:"name"`
	TaxNumber string   `json:"tax_number" bson:"tax_number"`
	Location  Location `json:"location" validate:"required"`
}

type Company struct {
	Owner       Owner       `json:"owner" validate:"required"`
	Information Information `json:"information" validate:"required"`
	Franchises  []Franchise `json:"franchises" validate:"required"`
}

type CreateCompanyDTO struct {
	Company Company `json:"company" validate:"required"`
}

func (c *CreateCompanyDTO) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

func (c *CreateCompanyDTO) ConvertDTOToModel() (models.Company, error) {
	owner, err := convertOwnerDTOToModel(c.Company.Owner)
	if err != nil {
		return models.Company{}, err
	}

	information := models.Information{
		Name:      c.Company.Information.Name,
		TaxNumber: c.Company.Information.TaxNumber,
		Location:  convertLocationDTOToModel(c.Company.Information.Location),
	}

	franchises := make([]models.Franchise, len(c.Company.Franchises))
	for i, f := range c.Company.Franchises {
		franchises[i] = models.Franchise{
			Name:     f.Name,
			URL:      f.URL,
			Location: convertLocationDTOToModel(f.Location),
		}
	}

	return models.Company{
		Owner:       owner,
		Information: information,
		Franchises:  franchises,
	}, nil
}

func convertOwnerDTOToModel(dto Owner) (models.Owner, error) {
	return models.Owner{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Contact: models.Contact{
			Email:    dto.Contact.Email,
			Phone:    dto.Contact.Phone,
			Location: convertLocationDTOToModel(dto.Contact.Location),
		},
	}, nil
}

func convertLocationDTOToModel(dto Location) models.Location {
	return models.Location{
		City:    dto.City,
		Country: dto.Country,
		Address: dto.Address,
		ZipCode: dto.ZipCode,
	}
}
