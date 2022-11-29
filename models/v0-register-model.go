package models

// need to fix the names for the user login models etc

// RegisterRequest example
type RegisterRequest struct {
	Username     string `json:"Username"  binding:"required"`
	UserPassword string `json:"UserPassword"  binding:"required"`
	Number       string `json:"Number"  binding:"required"`
	Email        string `json:"Email"  binding:"required"`
	FirstName    string `json:"FirstName"  binding:"required"`
	LastName     string `json:"LastName"  binding:"required"`
	Token        string `json:"Token"`
}

// RegisterResponse example

type RegisterResponse struct {
	Status string `json: "status"`
}

// type HotFixRegisterRequest struct {
// 	Username     string `json:"Username"  binding:"required"`
// 	UserPassword string `json:"UserPassword"  binding:"required"`
// 	Number       string `json:"Number"  binding:"required"`
// 	Email        string `json:"Email"  binding:"required"`
// 	FirstName    string `json:"FirstName"  binding:"required"`
// 	LastName     string `json:"LastName"  binding:"required"`
// }
