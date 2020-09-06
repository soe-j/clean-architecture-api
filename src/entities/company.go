package entities

type CompanyID int32

type Company struct {
	ID      CompanyID
	Code    string
	Name    string
	Address string
}
