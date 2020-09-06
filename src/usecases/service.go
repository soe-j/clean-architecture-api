package usecases

import (
	"github.com/soe-j/clean-architecture-api/src/entities"
)

// Service ...
type Service struct{}

// ServiceResponse ...
type ServiceResponse map[string]interface{}

// Index ...
func (s *Service) Index() ServiceResponse {
	company := &entities.Company{
		ID:      1,
		Code:    "GITHUB",
		Name:    "GitHub",
		Address: "USA",
	}
	user := &entities.User{
		ID:        1,
		Email:     "soe-j@github.com",
		Name:      "soe-j",
		CompanyID: company.ID,
	}

	return ServiceResponse{
		"company": company,
		"user":    user,
	}
}
