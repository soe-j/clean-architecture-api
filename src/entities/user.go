package entities

type UserID int32

type User struct {
	ID        UserID
	Email     string
	Name      string
	CompanyID CompanyID
}
