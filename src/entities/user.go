package entities

// UserID ...
type UserID int32

// User ...
type User struct {
	ID        UserID
	Email     string
	Name      string
	CompanyID CompanyID
}
