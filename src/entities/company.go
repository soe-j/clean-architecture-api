package entities

// CompanyID ...
type CompanyID int32

// Company ...
type Company struct {
	ID      CompanyID
	Code    string
	Name    string
	Address string
}
