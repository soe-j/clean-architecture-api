package usecases

import (
	"github.com/soe-j/clean-architecture-api/src/entities"
)

// userService ...
type userService struct{}

// Index ...
func (s *userService) Index() ServiceResponse {
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
